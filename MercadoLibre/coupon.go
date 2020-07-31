package main

import (
	"container/list"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//URL to connect API MLA.
const URL = "https://api.mercadolibre.com/items/"

var totalAmount float32

//GetValuesCoupon get list items to buy.
func getValuesCoupon(items *body) (int, *response) {
	itemsMap := getValues(items.ItemIds)
	if len(itemsMap) == 0 {
		return http.StatusNotFound, nil
	}

	itemsList := calculate(itemsMap, items.Amount)
	finalValue := &response{
		Items: itemsList,
		Total: totalAmount,
	}

	return http.StatusOK, finalValue
}

//Get values from API to items.
func getValues(itemIds []string) map[string]float32 {
	itemsMap := make(map[string]float32)
	for _, item := range itemIds {
		if _, ok := itemsMap[item]; ok {
			continue
		}

		if ok, price := getValueToMLA(item); ok {
			itemsMap[item] = price
		}
	}

	return itemsMap
}

//Get values from API MLA.
func getValueToMLA(item string) (bool, float32) {
	resp, err := http.Get(URL + item)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Printf("The HTTP request failed, for item:  %s\n", item)
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
func calculate(itemsMap map[string]float32, amount float32) []string {
	values := make([]float32, len(itemsMap))
	i := 0
	for _, value := range itemsMap {
		values[i] = value
		i++
	}

	var finalAmount, pos = getCoupon(amount, values, len(itemsMap))
	totalAmount = finalAmount
	var totalList []string
	//Recorre la lista de valores y los busca en el diccionario para obtener la llave.
	for e := pos.Front(); e != nil; e = e.Next() {
		for key, value := range itemsMap {
			if value == values[e.Value.(int)] {
				totalList = append(totalList, key)
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

	var final float32
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
