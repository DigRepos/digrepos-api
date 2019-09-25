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
		ListOptions: github.ListOptions{
			Page:    1,
			PerPage: 100,
		},
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
			Id:            *repo.ID,
			HTMLURL:       *repo.HTMLURL,
			FullName:      *repo.FullName,
			Star:          *repo.StargazersCount,
			ForksCount:    *repo.ForksCount,
			WatchersCount: *repo.WatchersCount,
			Owner: entity.Owner{
				Name:      repo.Owner.GetName(),
				AvatarUrl: repo.Owner.GetAvatarURL(),
			},
			Description: *repo.Description,
			Homepage:    repo.GetHomepage(),
			Topics:      repo.Topics,
			UpdatedAt:   repo.UpdatedAt.String(),
			Language:    repo.GetLanguage(),
			Size:        *repo.Size,
			License:     repo.License.GetName(),
		}
		repoDatas = append(repoDatas, summary)
	}
	return repoDatas, err
}
