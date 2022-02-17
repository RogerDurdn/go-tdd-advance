package blogrenderer_test

import (
	"bytes"
	"github.com/approvals/go-approval-tests"
	blogrenderer "github.com/go-tdd-advance/blogrender"
	"io"
	"testing"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		postRender, err := blogrenderer.NewPostRender()
		if err != nil {
			t.Fatal(err)
		}
		buf := bytes.Buffer{}
		err = postRender.Render(&buf, aPost)
		if err != nil {
			t.Fatal(err)
		}
		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)
	b.ResetTimer()
	postRender, err := blogrenderer.NewPostRender()
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		err := postRender.Render(io.Discard, aPost)
		if err != nil {
			b.Fatal(err)
		}
	}
}