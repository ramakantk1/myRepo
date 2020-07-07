//Model or entiry which contain the struct model of the data

package model

// Video is a representation of a person
type Video struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}
