package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/akmalhazim/motosikal/repository/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type WsMsg struct {
	Op   int    `json:"op"`
	Data string `json:"data"`
}

type CreateSurveyRequest struct {
	RespondentName  string `json:"respondentName"`
	RespondentEmail string `json:"respondentEmail"`
	RespondentPhone string `json:"respondentPhone"`
}

type CreateRecordRequest struct {
	Lat       float32    `json:"lat"`
	Lng       float32    `json:"lng"`
	Timestamp *time.Time `json:"timestamp"`
}

var (
	CreateRecordOp = 1

	ErrSocketClosed = errors.New("socket closed")
)

func main() {
	upgrader := websocket.Upgrader{}

	e := echo.New()
	e.HideBanner = true

	handleWsErr := func(conn *websocket.Conn, err error) {
		fmt.Fprintln(os.Stdout, "[websocket] error", err.Error())
	}

	e.GET("/ws", func(c echo.Context) error {
		conn, err := upgrader.Upgrade(c.Response().Writer, c.Request(), nil)
		if err != nil {
			return err
		}
		defer conn.Close()

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				handleWsErr(conn, ErrSocketClosed)
				break
			}

			wsMsg := new(WsMsg)
			if err := json.Unmarshal(msg, wsMsg); err != nil {
				handleWsErr(conn, err)
				break
			}

		L:
			switch wsMsg.Op {
			case CreateRecordOp:
				req := new(CreateRecordRequest)
				if err := json.Unmarshal([]byte(wsMsg.Data), req); err != nil {
					handleWsErr(conn, err)
					break L
				}

				fmt.Fprintln(os.Stdout, "[websocket] lat:", req.Lat)
				fmt.Fprintln(os.Stdout, "[websocket] lng:", req.Lng)

				break L
			default:
				handleWsErr(conn, err)
			}

		}

		return nil
	})

	// TODO make this platform agnostic
	db, err := sql.Open("mysql", "root@/motosikal")
	if err != nil {
		panic(err)
	}

	// TODO add proper interface to adhere with SOLID principles
	deviceRepo := mysql.NewDeviceRepo(db)

	e.GET("/devices", func(c echo.Context) error {
		ctx := c.Request().Context()

		devices, err := deviceRepo.List(ctx)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, devices)
	})

	e.POST("/surveys", func(c echo.Context) error {
		req := new(CreateSurveyRequest)
		if err := c.Bind(req); err != nil {
			return err
		}

		return c.JSON(http.StatusOK, req)
	})

	e.POST("/devices/:deviceID/records", func(c echo.Context) error {
		req := new(CreateRecordRequest)
		if err := c.Bind(req); err != nil {
			return err
		}

		return c.JSON(http.StatusOK, req)
	})

	e.Logger.Error(e.Start("0.0.0.0:8000"))
}
