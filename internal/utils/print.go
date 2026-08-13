package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/lanxre/kyokusu-cli/internal/models"
)

const (
	ansiReset  = "\x1b[0m"
	ansiBold   = "\x1b[1m"
	ansiDim    = "\x1b[2m"
	ansiCyan   = "\x1b[36m"
	ansiGreen  = "\x1b[32m"
	ansiYellow = "\x1b[33m"
	ansiRed    = "\x1b[31m"
)

type palette struct {
	bold, dim, cyan, green, yellow, red, reset string
}

func newPalette(w io.Writer) palette {
	if os.Getenv("NO_COLOR") != "" || !isTerminal(w) {
		return palette{}
	}
	return palette{
		bold:   ansiBold,
		dim:    ansiDim,
		cyan:   ansiCyan,
		green:  ansiGreen,
		yellow: ansiYellow,
		red:    ansiRed,
		reset:  ansiReset,
	}
}

func isTerminal(w io.Writer) bool {
	f, ok := w.(*os.File)
	if !ok {
		return false
	}
	info, err := f.Stat()
	if err != nil {
		return false
	}
	return info.Mode()&os.ModeCharDevice != 0
}

func PrintInit(w io.Writer, rootMessage models.Root, options []models.Definition, commands []models.Definition) {
	p := newPalette(w)

	printHeader(w, p, rootMessage)
	printUsage(w, p, rootMessage.Usage)
	printOptions(w, p, options)
	printCommands(w, p, commands)
}

func printHeader(w io.Writer, p palette, root models.Root) {
	if root.Long != "" {
		fmt.Fprintln(w, root.Long)
	}

	fmt.Fprintln(w)
}

func printUsage(w io.Writer, p palette, usage string) {
	fmt.Fprintf(w, "%sUsage:%s\n", p.bold, p.reset)
	fmt.Fprintf(w, "  %s%s%s %s[options] <command> [subcommand] [options]%s\n\n",
		p.cyan, usage, p.reset, p.dim, p.reset)
}

func printOptions(w io.Writer, p palette, options []models.Definition) {
	rows := make([]row, 0, len(options))
	for _, opt := range options {
		rows = append(rows, row{indent: 0, name: flagName(opt.Name), desc: opt.Short})
	}
	printSection(w, p, "Options", rows)
}

func printCommands(w io.Writer, p palette, commands []models.Definition) {
	if len(commands) == 0 {
		return
	}

	var rows []row
	for _, cmd := range commands {
		rows = append(rows, flattenCommand(cmd, 0)...)
	}

	printSection(w, p, "Commands", rows)
}

type row struct {
	indent int
	name   string
	desc   string
}

func flattenCommand(def models.Definition, indent int) []row {
	rows := []row{{indent: indent, name: def.Name, desc: def.Short}}

	for _, opt := range sorted(def.Options) {
		rows = append(rows, row{indent: indent + 2, name: flagName(opt.Name), desc: opt.Short})
	}

	for _, child := range sorted(def.Children) {
		rows = append(rows, flattenCommand(*child, indent+2)...)
	}

	return rows
}

func flagName(name string) string {
	if len(name) == 1 {
		return "-" + name
	}
	return "--" + name
}

func sorted(m map[string]*models.Definition) []*models.Definition {
	if len(m) == 0 {
		return nil
	}

	names := make([]string, 0, len(m))
	for name := range m {
		names = append(names, name)
	}
	sort.Strings(names)

	defs := make([]*models.Definition, 0, len(m))
	for _, name := range names {
		defs = append(defs, m[name])
	}
	return defs
}

func printSection(w io.Writer, p palette, title string, rows []row) {
	if len(rows) == 0 {
		return
	}

	fmt.Fprintf(w, "%s%s:%s\n", p.bold, title, p.reset)

	width := 0
	for _, r := range rows {
		if n := r.indent + len(r.name); n > width {
			width = n
		}
	}

	for _, r := range rows {
		label := strings.Repeat(" ", r.indent) + r.name
		fmt.Fprintf(w, "  %s%s%s%-*s  %s\n",
			p.green, label, p.reset,
			width-len(label), "",
			r.desc,
		)
	}

	fmt.Fprintln(w)
}

func PrintInputJSON(in models.Input) {
	bytes, err := json.MarshalIndent(in, "", "  ")
	if err != nil {
		return
	}
	fmt.Println(string(bytes))
}
