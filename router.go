package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func handlePing(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func handleGetUserByIstagramID(c *gin.Context) {
	id_str := c.Params.ByName("id")
	id, _ := strconv.Atoi(id_str)
	user := mongoGetUserByIstagramID(id)
	if user.Name == "" {
		c.String(http.StatusNotFound, "user not found")
		return
	}
	c.JSON(http.StatusOK, user)
}

func handleGetUser(c *gin.Context) {
	name := c.Params.ByName("name")
	user := FindUser(name)
	if user.Name == "" {
		c.String(http.StatusNotFound, "user not found")
		return
	}
	c.JSON(http.StatusOK, user)
}

func handleGetIGPostTs(c *gin.Context) {
	name := c.Param("name")
	r := mongoGetIGPostTs(name)
	ts := strconv.Itoa(r)
	c.String(http.StatusOK, ts)
}

func handleGetIGStoriesTs(c *gin.Context) {
	name := c.Param("name")
	r := mongoGetIGStoriesTs(name)
	ts := strconv.Itoa(r)
	c.String(http.StatusOK, ts)
}

func handleGetIGStoriesTsByID(c *gin.Context) {
	id_str := c.Param("id")
	id, _ := strconv.Atoi(id_str)
	r := mongoGetIGStoriesTsByID(id)
	ts := strconv.Itoa(r)
	c.String(http.StatusOK, ts)
}

func handleUpdateIGPost(c *gin.Context) {
	name := c.Param("name")
	ts, _ := strconv.Atoi(c.Param("timestamp"))
	//result := fmt.Sprintf("name: %v; timestamp: %v\n", name, ts)
	r := mongoUpdateIGPost(name, ts)

	c.JSON(http.StatusOK, r)
}

func handleUpdateIGStories(c *gin.Context) {
	name := c.Param("name")
	ts, _ := strconv.Atoi(c.Param("timestamp"))
	//result := fmt.Sprintf("name: %v; timestamp: %v\n", name, ts)
	r := mongoUpdateIGStories(name, ts)

	c.JSON(http.StatusOK, r)
}

func handleUpdateIGStoriesByID(c *gin.Context) {
	id_str := c.Param("id")
	id, _ := strconv.Atoi(id_str)
	ts, _ := strconv.Atoi(c.Param("timestamp"))
	//result := fmt.Sprintf("name: %v; timestamp: %v\n", name, ts)
	r := mongoUpdateIGStoriesByID(id, ts)

	c.JSON(http.StatusOK, r)
}

func handleGetAllUsers(c *gin.Context) {
	users := Find()
	//log.Printf("%v\n", user.Name)
	c.JSON(http.StatusOK, users)
}

func setupRouter() *gin.Engine {
	// gin.DisableConsoleColor()
	r := gin.Default()

	authorized := r.Group("/yarb", gin.BasicAuth(gin.Accounts{
		yarbBasicAuthUser: yarbBasicAuthPass,
	}))

	authorized.GET("/ping", handlePing)
	authorized.GET("/users", handleGetAllUsers)
	authorized.GET("/user/name/:name", handleGetUser)
	authorized.GET("/user/ig_id/:id", handleGetUserByIstagramID)
	authorized.GET("/user/name/:name/date/instagram_post", handleGetIGPostTs)
	authorized.GET("/user/name/:name/date/instagram_stories", handleGetIGStoriesTs)
	authorized.GET("/user/id/:id/date/instagram_stories", handleGetIGStoriesTsByID)
	authorized.GET("/user/name/:name/date/instagram_post/:timestamp", handleUpdateIGPost)
	authorized.GET("/user/name/:name/date/instagram_stories/:timestamp", handleUpdateIGStories)
	authorized.GET("/user/id/:id/date/instagram_stories/:timestamp", handleUpdateIGStoriesByID)
	authorized.GET("/user/name/:name/ig_post/:timestamp", handleUpdateIGStories)

	return r
}
