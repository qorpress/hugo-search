package main

import "testing"

const testHugoPath = "test"

// checks the number of pages built for the site (drafts don't count)
func TestReadSitePages(t *testing.T) {
	pages := readSitePages(testHugoPath)
	actual := pages.Len()

	/* hugo 0.18.1

	   ['title-page-3:page'
	    ':page'
	    'title-page-1:page'
	    'title-page-2:page'
	    'Search Results:page'
	    'Fail:section'
	    'Folder1:section'
	    'Tag1:taxonomy'
	    'Tag2:taxonomy'
	    'Tags:taxonomyTerm'
	    'hugo-search TEST site:home'] 11


	     hugo 0.31.1

	   ['title-page-3:page'
	    ':page'
	    'title-page-1:page'
	    'title-page-2:page'
	    'Search Results:page'
	    'Categories:taxonomyTerm'
	    'Fail:section'
	    'Folder1:section'
	    'Tag1:taxonomy'
	    'Tag2:taxonomy'
	    'Tags:taxonomyTerm'
	    'hugo-search TEST site:home'] 12
	*/
	expected := 12

	if expected != actual {
		var titles []string
		for _, page := range pages {
			titles = append(titles, "'"+page.Title()+":"+page.Kind+"'")
		}
		t.Errorf("Expected: %d, was: %d, pages returned:\n%s", expected, actual, titles)
	}
}

// checks that pages with no title are correctly detected
func TestPageHasTitle(t *testing.T) {
	pages := readSitePages(testHugoPath)
	var a, b bool
	for _, page := range pages {
		if page.Title() == "title-page-1" {
			a = pageHasTitle(page)
		} else if page.Title() == "" {
			b = !pageHasTitle(page)
		}
	}
	if !(a && b) {
		t.Errorf("Expected: has title==(true && false), was: (%v && %v)", a, b)
	}
}

// checks that valid pages are detected and special pages discarded
func TestPageHasValidContent(t *testing.T) {
	cases := map[string]bool{
		"title-page-3":          true,
		"":                      true,
		"title-page-1":          true,
		"title-page-2":          true,
		"Search Results":        false,
		"Categories":            false,
		"Fail":                  true,
		"Folder1":               true,
		"Tag1":                  false,
		"Tag2":                  false,
		"Tags":                  false,
		"hugo-search TEST site": true,
	}
	pages := readSitePages(testHugoPath)
	for _, page := range pages {
		expected := cases[page.Title()]
		actual := pageHasValidContent(page)
		if expected != actual {
			t.Errorf("Expected: %t, was: %t, page: '%s'", expected, actual, page.Title())
		}
	}
}
