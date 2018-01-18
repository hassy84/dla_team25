package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/renders/multitemplate"
	"fmt"
	"encoding/json"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"google.golang.org/appengine/log"
	"golang.org/x/net/context"
	"time"
	"strings"
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
}


func createMyRender() multitemplate.Render {
	r := multitemplate.New()
	r.AddFromFiles("Top", "templates/base.html", "templates/top.html", "templates/inputPart.html")
	r.AddFromFiles("Result", "templates/base.html", "templates/result.html", "templates/inputPart.html")

	return r
}

func HandleTop(gc *gin.Context) {
	gc.HTML(http.StatusOK, "Top", gin.H{})
}



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

	var youtubeResult YouTubeStruct
	err := json.Unmarshal([]byte(result), &youtubeResult)
	if err != nil {
		//fmt.Println(err)
		errMessage := fmt.Sprint(err)
		log.Errorf(aec, errMessage)
		return nil, fmt.Errorf("%s", errMessage)
	}

	vList := make([]VideoList, 0)
	allResult := youtubeResult.Items
	for i, eachResult := range allResult {
		if eachResult.ID.VideoID == "" {
			continue
		}
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

//ダミーのJSONを使う場合はここを切り替える
//	currentHostName, _ := appengine.ModuleHostname(aec, "", "", "")
//	url := "http://" + currentHostName + "/testJson"
//	vList, perr := GetVListFromDummy(aec, url)


	tempQString := gc.PostForm("qString")
	tempQString = strings.Replace(tempQString, "　", " ", -1) //全角スペースは半角に置き換え
	splitStrings := strings.Split(tempQString, " " ) //半角スペースで分割
	var qString string
	if len(splitStrings) == 1 {
		qString = splitStrings[0]
	}else{
		for i, eachString := range splitStrings{
			if i != 0 {
				qString = qString + ","
			}
			qString = qString + eachString
		}
	}
	log.Infof(aec, qString)
	url := "https://www.googleapis.com/youtube/v3/search?part=snippet&maxResults=6&&" +
		"key=AIzaSyD4HAyfiPbu4QxhMEKgyOO3iAc-Snb1kZw&q=" + qString
	vList, perr := GetVListFromYoutube(aec, url)


	if perr != nil {
		gc.String(http.StatusOK, fmt.Sprint("Error1: ", perr.Error()))
	} else {
		gc.HTML(http.StatusOK, "Result", vList) //テンプレートに変数を渡す場合
	}
}


func HandleTestJson(gc *gin.Context) {

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