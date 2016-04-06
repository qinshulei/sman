package sman

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
)

func showSnippets(slice SnippetSlice) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 4, 2, ' ', 0)
	sort.Sort(slice)
	for _, s := range slice {
		fmt.Fprintln(w, s.DisplayFile())
		DashLine()
		fmt.Fprintln(w, "\tName:\t"+s.Name)
		fmt.Fprintln(w, "\tDesc:\t"+s.DisplayDesc())
		fmt.Fprintln(w, "\tTags:\t"+s.DisplayTags())
		fmt.Fprintln(w, "\tDo:\t"+s.DisplayDo())
		fmt.Fprintln(w, "\tCommand:\t")
		fmt.Fprintln(w)
		for _, l := range strings.Split(s.DisplayCommand(), "\n") {
			fmt.Fprintln(w, "\t ", l)
		}
		fmt.Fprintln(w)
		for i, p := range s.Placeholders {
			i++
			n := fmt.Sprintf("\t\t\t[%v] %s", i, p.DisplayName())
			fmt.Fprintln(w, n)
			fmt.Fprintln(w, "\t\t\t\t\tOptions:\t"+p.DisplayOptions())
			fmt.Fprintln(w, "\t\t\t\t\tDesc:\t"+p.DisplayDesc())
		}
	}
	w.Flush()
}

func show(name string) {
	c := GetConfig()
	snippets := GetSnippets(name, fileFlag, c.SnippetDir, tagFlag)
	matchedSnippets := FSearchSnippet(snippets, name)
	showSnippets(matchedSnippets)
}

var showCmd = &cobra.Command{
	Use:     "show <snippet>",
	Aliases: []string{"s"},
	Short:   "Show snippet details",
	Long: `
Show all snippet details matching name or file.

Use tags flag to filter snippets by tag. Seperate multiple tags with comma.
Used when file is not defined and multiple snippets exists with the same name.

Examples:
s show alias/add -t bash

`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Show what?")
			os.Exit(1)
		}
		show(args[0])
	},
}

func init() {
	RootCmd.AddCommand(showCmd)
}
