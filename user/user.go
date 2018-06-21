package user

import (
	"encoding/json"
	"fmt"

	"github.com/ckahi/wechat/context"
	"github.com/ckahi/wechat/util"
)

const (
	userInfoURL      = "https://api.weixin.qq.com/cgi-bin/user/info"
	userListURL      = "https://api.weixin.qq.com/cgi-bin/user/get"
	batchUserInfoURL = "https://api.weixin.qq.com/cgi-bin/user/info/batchget"
)

//User 用户管理
type User struct {
	*context.Context
}

//NewUser 实例化
func NewUser(context *context.Context) *User {
	user := new(User)
	user.Context = context
	return user
}

//Info 用户基本信息
type Info struct {
	util.CommonError

	Subscribe     int32   `json:"subscribe"`
	OpenID        string  `json:"openid"`
	Nickname      string  `json:"nickname"`
	Sex           int32   `json:"sex"`
	City          string  `json:"city"`
	Country       string  `json:"country"`
	Province      string  `json:"province"`
	Language      string  `json:"language"`
	Headimgurl    string  `json:"headimgurl"`
	SubscribeTime int32   `json:"subscribe_time"`
	UnionID       string  `json:"unionid"`
	Remark        string  `json:"remark"`
	GroupID       int32   `json:"groupid"`
	TagidList     []int32 `json:"tagid_list"`
}

type InfoList struct {
	util.CommonError
	UserInfoList []Info `json:"user_info_list"`
}

type UserList struct {
	util.CommonError
	Total int64 `json:"total"`
	Count int64 `json:"count"`
	Data  struct {
		OpenID []string `json:"openid"`
	} `json:"data"`
	NextOpenID string `json:"next_openid"`
}

//GetUserInfo 获取用户基本信息
func (user *User) GetUserInfo(openID string) (userInfo *Info, err error) {
	var accessToken string
	accessToken, err = user.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s&openid=%s&lang=zh_CN", userInfoURL, accessToken, openID)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	userInfo = new(Info)
	err = json.Unmarshal(response, userInfo)
	if err != nil {
		return
	}
	if userInfo.ErrCode != 0 {
		err = fmt.Errorf("GetUserInfo Error , errcode=%d , errmsg=%s", userInfo.ErrCode, userInfo.ErrMsg)
		return
	}
	return
}

func (user *User) GetUserLists(nextOpenID string) (userList *UserList, err error) {
	var accessToken string
	accessToken, err = user.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s&next_openid=%s&lang=zh_CN", userListURL, accessToken, nextOpenID)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	userList = new(UserList)
	err = json.Unmarshal(response, userList)
	if err != nil {
		return
	}
	if userList.ErrCode != 0 {
		err = fmt.Errorf("GetUserInfo Error , errcode=%d , errmsg=%s", userList.ErrCode, userList.ErrMsg)
		return
	}
	return
}

func (user *User) GetUserInfoList(openids []string) (infoList *InfoList, err error) {
	var accessToken string
	accessToken, err = user.GetAccessToken()
	if err != nil {
		return
	}
	params := map[string][]map[string]string{
		"user_list": make([]map[string]string, 0),
	}
	for _, openid := range openids {
		params["user_list"] = append(params["user_list"], map[string]string{
			"openid": openid,
		})
	}

	uri := fmt.Sprintf("%s?access_token=%s", batchUserInfoURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, params)
	if err != nil {
		return
	}
	infoList = new(InfoList)
	err = json.Unmarshal(response, infoList)
	if err != nil {
		return
	}
	if infoList.ErrCode != 0 {
		err = fmt.Errorf("GetUserInfo Error , errcode=%d , errmsg=%s", infoList.ErrCode, infoList.ErrMsg)
		return
	}
	return
}
