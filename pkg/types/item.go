package types

type I struct {
	Id    int    `json:"id,omitempty"`
	Group int    `json:"group,omitempty"`
	Name  string `json:"name,omitempty"`
	Flag  int    `json:"flag,omitempty"`
}
