package message

//小程序 link消息
type Link struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	ThumbURL    string `json:"thumb_url,omitempty"`
	URL         string `json:"url,omitempty"`
}

//NewText 初始化文本消息
func NewLink(title, description, thumbURL, url string) *Link {
	link := new(Link)
	link.Title = title
	link.Description = description
	link.ThumbURL = thumbURL
	link.URL = url
	return link
}
