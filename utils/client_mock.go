package utils

import (
	"github.com/go-resty/resty/v2"
)

var ClientTest *resty.Client = resty.New()
