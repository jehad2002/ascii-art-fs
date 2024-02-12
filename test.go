// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func moaz(text string) string {
// 	var s int
// 	fmt.Println("choose a font 😊")
// 	fmt.Println("standard = 1")
// 	fmt.Println("shadow = 2")
// 	fmt.Println("thinkertoy = 3")
// 	fmt.Scan(&s)
// 	var m2 []string
// 	var m1 []string
// 	for i := 0; i < len(text); i++ {
// 		switch s {
// 		case 1:
// 			m2 = standard(text[i])
// 		case 2:
// 			m2 = shadow(text[i])
// 		case 3:
// 			m2 = thinkertoy(text[i])
// 		}
// 		if i == 0 {
// 			m1 = m2
// 		} else {
// 			for j := 0; j < len(m1); j++ {
// 				m1[j] += m2[j]
// 			}
// 		}
// 	}
// 	return strings.Join(m1, "\n")
// }

// func main() {
// 	if len(os.Args) != 2 {
// 		fmt.Println("Usage: go run . \"{YourText}\"")
// 		return
// 	}
// 	inputText := os.Args[1]
// 	lines := strings.Split(strings.TrimSpace(inputText), "\\n")
// 	asciiArt := []string{}
// 	for _, line := range lines {
// 		asciiArt = append(asciiArt, moaz(line))
// 	}
// 	asciiArtText := strings.Join(asciiArt, "\n")
// 	fmt.Println(asciiArtText)
// }

package main

import (
	"fmt"
	"os"
	"strings"
)

func moaz(text string, fontType string) (string, string) {
	var s int
	switch fontType {
	case "standard":
		s = 1
	case "shadow":
		s = 2
	case "thinkertoy":
		s = 3
	default:
		fmt.Println("Invalid font type. Using standard font.")
		s = 1
	}

	var m2 []string
	var m1 []string
	for i := 0; i < len(text); i++ {
		switch s {
		case 1:
			m2 = standard(text[i])
		case 2:
			m2 = shadow(text[i])
		case 3:
			m2 = thinkertoy(text[i])
		}
		if i == 0 {
			m1 = m2
		} else {
			for j := 0; j < len(m1); j++ {
				m1[j] += m2[j]
			}
		}
	}
	return strings.Join(m1, "\n"), fontType
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . \"{YourText}\" {FontType}")
		return
	}
	inputText := os.Args[1]
	fontType := os.Args[2]

	art, chosenFontType := moaz(inputText, fontType)
	fmt.Println("Font Type:", chosenFontType)
	fmt.Println(art)
}
