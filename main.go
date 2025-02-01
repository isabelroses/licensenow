package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/template"
	"time"

	"github.com/charmbracelet/huh"
)

var (
	outfile = "LICENSE"
	license string
	year    string
	name    string
)

func main() {
	// If no arguments are passed, show a form
	if len(os.Args) == 1 {
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("What's your name?").
					Value(&name),
				huh.NewInput().
					Title("What year is it?").
					Value(&year),
				huh.NewSelect[string]().
					Title("Pick a license.").
					OptionsFunc(func() []huh.Option[string] {
						var options []huh.Option[string]
						for v := range licenseNames {
							li := licenseNames[v]
							options = append(options, huh.NewOption(li, li))
						}
						return options
					}, &license).
					Value(&license),
			),
		).WithTheme(huh.ThemeCatppuccin())

		err := form.Run()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// otherwise parse the flags
		flag.StringVar(&license, "licence", "", "License to generate")
		flag.StringVar(&outfile, "outfile", outfile, "Output file")
		flag.StringVar(&outfile, "o", "LICENSE", "Output file")
		flag.StringVar(&year, "year", "", "Year to use in the license")
		flag.StringVar(&year, "y", "", "Year to use in the license")
		flag.StringVar(&name, "name", "", "Name to use in the license")
		flag.StringVar(&name, "n", "", "Name to use in the license")
		flag.Parse()
	}

	// We now parse the license as a template now
	tmpl, err := template.New("license").Parse(licenses[license])
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(outfile)
	if err != nil {
		log.Fatal(err)
	}

	// Data to pass to the template
	data := struct {
		Name string
		Year string
	}{
		Name: getName(),
		Year: getYear(),
	}

	err = tmpl.Execute(f, data)
	if err != nil {
		log.Print("execute: ", err)
		return
	}

	f.Close()
}

func getName() string {
	if name != "" {
		return name
	}
	return os.Getenv("USER")
}

func getYear() string {
	if year != "" {
		return year
	}
	return fmt.Sprint(time.Now().Year())
}
