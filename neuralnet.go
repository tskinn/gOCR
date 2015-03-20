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

////////////////////////////////////////////////////////////////////////////////
//  MESSAGE
////////////////////////////////////////////////////////////////////////////////
type Message struct {
	Message string          `json:"message"`
	IsInitialized bool      `json:"isInitialized"`
	NeuralNet [][][]float64 `json:"neuralNet"`
	Letters [][][]int       `json:"letters"`
	Winners []string        `json:"winners"`
	NumberOfLetters int     `json:"numLetters"`
	LearningRate float64    `json:"learningRate"`
	TotalIterations int     `json:"totalIterations"`
	CurrentIteration int    `json:"currentIteration"`
	UpdateInterval int      `json:"updateInterval"`
	NeighborEffect float64  `json:"neighborEffect"`
}


////////////////////////////////////////////////////////////////////////////////
//  LETTER
////////////////////////////////////////////////////////////////////////////////
type Letter struct {
	Value   string  `json:"letter"`
	Version string  `json:"version"`
	Rows    int     `json:"rows"`
	Columns int     `json:"cols"`
	Pixels  [][]int `json:"pixels"`
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

// Convert integer pixel representation to float64
func (letter Letter) getPixelsAsFloat() [][]float64 {
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

func (message *Message) loadLetters(letters []Letter) {
	message.Letters = make([][][]int, len(letters))
	message.Winners = make([]string, len(letters))
	for i := range letters {
		message.Letters[i] = make([][]int, len(letters[i].Pixels))
		for j:= range letters[i].Pixels {
			message.Letters[i][j] = make([]int , len(letters[i].Pixels[j]))
			for k := range letters[i].Pixels[j] {
				message.Letters[i][j][k] = letters[i].Pixels[j][k]
			}
		}
		message.Winners[i] += letters[i].Value
	}
}

// Initialize WeightMap struct
func (message *Message) init(letters, rows, cols int, seed int64) {
	message.NumberOfLetters = letters    
	r := rand.New(rand.NewSource(seed))
	message.NeuralNet = make([][][]float64, letters)
	for i := 0; i < letters; i++ {
		(message.NeuralNet)[i] = make([][]float64, rows)
		for j := 0; j < rows; j++ {
			(message.NeuralNet)[i][j] = make([]float64, cols)
			for k := 0; k < cols; k++ {
				(message.NeuralNet)[i][j][k] = r.Float64()
			}
		}
	}
	message.IsInitialized = true;
}

// Calculate which output node is closest match to input
func (message *Message) getWinner(letter [][]float64) int {
	winner := 0
	top := 99999.9
	for i := range message.NeuralNet { //i := 0; i < NUMLETTERS; i++ {
		distance := 0.0
		for j := range (message.NeuralNet)[i] { //0; j < NUMPIXELS; j++ {
			for k := range (message.NeuralNet)[i][j] {
				distance += (letter[j][k] - (message.NeuralNet)[i][j][k]) * (letter[j][k] - (message.NeuralNet)[i][j][k])
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
func getUpdatedWeight(distance, dRange, lWeight int, rate, wWeight float64) float64 {
	// Most pixels will be within half of the lattice length of active pixel
	longestDist := dRange / 2
	
	// some cases there will be pixels more than half the lattice away from active pixel
	if distance >= longestDist {
		distance = longestDist
	}

	// math.Acose(0.0) ====== 1.5707963267948966
	//distanceAdjustment := math.Cos(float64(distance) / ((1.5) / math.Acos(0.0)))
	// distanceAdjustment := math.Cos(float64(distance) / ((float64(longestDist) / 2.0) / math.Acos(0.0)))
	distanceAdjustment := math.Cos(float64(distance) / ((1.001) / math.Acos(0.0)))
	
	if distance >= 2 { // was 3
		distanceAdjustment = -1.0
	}
	// difference from 0.0 to weight is just the weight
	difference := wWeight
	
	// if distAdjust is negative (adjusting towards zero) or difference is between weight and 1.0
	if distanceAdjustment > 0 {
		difference = 1.0 - wWeight
	}
	result := float64(wWeight + (rate * difference * distanceAdjustment))
	//log.Printf("Distance Adjustment:  %f\tDifference:  %f\tRate:  %f\tWeight: %f\tResult: %f", distanceAdjustment, difference, rate, result)
	//log.Printf("Weight: %f + Adjustment: %f", wWeight, (rate * difference * distanceAdjustment))
	return result
}

// Update the weightmap according to the winner of training iteration and learning rate
func (message *Message) UpdateWinner(letter [][]int, winner int) {

	// (only created to shorten lines and not really necesary)
	curIt := float64(message.CurrentIteration)
	totIt := float64(message.TotalIterations)

	// as iterations go by, lessen the rate of change in the updating of the weights
	curRate := message.LearningRate *  math.Exp(-(curIt) * 2 / totIt)
	//curRate := message.LearningRate * math.Exp(-(curIt) / totIt)

	for i := range letter {
		for j := range letter[i] {
			distance := distToNeigh(letter, i, j)
			weight := message.NeuralNet[winner][i][j]
			lWeight := letter[i][j]
			message.NeuralNet[winner][i][j] = getUpdatedWeight(distance, len(letter), lWeight, curRate, weight)
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

func (message *Message) train(lettersJSON []Letter) {
	if message.CurrentIteration >= message.TotalIterations {
		message.Message = "done"
		return
	}
	
	itersToStop := message.CurrentIteration + message.UpdateInterval
	
	if message.UpdateInterval < 1 || itersToStop > message.TotalIterations {
		itersToStop = message.TotalIterations
	}
	//log.Printf("Training from iterations %d to %d", message.CurrentIteration, itersToStop)
	for i := message.CurrentIteration; i <= itersToStop; i++ { // why less than or equal to?
		for j:= range message.Letters {
			winner := message.getWinner(lettersJSON[j].getPixelsAsFloat())
			message.UpdateWinner(lettersJSON[winner].Pixels, winner)
		}
		message.CurrentIteration++
	}
	log.Println("Successfully trained")
}


// func main () {
// 	fmt.Println("lehll")
// }
