package issue

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func InitRouting(e *echo.Echo) {
	e.GET("/issue", IssueRead)
	e.GET("/issue/all", IssueAllRead)
	e.PUT("issue/update/title", titleUpdate);
	e.PUT("issue/update/comment", commentUpdate)
	e.POST("issue/update/closed", closedUpdate)
	e.POST("issue/create", IssueCreate)
}

//1件読み込み
func IssueRead(c echo.Context) error {
	strID := c.QueryParam("id")
	// 文字列→数値
	i, _ := strconv.Atoi(strID)
	res := Read(i)
	jsonStr, err := json.Marshal(res)
	if err != nil {
		fmt.Println("JSON marshal error: ", err)
        return c.String(http.StatusBadRequest, "エラー")
	}
	return c.JSON(http.StatusOK, string(jsonStr))
}


// 全ての読み込み
func IssueAllRead(c echo.Context) error {

	res := make(map[string]interface{})

	if c.QueryParam("isPath") == "true" {
		res["path"] = AllRead()
	}else{
		oRes := OpenRead()
		cRes := ClosedRead()
		res["open"] = oRes
		res["closed"] = cRes
	}

	jsonStr, err := json.Marshal(res)
	if err != nil {
		fmt.Println("JSON marshal error: ", err)
        return c.String(http.StatusBadRequest, "エラー")
	}
	return c.JSON(http.StatusOK, string(jsonStr))
}

type TitleUpdate struct {
	Id string `json:"id"`
	Title string `json:"title"`
}

func titleUpdate(c echo.Context) error {
	var i TitleUpdate
	if err := c.Bind(&i);err != nil {
		fmt.Println(err)
		return c.String(http.StatusOK, err.Error())
	}

	if err := UpdTitle(i.Id, i.Title);err != nil {
		fmt.Println(err)
		return c.String(http.StatusOK, err.Error())
	}

	return c.String(http.StatusOK, "echo")
}

type IssUpdate struct {
    Id string `json:"id"`
	Body string `json:"body"`
}

// idを指定してupdate
// applcation/json
func commentUpdate(c echo.Context) error {
	var i IssUpdate
	if err := c.Bind(&i);err != nil {
		fmt.Println(err)
		return c.String(http.StatusOK, err.Error())
	}

	if err := UpdComment(i.Id, i.Body);err != nil {
		fmt.Println(err)
		return c.String(http.StatusOK, err.Error())
	}

	return c.String(http.StatusOK, "echo")
}

type CloseUpdate struct {
	Id string `json:"id"`
	IsClosed bool `json:"isClosed"`
}

func closedUpdate(c echo.Context) error {
	var i CloseUpdate
	if err := c.Bind(&i);err != nil {
		fmt.Println(err)
		return c.String(http.StatusOK, err.Error())
	}
	if err := UpdClose(i.Id, i.IsClosed);err != nil {
		fmt.Println(err)
		return c.String(http.StatusOK, err.Error())
	}

	return c.String(http.StatusOK, "echo")
}

type IssCreate struct {
	Title string `json:"title"`
	Body string `json:"body"`
}

func IssueCreate(c echo.Context) error {
	var i IssCreate
	if err := c.Bind(&i);err != nil {
		fmt.Println(err)
		return c.String(http.StatusOK, err.Error())
	}

	if err := Create(i.Title, i.Body);err != nil {
		fmt.Println(err)
		return c.String(http.StatusOK, err.Error())
	}

	return c.String(http.StatusOK, "echo")
}