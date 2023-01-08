package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var WxSDK WeChat

func init() {
	WxSDK.AppId = "wxeafdff978daf544b"
	WxSDK.AppSecret = "e9b38699eb2a42f9d2df46bf7029c4be"
}

type WeChat struct {
	AppId       string `json:"app_id"`
	AppSecret   string `json:"app_secret"`
	AccessToken string `json:"access_token"`
}

type SnsOauth2 struct {
	Openid  string `json:"openid"`
	Unionid string `json:"unionid"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type AccessTokenErrorResponse struct {
	//{"errcode":40029,"errmsg":"invalid code, rid: 63969311-1fb0f5f8-0997410e"}
	ErrMsg  string `json:"errmsg"`
	ErrCode int    `json:"errcode"`
}

// GetWxOpenIdFromOauth2 通过code换取网页授权access_token
func (weChat *WeChat) GetWxOpenIdFromOauth2(code string) (*SnsOauth2, error) {

	// 微信登录文档
	// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html
	requestLine := strings.Join([]string{
		"https://api.weixin.qq.com/sns/jscode2session",
		"?appid=", weChat.AppId,
		"&secret=", weChat.AppSecret,
		"&js_code=", code,
		"&grant_type=authorization_code"}, "")

	resp, err := http.Get(requestLine)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("发送get请求获取 openid 读取返回body错误", err)
		return nil, err
	}
	if bytes.Contains(body, []byte("errmsg")) {
		ater := AccessTokenErrorResponse{}
		err = json.Unmarshal(body, &ater)
		if err != nil {
			fmt.Printf("发送get请求获取 openid 的错误信息 %+v\n", ater)
			return nil, err
		}
		return nil, fmt.Errorf("%s", ater.ErrMsg)
	} else {
		atr := SnsOauth2{}
		err = json.Unmarshal(body, &atr)
		if err != nil {
			fmt.Println("发送get请求获取 openid 返回数据json解析错误", err)
			return nil, err
		}
		return &atr, nil
	}
}
