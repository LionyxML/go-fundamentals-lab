package main

import "fmt"

type distanceMiles float64
type distanceKm float64

func (miles distanceMiles) ToKm() distanceKm {
	return distanceKm(1.60934 * miles)
}

func (km distanceKm) ToMiles() distanceMiles {
	return distanceMiles(km / 1.60934)
}

func main() {
	d := distanceMiles(4.5)

	fmt.Println(d)

	fmt.Println(d.ToKm())
	fmt.Println(d.ToKm().ToMiles())
}
