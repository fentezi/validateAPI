package validateAPI

import (
	"fmt"
)

type PhoneValidate struct {
	Phone    string  `json:"phone"`
	Valid    bool    `json:"valid"`
	Format   Format  `json:"format"`
	Country  Country `json:"country"`
	Location string  `json:"location"`
	Type     string  `json:"type"`
	Carrier  string  `json:"carrier"`
}

type Format struct {
	International string `json:"international"`
	Local         string `json:"local"`
}

type Country struct {
	Name   string `json:"name"`
	Code   string `json:"code"`
	Prefix string `json:"prefix"`
}

func (p PhoneValidate) Info() string {
	if !p.Valid {
		return "Phone not valid"
	}

	return fmt.Sprintf(
		"Phone: %s\nFormat: %s | %s \nCountry: %s | %s | %s \nLocation: %s\nType: %s\nCarrier: %s",
		p.Phone, p.Format.International, p.Format.Local, p.Country.Name, p.Country.Code, p.Country.Prefix, p.Location, p.Type, p.Carrier,
	)
}
