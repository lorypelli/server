package utils

import (
	"net/url"

	"github.com/a-h/templ"
)

type View string

const (
	GridView View = "grid"
	ListView View = "list"
)

func ParseView(raw string) View {
	if View(raw) == ListView {
		return ListView
	}
	return GridView
}

func (v View) Toggle() View {
	if v == GridView {
		return ListView
	}
	return GridView
}

func (v View) Href() templ.SafeURL {
	return templ.URL((&url.URL{RawQuery: v.params().Encode()}).String())
}

func (v View) params() url.Values {
	return url.Values{"view": {string(v)}}
}

func (v View) location(dir string) templ.SafeURL {
	u := (&url.URL{Path: dir}).JoinPath("/")
	if v != GridView {
		u.RawQuery = v.params().Encode()
	}
	return templ.URL(u.String())
}
