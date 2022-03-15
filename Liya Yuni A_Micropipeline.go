package main

import (
	"fmt"
)

func XOR(a chan byte, b chan byte, c chan byte) {
	for {
		a1 := <-a
		a1 = a1 & 1
		fmt.Println("a  =", a1)
		b1 := <-b
		fmt.Println("b  =", b1)
		b1 = b1 & 1
		c1 := a1 ^ b1
		c <- c1
	} //end for
} //end xor

func Cmuller(d chan byte, e chan byte, f chan byte) {
	for {
		d1 := <-d
		d1 = d1 ^ 0
		fmt.Println("d  =", d1)
		e1 := <-e
		e1 = e1 ^ 0
		fmt.Println("e  =", e1)
		f1 := d1 & e1
		f <- f1
	} //end for
} //end Cmuller

func TOGGLE(in chan bool, out chan bool) {
	var PREV = true
	for {
		in1 := <-in
		fmt.Println("in    =", in1)
		fmt.Println("prev  =", PREV)
		if in1 == false {
		} else {
			PREV = !PREV
		}
		out <- PREV
	} //end for
} //end func toggle

func SELECT(g chan int, h chan int, i chan int, k chan int) {
	var g1, h1 int
	for {
		g1 = <-g
		fmt.Println("g  =", g1)
		h1 = <-h
		fmt.Println("h  =", h1)
		if g1 == 1 {
			k <- h1
			fmt.Println("Jalur k")
		} else {
			i <- h1
			fmt.Println("Jalur i")
		} //end else
	} //end for
} // end func select

func ARBITER(req1 chan int, req2 chan int, gnt1 chan int, gnt2 chan int) {
	var state = 0
	var r1, r2 int
	for {
		r1 = <-req1
		fmt.Println("req1 =", r1)
		r2 = <-req2
		fmt.Println("req2 =", r2)
		if r1 == state && r2 != state {
			fmt.Println("req1 diterima")
			gnt1 <- r1
			fmt.Println("req2 diterima")
			gnt2 <- r2
		} else {
			fmt.Println("req1 diterima")
			gnt2 <- r1
			fmt.Println("req2 diterima")
			gnt1 <- r2
		} //end else
	} //end for
} //end func arbitter

func CALL(m1 chan int, m2 chan int, outm1 chan int, outm2 chan int) {
	var kondisi = 0
	var m11, m22 int
	for {
		m11 = <-m1
		fmt.Println("req1  =", m11)
		m22 = <-m2
		fmt.Println("req2  =", m22)
		if m11 == kondisi {
			outm1 <- m11
			fmt.Println("Diterima req1")
		} else {
			outm2 <- m22
			fmt.Println("Diterima req2")
		} //end else
	} //end for
} // end func call

func main() {
	a := make(chan byte) //XOR
	b := make(chan byte)
	c := make(chan byte)
	d := make(chan byte) //CMUller
	e := make(chan byte)
	f := make(chan byte)
	in := make(chan bool) //Toggle
	out := make(chan bool)
	g := make(chan int) //Select
	h := make(chan int)
	i := make(chan int)
	k := make(chan int)
	req1 := make(chan int) //Arbiter
	req2 := make(chan int)
	gnt1 := make(chan int)
	gnt2 := make(chan int)
	m1 := make(chan int) //Arbiter
	m2 := make(chan int)
	outm1 := make(chan int)
	outm2 := make(chan int)

	go XOR(a, b, c)
	fmt.Println("\nXOR1")
	a <- 0
	b <- 0
	hasilXOR := <-c
	fmt.Println("Hasil XOR1 =", hasilXOR)
	fmt.Println("XOR2")
	a <- 0
	b <- 1
	hasilXOR = <-c
	fmt.Println("Hasil XOR2 =", hasilXOR)
	fmt.Println("XOR3")
	a <- 1
	b <- 0
	hasilXOR = <-c
	fmt.Println("Hasil XOR3 =", hasilXOR)
	fmt.Println("XOR4")
	a <- 1
	b <- 1
	hasilXOR = <-c
	fmt.Println("Hasil XOR4 =", hasilXOR)

	go Cmuller(d, e, f)
	fmt.Println("\nCmuller1")
	d <- 0
	e <- 0
	hasilCmuller := <-f
	fmt.Println("Hasil Cmuller1 =", hasilCmuller)
	fmt.Println("Cmuller2")
	d <- 0
	e <- 1
	hasilCmuller = <-f
	fmt.Println("Hasil Cmuller2 =", hasilCmuller)
	fmt.Println("Cmuller3")
	d <- 1
	e <- 0
	hasilCmuller = <-f
	fmt.Println("Hasil Cmuller3 =", hasilCmuller)
	fmt.Println("Cmuller4")
	d <- 1
	e <- 1
	hasilCmuller = <-f
	fmt.Println("HasilCmuller4 =", hasilCmuller)

	go TOGGLE(in, out)
	fmt.Println("\nTOGGLE1")
	in <- false
	hasilTOGGLE := <-out
	fmt.Printf("Hasil Toggle1= %t\n", hasilTOGGLE)
	fmt.Println("TOGGLE2")
	in <- true
	hasilTOGGLE = <-out
	fmt.Printf("Hasil Toggle2= %t\n", hasilTOGGLE)
	fmt.Println("TOGGLE3")
	in <- true
	hasilTOGGLE = <-out
	fmt.Printf("Hasil Toggle3= %t\n", hasilTOGGLE)

	go SELECT(g, h, i, k)
	fmt.Println("\nSELECT1")
	g <- 1
	h <- 1
	hasilSELECT := <-k
	fmt.Println("Hasil Select1 =", hasilSELECT)
	fmt.Println("SELECT2")
	g <- 0
	h <- 0
	hasilSELECT = <-i
	fmt.Println("Hasil Select2 =", hasilSELECT)
	fmt.Println("SELECT3")
	g <- 1
	h <- 0
	hasilSELECT = <-k
	fmt.Println("Hasil Select3 =", hasilSELECT)
	fmt.Println("SELECT4")
	g <- 0
	h <- 1
	hasilSELECT = <-i
	fmt.Println("Hasil Select4 =", hasilSELECT)

	go ARBITER(req1, req2, gnt1, gnt2)
	fmt.Println("\nARBITTER1")
	req1 <- 0
	req2 <- 1
	hasilARBITER := <-gnt1
	fmt.Println("Hasil Arbitter1=", hasilARBITER)
	hasilARBITER1 := <-gnt2
	fmt.Println("Hasil Arbitter2=", hasilARBITER1)
	fmt.Println("ARBITTER2")
	req1 <- 1
	req2 <- 0
	hasilARBITER = <-gnt2
	fmt.Println("Hasil Arbitter1=", hasilARBITER)
	hasilARBITER1 = <-gnt1
	fmt.Println("Hasil Arbitter2=", hasilARBITER1)
	fmt.Println("ARBITTER3")
	req1 <- 0
	req2 <- 0
	hasilARBITER = <-gnt2
	fmt.Println("Hasil Arbitter1=", hasilARBITER)
	hasilARBITER1 = <-gnt1
	fmt.Println("Hasil Arbitter2=", hasilARBITER1)
	fmt.Println("ARBITTER4")
	req1 <- 1
	req2 <- 1
	hasilARBITER = <-gnt2
	fmt.Println("Hasil Arbitter1=", hasilARBITER)
	hasilARBITER1 = <-gnt1
	fmt.Println("Hasil Arbitter2=", hasilARBITER1)

	go CALL(m1, m2, outm1, outm2)
	fmt.Println("\nCALL1")
	m1 <- 0
	m2 <- 1
	hasilCALL := <-outm1
	fmt.Println("Hasil CALL1=", hasilCALL)
	fmt.Println("CALL2")
	m1 <- 1
	m2 <- 1
	hasilCALL = <-outm2
	fmt.Println("Hasil CALL2=", hasilCALL)
	fmt.Println("CALL3")
	m1 <- 1
	m2 <- 0
	hasilCALL = <-outm2
	fmt.Println("Hasil CALL3=", hasilCALL)
	fmt.Println("CALL4")
	m1 <- 0
	m2 <- 0
	hasilCALL = <-outm1
	fmt.Println("Hasil CALL2=", hasilCALL)

	fmt.Println("\nTekan Enter Untuk Mengakhiri Proses")
	fmt.Scanln()
}
