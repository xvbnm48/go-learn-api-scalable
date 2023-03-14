package helper

import "github.com/gin-gonic/gin"

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func HandleErr(ctx *gin.Context, err error) {
	if err != nil {
		//PanicIfError(err)
		ctx.JSON(404, gin.H{
			"message": err.Error(),
		})
	}
}
