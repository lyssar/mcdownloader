package minecraft

import (
	"fmt"
	"github.com/chzyer/readline"
	"github.com/manifoldco/promptui"
	"strings"
)

type noBellStdout struct{}

func (n *noBellStdout) Write(p []byte) (int, error) {
	if len(p) == 1 && p[0] == readline.CharBell {
		return 0, nil
	}
	return readline.Stdout.Write(p)
}

func (n *noBellStdout) Close() error {
	return readline.Stdout.Close()
}

var NoBellStdout = &noBellStdout{}

func (metaApi *MetaApi) RenderSelect() (Version, error) {

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F336 {{ .ID | cyan }} ({{ .Type | red }})",
		Inactive: "  {{ .ID | cyan }} ({{ .Type | red }})",
		Selected: "\U0001F336 {{ .ID | red | cyan }}",
	}

	searcher := func(input string, index int) bool {
		versions := metaApi.Versions.Versions[index]
		name := strings.Replace(strings.ToLower(versions.ID), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Stdout:    NoBellStdout,
		Label:     "Select minecraft version",
		Items:     metaApi.Versions.Versions,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil, err
	}

	return metaApi.Versions.Versions[i], nil
}
