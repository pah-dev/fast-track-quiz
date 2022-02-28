package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pah-dev/fast-track-quiz/api/routes"
	"github.com/pah-dev/fast-track-quiz/api/utils"
)

func main() {
    router := gin.Default()
    
    routes.Quiz(router)
    
    router.Run(":" + utils.GodotEnv("GO_PORT"))
}
