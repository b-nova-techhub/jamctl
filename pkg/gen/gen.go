package gen

import (
	"github.com/b-nova-techhub/jamctl/pkg/util"
	"github.com/gomarkdown/markdown"
	"github.com/gosimple/slug"
	"github.com/spf13/viper"
	"strings"
)

type StaticPage struct {
	Title        string `json:"title"`
	Permalink    string `json:"permalink"`
	Author       string `json:"author"`
	Tags         string `json:"tags"`
	Categories   string `json:"categories"`
	PublishDate  string `json:"publishDate"`
	Description  string `json:"description"`
	ShowComments string `json:"showComments"`
	IsPublished  string `json:"isPublished"`
	Body         string `json:"body"`
}

type Generating struct {
	Success bool     `json:"success"`
	Errors  []string `json:"errors"`
}

var (
	GeneratedPages []StaticPage
	Generated      *Generating
)

func Generate(files []string) *Generating {
	pages := make([]StaticPage, 0)

	for _, f := range files {
		var p StaticPage
		frontMatter := getFrontMatter(f)
		p.Title = frontMatter["title"]
		p.Permalink = getPermalink(frontMatter["title"])
		p.Author = frontMatter["author"]
		p.Tags = frontMatter["tags"]
		p.Categories = frontMatter["categories"]
		p.PublishDate = frontMatter["date"]
		p.Description = frontMatter["description"]
		p.ShowComments = frontMatter["showComments"]
		p.IsPublished = frontMatter["publish"]
		p.Body = convertBodyToMarkdown(f)
		pages = append(pages, p)
	}

	GeneratedPages = pages
	Generated = new(Generating)
	Generated.Success = true
	Generated.Errors = make([]string, 0)
	return Generated
}

func getFrontMatter(file string) map[string]string {
	header := getHeader(file)
	return getKeyValues(header, ":")
}

func getHeader(file string) string {
	header := util.SubstringBetween(file, "<"+viper.GetString("delimiter")+">", "</"+viper.GetString("delimiter")+">")
	return string(header)
}

func getKeyValues(s, sep string) map[string]string {
	lines := strings.Split(s, "\n")
	frontMatter := make(map[string]string)
	for _, pair := range lines {
		if util.StringNotEmpty(pair) {
			kv := strings.Split(pair, sep)
			frontMatter[kv[0]] = strings.TrimSpace(kv[1])
		}
	}
	return frontMatter
}

func getPermalink(s string) string {
	return "b-nova.com/home/content/" + slug.Make(s)
}

func getSlug(s string) string {
	slices := strings.Split(s, "/")
	return slices[len(slices)-1]
}

func convertBodyToMarkdown(s string) string {
	return string(markdown.ToHTML([]byte(getBody(s)), nil, nil))
}

func getBody(s string) string {
	return util.SubstringAfter(s, "</"+viper.GetString("delimiter")+">")
}
