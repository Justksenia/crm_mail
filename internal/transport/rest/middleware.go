package rest

import "github.com/gin-gonic/gin"

func (h *Handler) corsMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

}

func (h *Handler) optionsMessage(c *gin.Context) {
	h.corsMiddleware(c)
	c.Status(200)
}