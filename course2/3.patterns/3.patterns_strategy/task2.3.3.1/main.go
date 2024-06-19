package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
	"log"
	"studentgit.kata.academy/Alkolex/go-kata/course2/3.patterns/3.patterns_strategy/task2.3.3.1/models"
)

//go:generate go run github.com/vektra/mockery/v2@v2.20.2 --name=GithubLister
type GithubLister interface {
	GetItems(ctx context.Context, username string) ([]models.Item, error)
}

//---------------------- Gist

type GithubGist struct {
	client *github.Client
}

func NewGithubGist(c *github.Client) *GithubGist {
	return &GithubGist{client: c}
}

func (g *GithubGist) GetItems(ctx context.Context, username string) ([]models.Item, error) {
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

//---------------------- Repo

type GithubRepo struct {
	client *github.Client
}

func NewGithubRepo(c *github.Client) *GithubRepo {
	return &GithubRepo{client: c}
}

func (i *GithubRepo) GetItems(ctx context.Context, username string) ([]models.Item, error) {
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

// ------------------------------------ General

//go:generate go run github.com/vektra/mockery/v2@v2.20.2 --name=GeneralGithubLister
type GeneralGithubLister interface {
	GetItems(ctx context.Context, username string, strategy GithubLister) ([]models.Item, error)
}

type GeneralGithub struct {
	gLister GithubLister
	client  *github.Client
}

func NewGeneralGithub(c *github.Client) *GeneralGithub {
	return &GeneralGithub{client: c}
}

func (g *GeneralGithub) GetItems(ctx context.Context, username string, strgy GithubLister) ([]models.Item, error) {
	g.gLister = strgy
	return g.gLister.GetItems(ctx, username)
}

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "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"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	gist := NewGithubGist(client)
	repo := NewGithubRepo(client)

	gg := NewGeneralGithub(client)

	/*data*/
	_, err := gg.GetItems(context.Background(), "ihontin", gist)
	if err != nil {
		log.Println(err)
	}
	//fmt.Println(data)
	/*data*/
	_, err = gg.GetItems(context.Background(), "ihontin", repo)
	if err != nil {
		log.Println(err)
	}
	//fmt.Println(data)
	fmt.Print("data")
}
