package extensions

import (
	"github.com/golodash/galidator"
)

var g = galidator.G()

func Validator(s interface{}) galidator.Validator {
	return g.Validator(s)
}
