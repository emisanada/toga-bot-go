package exchange

import (
	"fmt"
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
	Price              int
	Volume             int
	Timestamp          int
	Snapping           int
	LastKnownPrice     int
	LastKnownTimestamp int
	Change1Day         bool
	Change3Day         bool
	Change7Day         bool
	VChange1Day        bool
	VChange3Day        bool
	VChange7Day        bool
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
		FormatJson(string(data))
		return string(data)
	}
}

// FormatString change all spaces from a string to underscores
func FormatString(s string) string {
	return strings.Replace(s, " ", "_", -1)
}

// FormatJson
func FormatJson(json string) {
	var data ExchangeData
	json.Unmarshal([]byte(json), &data)
	fmt.Printf(data)
	// return data
}
