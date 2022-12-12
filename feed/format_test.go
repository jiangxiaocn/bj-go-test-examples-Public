package feed

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAcceptsValidFeedItem(t *testing.T) {
	// Given
	subject := aFormat()
	source := aValidItem()

	// When
	_, messages := subject.Translate(&source)

	// Then
	assert.Equal(t, []string{"OK"}, messages)
}

func TestTranslatesTitle(t *testing.T) {
	// Given
	subject := aFormat()
	source := aValidItem()
	source.title = "Software Engineer"

	// When
	actual, messages := subject.Translate(&source)

	// Then
	assert.Equal(t, "Software Engineer", actual.title)
	assert.Equal(t, []string{"OK"}, messages)
}

func TestDetectsMissingTitle(t *testing.T) {
	// Given
	subject := aFormat()
	source := aValidItem()
	source.title = ""

	// When
	actual, messages := subject.Translate(&source)

	// Then
	assert.Equal(t, "", actual.title)
	assert.Equal(t, []string{"Title is missing or empty"}, messages)
}

func TestTranslatesCategories(t *testing.T) {
	// Given
	subject := aFormat()
	subject.categories = map[string]string{"c1": "C1", "c2": "C2"}

	source := aValidItem()
	source.categories = []string{"c1", "c2"}

	// When
	actual, messages := subject.Translate(&source)

	// Then
	assert.Equal(t, []string{"C1", "C2"}, actual.categories)
	assert.Equal(t, []string{"OK"}, messages)
}

func TestDetectsUnknownCategories(t *testing.T) {
	// Given
	subject := aFormat()
	subject.categories = map[string]string{"c1": "C1", "c2": "C2"}

	source := aValidItem()
	source.categories = []string{"c1", "u1", "c2", "u2"}

	// When
	actual, messages := subject.Translate(&source)

	// Then
	assert.Equal(t, []string{"C1", "C2"}, actual.categories)
	assert.Equal(t, []string{"Unknown categories: [u1 u2]"}, messages)
}

func TestAcceptsMissingCategories(t *testing.T) {
	// Given
	subject := aFormat()
	subject.categories = map[string]string{"c1": "C1", "c2": "C2"}

	source := aValidItem()
	source.categories = nil

	// When
	actual, messages := subject.Translate(&source)

	// Then
	assert.Equal(t, []string{}, actual.categories)
	assert.Equal(t, []string{"OK"}, messages)
}

func aFormat() Format {
	return Format{}
}

func aValidItem() Item {
	return Item{
		title: "Dummy title",
	}
}
