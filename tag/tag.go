package tag

import (
	"encoding/json"
	"fmt"

	"github.com/ckahi/wechat/context"
	"github.com/ckahi/wechat/util"
)

func NewMpTag(context *context.Context) *Tag {
	tag := new(Tag)
	tag.Context = context
	return tag
}

//获取公众号的标签列表
func (tag *Tag) GetTags() (tags RespTagsList, err error) {
	accessToken, err := tag.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s%s?access_token=%s", baseApi, tagsGetURL, accessToken)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &tags)
	if err != nil {
		return
	}
	if tags.ErrCode != 0 {
		err = fmt.Errorf("GetTags Error , errcode=%d , errmsg=%s", tags.ErrCode, tags.ErrMsg)
		return
	}
	return
}

//添加标签
func (tag *Tag) AddTag(tagName string) (tagId int64, err error) {
	accessToken, err := tag.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s%s?access_token=%s", baseApi, tagsCreateURL, accessToken)
	mpTag := MpTag{
		Name: tagName,
	}
	reqCreateTag := reqCreateTag{
		Tag: mpTag,
	}

	var response []byte
	response, err = util.PostJSON(uri, reqCreateTag)
	if err != nil {
		return
	}
	var respCreateTag RespCreateTag
	err = json.Unmarshal(response, &respCreateTag)
	if err != nil {
		return
	}
	if respCreateTag.ErrCode != 0 {
		err = fmt.Errorf("AddTag Error , errcode=%d , errmsg=%s", respCreateTag.ErrCode, respCreateTag.ErrMsg)
		return
	}
	tagId = respCreateTag.Tag.Id
	return
}

//批量给用户打标签
func (tag *Tag) BatchTagging(openids []string, tagid int64) (err error) {
	accessToken, err := tag.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s%s?access_token=%s", baseApi, batchTaggingURL, accessToken)
	var (
		response []byte
		resp     util.CommonError
	)

	batchTaggingParams := reqBatchTagging{
		OpenidList: openids,
		TagId:      tagid,
	}
	response, err = util.PostJSON(uri, batchTaggingParams)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &resp)
	if err != nil {
		return
	}
	if resp.ErrCode != 0 {
		err = fmt.Errorf("BatchTagging Error , errcode=%d , errmsg=%s", resp.ErrCode, resp.ErrMsg)
		return
	}
	return
}

//批量为用户取消
func (tag *Tag) BatchUnTagging(openids []string, tagid int64) (err error) {
	accessToken, err := tag.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s%s?access_token=%s", baseApi, batchUnTaggingURL, accessToken)
	var (
		response []byte
		resp     util.CommonError
	)

	batchUnTaggingParams := reqBatchTagging{
		OpenidList: openids,
		TagId:      tagid,
	}
	response, err = util.PostJSON(uri, batchUnTaggingParams)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &resp)
	if err != nil {
		return
	}
	if resp.ErrCode != 0 {
		err = fmt.Errorf("BatchUnTagging Error , errcode=%d , errmsg=%s", resp.ErrCode, resp.ErrMsg)
		return
	}
	return
}

//获取用户身上标签列表
func (tag *Tag) GetUserTags(openid string) (userTags RespUserTags, err error) {
	accessToken, err := tag.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s%s?access_token=%s", baseApi, userTagsGetURL, accessToken)

	params := map[string]string{
		"openid": openid,
	}
	var response []byte
	response, err = util.PostJSON(uri, params)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &userTags)
	if err != nil {
		return
	}
	if userTags.ErrCode != 0 {
		err = fmt.Errorf("BatchUnTagging Error , errcode=%d , errmsg=%s", userTags.ErrCode, userTags.ErrMsg)
		return
	}
	return
}
