package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

var engine *gin.Engine

func init() {
	engine = gin.New()
	engine.GET("/github/image/pulls", func(context *gin.Context) {
		response := map[string]any{
			"message": "120",
		}
		marshal, _ := json.Marshal(response)
		_, _ = context.Writer.Write(marshal)
	})
}

func Index(w http.ResponseWriter, r *http.Request) {
	engine.ServeHTTP(w, r)
}
