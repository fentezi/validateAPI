package validateAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	client *http.Client
}

func NewClient() *Client {
	return &Client{
		client: &http.Client{},
	}
}

func (c Client) GetValidatePhone(phone string) (PhoneValidate, error) {
	url := fmt.Sprintf("https://phonevalidation.abstractapi.com/v1?api_key=c762172f4d8f49d5a17b284bfd6701fd&phone=%s", phone)
	resp, err := c.client.Get(url)
	if err != nil {
		return PhoneValidate{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PhoneValidate{}, err
	}
	var phoneNumber PhoneValidate
	if err := json.Unmarshal(body, &phoneNumber); err != nil {
		return PhoneValidate{}, err
	}
	return phoneNumber, nil
}
