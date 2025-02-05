package memes

import (
	"context"

	lorem "github.com/derektata/lorem/ipsum"
	"github.com/ger9000/memes-as-a-service/internal/domain/meme"
)

type GetAllAction struct {
}

func NewGetAllAction() *GetAllAction {
	return &GetAllAction{}
}

func (a *GetAllAction) Do(ctx context.Context, latitude, longitude float64, query string) []meme.Meme {
	g := lorem.NewGenerator()
	g.WordsPerSentence = 10
	g.SentencesPerParagraph = 5
	g.CommaAddChance = 3
	memesCount := 10
	if query != "" {
		memesCount = len(query)
	}

	memes := []meme.Meme{}
	for i := 0; i < memesCount; i++ {
		memes = append(memes, meme.Meme{
			Body: g.Generate(100),
			Tags: "",
		})
	}

	return memes
}
