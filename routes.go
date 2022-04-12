package main

import (
	"CopyFelix/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {

	r.POST("/bing", controller.Register)
	return r

}
