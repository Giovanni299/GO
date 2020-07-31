package main

//request received from customer.
type body struct {
	ItemIds []string `json:"item_ids" xml:"item_ids" form:"item_ids" query:"item_ids"`
	Amount  float32  `json:"amount"`
}

//response to send to client.
type response struct {
	Items []string `json:"item_ids"`
	Total float32  `json:"total"`
}

//itemML response to API MLA.
type itemML struct {
	ID    string  `json:"id"`
	Price float32 `json:"price"`
}