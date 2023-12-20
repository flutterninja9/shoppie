package productsdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

func (ps *ProductSdk) GetProductById(id string, token string) (*ProductEntity, error) {
	url := ps.BaseUrl + "/" + id
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, errors.New("unable to form request")
	}

	request.Header.Add("Authorization", "Bearer "+token)
	client := http.DefaultClient
	res, reqErr := client.Do(request)

	if reqErr != nil {
		return nil, errors.New("unable to request service")
	}
	defer res.Body.Close()

	entity := new(ProductEntity)
	decodeErr := json.NewDecoder(res.Body).Decode(entity)

	if decodeErr != nil {
		return nil, decodeErr
	}

	return entity, nil
}

// / returns true if update request is successfull
func (ps *ProductSdk) UpdateProduct(id string, token string, body UpdateProductRequest) (bool, error) {
	url := ps.BaseUrl + "/" + id
	requestBody, encodeErr := json.Marshal(body)
	if encodeErr != nil {
		return false, errors.New("unable to marsahll request body")
	}

	request, err := http.NewRequest("PUT", url, bytes.NewBuffer(requestBody))

	if err != nil {
		return false, errors.New("unable to form request")
	}

	request.Header.Add("Authorization", "Bearer "+token)
	client := http.DefaultClient
	res, reqErr := client.Do(request)

	if reqErr != nil {
		return false, errors.New("unable to request service")
	}
	defer res.Body.Close()

	return true, nil
}
