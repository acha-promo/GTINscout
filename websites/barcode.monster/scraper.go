package barcodemonster

import (
	"fmt"
	"strings"

	core "achapromo.com/gtinscout"
)

const (
	URL = "https://barcode.monster"
)

type (
	Scraper struct {
		HttpClient *core.HttpClient
	}

	ProductData struct {
		Class       string `json:"class"`
		Code        string `json:"code"`
		Description string `json:"description"`
		ImageURL    string `json:"image_url"`
		Size        string `json:"size"`
		Status      string `json:"status"`
	}
)

func (s *Scraper) Scrape(gtin string) ([]core.Product, error) {
	productData, err := s.fetchProductData(gtin)
	if err != nil {
		return nil, err
	}

	product := core.Product{
		Name: sanitize(productData.Description),
		GTIN: productData.Code,
		URL:  fmt.Sprintf("%s/%s", URL, gtin),
	}

	return []core.Product{product}, nil
}

func (s *Scraper) Info() core.Website {
	return core.Website{URL: URL}
}

func sanitize(name string) string {
	if i := strings.Index(name, " (from barcode.monster)"); i != -1 {
		name = name[:i]
	}
	return name
}

func (s *Scraper) fetchProductData(barcode string) (*ProductData, error) {
	url := fmt.Sprintf("%s/api/%s", URL, barcode)
	var productData ProductData
	err := s.HttpClient.GetJSON(url, &productData)
	if err != nil {
		return nil, err
	}
	return &productData, nil
}