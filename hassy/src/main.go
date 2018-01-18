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
	"golang.org/x/net/context"
	"time"
)

func init() {
	router := SetupRouter()
	http.Handle("/", router)
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.HTMLRender = createMyRender()
	router.GET("/", HandleTop)
	router.POST("/result", HandleResult)
	router.GET("/testJson", HandleTestJson)
	router.GET("/test", func(gc *gin.Context) {
		c := appengine.NewContext(gc.Request)
		returnString, _ := appengine.ModuleHostname(c, "", "", "")
		gc.String(http.StatusOK, fmt.Sprint(returnString))
	})

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
	//r.AddFromFiles("list", "templates/list.html")
	//r.AddFromFiles("main", "templates/base.html", "templates/main.html", "templates/footer.html")
	//r.AddFromFiles("test", "templates/base.html", "templates/test.html", "templates/footer.html")
	r.AddFromFiles("Top", "templates/base.html", "templates/top.html", "templates/inputPart.html")
	r.AddFromFiles("Result", "templates/base.html", "templates/Result.html", "templates/inputPart.html")

	return r
}

func HandleTop(gc *gin.Context) {
	//c := appengine.NewContext(gc.Request)
	//
	////url := "http://localhost:8080/testJson"
	////	url := "https://team25-demo.appspot.com/testJson"
	//returnString, _ := appengine.ModuleHostname(c, "", "", "")
	//url := "http://" + returnString + "/testJson"
	//
	//parseClient := urlfetch.Client(c)
	//parseRes, ParseErr := parseClient.Get(url)
	//if ParseErr != nil {
	//	gc.String(http.StatusOK, fmt.Sprint("Error1: ", ParseErr.Error()))
	//	log.Errorf(c, "Error1: %v", ParseErr.Error(), url)
	//	return
	//}
	//
	//defer parseRes.Body.Close()
	//result := make([]byte, parseRes.ContentLength)
	//parseRes.Body.Read(result)
	//
	//var vList []VideoList
	//err := json.Unmarshal([]byte(result), &vList)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	gc.HTML(http.StatusOK, "Top", gin.H{}) //テンプレートに変数を渡す場合
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

func GetVListFromDummy(aec context.Context, url string) ([]VideoList, error) {
	parseClient := urlfetch.Client(aec)
	parseRes, ParseErr := parseClient.Get(url)

	if ParseErr != nil {
		//gc.String(http.StatusOK, fmt.Sprint("Error1: ", ParseErr.Error()))
		errMessage := fmt.Sprint("Error1: %v", ParseErr.Error(), url)
		log.Errorf(aec, errMessage)
		return nil, fmt.Errorf("%s", errMessage)
	}

	defer parseRes.Body.Close()
	result := make([]byte, parseRes.ContentLength)
	parseRes.Body.Read(result)

	var vList []VideoList
	err := json.Unmarshal([]byte(result), &vList)
	if err != nil {
		//fmt.Println(err)
		errMessage := fmt.Sprint(err)
		log.Errorf(aec, errMessage)
		return nil, fmt.Errorf("%s", errMessage)
	}
	return vList, nil
}

func GetVListFromYoutube(aec context.Context, url string) ([]VideoList, error) {
	//log.Infof(aec, url)

	parseClient := urlfetch.Client(aec)
	parseRes, ParseErr := parseClient.Get(url)

	if ParseErr != nil {
		//gc.String(http.StatusOK, fmt.Sprint("Error1: ", ParseErr.Error()))
		errMessage := fmt.Sprint("Error1: %v", ParseErr.Error(), url)
		log.Errorf(aec, errMessage)
		return nil, fmt.Errorf("%s", errMessage)
	}

//	log.Infof(aec, fmt.Sprint(parseRes))

	defer parseRes.Body.Close()
	result := make([]byte, parseRes.ContentLength)
	parseRes.Body.Read(result)

	var youtubeResult YouTubeStruct
	err := json.Unmarshal([]byte(result), &youtubeResult)
	if err != nil {
		//fmt.Println(err)
		errMessage := fmt.Sprint(err)
		log.Errorf(aec, errMessage)
		return nil, fmt.Errorf("%s", errMessage)
	}

//	log.Infof(aec, fmt.Sprint(youtubeResult.Items[0].Snippet.Description))

	vList := make([]VideoList, 0)
	allResult := youtubeResult.Items
	for i, eachResult := range allResult {
		var tempVList VideoList
		log.Infof(aec, fmt.Sprint(i, " / ", eachResult.Snippet.Title))
		tempVList.VideoId = eachResult.ID.VideoID
		tempVList.ThumbnailUrl = eachResult.Snippet.Thumbnails.High.URL
		tempVList.Title = eachResult.Snippet.Title
		tempVList.Description = eachResult.Snippet.Description
		vList = append(vList, tempVList)
	}
	return vList, nil
}

func HandleResult(gc *gin.Context) {
	aec := appengine.NewContext(gc.Request)

	qString := gc.PostForm("qString")
//	log.Infof(aec, qString)

	//currentHostName, _ := appengine.ModuleHostname(aec, "", "", "")
	//url := "http://" + currentHostName + "/testJson"
	//vList, perr := GetVListFromDummy(aec, url)

	url := "https://www.googleapis.com/youtube/v3/search?part=snippet&maxResults=6&&" +
		"key=AIzaSyD4HAyfiPbu4QxhMEKgyOO3iAc-Snb1kZw&q=" + qString
	vList, perr := GetVListFromYoutube(aec, url)

	if perr != nil {
		gc.String(http.StatusOK, fmt.Sprint("Error1: ", perr.Error()))
	} else {
		gc.HTML(http.StatusOK, "Result", vList) //テンプレートに変数を渡す場合
	}

	//parseClient := urlfetch.Client(aec)
	//parseRes, ParseErr := parseClient.Get(url)
	//if ParseErr != nil {
	//	gc.String(http.StatusOK, fmt.Sprint("Error1: ", ParseErr.Error()))
	//	log.Errorf(aec, "Error1: %v", ParseErr.Error(), url)
	//	return
	//}
	//defer parseRes.Body.Close()
	//result := make([]byte, parseRes.ContentLength)
	//parseRes.Body.Read(result)
	//
	//var vList []VideoList
	//err := json.Unmarshal([]byte(result), &vList)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//gc.HTML(http.StatusOK, "Result", vList) //テンプレートに変数を渡す場合
}

//func HandleTest(gc *gin.Context) {
//	//	gc.String(http.StatusOK, fmt.Sprint("Hello World from Gin"))
//
//	gc.HTML(http.StatusOK, "index", gin.H{
//		"title": "myName",
//	})
//}

func HandleTestJson(gc *gin.Context) {
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
"description" : "dummy Description dummy Description dummy Description dummy Description dummy Description dummy Description dummy Description dummy Description dummy Description dummy Description dummy Description dummy Description dummy Description dummy Description dummy Description "
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
"title" : "日本語タイトル",
"description" : "日本語の説明、説明。日本語の説明、説明。日本語の説明、説明。日本語の説明、説明。日本語の説明、説明。日本語の説明、説明。日本語の説明、説明。日本語の説明、説明。日本語の説明、説明。日本語の説明、説明。日本語の説明、説明。"
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


type YouTubeStruct struct {
	Kind          string `json:"kind"`
	Etag          string `json:"etag"`
	NextPageToken string `json:"nextPageToken"`
	RegionCode    string `json:"regionCode"`
	PageInfo      struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
	Items []struct {
		Kind string `json:"kind"`
		Etag string `json:"etag"`
		ID   struct {
			Kind    string `json:"kind"`
			VideoID string `json:"videoId"`
		} `json:"id"`
		Snippet struct {
			PublishedAt time.Time `json:"publishedAt"`
			ChannelID   string    `json:"channelId"`
			Title       string    `json:"title"`
			Description string    `json:"description"`
			Thumbnails  struct {
				Default struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"default"`
				Medium struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"medium"`
				High struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"high"`
			} `json:"thumbnails"`
			ChannelTitle         string `json:"channelTitle"`
			LiveBroadcastContent string `json:"liveBroadcastContent"`
		} `json:"snippet"`
	} `json:"items"`
}