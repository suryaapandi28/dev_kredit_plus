package handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/suryaapandi28/dev_kredit_plus/internal/entity"
	"github.com/suryaapandi28/dev_kredit_plus/internal/http/binder"
	"github.com/suryaapandi28/dev_kredit_plus/internal/service"
	"github.com/suryaapandi28/dev_kredit_plus/pkg/response"
	"golang.org/x/crypto/bcrypt"
)

type AccountHandler struct {
	accountService service.AccountService
}

func NewAccountHandler(accountService service.AccountService) AccountHandler {
	return AccountHandler{accountService: accountService}
}
func generateResetCode() string {
	rand.Seed(time.Now().UnixNano())
	code := fmt.Sprintf("%06d", rand.Intn(1000000))
	return code
}

type Formattime struct {
	ExpiredTimeVerify time.Time
}

func (h *AccountHandler) Registered(c echo.Context) error {
	req := binder.AccountreateRequest{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "there is an input error"))
	}
	// Validate check for empty fields
	if req.Kode_konsumen == "" || req.Email == "" ||
		req.Password == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Cannot be empty"))
	}

	GeneretedKodeAccount := uuid.New().String()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Genereted password invalid"))
	}

	Status_account := "No_Active"
	verification_status := "No_Verify"
	resetCode := generateResetCode()

	expiredTime := time.Now().Local().Add(5 * time.Minute)

	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err)
	}

	expiredTimeInJakarta := expiredTime.In(location)

	account := &entity.Accounts{
		Kode_account:         GeneretedKodeAccount,
		Kode_konsumen:        req.Kode_konsumen,
		Email:                req.Email,
		Password:             string(hashedPassword),
		Status_account:       Status_account,
		Verification_status:  verification_status,
		Verification_code:    resetCode,
		Verification_expired: expiredTimeInJakarta,
		Device_id:            req.Device_id,
		Auditable:            entity.NewAuditable(),
	}

	createdAccount, err := h.accountService.Registered(account)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Failed to create data konsumen"))
	}

	datasuccess := map[string]interface{}{
		"kode_account": []string{createdAccount.Kode_account},
	}
	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success update admin", datasuccess))
}
func (h *AccountHandler) Verification_account(c echo.Context) error {
	req := binder.AccountreateRequest{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "there is an input error"))
	}
	// Validate check for empty fields
	if req.Kode_account == "" || req.Verification_code == "" ||
		req.Device_id == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Cannot be empty"))
	}

	timeString := "0000-00-00 00:00:00"
	layout := "2006-01-02 15:04:05"

	timenil, err := time.Parse(layout, timeString)
	if err == nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Error parsing time"))
	}

	accountdata, err := h.accountService.FindAccountByID(req.Kode_account)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if accountdata.Device_id != req.Device_id {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid Device ID"))
	}

	// Waktu pertama
	time1, err := time.Parse("2006-01-02 15:04:05", accountdata.Verification_expired.Format("2006-01-02 15:04:05"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Error parsing time 1"))
	}

	// Waktu kedua
	time2, err := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Error parsing time 1"))
	}

	if time2.After(time1) {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "kedua lebih besar dari pertama"))
	} else if time2.Before(time1) {

		if req.Verification_code != accountdata.Verification_code {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid code verification"))
		}

		verifyaccount := &entity.Accounts{
			Kode_account:         req.Kode_account,
			Status_account:       "Active",
			Verification_status:  "Verify",
			Verification_code:    " ",
			Verification_expired: timenil,
			Device_id:            req.Device_id,
		}
		verifyaccountdata, err := h.accountService.Verification_account(verifyaccount)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
		}
		user, err := h.accountService.LoginAccount(verifyaccount.Kode_account, verifyaccountdata.Email, verifyaccountdata.Password)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
		}
		return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "Success Verification Account", user))
	} else {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Expired Verification account"))
	}

}
