package main

/*
   @Auth: menah3m
   @Desc:
*/
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.StaticFS("/showfiles", http.Dir("./"))
	r.Run()
}
