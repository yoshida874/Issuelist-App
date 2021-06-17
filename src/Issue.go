package issue

import (
	"net/http"
	"encoding/json"
	"fmt"

	"github.com/issue-list/dbProvider"
	"github.com/labstack/echo/v4"
)

func IssueRead(c echo.Context) error {
	res := dbProvider.Read()
	jsonStr, err := json.Marshal(res)
	if err != nil {
		fmt.Println("JSON marshal error: ", err)
        return c.String(http.StatusBadRequest, "エラー")
	}
	return c.JSON(http.StatusOK, string(jsonStr))
}