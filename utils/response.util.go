package utils

import "github.com/gin-gonic/gin"

func GenerateResponse(status bool, message string, optionalData ...any) gin.H {
	var data any
	if len(optionalData) > 0 {
		data = optionalData[0]
	}

	return gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	}
}
