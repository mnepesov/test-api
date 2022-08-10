package handlers

import (
	"back/internal/domains"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	queryParamDate = "date"
	timeLayout     = "2006-01-02"
)

// GetPicturesHandler godoc
// @Summary      Get APODs
// @Description	 Get all astronomy pictures and by date
// @Tags         APOD
// @Accept       json
// @Produce      json
// @Param 		 date  query     string false "Filter by date. Date format - YYYY-MM-DD"
// @Success      200  {object}  domains.APOD "Object containing list of APODs"
// @Failure      400  {object}  domains.ErrorResponse "invalid date param"
// @Failure      500  {object}  domains.ErrorResponse
// @Router       /pictures [get]
func (h *Handler) GetPicturesHandler(ctx *gin.Context) {
	dateParam := ctx.Query(queryParamDate)
	var apods []domains.APOD
	if len(dateParam) > 0 {
		date, err := time.Parse(timeLayout, dateParam)
		if err != nil {
			sendErrorResponse(ctx, "invalid date param", http.StatusBadRequest)
			return
		}
		
		apods, err = h.useCase.GetAPODsByDate(ctx, date.Format(timeLayout))
		if err != nil {
			sendErrorResponse(ctx, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		var err error
		apods, err = h.useCase.GetAPODs(ctx)
		if err != nil {
			sendErrorResponse(ctx, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	
	sendSuccessResponse(ctx, apods, http.StatusOK)
}
