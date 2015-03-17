package main

import (
	//"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"net/http"
	//"text/template"
)


var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
	//homeTempl = template.Must(template.ParseFiles("console.html"))
)


// type WeightMap struct {
// 	Message string `json:"message"`
// 	Map [][][]float64   `json:"weightmap"`
// 	Winners []string  `json:"winners"`
// 	Letters [][][]int `json:"letters"`
// 	NumberOfLetters int       `json:"numLetters"`
// 	LearningRate float64 `json:"learningRate"`
// 	TotalIterations int `json:"totalIterations"`
// 	CurrentIteration int `json:"currentIteration"`
// }



type WeightMap [][][]float64

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

func serveWS(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	//profile := Profile{"Alex", []string{"snowboarding", "programming"}}
	weights := WeightMap{}
	weights.init(26, 9, 9)
	//js, err := json.Marshal(weights)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lettersJSON := getLettersJSON("singleletterset.json")
	
	
	message := Message{}
	message.Message = "update"
	message.loadLetters(lettersJSON)
	message.init(26, 9, 9)
	conn.WriteJSON(message)
	
	//defer conn.Close()
	//w.Header().Set("Content-Type", "application/json")
	//w.Write(js)

	for {
		receiveMessage := &Message{}
		err := conn.ReadJSON(receiveMessage)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Message Received: " + receiveMessage.Message)
		if receiveMessage.Message == "train" {
			message.TotalIterations = receiveMessage.TotalIterations
			message.UpdateInterval = receiveMessage.UpdateInterval
			message.LearningRate = receiveMessage.LearningRate
			message.train(lettersJSON)
			conn.WriteJSON(message)
			
		}else if receiveMessage.Message == "continue" {
			message.train(lettersJSON)
			
			conn.WriteJSON(message)
		} 
		//log.Println(receiveMessage.Message)
		log.Println("Trained")
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("web")))
	//http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", serveWS)
	log.Println("Listening...")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
