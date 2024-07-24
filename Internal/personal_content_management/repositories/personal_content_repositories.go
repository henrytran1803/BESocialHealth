package personalcontentrepositories

import (
	personalcontentmodels "BESocialHealth/Internal/personal_content_management/models"
	userrepositories "BESocialHealth/Internal/user_management/repositories"
	"errors"
	"gorm.io/gorm"
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
func (r *PersonalContentRepository) UpdatePostById(post *personalcontentmodels.CreatePost) error {
	return r.DB.Table(personalcontentmodels.Post{}.TableName()).Where("id = ?", post.ID).Updates(post).Error
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

func (r *PersonalContentRepository) DeleteCommentByPostId(postId int) error {
	if err := r.DB.Table(personalcontentmodels.Comment{}.TableName()).
		Where("post_id = ?", postId).
		Delete(&personalcontentmodels.Comment{}).Error; err != nil {
		return err
	}
	return nil
}
func (r *PersonalContentRepository) DeletePostById(postId int) error {
	if err := r.DeletePhotosByPostId(postId); err != nil {
		return err
	}
	if err := r.DB.Table(personalcontentmodels.Like{}.TableName()).Where("post_id = ?", postId).Delete(&personalcontentmodels.Like{}).Error; err != nil {
		return err
	}
	if err := r.DeleteCommentByPostId(postId); err != nil {
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
func (r *PersonalContentRepository) DeletePhotosByPostId(postId int) error {
	// Xóa các ảnh liên quan đến bình luận
	if err := r.DB.Table(personalcontentmodels.Photo{}.TableName()).
		Where("comment_id IN (SELECT id FROM comments WHERE post_id = ?)", postId).
		Delete(&personalcontentmodels.Photo{}).Error; err != nil {
		return err
	}

	// Xóa các ảnh không liên kết với bất kỳ bình luận nào
	if err := r.DB.Table(personalcontentmodels.Photo{}.TableName()).
		Where("post_id = ?", postId).
		Delete(&personalcontentmodels.Photo{}).Error; err != nil {
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

		var photo *personalcontentmodels.Photo
		if err := r.DB.Table(personalcontentmodels.Photo{}.TableName()).Where("comment_id = ?", comment.Id).First(&photo).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				photo = nil
			} else {
				return nil, err
			}
		}

		userRepo := userrepositories.NewUserRepository(r.DB)
		userInfo, err := userRepo.GetUserById(int(comment.UserId))
		if err != nil {
			return nil, err
		}
		commentsFull = append(commentsFull, personalcontentmodels.GetComment{
			ID:     postId,
			Body:   comment.Body,
			UserId: comment.UserId,
			User:   userInfo,
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
	userRepo := userrepositories.NewUserRepository(r.DB)
	user, err := userRepo.GetUserById(int(post.UserId))
	if err != nil {
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
		User:           user,
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
		userRepo := userrepositories.NewUserRepository(r.DB)
		user, err := userRepo.GetUserById(int(post.UserId))
		getPosts = append(getPosts, personalcontentmodels.GetPost{
			ID:             post.Id,
			Title:          post.Title,
			Body:           post.Body,
			UserId:         post.UserId,
			Photos:         photos,
			Count_likes:    int(countLikes),
			Count_comments: int(countComments),
			User:           user,
		})
	}
	return getPosts, nil
}
func (r *PersonalContentRepository) GetAllPostsByUserId(id string) ([]personalcontentmodels.GetPost, error) {
	var posts []personalcontentmodels.Post
	if err := r.DB.Table(personalcontentmodels.Post{}.TableName()).Order("created_at asc").Where("user_id = ?", id).Find(&posts).Error; err != nil {
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
		userRepo := userrepositories.NewUserRepository(r.DB)
		user, err := userRepo.GetUserById(int(post.UserId))
		getPosts = append(getPosts, personalcontentmodels.GetPost{
			ID:             post.Id,
			Title:          post.Title,
			Body:           post.Body,
			UserId:         post.UserId,
			Photos:         photos,
			Count_likes:    int(countLikes),
			Count_comments: int(countComments),
			User:           user,
		})
	}
	return getPosts, nil
}
func (r *PersonalContentRepository) CheckIsLike(postID string, userID string) (bool, error) {
	var count int64
	if err := r.DB.Table(personalcontentmodels.Like{}.TableName()).
		Where("post_id = ? AND user_id = ?", postID, userID).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
