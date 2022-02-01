package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Covid struct {
	ErrorInfo struct {
		ErrorFlag    string      `json:"errorFlag"`
		ErrorCode    interface{} `json:"errorCode"`
		ErrorMessage interface{} `json:"errorMessage"`
	} `json:"errorInfo"`
	ItemList []struct {
		Date      string `json:"date"`
		NameJp    string `json:"name_jp"`
		Npatients string `json:"npatients"`
	} `json:"itemList"`
}

func main() {
	url := "https://opendata.corona.go.jp/api/Covid19JapanAll?dataName=%E6%B2%96%E7%B8%84%E7%9C%8C"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data Covid

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	} else {
		decoder := json.NewDecoder(res.Body)
		err = decoder.Decode(&data)
	}

	for i := 0; i < len(data.ItemList); i++ {
		var npatients int
		if i == len(data.ItemList)-1 {
			npatients, _ = strconv.Atoi(data.ItemList[i].Npatients)
		} else {
			var today, yesterday int
			today, _ = strconv.Atoi(data.ItemList[i].Npatients)
			yesterday, _ = strconv.Atoi(data.ItemList[i+1].Npatients)
			npatients = today - yesterday
		}
		fmt.Printf("%v / %d\n", data.ItemList[i].Date, npatients)
	}
}
