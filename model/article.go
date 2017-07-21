package model

import "time"

// Article 文章
type Article struct {
    ID             uint               `gorm:"primary_key" json:"id"`
    CreatedAt      time.Time          `json:"createdAt"`
    UpdatedAt      time.Time          `json:"updatedAt"`
    DeletedAt      *time.Time         `sql:"index" json:"deletedAt"`
    Name           string             `json:"name"`
    BrowseCount    int                `json:"browseCount"`
    Status         int                `json:"status"`
    Content        string             `json:"content"`
    Categories     []Category         `gorm:"many2many:article_category;ForeignKey:ID;AssociationForeignKey:ID" json:"categories"`
}

const (
    // ArticleVerifying 审核中
    ArticleVerifying      = 1

    // ArticleVerifySuccess 审核通过
    ArticleVerifySuccess  = 2

    // ArticleVerifyFail 审核未通过
    ArticleVerifyFail     = 3
)