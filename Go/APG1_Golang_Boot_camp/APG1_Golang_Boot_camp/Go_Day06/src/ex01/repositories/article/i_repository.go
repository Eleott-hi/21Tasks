package article

import "ex01/models"

type IRepository interface {
	Create(article *models.Article) error
	Get(id uint) (*models.Article, error)
	GetAll() ([]*models.Article, error)
	Update(article *models.Article) error
	Delete(id uint) error
}
