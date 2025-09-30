package seeders

import (
	"myblog/app/models"

	"github.com/goravel/framework/facades"
)

type PostSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *PostSeeder) Signature() string {
	return "PostSeeder"
}

// Run executes the seeder logic.
func (s *PostSeeder) Run() error {
	var posts []models.Post
	return facades.Orm().Factory().Count(10).Create(&posts)
}
