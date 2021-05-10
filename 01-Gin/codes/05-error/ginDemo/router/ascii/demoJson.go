package ascii

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseAsciiJson(c *gin.Context) {

	data := map[string]interface{}{
		"lang": "GO语言",
		"tag":  "<br>",
		"code": 200,
	}

	c.PureJSON(http.StatusOK, data)
}
