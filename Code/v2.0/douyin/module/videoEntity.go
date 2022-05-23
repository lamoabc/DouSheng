package module

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	VideoTitle    string `json:"title"`
}
type VideoTable struct {
	VideoId    int64  `gorm:"column:video_id"`
	AuthorId   int64  `gorm:"column:author_id"`
	PlayUrl    string `gorm:"column:play_url"`
	CoverUrl   string `gorm:"column:cover_url"`
	VideoTitle string `gorm:"column:video_title"`
	FavCount   int64  `gorm:"column:favourite_count"`
	ComCount   int64  `gorm:"column:comment_count"`
	UploadDate string `gorm:"column:upload_date"`
}
