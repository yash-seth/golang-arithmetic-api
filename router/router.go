package router

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// initialize router object and handlers for various endpoints
func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", helloWorld)
	r.GET("/home", helloHome)
	r.GET("/add", add)
	r.GET("/sub", sub)
	r.GET("/mul", mul)
	r.GET("/div", div)
	return r
}

// default endpoint handler
func helloWorld(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Hello World",
	})
}

// home endpoint handler
func helloHome(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Hello Home",
	})
}

// mathematical operation endpoint handlers
func add(ctx *gin.Context) {
	queryParams := ctx.Request.URL.Query()
	var sum int64 = 0
	for _, values := range queryParams {
		v, _ := strconv.ParseInt(values[0], 10, 64)
		sum += v
	}
	ctx.JSON(http.StatusOK, gin.H{
		"sum": sum,
	})
}

func sub(ctx *gin.Context) {
	queryParams := ctx.Request.URL.Query()
	var sub int64 = 0
	for _, values := range queryParams {
		v, _ := strconv.ParseInt(values[0], 10, 64)
		if sub == 0 {
			sub = v
		} else {
			sub -= v
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"sub": sub,
	})
}

func mul(ctx *gin.Context) {
	queryParams := ctx.Request.URL.Query()
	var mul int64 = 1
	for _, values := range queryParams {
		v, _ := strconv.ParseInt(values[0], 10, 64)
		mul *= v
	}
	ctx.JSON(http.StatusOK, gin.H{
		"mul": mul,
	})
}

func div(ctx *gin.Context) {
	queryParams := ctx.Request.URL.Query()
	var div float64 = 0
	for _, values := range queryParams {
		v, _ := strconv.ParseFloat(values[0], 64)
		if div == 0 {
			div = v
		} else {
			div /= v
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"div": div,
	})
}
