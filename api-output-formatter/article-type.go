package apiOutputFormatter

type ArticleType struct {
	ArticleTypeArticleType  []ArticleTypeArticleType `json:"ArticleTypeArticleType"`
	ArticleTypeText         []ArticleTypeText        `json:"ArticleTypeText"`
	Accepter                []string                 `json:"Accepter"`
}

type ArticleTypeArticleType struct {
	ArticleType            string `json:"ArticleType"`
}

type ArticleTypeText struct {
	ArticleType            string `json:"ArticleType"`
	Language               string `json:"Language"`
	ArticleTypeName        string `json:"ArticleTypeName"`
}
