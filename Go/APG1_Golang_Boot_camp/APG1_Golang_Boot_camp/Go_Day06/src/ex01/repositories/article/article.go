package article

import (
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

func (a *ArticleRepository) GetAll() ([]*models.Article, error) {
	var articles []*models.Article

	a.database.Find(&articles)

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
