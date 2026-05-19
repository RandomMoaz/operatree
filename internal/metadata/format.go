package metadata

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func FormatName(name string) string {
	// convert spaces to -
	if len(name) == 0 {
		return ""
	}

	words := strings.Fields(name)
	return strings.Join(words, "-")
}

// converts string tags and returns correct tags array
func ParseTags(tags string) []string {
	if len(tags) == 0 {
		return []string{}
	}

	s := strings.Split(tags, ",")

	for i, v := range s {
		t := v
		t = strings.Trim(v, " ")
		t = strings.ToLower(t)
		s[i] = t
	}

	return s
}

// converts string tags and returns correct tags array
func ParseParticipants(participants string) []string {

	if len(participants) == 0 {
		return []string{}
	}

	s := strings.Split(participants, ",")

	for i, v := range s {
		t := v
		t = strings.Trim(v, " ")
		t = cases.Title(language.English, cases.Compact).String(t)

		s[i] = t
	}

	return s
}

func ParsePersonName(name string) string {
	return cases.Title(language.English, cases.Compact).String(name)
}
