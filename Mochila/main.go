package main

import (
	"container/list"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

//URL to connect API MLA.
const URL = "https://api.mercadolibre.com/items/"

var totalAmount float32 = 0

type itemML struct {
	ID    string  `json:"id"`
	Price float32 `json:"price"`
}

func main() {
	t0 := time.Now()
	//TODO: Borrar estas 2 variable.
	itemIds := "['MLA710902496', 'MLA739047002', 'MLA621847666']"
	var amount float32 = 2500

	itemsMap := getValues(itemIds)
	var itemsList = calculate(itemsMap, amount)

	println(itemsList)
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
}

//Get values from API MLA.
func getValues(itemIds string) map[string]float32 {

	//TODO: Revisar como optimizar esto.
	itemIds1 := strings.Replace(itemIds, "['", "", -1)
	itemIds2 := strings.Replace(itemIds1, "']", "", -1)
	regex := regexp.MustCompile(`\', '`)
	items := regex.Split(itemIds2, -1)

	itemsMap := make(map[string]float32)
	for _, item := range items {
		if ok, price := getValueToMLA(item); ok {
			itemsMap[item] = price
		}
	}

	return itemsMap
}

func getValueToMLA(item string) (bool, float32) {
	resp, err := http.Get(URL + item)
	if err != nil || resp.StatusCode != 200 {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return false, 0
	}

	body, _ := ioutil.ReadAll(resp.Body)
	itemML2 := itemML{}
	jsonErr := json.Unmarshal(body, &itemML2)
	if jsonErr != nil {
		return false, 0
	}

	return true, itemML2.Price
}

//Calculate the list of items.
func calculate(itemsMap map[string]float32, amount float32) *list.List {
	values := make([]float32, len(itemsMap))
	i := 0
	for _, value := range itemsMap {
		values[i] = value
		i++
	}

	var finalAmount, pos = getCoupon(amount, values, len(itemsMap))
	totalAmount = finalAmount
	totalList := list.New()
	//Recorre la lista de valores y los busca en el diccionario para obtener la llave.
	for e := pos.Front(); e != nil; e = e.Next() {
		fmt.Println(values[e.Value.(int)])
		for key, value := range itemsMap {
			if value == values[e.Value.(int)] {
				totalList.PushBack(key)
				break
			}
		}
	}

	return totalList
}

func getCoupon(amount float32, values []float32, n int) (float32, *list.List) {
	var pos = list.New()
	if n == 0 || amount == 0 {
		return 0, pos
	}

	if values[n-1] > amount {
		return getCoupon(amount, values, n-1)
	}

	var final float32 = 0
	var val1, pos1 = getCoupon(amount-values[n-1], values, n-1)
	var val2, pos2 = getCoupon(amount, values, n-1)
	if values[n-1]+val1 >= val2 {
		final = values[n-1] + val1
		pos1.PushBack(n - 1)
		pos = pos1
	} else {
		final = val2
		pos = pos2
	}

	return final, pos
}
