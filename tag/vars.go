package tag

import (
	"github.com/ckahi/wechat/context"

	"github.com/ckahi/wechat/util"
)

const (
	baseApi           = "https://api.weixin.qq.com/cgi-bin"
	tagsCreateURL     = "/tags/create"
	tagsGetURL        = "/tags/get"
	tagsUpdateURL     = "/tags/update"
	tagsDeleteURL     = "/tags/delete"
	batchTaggingURL   = "/tags/members/batchtagging"
	batchUnTaggingURL = "/tags/members/batchuntagging"
	userTagsGetURL    = "/tags/getidlist"
)

type reqCreateTag struct {
	Tag MpTag `json:"tag"`
}

type RespCreateTag struct {
	util.CommonError
	Tag MpTag `json:"tag"`
}

type RespTagsList struct {
	util.CommonError
	Tags []MpTag `json:"tags"`
}

type MpTag struct {
	Id    int64  `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Count int64  `json:"count,omitempty"`
}

type Tag struct {
	*context.Context
}

type reqBatchTagging struct {
	OpenidList []string `json:"openid_list"`
	TagId      int64    `json:"tagid"`
}

type RespUserTags struct {
	util.CommonError
	TagIds []int `json:"tagid_list"`
}
