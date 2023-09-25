package log_stash

import (
	"blog_server/global"
	"blog_server/utils"
	"blog_server/utils/jwts"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Log struct {
	ip     string `json:"ip"`
	addr   string `json:"addr"`
	userId uint   `json:"user_id"`
}

func New(ip string, token string) *Log {
	claims, err := jwts.ParseToken(token)
	var userID uint
	if err == nil {
		userID = claims.UserID
	}

	addr := utils.GetAddr(ip)

	return &Log{
		ip:     ip,
		addr:   addr,
		userId: userID,
	}
}

func NewLogByGin(c *gin.Context) *Log {
	ip := c.ClientIP()
	token := c.Request.Header.Get("token")
	return New(ip, token)
}

func (l Log) Debug(content string) {
	l.send(DebugLevel, content)
}
func (l Log) Info(content string) {
	l.send(InfoLevel, content)
}
func (l Log) Warn(content string) {
	l.send(WarnLevel, content)
}
func (l Log) Error(content string) {
	l.send(ErrorLevel, content)
}

func (l Log) send(level Level, content string) {
	err := global.DB.Create(&LogStashModel{
		IP:      l.ip,
		Addr:    l.addr,
		Level:   level,
		Content: content,
		UserID:  l.userId,
	}).Error
	if err != nil {
		logrus.Error(err)
	}
}
