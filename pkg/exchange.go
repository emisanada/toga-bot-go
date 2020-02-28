package exchange

import (
	"io/ioutil"
	"net/http"
)

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
		return string(data)
	}
}
