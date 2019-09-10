package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func init() {
	e := GetEngine()
	e.POST("/version", InsertVersion)
	e.POST("/detailversions", InsertDetailVersion)
	e.GET("/version/:version/:subversion", GetVersion)
	e.GET("/edc/:filename", streamBMP)
}

func streamBMP(c *gin.Context) {
	filepath := "assets/" + c.Param("filename")
	c.File(filepath)
}

//InsertVersion is
func InsertVersion(c *gin.Context) {
	var version Version
	c.Bind(&version)
	if !insert(&version) {
		c.JSON(200, gin.H{"ACK": "NOK", "message": "tidak dapat insert to database"})
	} else {
		c.JSON(200, version)
	}
}

//InsertDetailVersion is
func InsertDetailVersion(c *gin.Context) {
	var detailversion Detailversion
	c.Bind(&detailversion)
	if !insert(&detailversion) {
		c.JSON(200, gin.H{"ACK": "NOK", "message": "tidak dapat insert to database"})
	} else {
		c.JSON(200, detailversion)
	}
}

//GetVersion is
func GetVersion(c *gin.Context) {
	var version Version

	version.VersionName = c.Param("version")
	subversion, _ := strconv.Atoi(c.Param("subversion"))
	if err := Select(&version); err != nil {
		c.JSON(500, err)
	} else {
		if err := db().Model(&version).Related(&version.Detailversions).Error; err != nil {
			c.JSON(500, err)
		} else {
			//c.JSON(200, version)
			i := 0
			var temp string
			response := map[string]string{}
			for i < len(version.Detailversions)-subversion {
				detail := version.Detailversions[i+subversion]
				i++
				temp = fmt.Sprintf("from%d", i)
				response[temp] = detail.From
				temp = fmt.Sprintf("to%d", i)
				response[temp] = detail.To
				temp = fmt.Sprintf("size%d", i)
				response[temp] = fmt.Sprintf("%d", detail.Filesize)
			}

			response["total"] = fmt.Sprintf("%d", len(version.Detailversions)-subversion)
			var newVersion Version

			if err := db().Where("id > ?", version.ID).First(&newVersion).Error; err != nil {
				response["version"] = version.VersionName

			} else {
				response["version"] = newVersion.VersionName
			}
			c.JSON(200, response)

		}
	}
}
