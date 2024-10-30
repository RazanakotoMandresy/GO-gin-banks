package adminbank

import (
	"net/http"

	"github.com/RazanakotoMandresy/go-gin-banks/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func (h handler) GetAdminInfo(ctx *gin.Context) {
	uuid, err := middleware.ExtractTokenUUID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": err.Error()})
		return
	}
	admin, err := middleware.Admin.Admin(middleware.Admin{Db: h.DB, UuidToFind: uuid})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"res": &admin})
}
