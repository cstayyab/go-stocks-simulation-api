package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

const (
	HOST = "0.0.0.0"
	PORT = "9000"
	TYPE = "tcp"
)

type Stock struct {
	Time   string  `json:"time"`
	Symbol string  `json:"symbol"`
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
	Volume int     `json:"volume"`
}

type stockTicks []Stock

var stocksValue = stockTicks{}

var openConnections = []net.Conn{}

func randomStockSymbol(length int) string {
	rand.Seed(time.Now().UnixNano())

	var alphabet string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var sb strings.Builder

	l := len(alphabet)

	for i := 0; i < length; i++ {
		c := alphabet[rand.Intn(l)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func generateInitialStockTicks() {
	// generate 10 stocks with default value of 100 for open, high, low and close price with 10000 volume and current time string
	for i := 0; i < 10; i++ {
		stocksValue = append(stocksValue, Stock{
			Time:   time.Now().Format(time.RFC3339),
			Symbol: randomStockSymbol(4),
			Open:   100,
			High:   100,
			Low:    100,
			Close:  100,
			Volume: 10000,
		})
	}
}

func main() {
	// generate initial stock ticks
	generateInitialStockTicks()

	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	} else {
		log.Println("Server started on port " + PORT)
	}
	go publishTicks()
	// close listener
	defer listen.Close()
	// Publish ticks to all open connections every second
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		for _, tick := range stocksValue {
			jsonTick, _ := json.MarshalIndent(tick, "", " ")
			jsonTick = append(jsonTick, '\n')
			conn.Write(jsonTick)
		}
		openConnections = append(openConnections, conn)
	}
}
func publishTicks() {
	for {
		// iterate over all open connections
		for _, conn := range openConnections {
			// choose random stock tick
			rand.Seed(time.Now().UnixNano())
			index := rand.Intn(len(stocksValue))
			// update stock time to current
			stocksValue[index].Time = time.Now().Format(time.RFC3339)
			// find 10% of close price
			fluctuation := stocksValue[index].Close * 0.1
			// generate random number between positive and negative fluctuation of close price
			max := stocksValue[index].Close + fluctuation
			min := stocksValue[index].Close - fluctuation
			rand.Seed(time.Now().UnixNano())
			stocksValue[index].Close = (rand.Float64() * (max - min)) + min
			newClose := stocksValue[index].Close
			if newClose > stocksValue[index].High {
				stocksValue[index].High = newClose
			} else if newClose < stocksValue[index].Low {
				stocksValue[index].Low = newClose
			}
			// Add a random number between 0 and 1000 to volume
			stocksValue[index].Volume += rand.Intn(1000)
			// publish tick to all open connections as json
			jsonTick, _ := json.MarshalIndent(stocksValue[index], "", " ")
			jsonTick = append(jsonTick, '\n')
			conn.Write(jsonTick)
		}
		// sleep for 100 milliseconds
		time.Sleep(100 * time.Millisecond)
	}
}
