package helpers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Status int
	Meta   interface{}
	Data   interface{}
}

func RespondJSON(ctx *gin.Context, status int, payload interface{}, err error) {
	fmt.Println("status ", status)

	if status != 200 {
		payload = gin.H{"error": err.Error()}
	} else {
		payload = gin.H{"data": payload}
	}

	ctx.JSON(status, payload)
}