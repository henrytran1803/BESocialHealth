package personalcontentmodels

import "BESocialHealth/comon"

type Post struct {
	comon.SQLModel
	Title  string `json:"title" gorm:"column:title"`
	Body   string `json:"body" gorm:"column:body"`
	UserId int64  `json:"user_id" gorm:"column:user_id"`
}
type Like struct {
	comon.SQLModel
	UserId int64 `json:"user_id" gorm:"column:user_id"`
	PostId int64 `json:"post_id" gorm:"column:post_id"`
}
type Comment struct {
	comon.SQLModel
	Body   string `json:"body" gorm:"column:body"`
	UserId int64  `json:"user_id" gorm:"column:user_id"`
	PostId int64  `json:"post_id" gorm:"column:post_id"`
}
type Photo struct {
	comon.SQLModel
	Photo_type string `json:"photo_type" gorm:"column:photo_type"`
	Image      []byte `json:"image" gorm:"column:image"`
	Url        string `json:"url" gorm:"column:url"`
	Post_id    string `json:"post_id" gorm:"column:post_id"`
}

func (Like) TableName() string    { return "likes" }
func (Comment) TableName() string { return "comments" }
func (Post) TableName() string    { return "posts" }
func (Photo) TableName() string   { return "photos" }

type CreatePostFull struct {
	Title       string        `json:"title" gorm:"column:title" form:"title" `
	Body        string        `json:"body" gorm:"column:body" form:"body" `
	UserId      int64         `json:"user_id" gorm:"column:user_id" form:"user_id" `
	CreatePhoto []CreatePhoto `json:"photos" form:"photos" `
}
type CreatePost struct {
	ID     int    `json:"id" gorm:"column:id"`
	Title  string `json:"title" gorm:"column:title"`
	Body   string `json:"body" gorm:"column:body"`
	UserId int64  `json:"user_id" gorm:"column:user_id"`
}
type CreatePhoto struct {
	Photo_type string  `json:"photo_type" gorm:"column:photo_type"`
	Image      []byte  `json:"image" gorm:"column:image"`
	Url        string  `json:"url" gorm:"column:url"`
	Post_id    *string `json:"post_id" gorm:"column:post_id"`
	Comment_id *string `json:"comment_id" gorm:"column:comment_id"`
}
type CreateCommentFull struct {
	Body        string      `json:"body" gorm:"column:body"`
	UserId      int64       `json:"user_id" gorm:"column:user_id"`
	PostId      int64       `json:"post_id" gorm:"column:post_id"`
	CreatePhoto CreatePhoto `json:"photos"`
}
type CreateComment struct {
	ID     int    `json:"id" gorm:"column:id"`
	Body   string `json:"body" gorm:"column:body"`
	UserId int64  `json:"user_id" gorm:"column:user_id"`
	PostId int64  `json:"post_id" gorm:"column:post_id"`
}

type CreateLike struct {
	UserId int64 `json:"user_id" gorm:"column:user_id"`
	PostId int64 `json:"post_id" gorm:"column:post_id"`
}
type GetPost struct {
	ID             int     `json:"id" gorm:"column:id"`
	Title          string  `json:"title" gorm:"column:title"`
	Body           string  `json:"body" gorm:"column:body"`
	UserId         int64   `json:"user_id" gorm:"column:user_id"`
	Photos         []Photo `json:"photos"`
	Count_likes    int     `json:"count_likes" `
	Count_comments int     `json:"count_comments" `
}
type GetComment struct {
	ID     int    `json:"id" gorm:"column:id"`
	Body   string `json:"body" gorm:"column:body"`
	UserId int64  `json:"user_id" gorm:"column:user_id"`
	Name   string `json:"name" `
	Photo  Photo  `json:"photos"`
}
