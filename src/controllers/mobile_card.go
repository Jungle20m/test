package controllers

import (
	"net/http"
	"strconv"
	"time"

	mobile_card_model "golang-structure/src/models/mobile_card"

	mobile_card_schema "golang-structure/src/schemas/mobile_card"

	mobile_card_repository "golang-structure/src/repositories/mobile_card"

	"golang-structure/src/common/logger"

	"golang-structure/src/common"

	"github.com/gofiber/fiber/v2"
)

/*
********************************************************
- Lấy chi tiết giao dịch đổi thẻ điện thoại theo khoảng ngày
********************************************************
*/
func GetAllMobileCardExchange(c *fiber.Ctx) error {
	logger.ProcessLog.Info().Msg("GetAllMobileCardExchange")
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
	total, err := mobile_card_repository.CountBeetwenDate(from_date, to_date)
	if err != nil {
		var error_object []*common.Response = nil
		error_object = append(error_object, &common.Response{
			Code:    http.StatusInternalServerError,
			Message: "cannot count total mobile cards",
			Data:    nil,
		})
		return c.Status(http.StatusInternalServerError).JSON(common.HttpErrorResponse(error_object))
	}
	records, err := mobile_card_repository.GetAll(limit, offset, from_date, to_date)
	if err != nil {
		var error_object []*common.Response = nil
		error_object = append(error_object, &common.Response{
			Code:    http.StatusInternalServerError,
			Message: "cannot get total mobile cards",
			Data:    nil,
		})
		return c.Status(http.StatusInternalServerError).JSON(common.HttpErrorResponse(error_object))
	}
	response_data := mobile_card_schema.GetAllDataResponse{
		Total:  total,
		Limit:  limit,
		Offset: offset,
		Body:   mapMobileCardsModelToSchemas(records),
	}
	response := common.HttpResponse(http.StatusOK, "success", response_data)
	return c.JSON(response)
}

/*
********************************************
- Lấy gmv của từng nhà mạng theo khoảng ngày
********************************************
*/
func GetMobileCardGMV(c *fiber.Ctx) error {
	logger.ProcessLog.Info().Msg("GetMobileCardGMV")
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
	records, err := mobile_card_repository.GetTotalPriceOfBrand(from_date, to_date)
	if err != nil {
		var error_object []*common.Response = nil
		error_object = append(error_object, &common.Response{
			Code:    http.StatusInternalServerError,
			Message: "cannot get gmv",
			Data:    nil,
		})
		return c.Status(http.StatusInternalServerError).JSON(common.HttpErrorResponse(error_object))
	}
	response_data := mobile_card_schema.MobileCardGMVResponse{
		Body: mapMobileCardGMVModelToSchemas(records),
	}
	response := common.HttpResponse(http.StatusOK, "success", response_data)
	return c.JSON(response)
}

/*
 Private function
*/

func mapMobileCardModelToSchema(model *mobile_card_model.MobileCardModel) *mobile_card_schema.MobileCardSchema {
	return &mobile_card_schema.MobileCardSchema{
		CustomerID:   model.CustomerID,
		PhoneNumber:  model.PhoneNumber,
		Name:         model.Name,
		Brand:        model.Brand,
		Price:        model.Price,
		Point:        model.Point,
		Quantity:     model.Quantity,
		ExchangeTime: model.ExchangeTime.Format("2006-01-02 15:04:05"),
	}
}

func mapMobileCardsModelToSchemas(models []*mobile_card_model.MobileCardModel) []*mobile_card_schema.MobileCardSchema {
	var schemas []*mobile_card_schema.MobileCardSchema
	for index := 0; index < len(models); index++ {
		schemas = append(schemas, mapMobileCardModelToSchema(models[index]))
	}
	return schemas
}

func mapMobileCardGMVModelToSchema(model *mobile_card_model.MobileCardGMVModel) *mobile_card_schema.MobileCardGMVSchema {
	return &mobile_card_schema.MobileCardGMVSchema{
		Brand: model.ID,
		GMV:   model.GMV,
	}
}

func mapMobileCardGMVModelToSchemas(models []*mobile_card_model.MobileCardGMVModel) []*mobile_card_schema.MobileCardGMVSchema {
	var schemas []*mobile_card_schema.MobileCardGMVSchema
	for index := 0; index < len(models); index++ {
		schemas = append(schemas, mapMobileCardGMVModelToSchema(models[index]))
	}
	return schemas
}
