package msjwt

type Config struct {
	Audience string `json:"audience"`
	Issuer   string `json:"issuer"`
	Secret   string `json:"secret"`
}
