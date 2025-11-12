package handlers

import (
	"admin-app/search/business"
	"admin-app/search/models"
	"encoding/json"
	// "strings"

	genericConstants "omnenest-backend/src/constants"
	genericModel "omnenest-backend/src/models"
	"omnenest-backend/src/utils/logger"
	"omnenest-backend/src/utils/responseUtils"
	"omnenest-backend/src/utils/tracer"
	"omnenest-backend/src/utils/validations"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type OrderColumnCustomGetController struct {
	service *business.OrderColumnCustomGetService
}

func NewOrderColumnCustomGetController(service *business.OrderColumnCustomGetService) *OrderColumnCustomGetController {
	return &OrderColumnCustomGetController{
		service: service,
	}
}

// @Summary The user will get the columnId after he put the UserId and AdvanceOptionType
// @Tag Get
// @Description
// @Accept json
// @Produce json
// @Param source header string true "Source (MOB or WEB)" default(MOB)
// @Param request body models.BFFOrderColumnCustomGetRequest true "Order Column Customization Get Request JSON"
// @Success 200 {object} models.BFFOrderColumnCustomGetResponse "ColumnId as Response"
// @Failure 400 {object} models.ErrorAPIResponse "Bad Request: Invalid input data or validation error"
// @Failure 500 {object} models.ErrorAPIResponse "The server encountered an unexplained problem which has prevented it from executing the given request"
// @Router /api/orders/get/columns [post]
func (controller *OrderColumnCustomGetController) HandleOrderGetColumn(ctx *gin.Context) {
	spanCtx, span := tracer.AddToSpan(ctx.Request.Context(), "HandleOrderGetColumn")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	log := logger.GetLogger(ctx)

	var bffOrderColumnCustomGetRequest models.BFFOrderColumnCustomGetRequest
	if err := ctx.ShouldBindJSON(&bffOrderColumnCustomGetRequest); err != nil {
		errorMsgs := genericModel.ErrorMessage{
			Key:          err.(*json.UnmarshalTypeError).Field,
			ErrorMessage: genericConstants.JsonBindingFieldError,
		}
		log.With(zap.Error(err)).Error(err.Error())
		responseUtils.SendBadRequest(ctx, []genericModel.ErrorMessage{errorMsgs})
		return
	}

	if err := validations.GetBFFValidator(spanCtx).Struct(bffOrderColumnCustomGetRequest); err != nil {
		validationErrors, validationErrorsStr := validations.FormatValidationErrors(spanCtx, err.(validator.ValidationErrors))
		log.With(zap.Error(err)).Error(validationErrorsStr)
		responseUtils.SendBadRequest(ctx, validationErrors)
		return
	}

	columnIds, err := controller.service.ServiceOrderColumnGet(ctx, spanCtx, bffOrderColumnCustomGetRequest)
	if err != nil {
		log.With(zap.Error(err)).Error(err.Error())
		responseUtils.SendInternalServerError(ctx, err)
		return
	}
	responseUtils.SendStatusOK(ctx, genericConstants.BFFResponseSuccessMessage, columnIds)
}
