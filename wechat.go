package wechat

import (
	"net/http"
	"sync"

	"github.com/ckahi/wechat/cache"
	"github.com/ckahi/wechat/context"
	"github.com/ckahi/wechat/group"
	"github.com/ckahi/wechat/js"
	"github.com/ckahi/wechat/material"
	"github.com/ckahi/wechat/menu"
	"github.com/ckahi/wechat/message"
	"github.com/ckahi/wechat/oauth"
	"github.com/ckahi/wechat/qrcode"
	"github.com/ckahi/wechat/server"
	"github.com/ckahi/wechat/template"
	"github.com/ckahi/wechat/user"
)

// Wechat struct
type Wechat struct {
	Context *context.Context
}

// Config for user
type Config struct {
	AppID          string
	AppSecret      string
	Token          string
	EncodingAESKey string
	Cache          cache.Cache
	Strategy       context.Strategy
}

// NewWechat init
func NewWechat(cfg *Config) *Wechat {
	context := new(context.Context)
	copyConfigToContext(cfg, context)
	return &Wechat{context}
}

func copyConfigToContext(cfg *Config, context *context.Context) {
	context.AppID = cfg.AppID
	context.AppSecret = cfg.AppSecret
	context.Token = cfg.Token
	context.EncodingAESKey = cfg.EncodingAESKey
	context.Cache = cfg.Cache
	context.Strategy = cfg.Strategy
	context.SetAccessTokenLock(new(sync.RWMutex))
	context.SetJsAPITicketLock(new(sync.RWMutex))
}

// GetServer 消息管理
func (wc *Wechat) GetServer(req *http.Request, writer http.ResponseWriter) *server.Server {
	wc.Context.Request = req
	wc.Context.Writer = writer
	return server.NewServer(wc.Context)
}

//GetAccessToken 获取access_token
func (wc *Wechat) GetAccessToken() (string, error) {
	return wc.Context.GetAccessToken()
}

// GetOauth oauth2网页授权
func (wc *Wechat) GetOauth() *oauth.Oauth {
	return oauth.NewOauth(wc.Context)
}

// GetMaterial 素材管理
func (wc *Wechat) GetMaterial() *material.Material {
	return material.NewMaterial(wc.Context)
}

// GetJs js-sdk配置
func (wc *Wechat) GetJs() *js.Js {
	return js.NewJs(wc.Context)
}

// GetMenu 菜单管理接口
func (wc *Wechat) GetMenu() *menu.Menu {
	return menu.NewMenu(wc.Context)
}

// GetMenu 分组管理接口
func (wc *Wechat) GetGroup() *group.Group {
	return group.NewGroup(wc.Context)
}

// GetUser 用户管理接口
func (wc *Wechat) GetUser() *user.User {
	return user.NewUser(wc.Context)
}

// GetTemplate 模板消息接口
func (wc *Wechat) GetTemplate() *template.Template {
	return template.NewTemplate(wc.Context)
}

// GetQrCode 二维码管理
func (wc *Wechat) GetQrCode() *qrcode.QrCode {
	return qrcode.NewQrCode(wc.Context)
}

// GetQrCode 二维码管理
func (wc *Wechat) GetCustom() *message.Custom {
	return message.NewCustom(wc.Context)
}
