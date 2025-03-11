package tests

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
)

func TestRoute() {
	router := gin.Default()
	fmt.Print(reflect.TypeOf(router))
}
