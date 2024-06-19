package main

import (
	"bytes"
	"context"
	"errors"
	"github.com/google/go-github/v53/github"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"os"
	"reflect"
	"studentgit.kata.academy/Alkolex/go-kata/course2/3.patterns/3.patterns_strategy/task2.3.3.1/mocks"
	"studentgit.kata.academy/Alkolex/go-kata/course2/3.patterns/3.patterns_strategy/task2.3.3.1/models"

	"testing"
)

func TestNewGithubRepo(t *testing.T) {
	httpClient := &http.Client{}
	ngbc := github.NewClient(httpClient)
	ngt := NewGithubRepo(ngbc)
	if reflect.TypeOf(ngt.client.Client()) != reflect.TypeOf(httpClient) {
		t.Errorf("expected = %v, got = %v", ngt.client.Client(), httpClient)
	}
}
func TestNewGithubGist(t *testing.T) {
	httpClient := &http.Client{}
	ngbc := github.NewClient(httpClient)
	ngt := NewGithubGist(ngbc)
	if reflect.TypeOf(ngt.client.Client()) != reflect.TypeOf(httpClient) {
		t.Errorf("expected = %v, got = %v", ngt.client.Client(), httpClient)
	}
}
func TestNewGeneralGithub(t *testing.T) {
	httpClient := &http.Client{}
	ngbc := github.NewClient(httpClient)
	ngt := NewGeneralGithub(ngbc)
	if reflect.TypeOf(ngt.client.Client()) != reflect.TypeOf(httpClient) {
		t.Errorf("expected = %v, got = %v", ngt.client.Client(), httpClient)
	}
}
func TestMainFunc(t *testing.T) {
	expected := "data"
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

func TestGithubGist_GithubRepo_GetItems(t *testing.T) {
	w := []models.Item{{Title: "1", Description: "2", Link: "3"}}
	gitMock := mocks.NewGithubLister(t)
	gitMock.On("GetItems", mock.Anything, "").
		Return(w, errors.New("git error"))
	got, err := gitMock.GetItems(context.Background(), "")

	if (err != nil) != true {
		t.Errorf("GetItems() error = %v, wantErr %v", err, false)
		return
	}
	if !reflect.DeepEqual(got, w) {
		t.Errorf("GetGists() got = %v, want %v", got, w)
	}
	g := mocks.NewGeneralGithubLister(t)
	var liste mocks.GithubLister
	g.On("GetItems", mock.Anything, "", mock.Anything).
		Return(w, errors.New("git error"))

	got, err = g.GetItems(context.Background(), "", liste)

	if (err != nil) != true {
		t.Errorf("GetGists() error = %v, wantErr %v", err, false)
		return
	}
	if !reflect.DeepEqual(got, w) {
		t.Errorf("GetGists() got = %v, want %v", got, w)
	}
	//got, err = g.GetItems(context.Background(), "", repo)

}
func TestGeneralGithub_GetItems(t *testing.T) {
	gitMock2 := new(mocks.GeneralGithubLister)
	gitMock2.On("GetItems", mock.Anything, "", mock.Anything).Return([]models.Item{}, nil)
	var liste mocks.GithubLister
	got, err := gitMock2.GetItems(context.Background(), "", liste)
	assert.Nil(t, err)
	assert.Equal(t, got, []models.Item{})
}
