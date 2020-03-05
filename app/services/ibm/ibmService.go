package ibm

import (
	"net/http"

	"github.com/eduwill/jarvis-api/app/base/config"
	"github.com/eduwill/jarvis-api/app/common/dbTemplate"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

var mapperPrefix = "ibm."

/*
	배너 목록조회
*/
func Banners(c *gin.Context) {
	params := make(map[string]interface{})
	svcCd := c.DefaultQuery("svcCd", "")
	db := config.GetGanagosiDB()

	if svcCd != "" {
		params["svcCd"] = svcCd
	}

	banners, _ := dbTemplate.SelectList(db, mapperPrefix+"banners", params)
	for _, banner := range banners {
		if banner["linkCount"] != "0" {
			links, _ := dbTemplate.SelectList(db, mapperPrefix+"links", banner)
			banner["links"] = links
		}
	}
	c.Render(http.StatusOK, render.IndentedJSON{Data: banners})
}

func Preview(c *gin.Context) {
	params := make(map[string]interface{})
	svcCd := c.DefaultQuery("svcCd", "")
	bnrNo := c.DefaultQuery("bnrNo", "0")
	db := config.GetGanagosiDB()

	if svcCd != "" && bnrNo != "" {
		params["svcCd"] = svcCd
		params["bnrNo"] = bnrNo
	}

	banners, _ := dbTemplate.SelectList(db, mapperPrefix+"banners-preview", params)
	for _, banner := range banners {
		if banner["linkCount"] != "0" {
			links, _ := dbTemplate.SelectList(db, mapperPrefix+"links", banner)
			banner["links"] = links
		}
	}
	c.Render(http.StatusOK, render.IndentedJSON{Data: banners})
}
