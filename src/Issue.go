package issue

import (
	"net/http"
	"encoding/json"
	"fmt"

	"github.com/issue-list/dbProvider"
	"github.com/labstack/echo/v4"
)

func InitRouting(e *echo.Echo) {
	 e.GET("/issue", IssueRead)
	e.GET("/issue/all", IssueAllRead)
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