package e926

import (
	"errors"
	"sort"
	"strings"
)

type ParsedTags struct {
	OrTags     []string
	NotTags    []string
	NormalTags []string
}

func (p ParsedTags) Normalized() string {
	tags := p.NotTags
	tags = append(tags, p.NormalTags...)
	tags = append(tags, p.OrTags...)
	sort.Strings(tags)

	out := strings.Join(tags, " ")
	return out
}

func (p ParsedTags) Matches(tags TagContainer) bool {
	var orSatisfied bool
	if len(p.OrTags) == 0 {
		orSatisfied = true
	}
	var satisfied = make(map[string]bool)

	for _, tag := range p.NormalTags {
		satisfied[tag] = false
	}

	for _, tag := range tags.All() {
		if contains(p.NotTags, "-"+tag) {
			return false
		}

		_, ok := satisfied[tag]
		if ok {
			satisfied[tag] = true
		}

		if contains(p.OrTags, "~"+tag) {

			orSatisfied = true
		}
	}

	if !orSatisfied {
		return false
	}

	for _, satisfied := range satisfied {
		if !satisfied {
			return false
		}
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
		if tag == "" {
			continue
		}
		if strings.Contains(tag, "*") && allowWildcards == false {
			return ParsedTags{}, ErrWildcardDisallowed
		}
		if tag[0:1] == "--" { // tags with repeating 'not' modifiers have no effect on the query
			continue
		} else if tag[0] == '-' {
			if tags[1] != '-' {
				notTags = append(notTags, tag)
			}
		} else if tag[0] == '~' {
			if tag[1] == '~' {
				return ParsedTags{}, ErrNoResultQuery
			}
			orTags = append(orTags, tag)

		} else {
			posTags = append(posTags, tag)
		}

	}

	for _, tag := range posTags {
		if contains(notTags, tag) {
			return ParsedTags{}, ErrNegateExistingTag
		}
	}
	return ParsedTags{
		orTags,
		notTags,
		posTags,
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
