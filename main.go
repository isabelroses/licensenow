package main

import (
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/charmbracelet/huh"
)

var (
	licence string
	outfile = "LICENSE"

	name = os.Getenv("USER")
	year = time.Now().Year()

	args = os.Args
)

func main() {
	if len(args) == 1 {
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title("Pick a licence.").
					Options(
						huh.NewOption(licnceOpt("MIT")),
						huh.NewOption(licnceOpt("GPLv3")),
						huh.NewOption(licnceOpt("cc by-nc-sa 4.0")),
					).
					Value(&licence),
			),
		).WithTheme(huh.ThemeCatppuccin())

		err := form.Run()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		args = os.Args[1:]

		for i := 0; i < len(args); i++ {
			switch args[i] {
			case "-o":
				i++
				outfile = args[i]
			case "-n":
				i++
				name = args[i]
			case "-y":
				i++
				year, _ = strconv.Atoi(args[i])
			default:
				licence = cleanLicence(args[i])
			}
		}
	}

	tmpl := template.Must(template.ParseFiles(licence))

	f, err := os.Create(outfile)
	if err != nil {
		log.Fatal(err)
	}

	data := struct {
		Name string
		Year int
	}{
		Name: name,
		Year: year,
	}

	err = tmpl.Execute(f, data)
	if err != nil {
		log.Print("execute: ", err)
		return
	}

	f.Close()
}

func licnceOpt(licence string) (string, string) {
	file := cleanLicence(licence)
	return licence, file
}

func cleanLicence(licence string) string {
	cleanLicence := strings.ToLower(licence)
	cleanLicence = strings.ReplaceAll(cleanLicence, " ", "-")
	cleanLicence = strings.ReplaceAll(cleanLicence, ".", "")
	return path.Join("./templates", cleanLicence+".tmpls")
}
