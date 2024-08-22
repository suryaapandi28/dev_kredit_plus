package handler

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/suryaapandi28/dev_kredit_plus/internal/entity"
	"github.com/suryaapandi28/dev_kredit_plus/internal/http/binder"
	"github.com/suryaapandi28/dev_kredit_plus/internal/service"
	"github.com/suryaapandi28/dev_kredit_plus/pkg/response"
)

type KonsumenHandler struct {
	konsumenService service.KonsumenService
}

func NewKonsumenHandler(konsumenService service.KonsumenService) KonsumenHandler {
	return KonsumenHandler{konsumenService: konsumenService}
}

func (h *KonsumenHandler) CreateUser(c echo.Context) error {
	req := binder.KonsumenCreateRequest{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "there is an input error"))
	}
	// Validate check for empty fields
	if req.NIK == "" || req.Full_name == "" ||
		req.Legal_name == "" || req.Tempat_lahir == "" || req.Tanggal_lahir == "" ||
		req.Gaji == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Cannot be empty"))
	}

	konsumennikdata, err := h.konsumenService.FindNIKKonsumen(req.NIK)
	convertnik := base64.StdEncoding.EncodeToString([]byte(req.NIK))

	if err != nil {
		return c.JSON(http.StatusFound, response.ErrorResponse(http.StatusFound, "We are sorry. Invalid find data"))
	}
	if konsumennikdata.NIK == convertnik {
		return c.JSON(http.StatusFound, response.ErrorResponse(http.StatusFound, "We are sorry. Your data has been registered"))
	}

	filektp, err := c.FormFile("foto_ktp")
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Failed to retrieve image KTP"))
	}

	fileselfie, err := c.FormFile("foto_selfie")
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Failed to retrieve image KTP"))
	}

	// Check image format
	chckFormatktp := strings.ToLower(filepath.Ext(filektp.Filename))
	if chckFormatktp != ".jpg" && chckFormatktp != ".jpeg" && chckFormatktp != ".png" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid image format. Only jpg, jpeg, and png are allowed"))
	}

	// Check image format
	chckFormatselfie := strings.ToLower(filepath.Ext(fileselfie.Filename))
	if chckFormatselfie != ".jpg" && chckFormatselfie != ".jpeg" && chckFormatselfie != ".png" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid image format. Only jpg, jpeg, and png are allowed"))
	}

	srcktp, err := filektp.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Failed to open image"))
	}
	defer srcktp.Close()

	srcselfie, err := fileselfie.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Failed to open image"))
	}
	defer srcselfie.Close()

	imageID := uuid.New()
	imageFilenamektp := fmt.Sprintf("%s%s", imageID, filepath.Ext(filektp.Filename))
	imagePathktp := filepath.Join("assets", "imagesktp", imageFilenamektp)

	imageFilenameselfie := fmt.Sprintf("%s%s", imageID, filepath.Ext(fileselfie.Filename))
	imagePathselfie := filepath.Join("assets", "imagesselfie", imageFilenameselfie)

	dstktp, err := os.Create(imagePathktp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Failed to create image file"))
	}
	defer dstktp.Close()

	dstselfie, err := os.Create(imagePathselfie)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Failed to create image file"))
	}
	defer dstselfie.Close()

	if _, err := io.Copy(dstktp, srcktp); err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Failed to copy image file"))
	}

	if _, err := io.Copy(dstselfie, srcselfie); err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Failed to copy image file"))
	}

	convertktpfile := base64.StdEncoding.EncodeToString([]byte(imageFilenamektp))
	convertselfiefile := base64.StdEncoding.EncodeToString([]byte(imageFilenameselfie))
	GeneretedKodeKonsumen := uuid.New().String()

	// var decodedByte, _ = base64.StdEncoding.DecodeString(convertktpfile)
	// var decodedString = string(decodedByte)

	konsumen := &entity.Konsumens{
		Kode_konsumen: GeneretedKodeKonsumen,
		NIK:           convertnik,
		Full_name:     req.Full_name,
		Legal_name:    req.Legal_name,
		Tempat_lahir:  req.Tempat_lahir,
		Tanggal_lahir: req.Tanggal_lahir,
		Gaji:          req.Gaji,
		Foto_ktp:      convertktpfile,
		Foto_selfie:   convertselfiefile,
		Kredit_score:  "A",
		Auditable:     entity.NewAuditable(),
	}

	createdKonsumen, err := h.konsumenService.CreateKonsumen(konsumen)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Failed to create data konsumen"))
	}

	datasuccess := map[string]interface{}{
		"kode_konsumen": []string{createdKonsumen.Kode_konsumen},
	}

	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "Konsumen created successfully", datasuccess))
}

func (h *KonsumenHandler) CheckLimitTenor(c echo.Context) error {
	req := binder.KonsumenCreateRequest{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "there is an input error"))
	}
	// Validate check for empty fields
	if req.Kode_konsumen == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Cannot be empty"))
	}

	konsumennikdata, err := h.konsumenService.CheckLimitTenor(req.Kode_konsumen)

	if err != nil {
		return c.JSON(http.StatusFound, response.ErrorResponse(http.StatusFound, "We are sorry. Invalid find data"))
	}
	if konsumennikdata.Kredit_score == "A" {
		Maxlimit := 5000000
		tenorenam := int(Maxlimit * 100 / 100)
		tenortiga := int(Maxlimit * 75 / 100)
		tenordua := int(Maxlimit * 60 / 100)
		tenorpertama := Maxlimit * 50 / 100
		tenorpertamafinal := strconv.Itoa(tenorpertama)
		tenorkeduafinal := strconv.Itoa(tenordua)
		tenortigafinal := strconv.Itoa(tenortiga)
		tenorenamfinal := strconv.Itoa(tenorenam)

		konsumenlimit := &entity.Tenor_limits{
			Kode_tenor:    uuid.New().String(),
			Kode_konsumen: req.Kode_konsumen,
			Nama_konsumen: konsumennikdata.Full_name,
			Tenor_pertama: tenorpertamafinal,
			Tenor_Kedua:   tenorkeduafinal,
			Tenor_Ketiga:  tenortigafinal,
			Tenor_keenam:  tenorenamfinal,
			Auditable:     entity.NewAuditable(),
		}

		limittenordata, err := h.konsumenService.FindLimitTenor(req.Kode_konsumen)

		if err != nil {
			return c.JSON(http.StatusFound, response.ErrorResponse(http.StatusFound, "We are sorry. Invalid find data"))
		}
		if limittenordata.Kode_konsumen != uuid.Nil.String() {
			createdKonsumenLimit, err := h.konsumenService.CreateKonsumenLimit(konsumenlimit)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Failed to create data limit tenor"))
			}
			datasuccess := map[string]interface{}{
				"nama_konsumen": []string{createdKonsumenLimit.Nama_konsumen},
				"tenor_pertama": []string{tenorpertamafinal},
				"tenor_kedua":   []string{tenorkeduafinal},
				"tenor_ketiga":  []string{tenortigafinal},
				"tenor_keenam":  []string{tenorenamfinal},
			}
			return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "Limit Tenor Update", datasuccess))
		}

	}

	if konsumennikdata.Kredit_score == "B" {
		Maxlimit := 4000000
		tenorenam := int(Maxlimit * 100 / 100)
		tenortiga := int(Maxlimit * 75 / 100)
		tenordua := int(Maxlimit * 60 / 100)
		tenorpertama := Maxlimit * 50 / 100
		tenorpertamafinal := strconv.Itoa(tenorpertama)
		tenorkeduafinal := strconv.Itoa(tenordua)
		tenortigafinal := strconv.Itoa(tenortiga)
		tenorenamfinal := strconv.Itoa(tenorenam)

		konsumenlimit := &entity.Tenor_limits{
			Kode_tenor:    uuid.New().String(),
			Kode_konsumen: req.Kode_konsumen,
			Nama_konsumen: konsumennikdata.Full_name,
			Tenor_pertama: tenorpertamafinal,
			Tenor_Kedua:   tenorkeduafinal,
			Tenor_Ketiga:  tenortigafinal,
			Tenor_keenam:  tenorenamfinal,
			Auditable:     entity.NewAuditable(),
		}

		limittenordata, err := h.konsumenService.FindLimitTenor(req.Kode_konsumen)

		if err != nil {
			return c.JSON(http.StatusFound, response.ErrorResponse(http.StatusFound, "We are sorry. Invalid find data"))
		}
		if limittenordata.Kode_konsumen != uuid.Nil.String() {
			createdKonsumenLimit, err := h.konsumenService.CreateKonsumenLimit(konsumenlimit)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Failed to create data limit tenor"))
			}
			datasuccess := map[string]interface{}{
				"nama_konsumen": []string{createdKonsumenLimit.Nama_konsumen},
				"tenor_pertama": []string{tenorpertamafinal},
				"tenor_kedua":   []string{tenorkeduafinal},
				"tenor_ketiga":  []string{tenortigafinal},
				"tenor_keenam":  []string{tenorenamfinal},
			}
			return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "Limit Tenor Update", datasuccess))
		}

	}
	if konsumennikdata.Kredit_score == "C" {
		Maxlimit := 3000000
		tenorenam := int(Maxlimit * 100 / 100)
		tenortiga := int(Maxlimit * 75 / 100)
		tenordua := int(Maxlimit * 60 / 100)
		tenorpertama := Maxlimit * 50 / 100
		tenorpertamafinal := strconv.Itoa(tenorpertama)
		tenorkeduafinal := strconv.Itoa(tenordua)
		tenortigafinal := strconv.Itoa(tenortiga)
		tenorenamfinal := strconv.Itoa(tenorenam)

		konsumenlimit := &entity.Tenor_limits{
			Kode_tenor:    uuid.New().String(),
			Kode_konsumen: req.Kode_konsumen,
			Nama_konsumen: konsumennikdata.Full_name,
			Tenor_pertama: tenorpertamafinal,
			Tenor_Kedua:   tenorkeduafinal,
			Tenor_Ketiga:  tenortigafinal,
			Tenor_keenam:  tenorenamfinal,
			Auditable:     entity.NewAuditable(),
		}

		limittenordata, err := h.konsumenService.FindLimitTenor(req.Kode_konsumen)

		if err != nil {
			return c.JSON(http.StatusFound, response.ErrorResponse(http.StatusFound, "We are sorry. Invalid find data"))
		}
		if limittenordata.Kode_konsumen != uuid.Nil.String() {
			createdKonsumenLimit, err := h.konsumenService.CreateKonsumenLimit(konsumenlimit)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Failed to create data limit tenor"))
			}
			datasuccess := map[string]interface{}{
				"nama_konsumen": []string{createdKonsumenLimit.Nama_konsumen},
				"tenor_pertama": []string{tenorpertamafinal},
				"tenor_kedua":   []string{tenorkeduafinal},
				"tenor_ketiga":  []string{tenortigafinal},
				"tenor_keenam":  []string{tenorenamfinal},
			}
			return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "Limit Tenor Update", datasuccess))
		}

	}
	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "Konsumen created successfully", konsumennikdata))
}
