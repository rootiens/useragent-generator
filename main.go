package main

import (
	"fmt"
	"os"
	"sync"
)

var (
	browserNames = []string{
		"Mozilla Firefox",
		"Google Chrome",
		"Apple Safari",
		"Microsoft Edge",
		"Opera",
	}

	osNames = []string{
		"Windows NT 10.0",
		"Windows NT 6.3",
		"Windows NT 6.2",
		"Windows NT 6.1",
		"Windows NT 6.0",
		"Windows NT 5.1",
		"Macintosh; Intel Mac OS X 10_15",
		"Macintosh; Intel Mac OS X 10_14",
		"Macintosh; Intel Mac OS X 10_13",
		"Macintosh; Intel Mac OS X 10_12",
		"X11; Linux x86_64",
		"X11; Ubuntu; Linux x86_64",
		"Android 11.0; Mobile",
		"Android 10.0; Mobile",
		"Android 9.0; Mobile",
		"Android 8.1; Mobile",
		"Android 7.1; Mobile",
	}

	extensions = []string{
		".NET CLR",
		"SV1",
		"Tablet PC",
		"Win64; IA64",
		"Win64; x64",
		"Win64; x86",
		"WOW64",
		"Trident/7.0",
		"Trident/6.0",
		"Trident/5.0",
		"Trident/4.0",
	}

	languages = []string{
		"en-US",
		"en-GB",
		"en-CA",
		"en-AU",
		"en",
		"fr-FR",
		"fr-CA",
		"de-DE",
		"es-ES",
		"it-IT",
		"ja-JP",
		"ko-KR",
		"pt-BR",
		"ru-RU",
		"zh-CN",
		"zh-TW",
	}

	versions = []string{
		"70.0",
		"71.0",
		"72.0",
		"73.0",
		"74.0",
		"75.0",
		"76.0",
		"77.0",
		"78.0",
		"79.0",
		"80.0",
		"81.0",
		"82.0",
		"83.0",
		"84.0",
		"85.0",
		"86.0",
		"87.0",
		"88.0",
		"89.0",
		"90.0",
		"91.0",
		"92.0",
		"93.0",
		"94.0",
		"95.0",
		"96.0",
		"97.0",
		"98.0",
		"99.0",
	}
)

func generateUserAgents() []string {
	var userAgents []string

	for _, browser := range browserNames {
		for _, os := range osNames {
			for _, extension := range extensions {
				for _, language := range languages {
					for _, version := range versions {
						userAgent := fmt.Sprintf("%s (%s; %s; %s) %s/%s", browser, os, extension, language, browser, version)
						userAgents = append(userAgents, userAgent)
					}
				}
			}
		}
	}

	return userAgents
}

func saveToFile(userAgents []string) {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, _ = file.WriteString("const uap = [\n")

	for _, userAgent := range userAgents {
		_, err := file.WriteString("\"" + userAgent + "\",\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	_, _ = file.WriteString("]")

	fmt.Println("User agents saved to output.txt")
}

func main() {
	userAgents := generateUserAgents()

	var wg sync.WaitGroup
	wg.Add(len(userAgents))

	for _, userAgent := range userAgents {
		go func(ua string) {
			defer wg.Done()
			fmt.Println(ua)
		}(userAgent)
	}

	wg.Wait()

	saveToFile(userAgents)
}
