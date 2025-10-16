package main

import "os"

func copy_file(filename string) {
	src := "template/" + filename
	dst := "output/" + filename
	if _, err := os.Stat(src); os.IsNotExist(err) {
		panic("File not found: " + src)
	}
	data, err := os.ReadFile(src)
	if err != nil {
		panic(err)
	}
	os.WriteFile(dst, data, 0644)
}

func init_() {
	// Ensure the output directory exists
	if err := os.MkdirAll("output", 0755); err != nil {
		panic(err)
	}

	// Copy static files
	copy_file("style.css")
	copy_file("profile.png")
	copy_file("profile.jpeg")
}

func dots(n int) string {
	const total = 5
	dotsHTML := ""
	for i := 0; i < total; i++ {
		if i < n {
			dotsHTML += `<span class="dot active">●</span>`
		} else {
			dotsHTML += `<span class="dot inactive">●</span>`
		}
	}
	return dotsHTML
}
