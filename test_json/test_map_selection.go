package main

import (
	"github.com/davecgh/go-spew/spew"

	"encoding/json"
	"fmt"
)

type SearchResult struct {
	Date        string      `json:"date"`
	IdCompany   int         `json:"idCompany"`
	Company     string      `json:"company"`
	IdIndustry  interface{} `json:"idIndustry"`
	Industry    string      `json:"industry"`
	IdContinent interface{} `json:"idContinent"`
	Continent   string      `json:"continent"`
	IdCountry   interface{} `json:"idCountry"`
	Country     string      `json:"country"`
	IdState     interface{} `json:"idState"`
	State       string      `json:"state"`
	IdCity      interface{} `json:"idCity"`
	City        string      `json:"city"`
} //SearchResult

func main() {
	sr := SearchResult{
		Date:        "1 thang 1 1970",
		IdCompany:   1,
		Company:     "Google",
		IdIndustry:  2,
		Industry:    "Tech",
		IdContinent: 3,
		Continent:   "N America",
		IdCountry:   4,
		Country:     "American",
		IdState:     5,
		State:       "Cali",
		IdCity:      6,
		City:        "Cali_City"}

	sr_mapp := make(map[string]interface{})

	sr_mapp["abc"] = "def"
	sr_mapp["def"] = "def_1"
	sr_mapp["SearchResult"] = sr

	spew.Dump(sr)
	spew.Dump(sr_mapp)

	json_out, _ := json.Marshal(sr_mapp)
	fmt.Printf("\r\n\tJsonOutput:\r\n%s\r\n", json_out)
}
