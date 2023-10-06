package chat_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/ctype"
	"blog_server/models/res"
	"blog_server/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/DanPlayer/randomname"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ChatUser struct {
	Conn     *websocket.Conn
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
}

var ConnGroupMap = map[string]ChatUser{}

const (
	InRoomMsg  ctype.MsgType = 1
	TextMsg    ctype.MsgType = 2
	ImageMsg   ctype.MsgType = 3
	VoiceMsg   ctype.MsgType = 4
	VideoMsg   ctype.MsgType = 5
	SystemMsg  ctype.MsgType = 6
	OutRoomMsg ctype.MsgType = 7
)

type GroupRequest struct {
	Content string        `json:"content"`
	MsgType ctype.MsgType `json:"msg_type"`
}
type GroupResponse struct {
	NickName    string        `json:"nick_name"`
	Avatar      string        `json:"avatar"`
	MsgType     ctype.MsgType `json:"msg_type"`
	Content     string        `json:"content"`
	OnlineCount int           `json:"online_count"`
	Date        time.Time     `json:"date"`
}

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

	// save user info and save user info to conn map
	addr := conn.RemoteAddr().String()                                 // get ip addr
	nickName := randomname.GenerateName()                              // get a random nickname
	nickNameFirst := string([]rune(nickName)[0])                       // get fisrt chinese character
	avatar := fmt.Sprintf("uploads/chat_avatar/%s.png", nickNameFirst) // get the avatar

	chatUser := ChatUser{
		Conn:     conn,
		NickName: nickName,
		Avatar:   avatar,
	}
	ConnGroupMap[addr] = chatUser

	global.Logger.Infof("%s connect successful", addr)
	for {
		// receive message
		_, p, err := conn.ReadMessage()
		if err != nil {
			// user leave chat room
			SendGroupMsg(conn, GroupResponse{
				NickName:    chatUser.NickName,
				Avatar:      chatUser.Avatar,
				MsgType:     OutRoomMsg,
				Content:     fmt.Sprintf("%s left chat room", chatUser.NickName),
				Date:        time.Now(),
				OnlineCount: len(ConnGroupMap) - 1,
			})
			break
		}

		// bind request info
		var request GroupRequest
		err = json.Unmarshal(p, &request)
		if err != nil {
			SendMsg(addr, GroupResponse{
				NickName:    chatUser.NickName,
				Avatar:      chatUser.Avatar,
				MsgType:     SystemMsg,
				Content:     "Failed to bind info",
				OnlineCount: len(ConnGroupMap),
			})
			continue
		}

		// check msg type and handle
		switch request.MsgType {
		case TextMsg:
			if strings.TrimSpace(request.Content) == "" {
				SendMsg(addr, GroupResponse{
					NickName:    chatUser.NickName,
					Avatar:      chatUser.Avatar,
					MsgType:     SystemMsg,
					Content:     "Message can't be empty",
					OnlineCount: len(ConnGroupMap),
				})
				continue
			}
			SendGroupMsg(conn, GroupResponse{
				NickName:    chatUser.NickName,
				Avatar:      chatUser.Avatar,
				Content:     request.Content,
				MsgType:     TextMsg,
				Date:        time.Now(),
				OnlineCount: len(ConnGroupMap),
			})
		case InRoomMsg:
			SendGroupMsg(conn, GroupResponse{
				NickName:    chatUser.NickName,
				Avatar:      chatUser.Avatar,
				Content:     fmt.Sprintf("%s entered chat room", chatUser.NickName),
				Date:        time.Now(),
				OnlineCount: len(ConnGroupMap),
				MsgType:     InRoomMsg,
			})
		default:
			SendMsg(addr, GroupResponse{
				NickName:    chatUser.NickName,
				Avatar:      chatUser.Avatar,
				MsgType:     SystemMsg,
				Content:     "Message type incorrect",
				OnlineCount: len(ConnGroupMap),
			})
		}

	}

	// clear setting
	defer conn.Close()

	// remove chatuser from conn map
	delete(ConnGroupMap, addr)
}

// group message
func SendGroupMsg(conn *websocket.Conn, response GroupResponse) {
	byteData, _ := json.Marshal(response)
	_addr := conn.RemoteAddr().String()
	ip, addr := getIPAndAddr(_addr)

	global.DB.Create(&models.ChatModel{
		NickName: response.NickName,
		Avatar:   response.Avatar,
		Content:  response.Content,
		IP:       ip,
		Addr:     addr,
		IsGroup:  true,
		MsgType:  response.MsgType,
	})
	for _, chatUser := range ConnGroupMap {
		chatUser.Conn.WriteMessage(websocket.TextMessage, byteData)
	}
}

// system message
func SendMsg(_addr string, response GroupResponse) {
	byteData, _ := json.Marshal(response)
	chatUser := ConnGroupMap[_addr]
	ip, addr := getIPAndAddr(_addr)
	global.DB.Create(&models.ChatModel{
		NickName: response.NickName,
		Avatar:   response.Avatar,
		Content:  response.Content,
		IP:       ip,
		Addr:     addr,
		IsGroup:  false,
		MsgType:  response.MsgType,
	})
	chatUser.Conn.WriteMessage(websocket.TextMessage, byteData)
}

func getIPAndAddr(_addr string) (ip string, addr string) {
	addrList := strings.Split(_addr, ":")
	ip = addrList[0]
	addr = utils.GetAddr(ip)
	return ip, addr // addrList[0] is ip, addrList[1] is port
}
