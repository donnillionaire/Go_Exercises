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
	var old string
	var new string

	var n int

	old = "1a(hex)"
	new = "5"
	n = -1


	old_2 = "101(bin)"
	new_2 = "5"
	n_2= -1
	//var strArr []string{}
	var i int
	//length := len(strArr)

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
		//new_string := strings.Replace(str, old, new, n)

		i++
	}

	for _, text := range strArr {
		//	fmt.Println(text)
		//fmt.Printf("%d %s\n", index, text)

		//new_string := strings.Replace(text, old, new, n)

		//fmt.Println(new_string)

		if text == "101(bin)" {
			fmt.Print()
			
		}

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

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
