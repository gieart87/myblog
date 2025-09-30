package models

import (
	"myblog/database/factories"

	"github.com/goravel/framework/contracts/database/factory"
	"github.com/goravel/framework/database/orm"
)

type Post struct {
	orm.Model
	Title  string
	Body   string
	Status string
}

func (p *Post) Factory() factory.Factory {
	return &factories.PostFactory{}
}
