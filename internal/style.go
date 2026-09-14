package internal

import (
	"fmt"
	"strings"
	"time"

	"github.com/pterm/pterm"
)

const badgeWidth = 7

type Printer struct {
	pterm.PrefixPrinter
}

var (
	Info    = badge(pterm.Info, "INFO")
	Success = badge(pterm.Success, "SUCCESS")
	Warning = badge(pterm.Warning, "WARNING")
	Error   = badge(pterm.Error, "ERROR")
)

var (
	Muted  = pterm.NewStyle(pterm.FgGray).Sprint
	Strong = pterm.NewStyle(pterm.Bold).Sprint
	Link   = pterm.NewStyle(pterm.FgLightCyan, pterm.Underscore).Sprint
)

func badge(base pterm.PrefixPrinter, label string) Printer {
	return Printer{*base.WithPrefix(pterm.Prefix{Text: center(label), Style: base.Prefix.Style})}
}

func Badge(label string, style *pterm.Style) Printer {
	return Printer{pterm.PrefixPrinter{
		Prefix:       pterm.Prefix{Text: center(label), Style: style},
		MessageStyle: pterm.NewStyle(),
	}}
}

func (p Printer) Log(format string, args ...any) {
	p.Println(Columns(Muted(time.Now().Format(time.TimeOnly)), fmt.Sprintf(format, args...)))
}

func center(label string) string {
	padding := max(badgeWidth-len(label), 0)
	left := padding / 2
	return fmt.Sprintf("%*s%s%*s", left, "", label, padding-left, "")
}

func Columns(cols ...string) string {
	return strings.Join(cols, Muted(" │ "))
}

func Duration(d time.Duration) string {
	switch {
	case d < time.Millisecond:
		return fmt.Sprintf("%dµs", d.Microseconds())
	case d < time.Second:
		return fmt.Sprintf("%.1fms", float64(d)/float64(time.Millisecond))
	default:
		return fmt.Sprintf("%.2fs", d.Seconds())
	}
}
