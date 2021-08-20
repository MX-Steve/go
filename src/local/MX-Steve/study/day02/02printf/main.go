package main

import (
	"fmt"
)

func main() {
	// a := true
	// b := false
	// fmt.Printf("%t\n",a)
	// fmt.Printf("%t\n",b)

	// c := "abc"
	// fmt.Printf("%p\n",&c)
	// fmt.Printf("%#p\n",&c)

	// d := 15
	// e := -10
	// fmt.Printf("%d\n",d)
	// fmt.Printf("%d\n",e)
	// fmt.Printf("%+d\n",d)
	// fmt.Printf("%+d\n",e)
	// fmt.Printf("|%5d|\n",d)
	// fmt.Printf("|%-5d|\n",d)
	// fmt.Printf("|%05d|\n",d)

	// f := 3.1415926535
	// fmt.Printf("%f\n",f)
	// fmt.Printf("%e\n",f)
	// fmt.Printf("%.3g\n",f)
	// fmt.Printf("%.3f\n",f)

	g := "abc\""
	fmt.Printf("|%5s|\n", g)
	fmt.Printf("|%-5s|\n", g)
	fmt.Printf("%q\n", g)
	fmt.Printf("%#q\n", g)
}
