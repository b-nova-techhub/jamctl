package gen

import (
	"b-nova-openhub/stapagen/pkg/header"
	"b-nova-openhub/stapagen/pkg/md"
	"b-nova-openhub/stapagen/pkg/url"
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

var GeneratedPages []StaticPage
var Generated *Generating

func Generate(files []string) *Generating {
	pages := make([]StaticPage, 0)

	for _, f := range files {
		var p StaticPage
		frontMatter := header.GetFrontMatter(f)
		p.Title = frontMatter["title"]
		p.Permalink = url.GetPermalink(frontMatter["title"])
		p.Author = frontMatter["author"]
		p.Tags = frontMatter["tags"]
		p.Categories = frontMatter["categories"]
		p.PublishDate = frontMatter["date"]
		p.Description = frontMatter["description"]
		p.ShowComments = frontMatter["showComments"]
		p.IsPublished = frontMatter["publish"]
		p.Body = md.ConvertBodyToMarkdown(f)
		pages = append(pages, p)
	}

	GeneratedPages = pages
	SetStatus()

	Generated = new(Generating)
	Generated.Success = true
	Generated.Errors = make([]string, 0)
	return Generated
}
