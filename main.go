package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/fogleman/gg"
)

func getArgs() float64 {

	var side int

	flag.IntVar(&side, "s", 1000, "Size of output image")
	flag.Usage = func() {
		fmt.Printf("Usage: \n")
		fmt.Printf("./go-non-linear -2 sidelength\n")
	}
	flag.Parse() // after declaring flags we need to call it

	return float64(side)
}

func rfloat(numsize float64) float64 {
	rf := rand.Float64() * numsize
	return rf
}

func nl(side float64) {
	midpt := side / 2

	rand.Seed(time.Now().UnixNano())
	dc := gg.NewContext(int(side), int(side))
	dc.SetRGB255(0, 0, 0)
	dc.Clear()

	a := 0.3 * math.Pi
	c := math.Cos(a)
	s := math.Sin(a)
	for i := 1; i <= int(midpt); i++ {
		rcol := rand.Intn(255)
		gcol := rand.Intn(255)
		bcol := rand.Intn(255)
		dc.SetRGB255(rcol, gcol, bcol)
		r := 8 * rfloat(side) / side
		Phi := 2 * math.Pi * (rfloat(side) / side)

		x := r * math.Cos(Phi)
		y := r * math.Sin(Phi)
		for j := 1; j <= int(side); j++ {
			xn := c*x - s*y - s*math.Sin(x)
			yn := s*x + c*y + c*math.Sin(x)
			x = xn
			y = yn
			xi := int(side*x/20 + midpt)
			yi := int(side*y/20 + midpt)
			//fmt.Fprintf(os.Stdout, "%d - %d\n", xi, yi)

			dc.SetPixel(xi, yi)
		}
	}
	fmt.Fprintf(os.Stdout, "\n\nSaving file…\n")
	dc.SavePNG("out.png")
}

func main() {
	fmt.Println("Plotting…")
	side := getArgs()
	nl(side)
}
