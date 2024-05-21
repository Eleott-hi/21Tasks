package article

import (
	"errors"
	"ex01/models"

	"gorm.io/gorm"
)

type ArticleRepository struct {
	database *gorm.DB
}

func New(database *gorm.DB) *ArticleRepository {
	return &ArticleRepository{database}
}

func (a *ArticleRepository) Create(article *models.Article) error {
	a.database.Create(article)

	return nil
}

func (a *ArticleRepository) Get(id uint) (*models.Article, error) {
	var article models.Article

	a.database.First(&article, id)

	if article.ID == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &article, nil
}

func (a *ArticleRepository) GetAll(filters ...int) ([]*models.Article, error) {
	var articles []*models.Article

	switch len(filters) {
	case 0:
		a.database.Find(&articles)
	case 1:
		a.database.Offset(filters[0]).Find(&articles)
	case 2:
		a.database.Offset(filters[0]).Limit(filters[1]).Find(&articles)
	default:
		return nil, errors.New("invalid number of filters")
	}

	return articles, nil
}

func (a *ArticleRepository) Update(article *models.Article) error {
	a.database.Save(article)

	return nil
}

func (a *ArticleRepository) Delete(id uint) error {
	a.database.Delete(&models.Article{}, id)
	return nil
}
