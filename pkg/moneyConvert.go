package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type exchangerate struct {
	Success bool `json:"success"`
	Query   struct {
		From   string
		To     string
		Amount float32
	}
	Info struct {
		Timestamp time.Duration
		Rate      float32
	}
	Date   string
	Result float32
}

func Convert(token, from, to, amount string) string {
	var exch exchangerate

	const base = "https://api.apilayer.com/exchangerates_data/convert"
	url := fmt.Sprintf("%v?to=%v&from=%v&amount=%v", base, to, from, amount)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("apikey", token)

	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	if res.Body != nil {
		defer res.Body.Close()
	}
	body, err := ioutil.ReadAll(res.Body)

	json.Unmarshal(body, &exch)

	return fmt.Sprint(exch.Result)
}
