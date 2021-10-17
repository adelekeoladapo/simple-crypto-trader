package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

//"wss://stream.binance.com:9443/ws/bnbbtc@kline_1m"

func main() {
	url := "wss://stream.binancefuture.com/ws/btcusdt@markPrice"
	log.Printf("connecting to %s", url)

	conn, res, err := websocket.DefaultDialer.Dial(url, nil)
	fmt.Println(conn)
	fmt.Println(res)
	fmt.Println(err)

	for {
		_, message, readErr := conn.ReadMessage()
		if readErr != nil {
			fmt.Println(readErr)
			return
		}
		fmt.Println(message)
	}

}


