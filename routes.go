package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type User struct {
	Id        int    `db:"id json:"id"`
	Firstname string `db:"firstname" json:"firstname"`
	Lastname  string `db:"lastname" json:"lastname"`
}

func setUp() {
	router := gin.Default()

	router.GET("/", index)

	v1 := router.Group("api/v1")
	{
		v1.GET("/users", GetUsers)
		v1.GET("/users/:id", GetUser)
		// v1.POST("/users", PostUser)
		// v1.PUT("/users/:id", UpdateUser)
		// v1.DELETE("/users/:id", DeleteUser)
	}
	router.LoadHTMLGlob("templates/*")
	// router.LoadHTMLFiles("views/index.html", "views/show.html")
	http.ListenAndServe(":2121", router)
	fmt.Println("Set Up complete:\n Listening on 2121")
}

func GetUsers(c *gin.Context) {
	type Users []User

	var users = Users{
		User{Id: 1, Firstname: "David", Lastname: "Stinnette"},
		User{Id: 2, Firstname: "Cuahuctemoc", Lastname: "Osorio"},
	}

	//var userTemplate = `{{ range $key, $value := . }} <li><strong>{{ $key }}</strong>: {{ $value }}</li> {{ end }}`
	//c.HTML(http.StatusOK, "users.tmpl", gin.H{"users": users})
	c.HTML(http.StatusOK, "show.tmpl", gin.H{})
	//c.JSON(200, users)
}

func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	user_id, _ := strconv.ParseInt(id, 0, 64)

	if user_id == 1 {
		content := gin.H{"id": user_id, "firstname": "David", "lastname": "Stinnette"}
		c.HTML(http.StatusOK, "show_single.tmpl", content)
		//c.JSON(200, content)
	} else if user_id == 2 {
		content := gin.H{"id": user_id, "firstname": "Cuahuctemoc", "lastname": "Osorio"}
		c.HTML(http.StatusOK, "show_single.tmpl", content)
		//c.JSON(200, content)
	} else {
		content := gin.H{"error": "user with id#" + id + "not found"}
		c.HTML(http.StatusOK, "show_single.tmpl", content)
		//c.JSON(200, content)
	}

}

func index(c *gin.Context) {
	//fmt.Println("We got here")
	c.HTML(http.StatusOK, "index.tmpl", gin.H{"title": "This is the Home Page"})
}
