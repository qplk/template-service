package profile

import (
	"net/http"

	"template-service/impl/auth"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UserOrdersHandler(ctx *gin.Context) {

	authUser := auth.GetAuthUserFromContext(ctx)

	if authUser == nil {
		ctx.JSON(http.StatusUnprocessableEntity, "")
		return
	}

	orders := h.personService.GetUserOrders(authUser)

	ctx.JSON(http.StatusOK, orders)
}
