package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v53/github"
	"net/http"
	"reflect"
	mock_main "studentgit.kata.academy/Alkolex/go-kata/course2/practice/task2/mocks"
	"studentgit.kata.academy/Alkolex/go-kata/course2/practice/task2/model"
	"testing"
)

func TestGithub_GetGists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGists := mock_main.NewMockGithuber(ctrl)
	mockGists.EXPECT().GetGists(gomock.Any(), "ihontin").
		Return([]model.Item{
			{
				Title:       "gist_title_1",
				Description: "description1",
				Link:        "url1",
			}}, nil,
		)

	ctx := context.Background()
	username := "ihontin"
	items, err := mockGists.GetGists(ctx, username)
	if err != nil {
		t.Fatalf("GetGists returned an error: %v", err)
	}
	expectedGists := []model.Item{
		model.Item{
			Title:       "get_title_1",
			Description: "description1",
			Link:        "url1",
		},
	}

	if len(items) != len(expectedGists) {
		t.Fatalf("GetGists returned incorrect number of items. Expected: %d, Got: %d", len(expectedGists), len(items))
	}

	for i, item := range items {
		if item.Title != fmt.Sprintf("gist_title_%d", i+1) {
			t.Errorf("GetGists returned incorrect title. Expected: %s, Got: %s", fmt.Sprintf("gist_title_%d", i+1), item.Title)
		}

		if item.Description != expectedGists[i].Description {
			t.Errorf("GetGists returned incorrect description. Expected: %s, Got: %s", expectedGists[i].Description, item.Description)
		}

		if item.Link != expectedGists[i].Link {
			t.Errorf("GetGists returned incorrect link. Expected: %s, Got: %s", expectedGists[i].Link, item.Link)
		}
	}
	testError := errors.New("this is test error")
	mockGists.EXPECT().GetGisList(gomock.Any(), "ihontin", gomock.Any()).
		Return(nil, nil, testError)

	_, _, err = mockGists.GetGisList(ctx, "ihontin", nil)
	if err == nil {
		t.Errorf("Error expected but got nil")
	}
}

func TestGithub_GetRepos(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepos := mock_main.NewMockGithuber(ctrl)
	mockRepos.EXPECT().GetRepos(gomock.Any(), "ihontin").
		Return([]model.Item{
			{
				Title:       "Repo_title_1",
				Description: "description1",
				Link:        "url1",
			},
			{
				Title:       "Repo_title_2",
				Description: "description2",
				Link:        "url2",
			},
		}, nil,
		)

	ctx := context.Background()
	username := "ihontin"
	items, err := mockRepos.GetRepos(ctx, username)
	if err != nil {
		t.Fatalf("GetRepos returned an error: %v", err)
	}
	expectedRepos := []model.Item{
		model.Item{
			Title:       "Repo_title_1",
			Description: "description1",
			Link:        "url1",
		},
		model.Item{
			Title:       "Repo_title_2",
			Description: "description2",
			Link:        "url2",
		},
	}
	if len(items) != len(expectedRepos) {
		t.Fatalf("GetRepos returned incorrect number of items. Expected: %d, Got: %d", len(expectedRepos), len(items))
	}

	for i, item := range items {
		if item.Title != expectedRepos[i].Title {
			t.Errorf("GetGists returned incorrect title. Expected: %s, Got: %s", expectedRepos[i].Title, item.Title)
		}

		if item.Description != expectedRepos[i].Description {
			t.Errorf("GetRepos returned incorrect description. Expected: %s, Got: %s", expectedRepos[i].Description, item.Description)
		}

		if item.Link != expectedRepos[i].Link {
			t.Errorf("GetRepos returned incorrect link. Expected: %s, Got: %s", expectedRepos[i].Link, item.Link)
		}
	}
}

func TestNewGithub(t *testing.T) {
	httpClient := &http.Client{}
	ngbc := github.NewClient(httpClient)
	ngt := NewGithub(ngbc)
	if reflect.TypeOf(ngt.client.Client()) != reflect.TypeOf(httpClient) {
		t.Errorf("expected = %v, got = %v", ngt.client.Client(), httpClient)
	}
}
