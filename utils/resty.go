package utils

import "github.com/go-resty/resty/v2"

var Client *resty.Client

func InitResty() {
	Client = resty.New()
}
