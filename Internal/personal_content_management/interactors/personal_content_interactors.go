package personalcontentinteractors

import (
	personalcontentmodels "BESocialHealth/Internal/personal_content_management/models"
	personalcontentrepositories "BESocialHealth/Internal/personal_content_management/repositories"
)

type PersonalContentInteractor struct {
	PersonalContentRepository *personalcontentrepositories.PersonalContentRepository
}

func NewPersonalContentInteractor(repo *personalcontentrepositories.PersonalContentRepository) *PersonalContentInteractor {
	return &PersonalContentInteractor{
		PersonalContentRepository: repo,
	}
}
func (i *PersonalContentInteractor) CreatePost(post *personalcontentmodels.CreatePostFull) error {
	if err := i.PersonalContentRepository.CreateFullPost(post); err != nil {
		return err
	}
	return nil
}
func (i *PersonalContentInteractor) UpdatePostById(id int, post *personalcontentmodels.CreatePost) error {
	if err := i.PersonalContentRepository.UpdatePostById(id, post); err != nil {
		return err
	}
	return nil
}
func (i *PersonalContentInteractor) DeletePostById(id int) error {
	if err := i.PersonalContentRepository.DeletePostById(id); err != nil {
		return err
	}
	return nil
}

//	func (i *PersonalContentInteractor) GetPostById(id int) (*personalcontentmodels.CreatePost, error) {
//		return nil
//	}
func (i *PersonalContentInteractor) CreateLike(like *personalcontentmodels.CreateLike) error {
	if err := i.PersonalContentRepository.CreateLike(like); err != nil {
		return err
	}
	return nil
}
func (i *PersonalContentInteractor) DeleteLikeByUserIDAndPostId(userId int, postId int) error {
	return i.PersonalContentRepository.DeleteLikeByUserIDAndPostId(userId, postId)
}
func (i *PersonalContentInteractor) CreateComent(coment *personalcontentmodels.CreateCommentFull) error {
	if err := i.PersonalContentRepository.CreateComentFull(coment); err != nil {
		return err
	}
	return nil
}
func (i *PersonalContentInteractor) GetPostById(id int) (*personalcontentmodels.GetPost, error) {
	post, err := i.PersonalContentRepository.GetPostById(id)
	if err != nil {
		return nil, err
	}

	return &post, nil
}
func (i *PersonalContentInteractor) GetAllComentByPostId(id int) (*[]personalcontentmodels.GetComment, error) {
	comments, err := i.PersonalContentRepository.GetAllCommentByPostId(id)
	if err != nil {
		return nil, err
	}
	return &comments, nil
}
func (i *PersonalContentInteractor) GetAllPost() (*[]personalcontentmodels.GetPost, error) {
	posts, err := i.PersonalContentRepository.GetAllPosts()
	if err != nil {
		return nil, err
	}
	return &posts, nil
}
