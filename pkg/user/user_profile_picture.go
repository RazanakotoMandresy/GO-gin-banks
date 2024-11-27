package user

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/RazanakotoMandresy/go-gin-banks/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func (h handler) UserPP(ctx *gin.Context) {
	uuid, err := middleware.ExtractTokenUUID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": err.Error()})
		return
	}
	// ge anle uuid anle tokny hovaina
	userUUidPP, err := middleware.User.User(middleware.User{UuidToFind: uuid, Db: h.DB})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	file, err := ctx.FormFile("filePP")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	splitedName := strings.Split(file.Filename, ".")
 	// rename le nom pour qu'il soit unique
	fileName := filepath.Base(splitedName[0] + fmt.Sprint(time.Now().Nanosecond()) + "." + splitedName[1])
	// destinantion
	destFile := fmt.Sprintf("upload/%v", fileName)
	if err := ctx.SaveUploadedFile(file, destFile); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	os.Remove(userUUidPP.Image)
	userUUidPP.Image = destFile
	h.DB.Save(userUUidPP)
	ctx.JSON(http.StatusCreated, gin.H{"user": userUUidPP})
}
