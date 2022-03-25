package controllers

import (
	"fmt"
	"golang-structure/src/common"
	"net/http"
	"strconv"
	"time"

	"golang-structure/src/common/logger"

	topup_repository "golang-structure/src/repositories/topup"

	topup_model "golang-structure/src/models/topup"

	topup_schema "golang-structure/src/schemas/topup"

	"github.com/gofiber/fiber/v2"
)

/*
**********************************************************
Lấy chi tiết tất cả giao dịch topup trong khoảng thời gian
**********************************************************
*/
func GetTopupDetails(c *fiber.Ctx) error {
	logger.ProcessLog.Info().Msg("GetTopupDetails")
	const (
		default_limit  int = 100
		default_offset int = 0
		layout_iso         = "2006-01-02"
	)
	// Validate params input
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = default_limit
	}
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		offset = default_offset
	}
	from_date, err := time.Parse(layout_iso, c.Query("from_date"))
	if err != nil {
		var error_object []*common.Response = nil
		error_object = append(error_object, &common.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return c.Status(http.StatusBadRequest).JSON(common.HttpErrorResponse(error_object))
	}
	to_date, err := time.Parse(layout_iso, c.Query("to_date"))
	if err != nil {
		var error_object []*common.Response = nil
		error_object = append(error_object, &common.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return c.Status(http.StatusBadRequest).JSON(common.HttpErrorResponse(error_object))
	}
	total, err := topup_repository.CountBeetwenDate(from_date, to_date)
	if err != nil {
		var error_object []*common.Response = nil
		error_object = append(error_object, &common.Response{
			Code:    http.StatusInternalServerError,
			Message: "cannot count total topup",
			Data:    nil,
		})
		return c.Status(http.StatusInternalServerError).JSON(common.HttpErrorResponse(error_object))
	}
	records, err := topup_repository.GetAll(limit, offset, from_date, to_date)
	if err != nil {
		fmt.Println(err)
	}
	response_data := topup_schema.GetTopupDetailsResponse{
		Total:  total,
		Limit:  limit,
		Offset: offset,
		Body:   mapTopupModelstoSchemas(records),
	}
	response := common.HttpResponse(http.StatusOK, "success", response_data)
	return c.JSON(&response)
}

/*
********************************************
- Lấy gmv topup
********************************************
*/
func GetTopupGMV(c *fiber.Ctx) error {
	logger.ProcessLog.Info().Msg("GetTopupGMV")
	const layout_iso = "2006-01-02"
	from_date, err := time.Parse(layout_iso, c.Query("from_date"))
	if err != nil {
		var error_object []*common.Response = nil
		error_object = append(error_object, &common.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return c.Status(http.StatusBadRequest).JSON(common.HttpErrorResponse(error_object))
	}
	to_date, err := time.Parse(layout_iso, c.Query("to_date"))
	if err != nil {
		var error_object []*common.Response = nil
		error_object = append(error_object, &common.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return c.Status(http.StatusBadRequest).JSON(common.HttpErrorResponse(error_object))
	}
	gmv, err := topup_repository.GetTotalAmount(from_date, to_date)
	if err != nil {
		var error_object []*common.Response = nil
		error_object = append(error_object, &common.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return c.Status(http.StatusBadRequest).JSON(common.HttpErrorResponse(error_object))
	}
	response_data := topup_schema.TopupGMVResponse{
		GMV: gmv,
	}
	response := common.HttpResponse(http.StatusOK, "success", response_data)
	return c.JSON(response)
}

/*
************************************
Private function
************************************
*/
func mapTopupModelToSchema(model *topup_model.TopupModel) *topup_schema.TopupSchema {
	return &topup_schema.TopupSchema{
		PhoneNumber:          model.PhoneNumber,
		RecipientPhoneNumber: model.RecipientPhoneNumber,
		Amount:               model.Amount,
		Brand:                model.Brand,
		PaymentMethod:        model.PaymentMethod,
		Status:               model.Status,
		Description:          model.Description,
		TopupTime:            model.TopupTime.Format("2006-01-02 15:04:05"),
	}
}

func mapTopupModelstoSchemas(models []*topup_model.TopupModel) []*topup_schema.TopupSchema {
	var schemas []*topup_schema.TopupSchema
	for index := 0; index < len(models); index++ {
		schemas = append(schemas, mapTopupModelToSchema(models[index]))
	}
	return schemas
}
