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
	"html/template"
	//"github.com/mjibson/goon"
	//"fmt"
)

func init() {
	router := SetupRouter()
	http.Handle("/", router)
}

func notfoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	template.Must(template.ParseFiles("templates/notfound.html")).Execute(w, nil)
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/test", func(c *gin.Context) {
		// テンプレート設定
		html := template.Must(template.ParseFiles("templates/base.html", "templates/test.html"))
		router.SetHTMLTemplate(html)
		c.HTML(http.StatusOK, "base.html", gin.H{})
	})

	router.GET("/main", func(c *gin.Context) {
		// テンプレート設定
		html := template.Must(template.ParseFiles("templates/base.html", "templates/main.html"))
		router.SetHTMLTemplate(html)
		c.HTML(http.StatusOK, "base.html", gin.H{})
	})


	router.GET("/", HandleTest)
	return router
}

func HandleTest(gc *gin.Context) {
	//	gc.String(http.StatusOK, fmt.Sprint("Hello World from Gin"))

	gc.HTML(http.StatusOK, "list.tmpl", gin.H{
		"title": "myName",
	})

}
