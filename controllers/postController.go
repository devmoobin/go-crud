package controllers

import (
	"devmoobin/go-crud/initializers"
	"devmoobin/go-crud/models"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {

	//Get data off req body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	//Create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// return it
	c.JSON(200, gin.H{
		"message": "Post data success",
		"data":    post,
	})
}

func PostsAll(c *gin.Context) {

	//Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//Respond with them

	c.JSON(200, gin.H{
		"message": "Get all data success",
		"data":    posts,
	})

}

func PostByIndex(c *gin.Context) {

	id := c.Param("id")

	//Get the posts
	var posts models.Post
	initializers.DB.First(&posts, id)

	//Respond with them

	c.JSON(200, gin.H{
		"message": "Get Id :" + id + "data success",
		"data":    posts,
	})

}

func PostUpdate(c *gin.Context) {

	//Get the id off the url
	id := c.Param("id")

	//Get the data off req body
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)
	//Find the post where update
	var post models.Post
	initializers.DB.First(&post, id)


	//update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title:body.Title,
		Body:body.Body,
	})


	//Create a post

	c.JSON(200, gin.H{
		"message": "Update Id :" + id + "data success",
		"data":    post,
	})
	
}

func PostDelete(c *gin.Context) {
    	//Get the id off the url
		id := c.Param("id")

		//Delete
		initializers.DB.Delete(&models.Post{}, id)

       c.Status(200)
}
