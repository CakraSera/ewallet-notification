package api

import (
	"net/http"

	"ewallet-notification/helpers"
	"ewallet-notification/internal/interfaces"

	"github.com/gin-gonic/gin"
)

type Healthcheck struct {
	HealthcheckServices interfaces.IHealthycheckServices
}

func (api *Healthcheck) HealthcheckHandlerHTTP(c *gin.Context) {
	msg, err := api.HealthcheckServices.HealthcheckServices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, msg, nil)
}
