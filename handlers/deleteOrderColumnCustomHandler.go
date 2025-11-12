package handlers

import (
	"admin-app/search/business"
	"admin-app/search/models"
	"encoding/json"
	genericConstants "omnenest-backend/src/constants"
	genericModel "omnenest-backend/src/models"
	"omnenest-backend/src/utils/logger"
	"omnenest-backend/src/utils/responseUtils"
	"omnenest-backend/src/utils/tracer"
	"omnenest-backend/src/utils/validations"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type OrderColumnDeleteCustomController struct {
	service *business.OrderColumnDeleteCustomService
}

func NewOrderColumnDeleteCustomController(service *business.OrderColumnDeleteCustomService) *OrderColumnDeleteCustomController {
	return &OrderColumnDeleteCustomController{
		service: service,
	}
}

// @Summary This API is used to get equity group by exchange name.
// @Description This API is used to get equity group by exchange name.
// @Tags Equity
// @Accept  json
// @Produce  json
// @Param source header string true "Source (MOB or WEB)" default(MOB)
// @Param request body models.BFFOrderColumnDeleteRequest true "Order Column Customization Delete Request JSON"
// @Success 200 {object} models.BFFOrderColumnDeleteResponse "Successful response"
// @Failure 400 {object} models.ErrorAPIResponse "Bad Request: Invalid input data or validation error"
// @Failure 404 {object} models.ErrorAPIResponse "No Data Found: Data is unavialable"
// @Failure 204 "No Content for the request"
// @Failure 500 {object} models.ErrorAPIResponse "The server encountered an unexplained problem which has prevented it from executing the given request"
// @Router /api/orders/columns/remove [post]
func (controller *OrderColumnDeleteCustomController) HandleOrderColumnDelete(ctx *gin.Context) {
	spanCtx, span := tracer.AddToSpan(ctx.Request.Context(), "HandleOrderColumnDelete")
	defer func() {
		if span != nil {
			span.End()
		}
	}()

	log := logger.GetLogger(ctx)

	var bffOrderColumnDeleteRequest models.BFFOrderColumnDeleteRequest
	if err := ctx.ShouldBindJSON(&bffOrderColumnDeleteRequest); err != nil {
		errorMsgs := genericModel.ErrorMessage{Key: err.(*json.UnmarshalTypeError).Field, ErrorMessage: genericConstants.JsonBindingFieldError}
		log.With(zap.Error(err)).Error(err.Error())
		responseUtils.SendBadRequest(ctx, []genericModel.ErrorMessage{errorMsgs})
		return
	}

	if err := validations.GetBFFValidator(spanCtx).Struct(bffOrderColumnDeleteRequest); err != nil {
		validationErrors, validationErrorsStr := validations.FormatValidationErrors(spanCtx, err.(validator.ValidationErrors))
		log.With(zap.Error(err)).Error(validationErrorsStr)
		responseUtils.SendBadRequest(ctx, validationErrors)
		return
	}

	err := controller.service.ServiceOrderColumnDelete(ctx, spanCtx, bffOrderColumnDeleteRequest)
	if err != nil {
		log.With(zap.Error(err)).Error(err.Error())
		if strings.Contains(strings.ToLower(err.Error()), genericConstants.NoDataFoundError) {
			responseUtils.SendNotFoundJSON(ctx, err)
			return
		}
		responseUtils.SendInternalServerError(ctx, err)
		return
	}
	
	orderColumnDeleteResponse := models.BFFOrderColumnDeleteResponse{
		Message: genericConstants.BFFResponseSuccessMessage,
	}
	responseUtils.SendStatusOK(ctx, genericConstants.BFFResponseSuccessMessage, orderColumnDeleteResponse)
}
