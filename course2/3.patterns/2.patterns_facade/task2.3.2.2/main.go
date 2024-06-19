package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
	"studentgit.kata.academy/Alkolex/go-kata/course2/3.patterns/2.patterns_facade/task2.3.2.2/models"
)

//go:generate go run github.com/vektra/mockery/v2@v2.20.2 --name=Githuber
type Githuber interface {
	GetGists(ctx context.Context, username string) ([]models.Item, error)
	GetRepos(ctx context.Context, username string) ([]models.Item, error)
}

//go:generate go run github.com/vektra/mockery/v2@v2.20.2 --name=AdapterGithuber
type AdapterGithuber interface {
	GetGists(ctx context.Context, username string) ([]models.Item, error)
	GetRepos(ctx context.Context, username string) ([]models.Item, error)
}

type Github struct {
	g      Githuber
	client *github.Client
}

func NewGithub(c *github.Client) *Github {
	return &Github{client: c}
}

func (g *Github) GetGists(ctx context.Context, username string) ([]models.Item, error) {
	userGiList, _, err := g.client.Gists.List(ctx, username, nil)
	if err != nil {
		return []models.Item{}, err
	}

	outItem := make([]models.Item, 0, len(userGiList))
	for _, val := range userGiList {
		outItem = append(outItem, models.Item{
			Title:       "Title_" + val.GetDescription(),
			Description: val.GetDescription(),
			Link:        val.GetHTMLURL(),
		})
	}
	return outItem, err
}

func (i *Github) GetRepos(ctx context.Context, username string) ([]models.Item, error) {
	reposes, _, err := i.client.Repositories.List(ctx, username, nil)
	if err != nil {
		return []models.Item{}, err
	}
	var outItem = make([]models.Item, 0, len(reposes))
	for _, repo := range reposes {
		outItem = append(outItem, models.Item{
			Title:       repo.GetName(),
			Description: repo.GetDescription(),
			Link:        repo.GetURL(),
		})
	}

	return outItem, err
}

//------------------------------------------------

type GithubProxy struct {
	git   Github
	cache map[string][]models.Item
}

func NewGithubProxy(c *github.Client) *GithubProxy {
	return &GithubProxy{
		git:   Github{client: c},
		cache: make(map[string][]models.Item, 2),
	}
}

func (g *GithubProxy) GetGists(ctx context.Context, username string) ([]models.Item, error) {
	cacheData := username + "GetGists"
	if item, ok := g.cache[cacheData]; ok {
		return item, nil
	}
	newItem, err := g.git.GetGists(ctx, username)
	if err != nil {
		return nil, err
	}

	g.cache[cacheData] = newItem
	return newItem, err
}
func (g *GithubProxy) GetRepos(ctx context.Context, username string) ([]models.Item, error) {
	cacheData := username + "GetRepos"
	if item, ok := g.cache[cacheData]; ok {
		return item, nil
	}

	newItem, err := g.git.GetRepos(ctx, username)
	if err != nil {
		return nil, err
	}

	g.cache[cacheData] = newItem
	return newItem, err
}

type GithubAdapter struct {
	gitAdapt AdapterGithuber
}

func (g *GithubAdapter) GetGists(ctx context.Context, username string) ([]models.Item, error) {
	return g.gitAdapt.GetGists(ctx, username)
}
func (g *GithubAdapter) GetRepos(ctx context.Context, username string) ([]models.Item, error) {
	return g.gitAdapt.GetRepos(ctx, username)
}
func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "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"},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	g := NewGithubProxy(client)
	_, _ = g.GetGists(context.Background(), "ihontin")
	_, _ = g.GetRepos(context.Background(), "ihontin")
	fmt.Print("not a bad test")
}
