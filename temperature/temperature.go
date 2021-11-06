package main

import "fmt"

type celsius float64
type kelvin float64
type fahrenheit float64

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func (c celsius) kelvin() kelvin {
	return kelvin(c - 273.15)
}

func (k kelvin) celsius() celsius {
	return celsius(k + 273.15)
}

func (k kelvin) fahrenheit() fahrenheit {
	return k.celsius().fahrenheit()
}

func (f fahrenheit) kelvin() kelvin {
	return f.celsius().kelvin()
}

func main() {
	var c celsius = 25
	var k kelvin = 450
	var f fahrenheit = 451

	fmt.Printf("%v °С is %v °K\n", c, c.kelvin())
	fmt.Printf("%v °С is %v °F\n", c, c.fahrenheit())

	fmt.Printf("%v °K is %v °C\n", k, k.celsius())
	fmt.Printf("%v °K is %v °F\n", k, k.fahrenheit())

	fmt.Printf("%v °F is %v °C\n", f, f.celsius())
	fmt.Printf("%v °F is %v °K\n", f, f.kelvin())
}
