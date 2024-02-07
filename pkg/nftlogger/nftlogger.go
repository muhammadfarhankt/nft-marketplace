package nftlogger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/muhammadfarhankt/nft-marketplace/pkg/utils"
)

type INftLogger interface {
	Print() INftLogger
	Save()
	SetQuery(c *fiber.Ctx)
	SetBody(c *fiber.Ctx)
	SetRespose(res any)
}

type nftLogger struct {
	Time       string `json:"time"`
	Ip         string `json:"ip"`
	Method     string `json:"method"`
	StatusCode int    `json:"status_code"`
	Path       string `json:"path"`
	Query      any    `json:"query"`
	Body       any    `json:"body"`
	Response   any    `json:"response"`
}

func InitNftLogger(c *fiber.Ctx, res any) INftLogger {
	log := &nftLogger{
		Time:       time.Now().Local().Format("2006-01-02 15:04:05"),
		Ip:         c.IP(),
		Method:     c.Method(),
		Path:       c.Path(),
		StatusCode: c.Response().StatusCode(),
	}
	log.SetQuery(c)
	log.SetBody(c)
	log.SetRespose(res)
	return log
}

func (n *nftLogger) Print() INftLogger {
	utils.Debug(n)
	return n
}

func (n *nftLogger) Save() {
	data := utils.Output(n)
	filename := fmt.Sprintf("./assets/logs/nftlogger_%v.txt", strings.ReplaceAll(time.Now().Local().Format("2006-01-02"), "-", ""))
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()
	file.WriteString(string(data) + "\n")
}

func (n *nftLogger) SetQuery(c *fiber.Ctx) {
	var body any
	if err := c.QueryParser(&body); err != nil {
		log.Printf("query parser error: %v", err)
	}

}

func (n *nftLogger) SetBody(c *fiber.Ctx) {
	var body any
	if err := c.BodyParser(&body); err != nil {
		log.Printf("body parser error: %v", err)
	}
	switch n.Path {
	case "v1/users/signup":
		n.Body = "never gonna give you up"
	}
}

func (n *nftLogger) SetRespose(res any) {
	n.Response = res
}
