package message

import (
	"encoding/json"
	"fmt"

	"github.com/ckahi/wechat/context"
	"github.com/ckahi/wechat/util"
)

const (
	customSendURL = "https://api.weixin.qq.com/cgi-bin/message/custom/send"
)

type Custom struct {
	*context.Context
}

type ReqSendCustom struct {
	ToUser  string  `json:"touser"`
	MsgType MsgType `json:"msgtype"`
	Image   *Image  `json:"image,omitempty"`
	Text    *Text   `json:"text"`
	News    *News   `json:"news,omitempty"`
	Link    *Link   `json:"link,omitempty"`
}

type RespSend struct {
	util.CommonError
}

func NewCustom(context *context.Context) *Custom {
	custom := new(Custom)
	custom.Context = context
	return custom
}

func (custom *Custom) Send(reqParams ReqSendCustom) (resp RespSend, err error) {
	var accessToken string
	accessToken, err = custom.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", customSendURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, reqParams)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return
	}

	if resp.ErrCode != 0 {
		err = fmt.Errorf("SendCustomMsg Error , errcode=%d , errmsg=%s", resp.ErrCode, resp.ErrMsg)
		return
	}
	return
}
