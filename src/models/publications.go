package models

// Publication struct for publications
type Publication struct {
	ID         uint64 `json:"id,omitempty"`
	Title      string `json:"title,omitempty"`
	Content    string `json:"content,omitempty"`
	AuthorID   uint64 `json:"author_id,omitempty"`
	AuthorNick uint64 `json:"author_nick,omitempty"`
	Likes      uint64 `json:"likes"`
	CreteAt    string `json:"create_at,omitempty"`
}
