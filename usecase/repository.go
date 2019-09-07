package usecase

import (
	"context"
	"fmt"

	entity "../entity"
	"github.com/google/go-github/github"
)

func Repositories(ctx context.Context, query string) ([]entity.RepositorySummary, error) {
	client := github.NewClient(nil)
	opts := &github.SearchOptions{
		Sort:  "stars",
		Order: "asc",
	}

	repoDatas := []entity.RepositorySummary{}
	result, _, err := client.Search.Repositories(ctx, query, opts)
	if err != nil {
		return repoDatas, err
	}
	repositories := result.Repositories
	fmt.Println("len repo", len(repositories))
	for _, repo := range repositories {
		fmt.Println(*repo.Name)
		summary := entity.RepositorySummary{
			Id:          *repo.ID,
			FullName:    *repo.FullName,
			Star:        *repo.StargazersCount,
			Owner:       repo.Owner.GetName(),
			Description: *repo.Description,
			UpdatedAt:   repo.UpdatedAt.String(),
			Language:    *repo.Language,
			Size:        *repo.Size,
			License:     repo.License.GetName(),
		}
		repoDatas = append(repoDatas, summary)
	}
	return repoDatas, err
}
