package article

import (
	"ex01/models"
	"ex01/repositories/article"
)

type ArticleService struct {
	repository article.IRepository
}

func New(repository article.IRepository) *ArticleService {
	return &ArticleService{
		repository: repository,
	}
}

func (a *ArticleService) Create(article *models.Article) error {
	return a.repository.Create(article)
}

func (a *ArticleService) Get(id uint) (*models.Article, error) {
	return a.repository.Get(id)
}

func (a *ArticleService) GetAll(filters ...int) ([]*models.Article, error) {
	return a.repository.GetAll(filters...)
}

func (a *ArticleService) Update(article *models.Article) error {
	return a.repository.Update(article)
}

func (a *ArticleService) Delete(id uint) error {
	return a.repository.Delete(id)
}
