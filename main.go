package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	token := os.Getenv("FINNHUB_TOKEN")
	w, _, err := websocket.DefaultDialer.Dial("wss://ws.finnhub.io?token="+token, nil)
	if err != nil {
		panic(err)
	}
	defer w.Close()

	symbols := []string{"AAPL", "AMZN", "BINANCE:BTCUSDT", "IC MARKETS:1"}
	for _, s := range symbols {
		msg, _ := json.Marshal(map[string]interface{}{"type": "subscribe", "symbol": s})
		w.WriteMessage(websocket.TextMessage, msg)
	}

	var msg struct {
		Type string `json:"type"`
		Data []struct {
			Symbol    string  `json:"s"`
			TimeStamp int64   `json:"t"`
			Price     float32 `json:"p"`
			Volume    float32 `json:"v"`
		} `json:"data"`
	}
	for {
		err := w.ReadJSON(&msg)
		if err != nil {
			panic(err)
		}
		for _, d := range msg.Data {
			fmt.Printf("---------\nSymbol %s \nPrice %.2f \nVolume %.2f\n", d.Symbol, d.Price, d.Volume)
		}
	}
}
