package sexcel

import (
	"fmt"
	"testing"
)

func Test_GetAxis(t *testing.T) {
	//for i := 0; i < 56; i++ {
	//
	//}
	fmt.Println(GetAxis(0, 0))
	fmt.Println(GetAxis(25, 1))
	fmt.Println(GetAxis(26, 1))
	fmt.Println(GetAxis(27, 1))
	fmt.Println(GetAxis(26+25, 1))
	fmt.Println(GetAxis(26+26, 1))
	fmt.Println(GetAxis(26+27, 1))
	fmt.Println(GetAxis(26*26+25, 1))

	fmt.Println(GetAxis(26*26+26, 1))
	fmt.Println(GetAxis(26*26+27, 1))
	fmt.Println(GetAxis(26*26*27+25, 1))
	fmt.Println(GetAxis(26*26*27+26, 1))
	fmt.Println(GetAxis(26*26*27+27, 1))
}
