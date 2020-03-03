package exchange

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

type ExchangeData struct {
	Success bool
	Item    Item `json:"data"`
}

type Item struct {
	ItemName string
	Data     Data
}

type Data struct {
	Price              int  `json:"price"`
	Volume             int  `json:"volume"`
	Timestamp          int  `json:"timestamp"`
	Snapping           int  `json:"snapping"`
	LastKnownPrice     int  `json:"last_known_price"`
	LastKnownTimestamp int  `json:"last_known_timestamp"`
	Change1Day         bool `json:"change1day"`
	Change3Day         bool `json:"change3day"`
	Change7Day         bool `json:"change7day"`
	VChange1Day        bool `json:"vchange1day"`
	VChange3Day        bool `json:"vchange3day"`
	VChange7Day        bool `json:"vchange7day"`
}

// GetPrice returns a string json with all the information about an item
func GetPrice(item string) string {
	response, err := http.Get("https://api-global.poporing.life/get_latest_price/" + item)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Warn("Something went wrong")
		return "What?"
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		log.WithFields(log.Fields{
			"item": item,
		}).Info("Sucessfully return data for item")
		FormatJson(string(data))
		return string(data)
	}
}

// FormatString change all spaces from a string to underscores
func FormatString(s string) string {
	return strings.Replace(s, " ", "_", -1)
}

// FormatJson WIP
func FormatJson(jsonData string) {
	var data ExchangeData
	json.Unmarshal([]byte(jsonData), &data)
	// fmt.Printf("%+v\n", data)
}
