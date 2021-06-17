package main

import (
	"io"
	"net/http"
	"os"


	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/issue-list/src"
)

func main() {
  // Echo instance
  e := echo.New()

  /*
  Logger: httpのrequestをログで出力
  Recover: 障害が出たときサーバーを落とさずエラーレスポンスを返す 
  */ 
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())
  
  e.POST("/save", Save)
  e.GET("/users/:id", GetUser)
  e.GET("/show", Show)
  e.GET("/issue", issue.IssueRead)
  // e.DELETE("/users/:id", deleteUser)

  // Start server
  e.Logger.Fatal(e.Start(":1323"))
}

// http://localhost:1323/users/Joe
func GetUser(c echo.Context) error {
  var id string = c.Param("id")
  return c.String(http.StatusOK, id)
}

// 複数クエリ
// http://localhost:1323/show?team=x-men&member=wolverine
func Show(c echo.Context) error {
  team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:" + team + ", member:" + member)
}

// // application/x-www-form-urlencoded
// func save(c echo.Context) error {
// 	name := c.FormValue("name")
// 	email := c.FormValue("email")
// 	return c.String(http.StatusOK, "name:" + name + ", email:" + email)
// }


// multipart/form-data
func Save(c echo.Context) error {
	// Get name
	name := c.FormValue("name")
	// Get avatar
  	avatar, err := c.FormFile("avatar")
  	if err != nil {
 		return err
 	}
 
 	// Source
 	src, err := avatar.Open()
 	if err != nil {
 		return err
 	}
 	defer src.Close()
 
 	// Destination
 	dst, err := os.Create(avatar.Filename)
 	if err != nil {
 		return err
 	}
 	defer dst.Close()
 
 	// Copy
 	if _, err = io.Copy(dst, src); err != nil {
  		return err
  	}

	return c.HTML(http.StatusOK, "<b>Thank you! " + name + "</b>")
}