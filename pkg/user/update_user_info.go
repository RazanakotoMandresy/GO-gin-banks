package user

import (
	"fmt"
	"net/http"
	"time"

	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/middleware"
	"github.com/gin-gonic/gin"
)

type UpdateRequest struct {
	AppUserName string `json:"AppUserName"`
	Residance   string `json:"residance"`
}

func (h handler) UpdateInfo(ctx *gin.Context) {
	body := new(UpdateRequest)
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
	user, err := h.GetUserSingleUserFunc(uuid)
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
		if result.Error.Error() == "ERROR: duplicate key value violates unique constraint \"uni_users_app_user_name\" (SQLSTATE 23505)" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": fmt.Sprintf("l'appUserName avec %v est deja utiliser par un autre utilisateur ", body.AppUserName)})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": result.Error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": &user})

}
