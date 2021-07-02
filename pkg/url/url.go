package url

import (
	"github.com/gosimple/slug"
	"strings"
)

func GetPermalink(s string) string {
	return "b-nova.com/home/content/" + slug.Make(s)
}

func GetSlug(s string) string {
	slices := strings.Split(s, "/")
	return slices[len(slices)-1]
}
