package main

import (
	"fmt"
	. "sz1/AppInit"
	"sz1/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("v1")
	{
		v1.Handle(HTTP_METHOD_GET, "/prods", func(ctx *gin.Context) {
			prods := []*models.Books{}
			GetDB().Limit(10).Order("book_id desc").Find(&prods)
			for _, v := range prods {
				fmt.Println(*v)
			}
			fmt.Println("!!!")
			ctx.JSON(200, prods)
		})
	}
	router.Run(SERVER_ADDRESS)
}
