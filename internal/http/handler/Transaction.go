package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/suryaapandi28/dev_kredit_plus/internal/http/binder"
	"github.com/suryaapandi28/dev_kredit_plus/internal/service"
	"github.com/suryaapandi28/dev_kredit_plus/pkg/response"
)

type TransactionHandler struct {
	konsumenService    service.KonsumenService
	transactionService service.TransactionService
}

func NewTransactionHandler(konsumenService service.KonsumenService, transactionService service.TransactionService) TransactionHandler {
	return TransactionHandler{konsumenService: konsumenService, transactionService: transactionService}
}

func (h *TransactionHandler) Createtransaction(c echo.Context) error {
	req := binder.TransactionCreateRequest{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "there is an input error"))
	}
	// Validate check for empty fields
	if req.Kode_konsumen == "" || req.OTR == "" ||
		req.Nama_aset == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Cannot be empty"))
	}

	limittenordata, err := h.konsumenService.FindLimitTenor(req.Kode_konsumen)

	if err != nil {
		return c.JSON(http.StatusFound, response.ErrorResponse(http.StatusFound, "We are sorry. Invalid find data"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "Konsumen created successfully", limittenordata))
}
