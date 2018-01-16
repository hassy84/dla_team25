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
	"fmt"
	"encoding/json"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"google.golang.org/appengine/log"
)

func init() {
	router := SetupRouter()
	http.Handle("/", router)
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.HTMLRender = createMyRender()
	//router.GET("/main", func(c *gin.Context) {
	//	c.HTML(200, "main", gin.H{})
	//})
	router.GET("/test", func(gc *gin.Context) {
		c := appengine.NewContext(gc.Request)
		returnString, _ := appengine.ModuleHostname(c, "", "", "")
		gc.String(http.StatusOK, fmt.Sprint(returnString))
	})
	router.GET("/", HandleMain)
	router.GET("/testJson", HandleJson)

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
	r.AddFromFiles("list", "templates/list.html")
	r.AddFromFiles("main", "templates/base.html", "templates/main.html", "templates/footer.html")
	r.AddFromFiles("test", "templates/base.html", "templates/test.html", "templates/footer.html")
	r.AddFromFiles("Main", "templates/base.html", "templates/top.html")
	return r
}

//func HandleToppage(gc *gin.Context) {
//	//gc.String(http.StatusOK, fmt.Sprint("Toppage from Gin"))
//
////	url := "http://localhost:8080/testJson"
////	url := "https://team25-demo.appspot.com/testJson"
//	url := "*/testJson"
//
//	c := appengine.NewContext(gc.Request)
//	parseClient := urlfetch.Client(c)
//	parseRes, ParseErr := parseClient.Get(url)
//	if ParseErr != nil {
//		gc.String(http.StatusOK, fmt.Sprint("Parseにしっぱいしました"))
//		return
//	}
//
//	defer parseRes.Body.Close()
//	result := make([]byte, parseRes.ContentLength)
//	parseRes.Body.Read(result)
//
//	var vList []VideoList
//	err := json.Unmarshal([]byte(result), &vList)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	//gc.JSON(http.StatusOK, vList) //単にJSONとして表示する場合
//	gc.HTML(http.StatusOK, "list", vList) //テンプレートに変数を渡す場合
//}

func HandleMain(gc *gin.Context) {
	c := appengine.NewContext(gc.Request)

	//url := "http://localhost:8080/testJson"
	//	url := "https://team25-demo.appspot.com/testJson"
	returnString, _ := appengine.ModuleHostname(c, "", "", "")
	url := "http://" + returnString + "/testJson"

	parseClient := urlfetch.Client(c)
	parseRes, ParseErr := parseClient.Get(url)
	if ParseErr != nil {
		gc.String(http.StatusOK, fmt.Sprint("Error1: ", ParseErr.Error()))
		log.Errorf(c, "Error1: %v", ParseErr.Error(), url)
		return
	}

	defer parseRes.Body.Close()
	result := make([]byte, parseRes.ContentLength)
	parseRes.Body.Read(result)

	var vList []VideoList
	err := json.Unmarshal([]byte(result), &vList)
	if err != nil {
		fmt.Println(err)
		return
	}

	gc.HTML(http.StatusOK, "Main", vList) //テンプレートに変数を渡す場合
}

func HandleTest(gc *gin.Context) {
	//	gc.String(http.StatusOK, fmt.Sprint("Hello World from Gin"))

	gc.HTML(http.StatusOK, "index", gin.H{
		"title": "myName",
	})
}

func HandleJson(gc *gin.Context) {
	//var jsonString VideoList // VideoList型、ストラクト、変数
	//jsonString.VideoId = "g2ag8t7AvX8"
	//jsonString.ThumbnailUrl = "https://i.ytimg.com/vi/g2ag8t7AvX8/hqdefault.jpg"
	//jsonString.Title = "title1"
	//jsonString.Description = "dummyDescription1"
	//gc.JSON(http.StatusOK, jsonString)

	var original_str = `
[
{
"videoId" : "g2ag8t7AvX8",
"thumbnailUrl" : "https://i.ytimg.com/vi/g2ag8t7AvX8/hqdefault.jpg",
"title" : "title1title1title1title1",
"description" : "dummyDescription1dummyDescription1dummyDescription1dummyDescription1dummyDescription1dummyDescription1dummyDescription1dummyDescription1dummyDescription1dummyDescription1dummyDescription1dummyDescription1dummyDescription1dummyDescription1dummyDescription1dummyDescription1dummyDescription1dummyDescription1dummyDescription1dummyDescription1dummyDescription1"
}
,
{
"videoId" : "blfDtisFTyM",
"thumbnailUrl" : "https://i.ytimg.com/vi/blfDtisFTyM/hqdefault.jpg",
"title" : "title2InShort",
"description" : "dummyDescription2InShortCase"
}
,
{
"videoId" : "blfDtisFTyM",
"thumbnailUrl" : "https://i.ytimg.com/vi/blfDtisFTyM/hqdefault.jpg",
"title" : "title2InShort",
"description" : "dummyDescription2InShortCase"
}
,
{
"videoId" : "blfDtisFTyM",
"thumbnailUrl" : "https://i.ytimg.com/vi/blfDtisFTyM/hqdefault.jpg",
"title" : "title2InShort",
"description" : "dummyDescription2InShortCase"
}
,
{
"videoId" : "blfDtisFTyM",
"thumbnailUrl" : "https://i.ytimg.com/vi/blfDtisFTyM/hqdefault.jpg",
"title" : "title2InShort",
"description" : "dummyDescription2InShortCase"
}
,
{
"videoId" : "blfDtisFTyM",
"thumbnailUrl" : "https://i.ytimg.com/vi/blfDtisFTyM/hqdefault.jpg",
"title" : "title2InShort",
"description" : "dummyDescription2InShortCase"
}
]
`
	var vList []VideoList
	err := json.Unmarshal([]byte(original_str), &vList)
	if err != nil {
		fmt.Println(err)
		return
	}

	gc.JSON(http.StatusOK, vList)
}

type VideoList struct {
	VideoId      string ` json:"videoId" binding:"required"`
	ThumbnailUrl string ` json:"thumbnailUrl" binding:"required"`
	Title        string ` json:"title" binding:"required"`
	Description  string ` json:"description" `
}
