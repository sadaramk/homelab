package main

import (
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/dotenv-org/godotenvvault"
	"github.com/ghodss/yaml"
	"golang.org/x/exp/slog"
)

func main() {

	//prepare logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	slog.SetDefault(logger)

	// Load Environment Variables
	dataFile, templateFile, readmeFile := loadFileLocations()

	// Load data from YAML file
	file, err := ioutil.ReadFile(dataFile)
	if err != nil {
		slog.Error("Error reading data yaml file - Error %s", err)
		os.Exit(1)
	}

	// Unmarshal the YAML data to the struct
	var data Data
	err = yaml.Unmarshal(file, &data)
	if err != nil {
		slog.Error("Error marshalling data from yaml to internal structure - Error %s", err)
		os.Exit(1)
	}

	// Transform the section names to generate the links
	var sections []Section
	for _, section := range data.Sections {
		link := strings.ToLower(strings.ReplaceAll(section.Title, " ", "-"))
		sections = append(sections, Section{Title: section.Title, Link: link, Content: section.Content})
	}

	// Create a new struct to pass to the template
	templateData := TemplateData{
		Title:       data.Title,
		Description: data.Description,
		Sections:    sections,
		Badges:      data.Badges,
		License:     data.License,
		Anchors:     data.Anchors,
		Authors:     data.Authors,
		References:  data.References,
	}

	slog.Info(data.Description)

	// Open the template file
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		slog.Error("Error reading template file - Error %s", err)
		os.Exit(1)
	}

	// Create a new README.md file
	newReadmeFile, err := os.Create(readmeFile)
	if err != nil {
		slog.Error("Error creating read me  file - Error %s", err)
		os.Exit(1)
	}
	defer newReadmeFile.Close()

	// Apply the data to the template and write the result to the file
	err = tmpl.Execute(newReadmeFile, templateData)
	if err != nil {
		slog.Error("Error applying data to the template and writing to the file - Error %s", err)
		os.Exit(1)
	}
}

func loadFileLocations() (string, string, string) {
	//
	err := godotenvvault.Load("../config.env")
	if err != nil {
		slog.Error("Error loading config.env file")
		os.Exit(1)
	}

	// Read environment variables
	dataFile := os.Getenv("HOME_README_DATA_FILE")
	templateFile := os.Getenv("README_TEMPLATE_FILE")
	readmeFile := os.Getenv("HOME_README_FILE")

	// Check if environment variables are set
	if dataFile == "" || templateFile == "" || readmeFile == "" {
		slog.Error("Please set HOME_README_DATA_FILE, TEMPLATE_FILE, and README_FILE environment variables")

		//exit if one of the input environment variables are missing
		os.Exit(1)
	}

	return dataFile, templateFile, readmeFile

}
