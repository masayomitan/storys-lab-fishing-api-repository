package action

import (
	"net/http"
	// "fmt"

	"github.com/gin-gonic/gin"
	"storys-lab-fishing-api/app/adapter/api/logging"
	"storys-lab-fishing-api/app/adapter/api/response"
	"storys-lab-fishing-api/app/adapter/logger"
	"storys-lab-fishing-api/app/usecase/article"
	"storys-lab-fishing-api/app/utils"
)

type FindOneArticleAction struct {
	uc  usecase.FindOneArticleUseCase
	log logger.Logger
}

type FindAllArticleAction struct {
	uc  usecase.FindAllArticleUseCase
	log logger.Logger
}

type FindAllArticleByArticleCategoryIDAction struct {
	uc  usecase.FindAllArticleByArticleCategoryIDUseCase
	log logger.Logger
}

func NewFindOneArticleAction(uc usecase.FindOneArticleUseCase, log logger.Logger) FindOneArticleAction {
	return FindOneArticleAction{
		uc:  uc,
		log: log,
	}
}

func NewFindAllArticleAction(uc usecase.FindAllArticleUseCase, log logger.Logger) FindAllArticleAction {
	return FindAllArticleAction{
		uc:  uc,
		log: log,
	}
}

func NewFindAllArticleByArticleCategoryIdAction(uc usecase.FindAllArticleByArticleCategoryIDUseCase, log logger.Logger) FindAllArticleByArticleCategoryIDAction {
	return FindAllArticleByArticleCategoryIDAction{
		uc:  uc,
		log: log,
	}
}

func (t FindOneArticleAction) FindOne(c *gin.Context) {
	const logKey = "find_one_article"

	output, err := t.uc.Execute(c.Request.Context(), utils.StrToInt(c.Param("id")))
	if err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when returning the article")

		return
	}
	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when returning article")

	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
}


func (t FindAllArticleAction) FindAll(c *gin.Context) {
	const logKey = "find_all_articles"

	output, err := t.uc.Execute(c.Request.Context())
	if err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when returning the article list")

		return
	}
	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when returning article list")

	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
}


func (t FindAllArticleByArticleCategoryIDAction) FindAllByArticleCategoryID(c *gin.Context) {
	const logKey = "find_all_articles"

	output, err := t.uc.ExecuteSecond(c.Request.Context(), utils.StrToInt(c.Param("id")))
	if err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when returning the article list")

		return
	}
	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when returning article list")

	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
}
