package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	queryParamAdd = "add"
)

// GetAPODHandler godoc
// @Summary      Get APOD
// @Description	 Get Astronomy Picture of the Day
// @Tags         APOD
// @Accept       json
// @Produce      json
// @Param 		 add  query     int true "Adding to database. Will be added - 1; Will not be added - 0."
// @Success      200  {object}  domains.NasaAPOD
// @Failure      400  {object}  domains.ErrorResponse
// @Failure      500  {object}  domains.ErrorResponse
// @Router       /apod [get]
func (h *Handler) GetAPODHandler(ctx *gin.Context) {
	addParam, err := strconv.Atoi(ctx.Query(queryParamAdd))
	if err != nil {
		sendErrorResponse(ctx, "invalid add param", http.StatusBadRequest)
		return
	}
	
	apod, err := h.useCase.GetAPODFromNasa(ctx)
	if err != nil {
		sendErrorResponse(ctx, err.Error(), http.StatusInternalServerError)
		return
	}
	
	if addParam == 1 {
		if err := h.useCase.AddAPOD(ctx, apod); err != nil {
			sendErrorResponse(ctx, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	
	sendSuccessResponse(ctx, apod, http.StatusOK)
}
