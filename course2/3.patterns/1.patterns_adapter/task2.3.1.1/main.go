package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v53/github"
	"studentgit.kata.academy/Alkolex/go-kata/course2/3.patterns/1.patterns_adapter/task2.3.1.1/model"
)

//go:generate mockgen -source=main.go -destination=mocks/mock.go

////go:generate go run github.com/vektra/mockery/v2@v2.20.2 --name=Githuber

type Githuber interface {
	GetGists(ctx context.Context, username string) ([]model.Item, error)
	GetRepos(ctx context.Context, username string) ([]model.Item, error)
}

type Github struct {
	client *github.Client
}

func NewGithub(c *github.Client) *Github {
	return &Github{
		client: c,
	}
}

func (i *Github) GetGists(ctx context.Context, username string) ([]model.Item, error) {
	gists, _, err := i.client.Gists.List(ctx, username, nil)
	if err != nil {
		return []model.Item{}, err
	}
	var outItem = make([]model.Item, 0, len(gists))
	var count int
	for _, gist := range gists {
		count++
		outItem = append(outItem, model.Item{
			Title:       fmt.Sprintf("gist_title_%d", count),
			Description: gist.GetDescription(),
			Link:        gist.GetHTMLURL(),
		})
	}
	return outItem, err
}

func (i *Github) GetRepos(ctx context.Context, username string) ([]model.Item, error) {
	reposes, _, err := i.client.Repositories.List(ctx, username, nil)
	if err != nil {
		return []model.Item{}, err
	}
	var outItem = make([]model.Item, 0, len(reposes))
	for _, repo := range reposes {
		outItem = append(outItem, model.Item{
			Title:       repo.GetName(),
			Description: repo.GetDescription(),
			Link:        repo.GetURL(),
		})
	}
	return outItem, err
}

//func main() {
//	ctx := context.Background()
//	ts := oauth2.StaticTokenSource(
//		&oauth2.Token{AccessToken: "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"},
//	)
//	tc := oauth2.NewClient(ctx, ts)
//	client := github.NewClient(tc)
//
//	g := NewGithub(client)
//	fmt.Println(g.GetGists(context.Background(), "ptflp"))
//	fmt.Println(g.GetRepos(context.Background(), "ptflp"))
//}
