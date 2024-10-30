package user

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/RazanakotoMandresy/go-gin-banks/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func (h handler) UpdateInfo(ctx *gin.Context) {
	body := new(updateRequest)
	if err := ctx.Bind(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	uuid, err := middleware.ExtractTokenUUID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}
	user, err := middleware.GetUserUUID(h.DB, uuid)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"err": err})
		return
	}
	if body.AppUserName == "" {
		body.AppUserName = user.AppUserName
	}
	user.AppUserName = body.AppUserName
	if body.Residance == "" {
		body.Residance = user.Residance
	}
	user.Residance = body.Residance
	user.Updated_at = time.Now()
	// save des modif
	result := h.DB.Save(&user)
	if result.Error != nil {
		strErr := result.Error.Error()
		// cannot use a real check cz the errors happen in differents languages
		if strings.ContainsAny(strErr, "23505") {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": fmt.Sprintf("cannot duplicate: -%v", strErr)})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": result.Error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": &user})

}
