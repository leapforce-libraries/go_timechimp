package timechimp

type Tag struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
	Type   int64  `json:"type"`
}
