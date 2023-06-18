package main

// Data holds the basic structure of the data read from the YAML file
type Data struct {
	Title       string
	Description string
	Sections    []Section
	Badges      []Badge
	License     string
	Copywrite   string
	Authors     []Authors
	Anchors     []Anchors
	References  []References
}

// Badge holds the structure for badge data
type Badge struct {
	Name  string
	Value string
	Color string
	URL   string
}

// Section holds the structure for section data
type Section struct {
	Title   string
	Content string
	Link    string
}

// Authors holds the structure for Anchor data
type Authors struct {
	Name string
	URL  string
}

// Anchor holds the structure for Anchor data
type Anchors struct {
	Name string
	URL  string
}

// References holds the structure for Anchor data
type References struct {
	Name        string
	Description string
	URL         string
}

// TemplateData holds the structure of the data used in the template
type TemplateData struct {
	Title       string
	Description string
	Sections    []Section
	Badges      []Badge
	License     string
	Authors     []Authors
	Anchors     []Anchors
	References  []References
}
