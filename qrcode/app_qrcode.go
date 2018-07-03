package qrcode

import (
	"encoding/json"
	"fmt"

	"github.com/ckahi/wechat/util"
)

const (
	unLimitAPI = "http://api.weixin.qq.com/wxa/getwxacodeunlimit"
	limitAPI   = "http://api.weixin.qq.com/wxa/getwxacode"
)

type ReqCreateUnLimit struct {
	Scene string `json:"scene"`
	Page  string `json:"page"`
	Width int    `json:"width"`
}

func (qrcode *QrCode) CreateUnLimitQrCode(scene, page string, width int) (response []byte, err error) {
	var accessToken string
	accessToken, err = qrcode.GetAccessToken()
	if err != nil {
		return
	}
	reqParams := ReqCreateUnLimit{
		scene, page, width,
	}
	uri := fmt.Sprintf("%s?access_token=%s", unLimitAPI, accessToken)
	response, err = util.PostJSON(uri, reqParams)
	if err != nil {
		return
	}
	var resp util.CommonError
	err = json.Unmarshal(response, &resp)
	if err != nil {
		err = nil
		return
	}
	if resp.ErrCode != 0 {
		err = fmt.Errorf("GetAppQrCode Error , errcode=%d , errmsg=%s", resp.ErrCode, resp.ErrMsg)
		return
	}
	return
}
