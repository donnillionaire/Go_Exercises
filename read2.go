package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//var strArr [68]string
	//var strArr []string{ ... }
	strArr := [...]string{}
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
		fmt.Print(strArr[i])
		i++
	}

	for i := 0; i < length; i++ {
		//fmt.Print(strArr[i])

		// if strArr[i] == "instead" {
		// 	fmt.Println("....INSTEAD")

		// }

		if strings.Compare(strArr[i], str1) == 0 {
			//fmt.Printf("The two strings are equal.")
			fmt.Printf("String 1: %s compares to provided string 2: %s ", strArr[i], str1)
			fmt.Printf("text before: %s\n", strArr[i-1])
			dec_str := strconv.Itoa(bin2int(strArr[i-1]))
			strArr[i-1] = dec_str

			//dec, _ := strconv.ParseInt(strArr[i-1], 16, 32)
			//fmt.Print(dec)
			//strArr[i-1] = strconv.Itoa(int(dec))
			//fmt.Print(strArr[i-1])
			//conv := strArr

			//fmt.Printf(strArr[i])
			//correctedSentence := strings.Join(strArr, " ")
			fmt.Print(strArr)
		} else {
			//fmt.Println("The two strings are not equal.")
			//fmt.Printf("String 1: %s does not compares to provided string 2: %s ", strArr[i], str1)

		}

		//correctedSentence := strings.Join(strArr[i-1], " ")

	}
	// sentence, err := ioutil.ReadFile("file.txt")
	// hex := "i"

	// myString := string(sentence)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// output := hexToDecimal(string(sentence))
	// fmt.Println(output)

	//func hexToDecimal(sentence string) string {
	//listOfWords := strings.Fields(myString)
	//fmt.Print(listOfWords)
	// for idx, _ := range listOfWords {
	// 	// if w == "(hex)" {
	// 	// 	dec, err := strconv.ParseInt(listOfWords[idx-1], 16, 32)
	// 	// 	//fmt.Print(dec)

	// 	// 	if err != nil {
	// 	// 		fmt.Errorf(err.Error())
	// 	// 	}

	// 	// 	listOfWords[idx-1] = strconv.Itoa(int(dec))
	// 	//	fmt.Print(listOfWords[idx-1])
	// 	//fmt.Print(listOfWords[idx])
	// 	//fmt.Print(w)

	// 	if listOfWords[idx] == hex {
	// 		fmt.Print("found")

	// 	}

	//correctedSentence := strings.Join(listOfWords, " ")

	//fmt.Print(correctedSentence)
	//}
}

//}

// func hexToDecimal(sentence string) string {
// 	listOfWords := strings.Fields(sentence)
// 	for idx, w := range listOfWords {
// 		if w == "1a(hex)" {
// 			dec, err := strconv.ParseInt(listOfWords[idx-1], 16, 32)
// 			//fmt.Print(dec)

// 			if err != nil {
// 				fmt.Errorf(err.Error())
// 			}

// 			listOfWords[idx-1] = strconv.Itoa(int(dec))
// 		}
// 	}

// 	correctedSentence := strings.Join(listOfWords, " ")
// 	return correctedSentence
// }

func bin2int(binStr string) int {

	// base 2 for binary
	result, _ := strconv.ParseInt(binStr, 2, 64)
	return int(result)
}
