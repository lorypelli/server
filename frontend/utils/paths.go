package utils

import (
	"fmt"
	"path"
	"strings"

	"github.com/a-h/templ"
)

type Crumb struct {
	Name string
	Href templ.SafeURL
}

func Parent(route string) templ.SafeURL {
	return Dir(path.Dir(path.Clean(route)))
}

func File(route, name string) templ.SafeURL {
	return templ.URL(path.Join(route, name))
}

func Folder(route, name string) templ.SafeURL {
	return Dir(path.Join(route, name))
}

func Dir(route string) templ.SafeURL {
	if route == "/" {
		return templ.URL(route)
	}
	return templ.URL(fmt.Sprintf("%s/", route))
}

func Crumbs(route string) []Crumb {
	segments := strings.Split(strings.Trim(path.Clean(route), "/"), "/")
	if segments[0] == "" {
		return nil
	}
	crumbs := make([]Crumb, len(segments))
	href := ""
	for i, segment := range segments {
		href = fmt.Sprintf("%s/%s", href, segment)
		crumbs[i] = Crumb{Name: segment, Href: Dir(href)}
	}
	return crumbs
}
