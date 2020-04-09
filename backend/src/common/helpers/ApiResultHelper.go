package helpers

import (
	"fmt"
	"reflect"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiReturn(context *gin.Context, output interface{}){
	var xType = reflect.TypeOf(output)
	var xValue = reflect.ValueOf(output)
	fmt.Println(xType, xValue)
	success := reflect.Indirect(xValue).FieldByName("Success")
	if success.Bool() {
		Ok(context, output)
	} else {
		BadRequest(context, output)
	}
}

func Ok(context *gin.Context, data interface{}){
	context.JSON(http.StatusOK, data)
}

func BadRequest(context *gin.Context, data interface{})  {
	context.JSON(http.StatusBadRequest, data)
}