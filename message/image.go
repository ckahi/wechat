package message

//Image 图片消息
type Image struct {
	CommonToken `json:"-"`
	MediaID     string `xml:"MediaId" json:"media_id"`
}

//NewImage 回复图片消息
func NewImage(mediaID string) *Image {
	image := new(Image)
	image.MediaID = mediaID
	return image
}
