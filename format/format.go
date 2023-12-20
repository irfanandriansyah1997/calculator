package format

import (
	"fmt"

	"github.com/irfanandriansyah1997/calculator/customerror"
)

func Format(value float64, err error) {
	if err != nil {
		switch finalErr := err.(type) {
		case customerror.InvalidArgError:
			fmt.Println(finalErr.Error())
		default:
			fmt.Println(finalErr.Error())
		}

		return
	}

	fmt.Println(value)
}
