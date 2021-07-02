package header

import (
	"b-nova-openhub/stapagen/pkg/config"
	"b-nova-openhub/stapagen/pkg/util"
	"strings"
)

func GetFrontMatter(file string) map[string]string {
	header := getHeader(file)
	return getKeyValues(header, ":")
}

func getHeader(file string) string {
	header := util.SubstringBetween(file, "<"+config.AppConfig.ContentDelimiterTag+">", "</"+config.AppConfig.ContentDelimiterTag+">")
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
