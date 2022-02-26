package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pah-dev/fast-track-quiz/api/routes"
)

func main() {
    router := gin.Default()
    
    routes.Quiz(router)
    
    router.Run("localhost:8089")
}
