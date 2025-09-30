package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250930173224CreatePostsTable struct{}

// Signature The unique signature for the migration.
func (r *M20250930173224CreatePostsTable) Signature() string {
	return "20250930173224_create_posts_table"
}

// Up Run the migrations.
func (r *M20250930173224CreatePostsTable) Up() error {
	if !facades.Schema().HasTable("posts") {
		return facades.Schema().Create("posts", func(table schema.Blueprint) {
			table.ID()
			table.String("title", 255)
			table.Text("body")
			table.String("status", 20)
			table.TimestampsTz()
			table.SoftDeletes()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250930173224CreatePostsTable) Down() error {
	return facades.Schema().DropIfExists("posts")
}
