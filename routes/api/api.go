package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
func Pong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
