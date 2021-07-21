package issue

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	dbprovider "github.com/issue-list/dbProvider"
	"github.com/labstack/echo/v4"
)

func InitRouting(e *echo.Echo) {
	e.GET("/issue", IssueRead)
	e.GET("/issue/all", IssueAllRead)
	e.PUT("issue/update", IssueUpdate)
	e.POST("issue/create", IssueCreate)
}

//1件読み込み
func IssueRead(c echo.Context) error {
	strID := c.QueryParam("id")
	i, _ := strconv.Atoi(strID)
	res := dbprovider.Read(i)
	jsonStr, err := json.Marshal(res)
	if err != nil {
		fmt.Println("JSON marshal error: ", err)
        return c.String(http.StatusBadRequest, "エラー")
	}
	return c.JSON(http.StatusOK, string(jsonStr))
}


// 全ての読み込み
func IssueAllRead(c echo.Context) error {
	oRes := dbprovider.OpenRead()
	cRes := dbprovider.ClosedRead()
	res := make(map[string]interface{})
	res["open"] = oRes
	res["closed"] = cRes
	jsonStr, err := json.Marshal(res)
	if err != nil {
		fmt.Println("JSON marshal error: ", err)
        return c.String(http.StatusBadRequest, "エラー")
	}
	return c.JSON(http.StatusOK, string(jsonStr))
}

type IssUpdate struct {
    Id string `json:"id"`
	Body string `json:"body"`
	IsClosed bool `json:"isClosed"`
}

// idを指定してupdate
// applcation/json
func IssueUpdate(c echo.Context) error {
	var i IssUpdate
	if err := c.Bind(&i);err != nil {
		fmt.Println(err)
		return c.String(http.StatusOK, err.Error())
	}

	if err := dbprovider.Update(i.Id, i.Body, i.IsClosed);err != nil {
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

	if err := dbprovider.Create(i.Title, i.Body);err != nil {
		fmt.Println(err)
		return c.String(http.StatusOK, err.Error())
	}

	return c.String(http.StatusOK, "echo")
}