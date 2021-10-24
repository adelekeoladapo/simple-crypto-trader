package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gitlab.com/dapo/crypto-trader/db"
	"gitlab.com/dapo/crypto-trader/service"
	"gitlab.com/dapo/crypto-trader/service/impl"
	"log"
	"os"
)

//"wss://stream.binance.com:9443/ws/bnbbtc@kline_1m"

func main() {
	/* Load configuration */
	if err := godotenv.Load(); err != nil {
		log.Panic("Could not load app configuration")
	}
	/* Connect to database */
	db.InitDB(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	defer db.Db.Close()
	/* Connect to Redis Server */
	db.InitRedisDb(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"), os.Getenv("REDIS_PASSWORD"), os.Getenv("REDIS_DATABASE"))
	defer db.RedisClient.Close()



	// Test Trade Service
	var tradeService service.TradeService
	tradeService = impl.GetTradeServiceImpl()
	/*  New Trade
	newTradeRequest := dto.NewTradeRequest{
		Quantity:            11.4,
		Product:             "ETHUSDT",
		CurrentPrice:        100,
		EntryPrice:          95,
		MinimumSellingPrice: 120.5,
	}
	if res, err := tradeService.CreateTrade(newTradeRequest); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Response: %v \n", res)
	} */

	if res, e := tradeService.StartTrade(1); e != nil {
		fmt.Println(e)
	} else {
		fmt.Println("Response: ", res)
	}



	//p := processor.GetIndicatorProcessor()
	//
	//p.Process()

}


