package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// func main() {
// 	WordbyWordScan()
// }
func main() {
	var strArr [65]string
	//var strArr []string{}

	str1 := "(bin)"

	var i int
	length := len(strArr)

	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		strArr[i] = scanner.Text()
		i++
	}

	// for i < length {
	// 	if strArr[i] == "101(bin)" {
	// 		//strArr = append(strArr[:i], strArr[i+1:]...)
	// 		//length--
	// 		fmt.Print("FOUND outfits")
	// 	} else {
	// 		i++
	// 	}
	// }
	//	x = x[:i]

	// for _, text := range strArr {
	// 	fmt.Println(text)
	// }

	// for index, text := range strArr {
	// 	//	fmt.Println(text)
	// 	fmt.Printf("%d %s\n", index, text)
	// }

	for i := 0; i < length; i++ {
		//fmt.Print(strArr[i])

		// if strArr[i] == "instead" {
		// 	fmt.Println("....INSTEAD")

		// }

		if strings.Compare(strArr[i], str1) == 0 {
			//fmt.Printf("The two strings are equal.")
			fmt.Printf("String 1: %s compares to provided string 2: %s ", strArr[i], str1)
			fmt.Printf("text before: %s", strArr[i-1])
			//fmt.Printf(strArr[i])
		} else {
			//fmt.Println("The two strings are not equal.")
			//fmt.Printf("String 1: %s does not compares to provided string 2: %s ", strArr[i], str1)

		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
