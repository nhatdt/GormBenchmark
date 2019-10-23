package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	echopprof "github.com/sevenNt/echo-pprof"
)

type (
	handler struct{}

	Message struct {
		Message string `json:"message"`
	}

	World struct {
		ID           uint16 `json:"id"`
		RandomNumber uint16 `gorm:"column:randomNumber" json:"randomNumber"`
	}
)

func (World) TableName() string {
	return "world"
}

const (
	worldRowCount  = 10000
	maxConnections = 1000
)

var (
	// Database
	DB *gorm.DB

	helloWorld = []byte("Hello, World!")
)

func (h *handler) query_rows() echo.HandlerFunc {
	return func(c echo.Context) error {
		n := getQueryCountRows(c.QueryParam("n"))
		worlds := make([]World, n)
		min := rand.Intn(worldRowCount - n)

		if err := DB.Model(&World{}).Limit(n).Offset(min).Find(&worlds).Error; err != nil {
			return err
		}

		c.Response().Header().Set("Server", "Echo")
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		return json.NewEncoder(c.Response()).Encode(worlds)
	}
}

func getQueryCountRows(q string) int {
	n, err := strconv.Atoi(q)
	if err != nil || n < 1 {
		return 1
	}
	if n > 1000 {
		return 1000
	}
	return n
}

func InitRoutes(e *echo.Echo) {
	h := new(handler)
	e.GET("/query", h.query_rows())
}

func main() {
	e := echo.New()
	dsn := "benchmarkdbuser:benchmarkdbpass@tcp(%s:3306)/hello_world"
	dbhost := "dbserver"

	var err error

	DB, err = gorm.Open("mysql", fmt.Sprintf(dsn, dbhost))

	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	DB.DB().SetMaxIdleConns(maxConnections)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	DB.DB().SetMaxOpenConns(maxConnections)
	defer DB.Close()
	InitRoutes(e)
	echopprof.Wrap(e)
	e.Start(":8080")
}
