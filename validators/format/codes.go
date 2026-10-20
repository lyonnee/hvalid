package format

import "strings"

func codeSet(list string) map[string]struct{} {
	set := make(map[string]struct{}, 64)
	for _, code := range strings.Fields(list) {
		set[code] = struct{}{}
	}
	return set
}
