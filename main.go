package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

var Lowercase = "abcdefghijklmnopqrstuvwxyz"
var Uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func contains(s []string, str string) bool {
	for _, v := range s {
		if strings.Contains(str, v) {
			return true
		}
	}

	return false
}

func bruteforceMsg(txtToBruteforce string) (string, int) {
	Vocabulary := []string{"The", "in", "if", "and", "by", "on"}

	//Defeniere maximale Rotation1
	var maxRotation = 24

	//Entschlüssel den Text mit einer rotation von 1 bis max.

	for substitution := 1; substitution < maxRotation; substitution++ {
		/*for y := 1; y < maxRotation; y++ {
			fmt.Println("düdüdüdü")

		}*/
		//txtToBruteforce = "The"
		//Prüfe bei jeder Rotation ob der Text matches mit vocabulary hat
		var newText = decryptText(txtToBruteforce, substitution)
		//Wenn es ein match gibt - liefere entschlüsselten Text und die Rotation zurück
		if contains(Vocabulary, newText) {
			fmt.Println("In this text we found some matches: ", newText)
			fmt.Println(contains(Vocabulary, newText))
			fmt.Println("This text was cracked with the ", substitution, "rotation.")

		} else {
			fmt.Println("Nothing was found")
		}

	}

	return "", 0
}

//getIndexOfLowerCaseLetter returns the index of the given character in the lowercase alphabet if not found: return -1
func getIndexOfLowerCaseLetter(char rune) int {
	for index, letter := range Lowercase {
		if char == letter {
			return index
		}

	}
	return -1
}

//getIndexOfUpperCaseLetter returns the index of the given character in the Uppercase alphabet if not found: return -1
func getIndexOfUpperCaseLetter(char rune) int {
	for index, letter := range Uppercase {
		if char == letter {
			return index
		}
	}
	return -1
}

//getNewRuneLowerCase takes the given index and returns the new rune based on given rotation number
func getNewRuneLowerCase(index int, rotation int) rune {
	newIndex := index + rotation
	for newIndex >= len(Lowercase) {
		newIndex = newIndex - (len(Lowercase))
	}

	return rune(Lowercase[newIndex])

}

//getNewRuneUpperCase takes the given index and returns the new rune based on given rotation number
func getNewRuneUpperCase(index int, rotation int) (newRune rune) {
	newIndex := index + rotation
	for newIndex >= len(Uppercase) {
		newIndex = newIndex - (len(Uppercase))
	}
	return rune(Uppercase[newIndex])

}

//encryptString takes given string and return encrypted string based on given rotation number
func encryptString(txtToEncrypt string, rotation int) string {
	var buffer bytes.Buffer
	for _, char := range txtToEncrypt {
		//Schritt 1: index von gegebenen character holen im kontext von LowerCase
		currentIndex := getIndexOfLowerCaseLetter(char)
		if currentIndex > -1 {
			//Schritt 2: addieren von der rotation
			newRune := getNewRuneLowerCase(currentIndex, rotation)
			//Schritt 3: Schreiben
			buffer.WriteRune(newRune)
		} else {
			currentIndex = getIndexOfUpperCaseLetter(char)
			if currentIndex > -1 {
				newRune := getNewRuneUpperCase(currentIndex, rotation)
				buffer.WriteRune(newRune)
			} else {
				buffer.WriteRune(char)
			}
		}
	}
	return buffer.String()

}

//getNewRuneLowerCaseDec takes the given index and returns the new rune based on given rotation number
func getNewRuneLowerCaseDec(index int, rotation int) rune {
	newIndex := index - rotation
	for newIndex < 0 {
		newIndex = newIndex + (len(Lowercase))
	}

	return rune(Lowercase[newIndex])

}

//getNewRuneUpperCaseDec takes the given index and returns the new rune based on given rotation number
func getNewRuneUpperCaseDec(index int, rotation int) (newRune rune) {
	newIndex := index - rotation
	for newIndex < 0 {
		newIndex = newIndex + (len(Uppercase))
	}
	return rune(Uppercase[newIndex])

}

//decryptText takes given string and returns decrypted message based on the rotation
func decryptText(txtToDecrypt string, rotation int) string {
	var buffer bytes.Buffer
	for _, char := range txtToDecrypt {
		currentIndex := getIndexOfLowerCaseLetter(char)
		if currentIndex > -1 {
			newRune := getNewRuneLowerCaseDec(currentIndex, rotation)
			buffer.WriteRune(newRune)
		} else {
			currentIndex = getIndexOfUpperCaseLetter(char)
			if currentIndex > -1 {
				newRune := getNewRuneUpperCaseDec(currentIndex, rotation)
				buffer.WriteRune(newRune)
			} else {
				buffer.WriteRune(char)
			}
		}
	}
	return buffer.String()

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var txtTDecode string
	var txtTEncode string
	var substitution int
	var mode int

	fmt.Println("What do you want? 1 for encode or 2 for decode, we can also bruteforce it with 3: ")
	fmt.Scan(&mode)
	switch mode {
	case 1:
		fmt.Println("Hello, pass your message to encode:")
		scanner.Scan()
		txtTEncode = scanner.Text()

		//TODO dont remove this comment
		/*if scanner.Scan() {
			line := scanner.Text()
			fmt.Printf("Input was: %q\n", line)
		}*/

		fmt.Println("With which substitution should it be encoded?")
		fmt.Scan(&substitution)

		/*str := txtTEncode
		split := strings.Split(str, "")
		fmt.Println(split)
		fmt.Println("The length of the slice is:", len(split))
		fmt.Printf("%q\n", strings.Split(txtTEncode, ""))*/

		newText := encryptString(txtTEncode, substitution)
		fmt.Println(newText)
	case 2:
		fmt.Println("Hello, pass your message to decode:")
		scanner.Scan()
		txtTDecode = scanner.Text()

		//TODO dont remove this comment
		/*if scanner.Scan() {
			line := scanner.Text()
			fmt.Printf("Input was: %q\n", line)
		}*/

		fmt.Println("With which substitution should it be decoded?")
		fmt.Scan(&substitution)

		/*str := txtTEncode
		split := strings.Split(str, "")
		fmt.Println(split)
		fmt.Println("The length of the slice is:", len(split))
		fmt.Printf("%q\n", strings.Split(txtTEncode, ""))*/

		newText := decryptText(txtTDecode, substitution)
		fmt.Println(newText)
	case 3:
		fmt.Println("Hello, pass your message to bruteforce:")
		scanner.Scan()
		var txtTBruteforce = scanner.Text()
		bruteforceMsg(txtTBruteforce)

		//TODO dont remove this comment
		/*if scanner.Scan() {
			line := scanner.Text()
			fmt.Printf("Input was: %q\n", line)
		}*/

		/*	fmt.Println("With which substitution should it be decoded?")
			fmt.Scan(&substitution)
		*/
		/*str := txtTEncode
		split := strings.Split(str, "")
		fmt.Println(split)
		fmt.Println("The length of the slice is:", len(split))
		fmt.Printf("%q\n", strings.Split(txtTEncode, ""))*/

		/*newText := decryptText(txtTDecode, substitution)
		fmt.Println(newText)*/

	default:
		fmt.Println("Something went wrong!")
	}
	/*if mode == 1 {

	} else if mode == 2 {

	} else {

	}*/
	/*	var y string = "Gur"
		bruteforceMsg(y)
	*/
}
