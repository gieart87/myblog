package factories

import (
	"math/rand"

	"github.com/bxcodec/faker/v4"
)

type PostFactory struct {
}

// Definition Define the model's default state.
func (f *PostFactory) Definition() map[string]any {
	statuses := []string{"DRAFT", "PUBLISHED", "ARCHIVED"}

	return map[string]any{
		"Title":  faker.Sentence(),
		"Body":   faker.Paragraph(),
		"Status": statuses[rand.Intn(len(statuses))],
	}
}
