package qq

import (
	"blog_server/global"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type QQInfo struct {
	Nickname string `json:"nickname"`
	Gender   string `json:"gender"`
	Avatar   string `json:"figureurl_qq"`
	OpenID   string `json:"open_id"`
}

type QQLogin struct {
	appID       string
	appKey      string
	redirect    string
	code        string
	accessToken string
	openID      string
}

func NewQQLogin(code string) (qqInfo QQInfo, err error) {
	qqLogin := &QQLogin{
		appID:    global.Config.QQ.AppID,
		appKey:   global.Config.QQ.Key,
		redirect: global.Config.QQ.Redirect,
		code:     code,
	}
	err = qqLogin.GetAccessToken()
	if err != nil {
		global.Logger.Error(err)
		return qqInfo, err
	}
	err = qqLogin.GetOpenID()
	if err != nil {
		global.Logger.Error(err)
		return qqInfo, err
	}
	qqInfo, err = qqLogin.GetUserInfo()
	if err != nil {
		global.Logger.Error(err)
		return qqInfo, err
	}
	qqInfo.OpenID = qqLogin.openID
	return qqInfo, nil
}

// 1. Get token with authorization code
func (q *QQLogin) GetAccessToken() error {
	params := url.Values{} // collect parameters
	params.Add("grant_type", "authorization_code")
	params.Add("client_id", q.appID)
	params.Add("client_secret", q.appKey)
	params.Add("code", q.code)
	params.Add("redirect_uri", q.redirect)
	urL, err := url.Parse("https://graph.qq.com/oauth2.0/token") // parse
	if err != nil {
		return err
	}

	urL.RawQuery = params.Encode() // encode and set collection urL.RawQuery

	res, err := http.Get(urL.String()) // get requesr
	if err != nil {
		return err
	}

	defer res.Body.Close()            // res.body implement io.closer and io.reader
	info, err := parseQuery(res.Body) // get a map
	if err != nil {
		return err
	}

	q.accessToken = info[`access_token`][0]

	return nil
}

// 2. Get openId with access token
func (q *QQLogin) GetOpenID() error {
	u, err := url.Parse(fmt.Sprintf("https://graph.qq.com/oauth2.0/me?access_token=%s", q.accessToken)) // parse
	if err != nil {
		return err
	}

	res, err := http.Get(u.String()) // get
	if err != nil {
		return err
	}
	defer res.Body.Close()

	openID, err := getOpenID(res.Body)
	if err != nil {
		return err
	}

	q.openID = openID
	return nil
}

// 3. Get user　info with openid
func (q *QQLogin) GetUserInfo() (qqInfo QQInfo, err error) {
	params := url.Values{}
	params.Add("access_token", q.accessToken)
	params.Add("oauth_consumer_key", q.appID)
	params.Add("openid", q.openID)
	urL, err := url.Parse("https://graph.qq.com/user/get_user_info")
	if err != nil {
		return qqInfo, err
	}

	urL.RawQuery = params.Encode() // encode

	res, _ := http.Get(urL.String())
	data, _ := io.ReadAll(res.Body)
	err = json.Unmarshal(data, &qqInfo)
	if err != nil {
		return qqInfo, err
	}

	return qqInfo, nil
}

func parseQuery(r io.Reader) (val map[string][]string, err error) {
	val, err = url.ParseQuery(readAll(r))
	if err != nil {
		return val, err
	}
	return val, nil
}

func getOpenID(r io.Reader) (string, error) {
	body := readAll(r)
	start := strings.Index(body, `"openid":"`) + len(`"openid":"`)
	if start == -1 {
		return "", fmt.Errorf("openid not found")
	}
	end := strings.Index(body[start:], `"`)
	if end == -1 {
		return "", fmt.Errorf("openid not found")
	}
	return body[start : start+end], nil
}

func readAll(r io.Reader) string {
	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
