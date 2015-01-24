package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"strings"
)

type WeightMap [][][]float64

type LetterStruct struct {
	Letter  string `json:"letter"`
	Version string `json:"version"`
	Rows    int    `json:"rows"`
	Columns int    `json:"cols"`
	Pixels  []int  `json:"pixels"`
}

func (weights *WeightMap) init(letters, rows, cols int) {
	r := rand.New(rand.NewSource(999))
	*weights = make([][][]float64, letters)
	for i := 0; i < letters; i++ {
		(*weights)[i] = make([][]float64, rows)
		for j := 0; j < rows; j++ {
			(*weights)[i][j] = make([]float64, cols)
			for k := 0; k < cols; k++ {
				(*weights)[i][j][k] = r.Float64()
			}
		}
	}
}

func (weights *WeightMap) getWinner(letter [][]float64) int {
	winner := 0
	top := 99999.9
	for i := range *weights { //i := 0; i < NUMLETTERS; i++ {
		distance := 0.0
		for j := range (*weights)[i] { //0; j < NUMPIXELS; j++ {
			for k := range (*weights)[i][j] {
				distance += (letter[j][k] - (*weights)[i][j][k]) * (letter[j][k] - (*weights)[i][j][k])
			}
		}
		if math.Sqrt(distance) < top {
			top = math.Sqrt(distance)
			winner = i
		}
	}
	fmt.Println(winner)
	return winner
}

func (weights *WeightMap) print() {
	for i := range *weights {
		for j := range (*weights)[i] {
			fmt.Println((*weights)[i][j])
		}
	}
}

func getLettersJSON(filename string) []LetterStruct {
	// Open file containing letters
	file, er := ioutil.ReadFile(filename)
	if er != nil {
		fmt.Println("File IO error!")
	}
	var allLetters []LetterStruct // list of letters structs
	dec := json.NewDecoder(strings.NewReader(string(file)))
	for {
		var m LetterStruct
		if err := dec.Decode(&m); err == io.EOF { // decode JSON
			break
		} else if err != nil {
			log.Fatal(err)
		}
		allLetters = append(allLetters[:], m) // append new letter
	}
	return allLetters
}

func print(letters []LetterStruct) {
	for _, letter := range letters {
		letter.print()
	}
}

func getNumberOfLetters(letters []LetterStruct) int {
	ls := make([]string, 0)
	var isIn bool
	for i := range letters {
		isIn = false
		for j := range ls {
			if letters[i].Letter == ls[j] {
				isIn = true
				break
			}
		}
		if !isIn {
			ls = append(ls, letters[i].Letter)
		}
	}
	return len(ls)
}

func (letter LetterStruct) print() {
	fmt.Println("Letter: \t", letter.Letter)
	fmt.Println("Version: \t", letter.Version)
	fmt.Println("Rows: \t\t", letter.Rows)
	fmt.Println("Columns: \t", letter.Columns)
}

// converts 1d int list of pixels to 2d float64
func (letter LetterStruct) getPixels() [][]float64 {
	newLetter := make([][]float64, letter.Rows)
	pixel := 0 // keep track of pixels in letterstruct are 1d array of ints
	for i := range newLetter {
		newLetter[i] = make([]float64, letter.Columns)
		for j := range newLetter[i] {
			if letter.Pixels[pixel] == 1 { // convert to floats
				newLetter[i][j] = 1.0
			} else {
				newLetter[i][j] = 0.0 // probably not necesarry as is initd to 0.0
			}
			pixel++
		}
	}
	return newLetter
}

func main() {
	lettersJSON := getLettersJSON("newletters.json")
	weightMap := WeightMap{}
	weightMap.init(getNumberOfLetters(lettersJSON), lettersJSON[0].Rows, lettersJSON[0].Columns)
	for i := range lettersJSON {
		weightMap.getWinner(lettersJSON[i].getPixels())
	}
}
