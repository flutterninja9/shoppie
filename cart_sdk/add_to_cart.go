package cartsdk

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Items []*CartItem

func (i *Items) ToJson() (io.Reader, error) {
	bufffer := new(bytes.Buffer)

	err := json.NewEncoder(bufffer).Encode(i)
	return bufffer, err
}

func (s *CartSdk) AddToCart(i Items, userId string, token string) (*AddToCartRespose, error) {
	url := s.baseUrl + "/" + userId + "/items"
	json, err := i.ToJson()

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, json)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return AddToCartResposeFromJson(res.Body)
}
