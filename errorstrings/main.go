package errorstring

import (
	"errors"
)

var err1 = errors.New("error.")
var err2 = errors.New("error:")
var err3 = errors.New("error!")

var err4 = errors.New("Error")  // ----min_confidence=0.6
var err4 = errors.New("Error!") // ----min_confidence=0.6

func main() {}
