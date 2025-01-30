package middlewares

import (
	"OneTix/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Validator[T any](source string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data T
		isSuccess := true

		switch source {
		case "body":
			handleRequestBody(c, &data, &isSuccess)
		case "query":
			handleRequestQuery(c, &data, &isSuccess)
		case "params":
			handleRequestParams(c, &data, &isSuccess)
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request input source"})
			c.Abort()
			return
		}

		if err := validate.Struct(data); err != nil && isSuccess {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "message": generateValidationErrors(err)})
			c.Abort()
			return
		}

		c.Set(fmt.Sprintf("validated_%s", source), data)
		c.Next()
	}
}

func generateValidationErrors(err error) []string {
	var validationErrors []string
	for _, err := range err.(validator.ValidationErrors) {
		message := fmt.Sprintf("Field '%s' is %s", utils.ToSnakeCase(err.Field()), err.ActualTag())

		if err.Param() != "" {
			message += fmt.Sprintf("='%s'", err.Param())
		}

		validationErrors = append(validationErrors, message)
	}

	return validationErrors
}

func handleRequestBody(c *gin.Context, data any, isSuccess *bool) {
	if c.Request.ContentLength == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request body is empty"})
		c.Abort()

		*isSuccess = false
		return
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Request body validation error": err.Error()})
		c.Abort()

		*isSuccess = false
	}

}

func handleRequestQuery(c *gin.Context, data any, isSuccess *bool) {
	if err := c.ShouldBindQuery(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Request query parameters validation error": err.Error()})
		c.Abort()

		*isSuccess = false
	}

}

func handleRequestParams(c *gin.Context, data any, isSuccess *bool) {
	if err := c.ShouldBindUri(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request parameters error", "details": err.Error()})
		c.Abort()

		*isSuccess = false
	}
}
