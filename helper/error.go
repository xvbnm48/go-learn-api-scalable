package helper

import "github.com/gin-gonic/gin"

// helper function for handle error

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

// HandleErr handle error
func HandleErr(ctx *gin.Context, err error) {
	if err != nil {
		//PanicIfError(err)
		ctx.JSON(404, gin.H{
			"message": err.Error(),
		})
	}
}
