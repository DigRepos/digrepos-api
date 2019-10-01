package usecase

import (
	"context"
	"fmt"

	entity "../entity"
	"github.com/google/go-github/github"
)

func Repositories(ctx context.Context, query string) ([]entity.RepositorySummary, error) {
	fmt.Println(query)
	client := github.NewClient(nil)
	opts := &github.SearchOptions{
		Sort:  "stars",
		Order: "desc",
		ListOptions: github.ListOptions{
			Page:    1,
			PerPage: 300,
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
			Id:            repo.GetID(),
			HTMLURL:       repo.GetHTMLURL(),
			FullName:      repo.GetFullName(),
			Star:          repo.GetStargazersCount(),
			ForksCount:    repo.GetForksCount(),
			WatchersCount: repo.GetWatchersCount(),
			Owner: entity.Owner{
				Name:      repo.Owner.GetName(),
				AvatarUrl: repo.Owner.GetAvatarURL(),
			},
			Description: repo.GetDescription(),
			Homepage:    repo.GetHomepage(),
			Topics:      repo.Topics,
			UpdatedAt:   repo.UpdatedAt.String(),
			Language:    repo.GetLanguage(),
			Size:        repo.GetSize(),
			License:     repo.License.GetName(),
		}
		repoDatas = append(repoDatas, summary)
	}
	return repoDatas, err
}
