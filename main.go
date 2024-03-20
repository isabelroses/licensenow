package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/charmbracelet/huh"
)

func main() {
	var url string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Pick a licence.").
				Options(
					huh.NewOption("MIT", "https://gist.githubusercontent.com/isabelroses/fa6f71651be564ada535bb56dec1e13b/raw/c766f746e94771d92bea73db881a567215b5fe77/MIT%2520License"),
					huh.NewOption("GPLv3", "https://www.gnu.org/licenses/gpl-3.0.txt"),
					huh.NewOption("cc by-nc-sa 4.0", "https://creativecommons.org/licenses/by-nc-sa/4.0/legalcode.txt"),
				).
				Value(&url),
		),
	).WithTheme(huh.ThemeCatppuccin())

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to fetch the URL: %v", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return
	}

	file, err := os.OpenFile("LICENSE", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Printf("Failed to open file: %v", err)
		return
	}
	defer file.Close()

	_, err = file.Write(body)
}
