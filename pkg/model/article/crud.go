package article

import (
	"goblog/pkg/model"
	"goblog/pkg/types"
)

// Get 通过 id 获取文章
func Get(idstr string) (Article, error) {
	var err error
	var article Article
	id := types.StringToUnit64(idstr)
	if err = model.DB.First(&article, id).Error; err != nil {
		return article, err
	}

	return article, err
}

func GetAll() ([]Article, error) {
	var err error
	var articles []Article

	if err = model.DB.Find(&articles).Error; err != nil {
		return articles, err
	}
	return articles, nil
}
