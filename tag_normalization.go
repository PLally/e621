package e621

import (
	"errors"
	"strings"
)



type ParsedTags struct {
	OrTags []string
	NotTags []string
	NormalTags []string
}

// TODO sort tags
func (p ParsedTags) Normalized() string{
	return strings.Join(p.NotTags, " ") + strings.Join(p.NormalTags, " " )+" "+strings.Join(p.OrTags, " ")
}

func (p ParsedTags) Matches(tags []string) bool {
	var orSatisfied bool

	var satisfied = make(map[string]bool)

	for _, tag := range p.NormalTags {
		satisfied[tag] = false
	}

	for _, tag := range tags {
		if contains(p.NotTags, tag) {
			return false
		}

		_, ok := satisfied[tag]
		if ok {
			satisfied[tag] = true
		}

		if contains(p.OrTags, tag) {
			orSatisfied = true
		}
	}

	if !orSatisfied { return false }

	for _, satisfied := range satisfied {
		if !satisfied { return false }
	}

	return true
}

var ErrNoResultQuery = errors.New("This Query will never return any results")
var ErrNegateExistingTag = errors.New("You can not negate an existing tag.")
var ErrWildcardDisallowed = errors.New("Wildcards are not allowed")

func ParseTags(tags string, allowWildcards bool) (ParsedTags, error) {
	posTags := make([]string, 0)
	notTags := make([]string, 0)
	orTags := make([]string, 0)

	for _, tag := range strings.Split(tags, " ") {
		if strings.Contains(tag,  "*") && allowWildcards == false {
			return ParsedTags{}, ErrWildcardDisallowed
		}
		if tag[0:1] == "--" { // tags with repeating 'not' modifiers have no effect on the query
			continue
		} else if tag[0] == '-' {
			if tags[1] != '-' {
				notTags = append(notTags, tag)
			}
		} else if tag[0] == '~' {
			if tag[1] == '~' { return ParsedTags{}, ErrNoResultQuery }
			orTags =  append(orTags, tag)

		}  else {
			posTags = append(posTags, tag)
		}

	}

	for _, tag := range posTags {
		if contains(notTags, tag) {
			return ParsedTags{}, ErrNegateExistingTag
		}
	}

	return ParsedTags{
		posTags,
		notTags,
		orTags,
	}, nil
}

func contains(r []string, s string) bool {
	for _, v := range r {
		if s == v {
			return true
		}
	}
	return false
}