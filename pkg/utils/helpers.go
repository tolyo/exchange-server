package utils

import (
	"strings"
	"text/template"
)

// Format helps with string concatenation. Example: utils.Format("Variable {{.}} ", data)
func Format(s string, v interface{}) string {
	t, b := new(template.Template), new(strings.Builder)
	template.Must(t.Parse(s)).Execute(b, v)
	return b.String()
}

func Each[T comparable](list []T, fn func(T)) {
	for _, value := range list {
		if fn != nil {
			fn(value)
		}
	}
}

func Map[T, V comparable](list []T, fn func(T) V) []V {
	res := make([]V, len(list))
	for i, value := range list {
		res[i] = fn(value)
	}
	return res
}

func Filter[T comparable](list []T, fn func(T) bool) []T {
	filteredList := make([]T, 0)
	for _, value := range list {
		if fn(value) {
			filteredList = append(filteredList, value)
		}
	}
	return filteredList
}
