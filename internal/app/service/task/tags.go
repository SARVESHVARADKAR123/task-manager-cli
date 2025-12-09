package task

import "strings"

func sanitizeTags(tags []string) []string {
    clean := make([]string, 0, len(tags))
    seen := map[string]struct{}{}

    for _, t := range tags {
        tt := strings.TrimSpace(strings.ToLower(t))
        if tt == "" {
            continue
        }
        if _, ok := seen[tt]; ok {
            continue
        }
        seen[tt] = struct{}{}
        clean = append(clean, tt)
    }
    return clean
}
