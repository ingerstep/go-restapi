package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gorestapi "github.com/ingerstep/go-restapi"
)

func (h *Handler) signUp(c *gin.Context) {
	var input gorestapi.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {

}
