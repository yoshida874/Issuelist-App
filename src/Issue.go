package issue

import (
	"encoding/json"
	"fmt"
	"net/http"

	dbprovider "github.com/issue-list/dbProvider"
	"github.com/labstack/echo/v4"
)

func InitRouting(e *echo.Echo) {
	e.GET("/issue", IssueRead)
	e.GET("/issue/all", IssueAllRead)
	e.PUT("issue/update", IssueUpdate)
}

//1件読み込み
func IssueRead(c echo.Context) error {
	res := dbprovider.Read()
	jsonStr, err := json.Marshal(res)
	if err != nil {
		fmt.Println("JSON marshal error: ", err)
        return c.String(http.StatusBadRequest, "エラー")
	}
	return c.JSON(http.StatusOK, string(jsonStr))
}

// 全ての読み込み
func IssueAllRead(c echo.Context) error {
	res := dbprovider.AllRead()
	jsonStr, err := json.Marshal(res)
	if err != nil {
		fmt.Println("JSON marshal error: ", err)
        return c.String(http.StatusBadRequest, "エラー")
	}
	return c.JSON(http.StatusOK, string(jsonStr))
}

type PutIssue struct {
    Id string `json:"id"`
	Body string `json:"body"`
}
// idを指定してupdate
// applcation/json
func IssueUpdate(c echo.Context) error {
	var p PutIssue
	if err := c.Bind(&p);err != nil {
		fmt.Println(err)
	}
	dbprovider.Update(p.Id, p.Body);

	return c.String(http.StatusOK, "echo")
}