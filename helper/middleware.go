package helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GinAuthSwap() gin.HandlerFunc {
	return func(c *gin.Context) {
		SwapHeaderAuth(c.Request)
	}
}

func SwapHeaderAuth(r *http.Request) {
	if r.Header["X-Lyrid-Authorization"] != nil {
		r.Header.Del("Authorization")
		r.Header.Add("Authorization", r.Header["X-Lyrid-Authorization"][0])
		r.Header.Del("X-Lyrid-Authorization")
	}
}
