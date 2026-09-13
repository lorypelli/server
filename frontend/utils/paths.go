package utils

import (
	"net/url"
	"path"
	"strings"

	"github.com/a-h/templ"
)

type Crumb struct {
	Name string
	Href templ.SafeURL
}

func Parent(route string, view View) templ.SafeURL {
	return Dir(path.Dir(path.Clean(route)), view)
}

func File(route, name string) templ.SafeURL {
	return templ.URL((&url.URL{Path: path.Join(route, name)}).String())
}

func Folder(route, name string, view View) templ.SafeURL {
	return Dir(path.Join(route, name), view)
}

func Dir(route string, view View) templ.SafeURL {
	return view.location(route)
}

func Crumbs(route string, view View) []Crumb {
	segments := strings.FieldsFunc(path.Clean(route), func(r rune) bool { return r == '/' })
	crumbs := make([]Crumb, len(segments))
	href := "/"
	for i, segment := range segments {
		href = path.Join(href, segment)
		crumbs[i] = Crumb{Name: segment, Href: Dir(href, view)}
	}
	return crumbs
}
