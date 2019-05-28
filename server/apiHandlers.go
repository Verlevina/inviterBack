package server

import "github.com/gin-gonic/gin"
import "api_example/entities"
import "encoding/json

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func getUser(c *gin.Context) {
	id := c.Query("id")

	user := new(entities.User);
	user.Id = id;
	user.Name = "Вася";

	c.JSON(200, gin.H{
		"user": user,
	})
}

func postUser(c *gin.Context) {
	user := struct {
		Id string `json:"id"`
		Name string `json:"name"`
	} {}

	c.BindJSON(&user)
	user.Name =  "Petia"

	c.JSON(200, gin.H{
		"user": user,
	})
}

func addTemplate(list *[]entities.Template, id int32, title string, link string) {
	s := *list;
	template1 := new(entities.Template);
	template1.Id = id;
	template1.Title = title;
	template1.Link = link;
	s = append(s, *template1);
}

func getEventTemplate(c *gin.Context) {

	var templates = []entities.Template {};
	addTemplate(&templates, 1, "wedding", "link");

	error, jsonResult := json.Marshal(templates);

	if error != nil {
		c.JSON(200, c.);
	} else {
		c.JSON(500, nil);
	}
}