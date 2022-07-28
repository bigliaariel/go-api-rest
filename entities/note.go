package entities

type Note struct {
	Id    int32  `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
	Color int32  `json:"color,omitempty"`
}
