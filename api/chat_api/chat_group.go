package chat_api

import (
	"blog_server/global"
	"blog_server/models/res"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var ConnGroupMap = map[string]*websocket.Conn{}

func (ChatApi) ChatGroupView(c *gin.Context) {
	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	// upgrade http to websocket
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	addr := conn.RemoteAddr().String()
	ConnGroupMap[addr] = conn
	global.Logger.Infof("%s connect successful", addr)
	for {
		// receive message
		_, p, err := conn.ReadMessage()
		if err != nil {
			break
		}

		// send message
		SendGroupMsg(string(p))
	}

	// clear setting
	defer conn.Close()
	delete(ConnGroupMap, addr)
}

func SendGroupMsg(text string) {
	for _, conn := range ConnGroupMap {
		conn.WriteMessage(websocket.TextMessage, []byte(text))
	}
}
