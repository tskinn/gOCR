package main

import (
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"net/http"
	"time"
)


var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
)


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
	weights := WeightMap{}
	weights.init(26, 9, 9)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lettersJSON := getLettersJSON("singleletterset.json")
	
	
	message := Message{}
	message.Message = "init"
	message.loadLetters(lettersJSON)
	message.init(26, 9, 9, time.Now().UTC().UnixNano())
	conn.WriteJSON(message)
	
	defer conn.Close()
	
	for {
		receiveMessage := &Message{}
		err := conn.ReadJSON(receiveMessage)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Message Received: " + receiveMessage.Message)
		if receiveMessage.Message == "start" {
			message.TotalIterations = receiveMessage.TotalIterations
			message.UpdateInterval = receiveMessage.UpdateInterval
			message.LearningRate = receiveMessage.LearningRate
			message.NeighborEffect = receiveMessage.NeighborEffect
			message.Message = "update"
			message.Letters = receiveMessage.Letters
			message.train()
			
			log.Println("Trained")
		} else if receiveMessage.Message == "continue" {
			message.UpdateInterval = receiveMessage.UpdateInterval
			message.Message = "update"
			message.train()
	
		} else if receiveMessage.Message == "reset" {
			message = Message{}
			message.Message = "init"
			message.loadLetters(lettersJSON)
			message.init(26, 9, 9, time.Now().UTC().UnixNano())		
		} else if receiveMessage.Message == "test" {
			message.Letters = receiveMessage.Letters
			message.test()
			message.Message = "results"
		}
		
		conn.WriteJSON(message)
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("web")))
	http.HandleFunc("/ws", serveWS)
	log.Println("Listening...")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
