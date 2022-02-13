package policies

import (
	"goblog/pkg/auth"
	"goblog/pkg/model/article"
)

func CanModifyArticle(_article article.Article) bool {
	return auth.User().ID == _article.UserID
}
