package ibm

import (
	"github.com/gin-gonic/gin"
)

func IbmRouter(r *gin.Engine) {
	ibm_v1 := r.Group("ibm")
	{
		ibm_v1.GET("/banners", Banners)
		ibm_v1.GET("/banners/preview", Preview)
	}
}
