package feed

import "fmt"

type Format struct {
	categories map[string]string
}

func (f *Format) Translate(source *Item) (Job, []string) {
	var messages []string
	result := Job{}

	messages = append(messages, f.translateTitle(source, &result)...)
	messages = append(messages, f.translateCategories(source, &result)...)

	if len(messages) == 0 {
		messages = append(messages, "OK")
	}

	return result, messages
}

func (f *Format) translateTitle(source *Item, result *Job) []string {
	if source.title != "" {
		result.title = source.title
		return nil
	} else {
		return []string{"Title is missing or empty"}
	}
}

func (f *Format) translateCategories(source *Item, result *Job) []string {
	var unknownCategories []string
	result.categories = []string{}

	for _, category := range source.categories {
		if translated, ok := f.categories[category]; ok {
			result.categories = append(result.categories, translated)
		} else {
			unknownCategories = append(unknownCategories, category)
		}
	}

	if unknownCategories == nil {
		return nil
	} else {
		return []string{fmt.Sprintf(
			"Unknown categories: %s", unknownCategories,
		)}
	}
}
