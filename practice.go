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

const NUMLETTERS = 4
const NUMPIXELS = 81

type LetterStruct struct {
	Letter     string `json:"letter"`
	Version    string `json:"version"`
	Dimensions int    `json:"dimensions"`
	Pixels     []int  `json:"pixels"`
}

func convertInput(letter []int) [NUMPIXELS]float64 {
	var newLetter [NUMPIXELS]float64
	for i := 0; i < NUMPIXELS; i++ {
		if letter[i] == 1 {
			newLetter[i] = 0.5
		} else {
			newLetter[i] = -0.5
		}
	}
	return newLetter
}

func printLetter(letter []int) {
	for i := 0; i < NUMPIXELS; i++ {
		if i%9 == 0 {
			fmt.Println()
		}
		fmt.Print(" ")
		fmt.Print(letter[i])
	}
	fmt.Println()
}

func printWeights(weights [NUMLETTERS][NUMPIXELS]float64) {
	for i := 0; i < 26; i++ {
		for j := 0; j < NUMPIXELS; j++ {
			if j%9 == 0 {
				fmt.Println()
			}
			fmt.Print(weights[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func getWinner(weights [NUMLETTERS][NUMPIXELS]float64, inputVector [NUMPIXELS]float64) int {
	winner := 0
	top := 99999.9
	for i := 0; i < NUMLETTERS; i++ {
		distance := 0.0
		for j := 0; j < NUMPIXELS; j++ {
			distance += (inputVector[j] - weights[i][j]) * (inputVector[j] - weights[i][j])
		}
		if math.Sqrt(distance) < top {
			top = math.Sqrt(distance)
			winner = i
		}
	}
	fmt.Println(winner)
	return winner
}

func updateWeights(weights [NUMLETTERS][NUMPIXELS]float64, inputVector [NUMPIXELS]int, winner int, rate float64) {
	for i := 0; i < NUMPIXELS; i++ {
		if inputVector[i] == 1 {
			weights[winner][i] = 0.5 - weights[winner][i]
		}
	}
}

func populateWeightMap(weightMap [NUMLETTERS][NUMPIXELS]float64) [NUMLETTERS][NUMPIXELS]float64 {
	r := rand.New(rand.NewSource(9))
	for i := 0; i < NUMLETTERS; i++ {
		for j := 0; j < NUMPIXELS; j++ {
			weightMap[i][j] = r.Float64()
		}
	}
	return weightMap
}

func main() {

	// Open file containing letters
	file, er := ioutil.ReadFile("newletters.json")
	if er != nil {
		fmt.Println("File IO error!")
	}

	var allLetters []LetterStruct // list of letters structs
	dec := json.NewDecoder(strings.NewReader(string(file)))
	for {
		var m LetterStruct
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		allLetters = append(allLetters[:], m)
		//fmt.Printf("%s: %d\n", m.Letter, m.Pixels)
	}

	var weightMap [NUMLETTERS][NUMPIXELS]float64 // map of weights
	weightMap = populateWeightMap(weightMap)

	getWinner(weightMap, convertInput(allLetters[0].Pixels))
}
