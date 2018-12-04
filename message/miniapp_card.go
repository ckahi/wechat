package message

//Image 图片消息
type MiniAppCard struct {
	CommonToken  `json:"-"`
	Title        string `json:"title"`
	Appid        string `json:"appid"`
	Pagepath     string `json:"pagepath"`
	ThumbMediaId string `json:"thumb_media_id"`
}

//NewImage 回复图片消息
func NewMiniAppCard(title, appid, pagepath, thumbMediaId string) *MiniAppCard {
	miniAppCard := new(MiniAppCard)
	miniAppCard.Title = title
	miniAppCard.Appid = appid
	miniAppCard.Pagepath = pagepath
	miniAppCard.ThumbMediaId = thumbMediaId
	return miniAppCard
}
