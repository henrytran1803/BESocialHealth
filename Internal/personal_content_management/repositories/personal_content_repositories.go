package personalcontentrepositories

import (
	personalcontentmodels "BESocialHealth/Internal/personal_content_management/models"
	"strconv"
)

func (r *PersonalContentRepository) CreatePost(post *personalcontentmodels.CreatePost) (int, error) {
	if err := r.DB.Table(personalcontentmodels.Post{}.TableName()).Create(post).Error; err != nil {
		return 0, err
	}
	return post.ID, nil
}
func (r *PersonalContentRepository) CreatePhoto(photo *personalcontentmodels.CreatePhoto) error {
	return r.DB.Table(personalcontentmodels.Photo{}.TableName()).Create(photo).Error
}
func (r *PersonalContentRepository) UpdatePostById(id int, post *personalcontentmodels.CreatePost) error {
	return r.DB.Table(personalcontentmodels.Post{}.TableName()).Where("id = ?", id).Updates(post).Error
}

func (r *PersonalContentRepository) CreateFullPost(post *personalcontentmodels.CreatePostFull) error {
	postOnly := personalcontentmodels.CreatePost{
		Title:  post.Title,
		Body:   post.Body,
		UserId: post.UserId,
	}
	id, err := r.CreatePost(&postOnly)
	idstr := strconv.Itoa(id)
	if err != nil {
		return err
	}
	photos := post.CreatePhoto
	for _, photo := range photos {

		photo.Post_id = &idstr
		photo.Photo_type = "3"
		photo.Comment_id = nil
		if err := r.CreatePhoto(&photo); err != nil {
			return err
		}
	}
	return nil
}

func (r *PersonalContentRepository) CreateLike(post *personalcontentmodels.CreateLike) error {
	if err := r.DB.Table(personalcontentmodels.Like{}.TableName()).Create(post).Error; err != nil {
		return err
	}
	return nil
}
func (r *PersonalContentRepository) CreateComment(post *personalcontentmodels.CreateComment) (int, error) {
	if err := r.DB.Table(personalcontentmodels.Comment{}.TableName()).Create(post).Error; err != nil {
		return 0, err
	}
	return post.ID, nil
}
func (r *PersonalContentRepository) CreateComentFull(post *personalcontentmodels.CreateCommentFull) error {
	coment := personalcontentmodels.CreateComment{
		Body:   post.Body,
		UserId: post.UserId,
		PostId: post.PostId,
	}

	id, err := r.CreateComment(&coment)
	if err != nil {
		return err
	}
	photo := post.CreatePhoto
	photo.Photo_type = "2"
	idStr := strconv.Itoa(id)
	photo.Comment_id = &idStr
	if err := r.CreatePhoto(&photo); err != nil {
		return err
	}
	return nil
}
func (r *PersonalContentRepository) DeleteLikeByUserIDAndPostId(userId int, postId int) error {
	if err := r.DB.Table(personalcontentmodels.Like{}.TableName()).Where("post_id = ? and user_id = ?", postId, userId).Delete(&personalcontentmodels.Like{}).Error; err != nil {
		return err
	}
	return nil
}
func (r *PersonalContentRepository) DeleteCommentByUserIDAndPostId(userId string, postId int) error {
	if err := r.DB.Table(personalcontentmodels.Comment{}.TableName()).Where("post_id = ? and user_id = ?", postId, userId).Delete(&personalcontentmodels.Comment{}).Error; err != nil {
		return err
	}
	return nil
}
func (r *PersonalContentRepository) DeletePostById(postId int) error {
	if err := r.DeletePhotoByPostId(postId); err != nil {
		return err
	}
	if err := r.DB.Table(personalcontentmodels.Post{}.TableName()).Where("id = ?", postId).Delete(&personalcontentmodels.Post{}).Error; err != nil {
		return err
	}
	return nil
}
func (r *PersonalContentRepository) DeletePhotoById(photoId int) error {
	if err := r.DB.Table(personalcontentmodels.Photo{}.TableName()).Where("id = ?", photoId).Delete(&personalcontentmodels.Photo{}).Error; err != nil {
		return err
	}
	return nil
}
func (r *PersonalContentRepository) DeletePhotoByPostId(postId int) error {
	if err := r.DB.Table(personalcontentmodels.Photo{}.TableName()).Where("post_id = ?", postId).Delete(&personalcontentmodels.Photo{}).Error; err != nil {
		return err
	}
	return nil
}
func (r *PersonalContentRepository) DeletePhotoByCommentId(postId int) error {
	if err := r.DB.Table(personalcontentmodels.Photo{}.TableName()).Where("coment_id = ?", postId).Delete(&personalcontentmodels.Photo{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *PersonalContentRepository) GetAllCommentByPostId(postId int) ([]personalcontentmodels.GetComment, error) {
	var comments []personalcontentmodels.Comment
	var commentsFull []personalcontentmodels.GetComment
	if err := r.DB.Table(personalcontentmodels.Comment{}.TableName()).Where("post_id = ?", postId).Find(&comments).Error; err != nil {
		return nil, err
	}

	for _, comment := range comments {
		var user struct {
			Name string `gorm:"column:name"`
		}
		if err := r.DB.Table("users").Where("id = ?", comment.UserId).Select("name").First(&user).Error; err != nil {
			return nil, err
		}
		var photo personalcontentmodels.Photo
		if err := r.DB.Table(personalcontentmodels.Photo{}.TableName()).Where("post_id = ?", postId).First(&photo).Error; err != nil {
			return nil, err
		}
		commentsFull = append(commentsFull, personalcontentmodels.GetComment{
			ID:     postId,
			Body:   comment.Body,
			UserId: comment.UserId,
			Name:   user.Name,
			Photo:  photo,
		})
	}

	return commentsFull, nil
}
func (r *PersonalContentRepository) GetPhotosByPostId(postId int) ([]personalcontentmodels.Photo, error) {
	var photos []personalcontentmodels.Photo
	if err := r.DB.Table(personalcontentmodels.Photo{}.TableName()).Where("post_id = ?", postId).Find(&photos).Error; err != nil {
		return nil, err
	}
	return photos, nil
}

func (r *PersonalContentRepository) GetPostById(postId int) (personalcontentmodels.GetPost, error) {
	var post personalcontentmodels.Post
	if err := r.DB.Table(personalcontentmodels.Post{}.TableName()).Where("id = ?", postId).First(&post).Error; err != nil {
		return personalcontentmodels.GetPost{}, err
	}

	photos, err := r.GetPhotosByPostId(postId)
	if err != nil {
		return personalcontentmodels.GetPost{}, err
	}

	var countLikes int64
	if err := r.DB.Table(personalcontentmodels.Like{}.TableName()).Where("post_id = ?", postId).Count(&countLikes).Error; err != nil {
		return personalcontentmodels.GetPost{}, err
	}

	var countComments int64
	if err := r.DB.Table(personalcontentmodels.Comment{}.TableName()).Where("post_id = ?", postId).Count(&countComments).Error; err != nil {
		return personalcontentmodels.GetPost{}, err
	}

	getPost := personalcontentmodels.GetPost{
		ID:             postId,
		Title:          post.Title,
		Body:           post.Body,
		UserId:         post.UserId,
		Photos:         photos,
		Count_likes:    int(countLikes),
		Count_comments: int(countComments),
	}

	return getPost, nil
}
func (r *PersonalContentRepository) GetAllPosts() ([]personalcontentmodels.GetPost, error) {
	var posts []personalcontentmodels.Post
	if err := r.DB.Table(personalcontentmodels.Post{}.TableName()).Order("created_at asc").Find(&posts).Error; err != nil {
		return nil, err
	}

	var getPosts []personalcontentmodels.GetPost
	for _, post := range posts {
		photos, err := r.GetPhotosByPostId(int(post.Id))
		if err != nil {
			return nil, err
		}

		var countLikes int64
		if err := r.DB.Table(personalcontentmodels.Like{}.TableName()).Where("post_id = ?", post.Id).Count(&countLikes).Error; err != nil {
			return nil, err
		}

		var countComments int64
		if err := r.DB.Table(personalcontentmodels.Comment{}.TableName()).Where("post_id = ?", post.Id).Count(&countComments).Error; err != nil {
			return nil, err
		}

		getPosts = append(getPosts, personalcontentmodels.GetPost{
			ID:             post.Id,
			Title:          post.Title,
			Body:           post.Body,
			UserId:         post.UserId,
			Photos:         photos,
			Count_likes:    int(countLikes),
			Count_comments: int(countComments),
		})
	}

	return getPosts, nil
}
