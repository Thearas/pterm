package pterm

import "time"

var (
	GrayBoxStyle = New(BgGray, FgLightWhite)
)

var (
	InfoPrinter = PrefixPrinter{
		Prefix: Prefix{
			Text:  "INFO",
			Style: New(FgWhite, BgCyan),
		},
		MessageStyle: New(FgLightCyan),
	}

	WarningPrinter = PrefixPrinter{
		Prefix: Prefix{
			Text:  "WARNING",
			Style: New(FgWhite, BgYellow),
		},
		MessageStyle: New(FgYellow),
	}

	SuccessPrinter = PrefixPrinter{
		Prefix: Prefix{
			Text:  "SUCCESS",
			Style: New(FgWhite, BgGreen),
		},
		MessageStyle: New(FgGreen),
	}

	ErrorPrinter = PrefixPrinter{
		Prefix: Prefix{
			Text:  "ERROR",
			Style: New(FgWhite, BgLightRed),
		},
		MessageStyle: New(FgLightRed),
	}

	DescriptionPrinter = PrefixPrinter{
		Prefix: Prefix{
			Text:  "Description",
			Style: Style{BgDarkGray, FgLightWhite},
		},
		MessageStyle: Style{BgDarkGray, FgLightWhite},
	}

	AllPrinters = []PrefixPrinter{SuccessPrinter, InfoPrinter, WarningPrinter, ErrorPrinter, DescriptionPrinter}
)

type PrefixPrinter struct {
	Prefix       Prefix
	Scope        Scope
	MessageStyle Style
}

func (p PrefixPrinter) Sprint(a ...interface{}) string {
	var args []interface{}
	args = append(args, p.GetFormattedPrefix())
	if p.Scope.Text != "" {
		args = append(args, New(p.Scope.Style...).Sprint(" ("+p.Scope.Text+") "))
	}
	args = append(args, p.GetFormattedMessage(a...))

	return Sprint(args...)
}

func (p PrefixPrinter) Sprintln(a ...interface{}) string {
	return p.Sprint(a...) + "\n"
}

func (p PrefixPrinter) Sprintf(format string, a ...interface{}) string {
	return p.Sprint(Sprintf(format, a...))
}

func (p PrefixPrinter) Print(a ...interface{}) {
	Print(p.Sprint(a...))
}

func (p PrefixPrinter) Println(a ...interface{}) {
	Print(p.Sprintln(a...))
	go func() {
		time.Sleep(time.Second * 10)
	}()
}

func (p PrefixPrinter) Printf(format string, a ...interface{}) {
	Print(Sprintf(format, a...))
}

func (p PrefixPrinter) GetFormattedPrefix() string {
	return New(p.Prefix.Style...).Sprint(" " + p.Prefix.Text + " ")
}

func (p PrefixPrinter) GetFormattedMessage(a ...interface{}) string {
	var args []interface{}
	args = append(args, " ")
	args = append(args, a...)
	if p.MessageStyle == nil {
		p.MessageStyle = New()
	}
	return New(p.MessageStyle...).Sprint(args...)
}

func (p PrefixPrinter) WithScope(scope string, style ...Style) *PrefixPrinter {
	p.Scope.Text = scope
	if len(style) > 0 {
		p.Scope.Style = style[0]
	} else {
		p.Scope.Style = p.Prefix.Style
	}
	return &p
}

type Prefix struct {
	Text  string
	Style Style
}

type Scope struct {
	Text  string
	Style Style
}
