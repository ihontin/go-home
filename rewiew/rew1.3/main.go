package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-yaml/yaml"
	"gitlab.com/ptflp/gomarkdown"
	"os"
	"regexp"
	"strings"
	"time"
)

type Post struct {
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Topics      []Topic        `json:"topics"`
	Tags        map[string]Tag `json:"tags"`
	Author      Author         `json:"author"`
	Date        time.Time      `json:"date"`
}

type Author struct {
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
}

type Tag struct {
	Name string `json:"name"`
}

type Topic struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Parser interface {
	Parse(filename string) Parser
	Sanitize(tag string, descriptionWord string) Parser
	ValidateError() error
	GetPosts() []Post
	SaveJSON(filename string) error
	SaveYAML(filename string) error
	SaveMD(filename string) error
}

type MyParser struct {
	Posts []Post
	md    gomarkdown.MarkDown
	Error error
}

func (p *MyParser) Parse(filename string) Parser {
	// set posts to empty slice
	p.Posts = []Post{}

	p.md.ParseFile(filename)

	if p.md.DocumentMD == nil {
		p.Error = fmt.Errorf("empty document")
		return p
	}

	if p.md.DocumentMD.Childs == nil {
		p.Error = fmt.Errorf("empty childs")
		return p
	}

	rawPosts := p.md.DocumentMD.Childs
	for _, rawPost := range rawPosts {
		post := parsePost(rawPost)
		if post.Title != "" {
			p.Posts = append(p.Posts, post)
		}
	}
	return p
}

func parsePost(rawPost *gomarkdown.Container) Post {
	if rawPost == nil {
		return Post{}
	}
	if rawPost.Type != gomarkdown.Header {
		return Post{}
	}
	var post Post
	post.Tags = make(map[string]Tag)
	post.Title = rawPost.Content
	layout := "3:04 PM 1/02/2006"
	for _, child := range rawPost.Childs {
		switch child.Type {
		case gomarkdown.Paragraph:
			post.Description = child.Content
		case gomarkdown.Header:
			var topic Topic
			if child.Level != 2 {
				continue
			}
			if child.Content == "" {
				continue
			}
			topic = Topic{Title: child.Content}
			for _, subChild := range child.Childs {
				if strings.Contains(subChild.Content, "type:") {
					for _, rawTag := range strings.Split(subChild.Content, "type:") {
						tags := strings.Split(rawTag, " ")
						if len(tags) > 0 && tags[0] != "user" {
							post.Tags[tags[0]] = Tag{
								Name: tags[0],
							}
							continue
						}
						if len(tags) > 0 && tags[0] == "user" {
							authorDatas := strings.Split(strings.ReplaceAll(rawTag, "user ", ""), "(")
							if len(authorDatas) > 1 {
								post.Author = Author{
									Name:     strings.TrimSpace(authorDatas[0]),
									Nickname: strings.ReplaceAll(authorDatas[1], ")", ""),
								}
							}
							continue
						}
					}
				}
				t, err := time.Parse(layout, subChild.Content)
				if err == nil {
					post.Date = t
					break
				}
				if subChild.Type == gomarkdown.Paragraph {
					topic.Description = subChild.Content
				}
			}
			post.Topics = append(post.Topics, topic)
		}
	}
	return post
}

// Ваш код для санитизации данных и валидации
func (p *MyParser) Sanitize(tag string, descriptionWord string) Parser {
	re, _ := regexp.Compile(descriptionWord)
	for val := range p.Posts {
		p.Posts[val].Description = re.ReplaceAllString(p.Posts[val].Description, "")
		p.Posts[val].Tags[tag] = Tag{"Unknown"}
	}
	return p
}

func (p *MyParser) ValidateError() error {
	myErr := errors.New("validation error")
	for _, po := range p.Posts {
		if po.Title == "" {
			return fmt.Errorf("%v, Title is empty", myErr)
		}
	}
	return nil
}

func (p *MyParser) GetPosts() []Post {
	return p.Posts
}

// Ваш код для сохранения в различных форматах

func (p *MyParser) SaveJSON(filename string) error {
	newErr := errors.New("string is empty")
	if len(filename) < 1 {
		return newErr
	}
	data, err := json.Marshal(p.Posts)
	if err != nil {
		return err
	}
	err = saveToFile(filename, data)
	return err
}

func saveToFile(filename string, data []byte) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.FileMode(0644))
	if err != nil {
		return err
	}
	defer file.Close()
	//writer := bufio.NewWriter(file) // create writer to the file
	//writer.Write(data)              // writing to the file
	//writer.Flush()
	_, err = file.Write(data)
	return err
}

func (p *MyParser) SaveYAML(filename string) error {
	// Ваш код для сохранения в формате YAML
	newErr := errors.New("string is empty")
	if len(filename) < 1 {
		return newErr
	}
	data, err := yaml.Marshal(p.Posts)
	if err != nil {
		return err
	}
	err = saveToFile(filename, data)
	return err
}

func (p *MyParser) SaveMD(filename string) error {
	var buffer bytes.Buffer
	p.md.RenderMD(&buffer, nil)
	return saveToFile(filename, buffer.Bytes())
}

func main() {
	parser := &MyParser{}
	parser.Parse("README.md")
	parser.SaveJSON("output.json")
	parser.SaveYAML("output.yaml")
	parser.SaveMD("output.md")
}
