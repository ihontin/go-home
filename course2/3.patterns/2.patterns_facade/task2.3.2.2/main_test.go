package main

import (
	"bytes"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"os"
	"reflect"
	"studentgit.kata.academy/Alkolex/go-kata/course2/3.patterns/2.patterns_facade/task2.3.2.2/mocks"
	"studentgit.kata.academy/Alkolex/go-kata/course2/3.patterns/2.patterns_facade/task2.3.2.2/models"
	"testing"
)

func TestGithubProxy_GetGists(t *testing.T) {

	w := []models.Item{{Title: "1", Description: "2", Link: "3"}}
	gitMock := mocks.NewAdapterGithuber(t)
	gitMock.On("GetGists", mock.Anything, "").
		Return(w, errors.New("git error"))
	g := &GithubAdapter{
		gitAdapt: gitMock,
	}

	got, err := g.GetGists(context.Background(), "")

	if (err != nil) != true {
		t.Errorf("GetGists() error = %v, wantErr %v", err, true)
		return
	}
	if !reflect.DeepEqual(got, w) {
		t.Errorf("GetGists() got = %v, want %v", got, w)
	}

	gitMock2 := new(mocks.Githuber)
	gitMock2.On("GetGists", mock.Anything, "").Return([]models.Item{}, nil)

	got, err = gitMock2.GetGists(context.Background(), "")
	assert.Nil(t, err)
	assert.Equal(t, got, []models.Item{})
}

func TestGithubProxy_GetRepos(t *testing.T) {

	w := []models.Item{{Title: "1", Description: "2", Link: "3"}}
	gitMock := mocks.NewAdapterGithuber(t)
	gitMock.On("GetRepos", mock.Anything, "").
		Return(w, errors.New("git error"))
	g := &GithubAdapter{
		gitAdapt: gitMock,
	}

	got, err := g.GetRepos(context.Background(), "")

	if (err != nil) != true {
		t.Errorf("GetGists() error = %v, wantErr %v", err, true)
		return
	}
	if !reflect.DeepEqual(got, w) {
		t.Errorf("GetGists() got = %v, want %v", got, w)
	}

	gitMock2 := new(mocks.Githuber)
	gitMock2.On("GetRepos", mock.Anything, "").Return([]models.Item{}, nil)

	got, err = gitMock2.GetRepos(context.Background(), "")
	assert.Nil(t, err)
	assert.Equal(t, got, []models.Item{})
}

func TestMainFunc(t *testing.T) {
	expected := "not a bad test"
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = old

	stdout := bytes.Buffer{}
	_, _ = stdout.ReadFrom(r)
	if expected != stdout.String() {
		t.Errorf("expected = %s, got = %s", expected, stdout.String())
	}

}
