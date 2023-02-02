package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hxx258456/pyramidel-chain-baas/internal/localconfig"
	"github.com/hxx258456/pyramidel-chain-baas/pkg/utils/logger"
	"net/http"
)

func SetUpRouter(port *localconfig.TopLevel) {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	err := r.Run(fmt.Sprintf(":%s", port.Port))
	if err != nil {
		return
	}
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "views/404.html", nil)
	})
}
