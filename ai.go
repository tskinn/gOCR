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

//type WeightMap [][][]float64

type WeightMap struct {
	Map [][][]float64   `json:"weightmap"`
	Winners []string  `json:"winners"`
	NumberOfLetters int       `json:"numLetters"`
	LearningRate float64 `json:"learningRate"`
	TotalIterations int `json:"totalIterations"`
	CurrentIteration int `json:"currentIteration"`
}

type Letter struct {
	Value   string  `json:"letter"`
	Version string  `json:"version"`
	Rows    int     `json:"rows"`
	Columns int     `json:"cols"`
	Pixels  [][]int `json:"pixels"`
}

// Initialize WeightMap struct
func (weights *WeightMap) init(letters, rows, cols int) {
	weights.NumberOfLetters = letters    
	r := rand.New(rand.NewSource(999))
	weights.Map = make([][][]float64, letters)
	for i := 0; i < letters; i++ {
		(weights.Map)[i] = make([][]float64, rows)
		for j := 0; j < rows; j++ {
			(weights.Map)[i][j] = make([]float64, cols)
			for k := 0; k < cols; k++ {
				(weights.Map)[i][j][k] = r.Float64()
			}
		}
	}
}

// Calculate which output node is closest match to input
func (weights *WeightMap) getWinner(letter [][]float64) int {
	winner := 0
	top := 99999.9
	for i := range weights.Map { //i := 0; i < NUMLETTERS; i++ {
		distance := 0.0
		for j := range (weights.Map)[i] { //0; j < NUMPIXELS; j++ {
			for k := range (weights.Map)[i][j] {
				distance += (letter[j][k] - (weights.Map)[i][j][k]) * (letter[j][k] - (weights.Map)[i][j][k])
			}
		}
		if math.Sqrt(distance) < top {
			top = math.Sqrt(distance)
			winner = i
		}
	}
	//fmt.Println(winner)
	return winner
}

// get updated individual weight
func getUpdWeight(distance, dRange, lWeight int, rate, wWeight float64) float64 {
	// Most pixels will be within half of the lattice length of active pixel
	longestDist := dRange / 2
	
	// some cases there will be pixels more than half the lattice away from active pixel
	if distance >= longestDist {
		distance = longestDist
	}

	// math.Acose(0.0) ====== 1.5707963267948966
	distanceAdjustment := math.Cos(float64(distance) / ((1.5) / math.Acos(0.0)))
	// distanceAdjustment := math.Cos(float64(distance) / ((float64(longestDist) / 2.0) / math.Acos(0.0)))

	if distance >= 3 {
		distanceAdjustment = -1.0
	}
	// difference from 0.0 to weight is just the weight
	difference := wWeight
	
	// if distAdjust is negative (adjusting towards zero) or difference is between weight and 1.0
	if distanceAdjustment > 0 {
		difference = 1.0 - wWeight
	}
	
	// if wWeight + rate * difference * distanceAdjustment > 1 {
	// 	fmt.Println("distance:       ", distance)
	// 	fmt.Println("dRange:         ", dRange)
	// 	fmt.Println("rate:           ", rate)
	// 	fmt.Println("wWeight:        ", wWeight)
	// 	fmt.Println("difference:     ", difference)
	// 	fmt.Println("distanceAdjust: ", distanceAdjustment)
	// 	fmt.Println("added:       ", rate * difference * distanceAdjustment)
	// 	fmt.Println("newWeight:      ", wWeight + rate * difference * distanceAdjustment)
	// 	fmt.Println()
	// }
	
	return wWeight + (rate * difference * distanceAdjustment)
}

// Update the weightmap according to the winner of training iteration and learning rate
func (weights *WeightMap) UpdateWinner(letter [][]int, winner int) {

	// (only created to shorten lines and not really necesary)
	curIt := float64(weights.CurrentIteration)
	totIt := float64(weights.TotalIterations)

	// as iterations go by, lessen the rate of change in the updating of the weights
	curRate := weights.LearningRate * math.Exp(-(curIt) / totIt)

	for i := range letter {
		for j := range letter[i] {
			distance := distToNeigh(letter, i, j)
			weight := weights.Map[winner][i][j]
			lWeight := letter[i][j]
			weights.Map[winner][i][j] = getUpdWeight(distance, len(letter), lWeight, curRate, weight)
		}
	}
}

// Finds the distance to closest black pixel in letter
// helper function to update weightmap training iteration winner
func distToNeigh(letter [][]int, x, y int) int {
	shortestDist := 999999
	if letter[x][y] == 1 {
		return 0
	}
	for i := range letter {
		for j := range letter[i] {
			dist := 99999
			if letter[i][j] == 1 {
				dist = int(math.Sqrt(math.Pow(float64((i-x)), 2.0) + math.Pow(float64((j-y)), 2.0)))
			}
			if dist < shortestDist {
				shortestDist = dist
			}
			if shortestDist == 0 {
				return 0
			}
		}
	}
	return shortestDist
}

// Print the weightmap
func (weights *WeightMap) print() {
	for i := range weights.Map {
		for j := range (weights.Map)[i] {
			for k := range weights.Map[i][j] {
				fmt.Print(int(weights.Map[i][j][k] * 10))
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

// Get all the letters to be used for training from json file
func getLettersJSON(filename string) []Letter {
	// Open file containing letters
	file, er := ioutil.ReadFile(filename)
	if er != nil {
		fmt.Println("File IO error!")
	}
	var allLetters []Letter // list of letters structs
	dec := json.NewDecoder(strings.NewReader(string(file)))
	for {
		var m Letter
		if err := dec.Decode(&m); err == io.EOF { // decode JSON
			break
		} else if err != nil {
			log.Fatal(err)
		}
		allLetters = append(allLetters[:], m) // append new letter
	}
	return allLetters
}

// Print letters??
func print(letters []Letter) {
	for _, letter := range letters {
		letter.print()
	}
}

// Calculate number of distinct letters
func getNumberOfLetters(letters []Letter) int {
	ls := make([]string, 0)
	var isIn bool
	for i := range letters {
		isIn = false
		for j := range ls {
			if letters[i].Value == ls[j] {
				isIn = true
				break
			}
		}
		if !isIn {
			ls = append(ls, letters[i].Value)
		}
	}
	return len(ls)
}

// Print letter struct
func (letter Letter) print() {
	fmt.Println("Letter: \t", letter.Value)
	fmt.Println("Version: \t", letter.Version)
	fmt.Println("Rows: \t\t", letter.Rows)
	fmt.Println("Columns: \t", letter.Columns)
}

// converts int list of pixels to float64
func (letter Letter) convertToFloat() [][]float64 {
	newLetter := make([][]float64, letter.Rows)
	for i := range newLetter {
		newLetter[i] = make([]float64, letter.Columns)
		for j := range newLetter[i] {
			if letter.Pixels[i][j] == 1 { // convert to floats
				newLetter[i][j] = 1.0
			} else {
				newLetter[i][j] = 0.0 // probably not necesarry as is initd to 0.0
			}
		}
	}
	return newLetter
}

func main() {
	lettersJSON := getLettersJSON("singleletterset.json")
	weightMap := WeightMap{}
	weightMap.TotalIterations = 100
	weightMap.CurrentIteration = 0
	weightMap.LearningRate = 0.1
	weightMap.init(getNumberOfLetters(lettersJSON), lettersJSON[0].Rows, lettersJSON[0].Columns)
	for j := 0; j < weightMap.TotalIterations; j++ {
		for i := range lettersJSON {
			winner := weightMap.getWinner(lettersJSON[i+4 % 4].convertToFloat())
			weightMap.UpdateWinner(lettersJSON[winner].Pixels, winner)
		}
		weightMap.CurrentIteration++
	}
	weightMap.print()

	//difference := math.Abs(float64(lWeight) - wWeight)
	//distanceAdjustment := math.Cos(float64(distance) / (float64(longestDist) / math.Acos(0.0)))
	
	// distance dRange lWeight, rate, wWeight
	// we := 0.38124651
	// fmt.Println("Before: ", we)
	// fmt.Println("After:  ", getUpdWeight(0, 9, 0, 0.2, we))
	fmt.Println(math.Acos(0.0))
}
