package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetOrgID(c *gin.Context) int64 {
	claims := c.MustGet("claims").(*Claims)
	orgID, _ := strconv.ParseInt(claims.OrgID, 10, 64)

	return orgID
}
