package minecraft

import (
	"github.com/lyssar/msdcli/utils"
	"github.com/manifoldco/promptui"
	"strings"
)

func (metaApi *MetaApi) RenderSelect(showAll bool) (Version, error) {

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0000279C {{ .ID | cyan }} ({{ .Type | yellow }})",
		Inactive: "  {{ .ID | black }} ({{ .Type | black }})",
		Selected: "{{ \"\u2771 Selected: \" | yellow }} {{ .ID | yellow | bold }}",
	}

	searcher := func(input string, index int) bool {
		versions := metaApi.Versions.Versions[index]
		name := strings.Replace(strings.ToLower(versions.ID), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	noBell := utils.NewNoBellStdout()
	listItems := metaApi.Versions.GetVersionsForType("release")
	if showAll {
		listItems = metaApi.Versions.Versions
	}
	prompt := promptui.Select{
		Stdout:    &noBell,
		Label:     "Select minecraft version",
		Items:     listItems,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()

	if err != nil {
		return Version{}, err
	}

	return listItems[i], nil
}
