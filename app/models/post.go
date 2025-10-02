package models

import (
	"myblog/database/factories"

	"github.com/goravel/framework/contracts/database/factory"
	"github.com/goravel/framework/database/orm"
)

type Post struct {
	orm.Model
	Title  string `json:"title"`
	Body   string `json:"body"`
	Status string `json:"status"`
}

func (p *Post) Factory() factory.Factory {
	return &factories.PostFactory{}
}
