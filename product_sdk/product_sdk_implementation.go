package productsdk

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func (ps *ProductSdk) GetProductById(id string, token string) (*GetProductByIdResponse, error) {
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

	return GetProductByIdResponseFromJson(res.Body)
}

// returns true if update request is successfull
func (ps *ProductSdk) UpdateProduct(id string, token string, body UpdateProductRequest) (bool, error) {
	url := ps.BaseUrl + "/" + id

	request, err := http.NewRequest("PUT", url, body.ToJson())

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

func (p *ProductSdk) GetAllProducts(token string) (*ProductsResult, error) {
	url := p.BaseUrl + "/"
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)
	client := http.Client{}

	res, resErr := client.Do(req)

	if resErr != nil {
		return nil, resErr
	}

	return ProductsResultFromJson(res.Body)
}

// Returns [ProductEntity] if the product is created
func (p *ProductSdk) CreateProduct(token string, r *CreateProductRequest) (*ProductEntity, error) {
	url := p.BaseUrl + "/"
	req, err := http.NewRequest(http.MethodPost, url, r.ToJson())

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	res, resErr := client.Do(req)

	if resErr != nil {
		return nil, resErr
	}

	return CreateProductResponseFromJson(res.Body)
}

// Returns true if the product is deleted
func (p *ProductSdk) DeleteProduct(token string, pId string) (bool, error) {
	url := p.BaseUrl + "/" + pId
	log.Println("Hitting url", url, " from sdk to delete product")
	req, err := http.NewRequest(http.MethodDelete, url, nil)

	if err != nil {
		return false, err
	}

	req.Header.Add("Authorization", "Bearer "+token)
	client := http.Client{}
	_, resErr := client.Do(req)

	if resErr != nil {
		log.Fatal(resErr)
		return false, resErr
	}

	log.Println("[SDK]: Deletion call success")
	return true, nil
}
