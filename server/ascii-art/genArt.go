// Ascii-art is a program which consists in receiving a string as an argument
// and outputting the string in a graphic representation using ASCII.
package asciiweb

import (
	"fmt"
	"os"
	"strings"

	// ascii "asciiweb/ascii-art/utilities"
	ascii "asciiweb/server/ascii-art/utilities"
)

func GenArt(userInput, banner string) string {
	asciiArt := ""

	userInput = ascii.HandleBackspace(userInput)
	userInput = strings.ReplaceAll(string(userInput), "\\t", "   ") // handling the tab sequence

	// Read the ascii art text file

	filename := "server/ascii-art/" + banner
	file, err := os.ReadFile(filename)
	if err != nil {
		return "500"
	}

	//
	fileData := ascii.StringContain(string(file))

	// Handling empty files
	if len(fileData) != 856 {
		fmt.Println("Error: >> Banner files  is empty with length of: ", len(file))
		return ""
	}

	userInput = strings.ReplaceAll(userInput, "\\n", "\n")
	userInput = strings.ReplaceAll(userInput, "\r", "\n")
	userInputParts, err := ascii.HandleNewLine(userInput)
	ascii.ErrHandler(err)

	for _, part := range userInputParts {
		if part == "" {
			fmt.Println() // Print a newline if the part is empty (i.e., consecutive newline characters)
			continue
		}
		for i := 0; i < 8; i++ { // this loop is responsible for the height of each character
			for _, char := range part { // iterates through each character of the current word
				startingIndex := ascii.GetStartingIndex(int(char)) // obtaining the starting position of the char
				if startingIndex >= 0 {
					asciiArt += fileData[startingIndex+i] // printing the character line by line
				}
			}
			asciiArt += "\n" // printing a new line after printing each line of the charcter
		}
	}
	return asciiArt
}
