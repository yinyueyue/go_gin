package image

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetImage(c *gin.Context)  {

	response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}

	reader := response.Body
	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")

	extraHeaders := map[string]string{
		//"Content-Disposition": `attachment; filename="gopher.png"`,  //下载图片
		"content-type": "image/png",   //显示在浏览器
	}

	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)



}
