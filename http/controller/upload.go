package controller

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"static_service/http/buz_code"
	"static_service/http/controller/request"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func BindJSON(ctx *gin.Context, req interface{}) (err error) {
	if err = ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": buz_code.CODE_INVALID_ARGS,
			"msg":  fmt.Sprintf("invalid params %s\n", err.Error()),
		})
	}
	return
}
func BindQuery(ctx *gin.Context, form interface{}) (err error) {
	if err = ctx.BindWith(form, binding.FormMultipart); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": buz_code.CODE_INVALID_ARGS,
			"msg":  fmt.Sprintf("invalid params %s\n", err.Error()),
		})
	}
	return
}
func UploadImg(c *gin.Context) {
	req := request.UploadImg{}
	err := BindQuery(c, &req)
	if err != nil {
		return
	}
	paths := map[string]string{}
	for _, f := range req.Imgs {
		fileExt := strings.ToLower(path.Ext(f.Filename))
		if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" {
			c.JSON(200, gin.H{
				"code": 400,
				"MSG":  "upload failed! Only PNG, JPG, GIF, JPEG files are allowed",
			})
			return
		}
		fileDir := os.Getenv("STATIC_ROOT") + "/img/"
		isExist := false
		if !isExist {
			os.Mkdir(fileDir, os.ModePerm)
		}
		filepath := fmt.Sprintf("%s%s%s", fileDir, f.Filename+"_"+req.AppID, fileExt)
		err := c.SaveUploadedFile(f, filepath)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": buz_code.CODE_SERVER_ERROR,
				"msg":  "server error",
			})
			return
		}
		paths[f.Filename] = filepath
	}

	c.JSON(200, gin.H{
		"code": buz_code.CODE_OK,
		"msg":  "ok",
		"data": paths,
	})
}
