package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID       string `json:"id" form:"id"`
	Nama     string `json:"nama" form:"nama"`
	HP       string `json:"hp" form:"hp"`
	Password string `json:"password" form:"password"`
}

func connectGorm() (*gorm.DB, error) {
	dsn := "root:@tcp(ec2-54-151-215-43:3306)/test_deploy_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

type RegisterFormat struct {
	Nama     string `json:"nama" form:"nama"`
	HP       string `json:"hp" form:"hp"`
	Password string `json:"password" form:"password"`
}

func Create(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RegisterFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		err := db.Save(input).Error
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("berhasil register", input))
	}
}

func GetAll(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var result []User
		err := db.Find(&result).Error
		if err != nil {
			fmt.Println("Error on Query", err.Error())
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("berhasil register", result))
	}
}

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func FailResponse(msg string) map[string]string {
	return map[string]string{
		"message": msg,
	}
}

func main() {
	//var errorStr string
	e := echo.New()
	db, err := connectGorm()
	if err != nil {
		log.Println(err.Error())
		//errorStr = "Error Database connection"
	}
	db.AutoMigrate(&User{})
	e.Use(middleware.Logger())

	e.GET("/users", GetAll(db))
	e.POST("/users", Create(db))

	// log.Print("\n\n\n")
	// log.Print("\n\n\n")
	// e.GET("/hello", func(c echo.Context) error {
	// 	return c.JSON(http.StatusOK, "hello world")
	// })
	// e.GET("/hello_gaes", func(c echo.Context) error {
	// 	return c.JSON(http.StatusOK, "hello gaes from wonderland")
	// })
	// e.GET("/hello_bang", func(c echo.Context) error {
	// 	return c.JSON(http.StatusOK, "hello bang makan bang")
	// })
	// e.GET("/hello_coy", func(c echo.Context) error {
	// 	return c.JSON(http.StatusOK, "hello coy makan kuy")
	// })
	e.Start(":8000")
}
