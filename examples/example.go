package main

import (
	"github.com/oguzhantasimaz/steam-web-api-go/swa"
	log "github.com/sirupsen/logrus"
)

func main() {
	client := swa.NewClient("YOUR_API_KEY")

	// Re-set api key
	client.SetAPIKey("YOUR_API_KEY2")

	// Get Item History for P250 Crimson Kimono Factory New

	ih, err := client.Items().GetItemPriceHistory("P250 | Crimson Kimono (Factory New)")
	if err != nil {
		panic(err)
	}

	log.Println(ih)

	//Get Friend list of a user

	friends, err := client.Steams().GetSteamFriends("STEAM_0:0:66078051")
	if err != nil {
		panic(err)
	}

	log.Println(friends)

	// Get Excange Rate

	er, err := client.Currencies().GetExchangeRate("USD", "TRY")
	if err != nil {
		panic(err)
	}

	log.Println(er)

	// Get Random Profile

	rp, err := client.Profiles().GetRandomProfiles(1)
	if err != nil {
		panic(err)
	}

	log.Println(rp)

	// Get someone who has high valued inventory

	hvi, err := client.Profiles().GetHighValueProfiles(1)
	if err != nil {
		panic(err)
	}

	log.Println(hvi)

	// Get currency list

	cl, err := client.Currencies().GetCurrencyList("USD")
	if err != nil {
		panic(err)
	}

	log.Println(cl)

	//Get someone's steam inventory

	inv, err := client.Inventories().GetSteamInventory("STEAM_0:0:66078051", "csgo")
	if err != nil {
		panic(err)
	}

	log.Println(inv)

}
