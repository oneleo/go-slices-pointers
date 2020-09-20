package main

import "fmt"

func main() {
	// String slice, and func call by value
	var a = []string{"1", "2", "3"}
	changeSliceA(a)

	// String slice, and func call by pointer
	b := make([]string, len(a))
	copy(b, a)
	changeSliceB(&b)

	// String pointer slice, and func call by value
	s1 := "1"
	s2 := "2"
	//s3 := "3"
	var c = []*string{&s1, &s2, new(string)}
	changeSliceC(c)

	// String pointer slice, and func call by pointer
	d := make([]*string, len(c))
	copy(d, c)
	changeSliceD(&d)

	// Result
	fmt.Print("a at ", fmt.Sprintf("%p", &a), "\ta = ", a, "\n")
	fmt.Print("b at ", fmt.Sprintf("%p", &b), "\tb = ", b, "\n")
	fmt.Print("c at ", fmt.Sprintf("%p", &c), "\tc = ", c, "\n")
	fmt.Print("d at ", fmt.Sprintf("%p", &d), "\td = ", d, "\n")
}

func changeSliceA(i []string) {
	i[0] = "3"
	i = append(i, "4")
}

func changeSliceB(i *([]string)) {
	(*i)[0] = "3"
	*i = append(*i, "4")
}

func changeSliceC(i [](*string)) {
	*(i[0]) = "3"
	s := "4"
	i = append(i, &s)
}

func changeSliceD(i *([](*string))) {
	*((*i)[0]) = "3"
	s := "4"
	*i = append(*i, &s)
}
