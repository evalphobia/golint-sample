package errorf

import (
	"errors"
	"fmt"
)

func main() {
	errors.New(fmt.Sprintf("error!"))
}
