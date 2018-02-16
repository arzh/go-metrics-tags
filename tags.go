package go_metrics_tags

import (
	"github.com/rcrowley/go-metrics"
	"strings"
)

const (
	tagSep = "|tags:"
)

func NameWithTags(name string, tags []string) string {
	if len(tags) == 0 {
		return name
	}

	return name + tagSep + strings.Join(tags, ",")
}

func HasTags(name string) bool {
	return strings.Contains(name, tagSep)
}

func GetTags(name string) (string, []string) {
	sp := strings.Split(name, tagSep)
	n := sp[0]
	tags := make([]string, len(sp)-1)
	if len(sp) > 1 {
		tags = strings.Split(sp[1], ",")
	}

	return n, tags
}

func TagMetric(registry metrics.Registry, name string, tags []string, m interface{}) interface{} {
	id := NameWithTags(name, tags)
	return registry.GetOrRegister(id, m)
}
