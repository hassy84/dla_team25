package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	//"fmt"
	//"strconv"
	//"io/ioutil"
	//"google.golang.org/appengine"
	//"google.golang.org/appengine/urlfetch"
	//"google.golang.org/appengine/log"
	//"encoding/json"
	//"math"
	//"github.com/mjibson/goon"
	//"bytes"
	//"google.golang.org/appengine/datastore"
	//"errors"
	//"html/template"
	"github.com/gin-gonic/contrib/renders/multitemplate"
	//"github.com/mjibson/goon"
	//"fmt"
)

func init() {
	router := SetupRouter()
	http.Handle("/", router)
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.HTMLRender = createMyRender()
	router.GET("/main", func(c *gin.Context) {
		c.HTML(200, "main", gin.H{})
	})
	router.GET("/test", func(c *gin.Context) {
		c.HTML(200, "test", gin.H{})
	})
	router.GET("/", HandleTest)
	return router

	//router.LoadHTMLGlob("templates/*")
	//router.GET("/test", func(c *gin.Context) {
	//	// テンプレート設定
	//	html := template.Must(template.ParseFiles("templates/base.html", "templates/test.html", "templates/footer.html"))
	//	router.SetHTMLTemplate(html)
	//	c.HTML(http.StatusOK, "base.html", gin.H{})
	//})
	//router.GET("/main", func(c *gin.Context) {
	//	// テンプレート設定
	//	html := template.Must(template.ParseFiles("templates/base.html", "templates/main.html", "templates/footer.html"))
	//	router.SetHTMLTemplate(html)
	//	c.HTML(http.StatusOK, "base.html", gin.H{})
	//})
	//router.GET("/", HandleTest)
	//return router
}

func createMyRender() multitemplate.Render {
	r := multitemplate.New()
	r.AddFromFiles("index", "templates/list.tmpl")
	r.AddFromFiles("main", "templates/base.html", "templates/main.html", "templates/footer.html")
	r.AddFromFiles("test", "templates/base.html", "templates/test.html", "templates/footer.html")
	return r
}





func HandleTest(gc *gin.Context) {
	//	gc.String(http.StatusOK, fmt.Sprint("Hello World from Gin"))

	gc.HTML(http.StatusOK, "index", gin.H{
		"title": "myName",
	})


}
