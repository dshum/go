package main

import (
	"fmt"
)

type celsius float64
type fahrenheit float64
type getRowFn func(row int) (string, string)

const (
	line         = "======================="
	rowFormat    = "| %8s | %8s |\n"
	numberFormat = "%.2f"
)

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func drawTable(fromName, toName string, from, to int, getRow getRowFn) {
	fmt.Println(line)
	fmt.Printf(rowFormat, fromName, toName)
	fmt.Println(line)
	for t := from; t <= to; t += 5 {
		cell1, cell2 := getRow(t)
		fmt.Printf(rowFormat, cell1, cell2)
	}
	fmt.Println(line)
}

func ctof(t int) (string, string) {
	c := celsius(t)
	f := c.fahrenheit()
	cell1 := fmt.Sprintf(numberFormat, c)
	cell2 := fmt.Sprintf(numberFormat, f)
	return cell1, cell2
}

func ftoc(t int) (string, string) {
	f := fahrenheit(t)
	c := f.celsius()
	cell1 := fmt.Sprintf(numberFormat, f)
	cell2 := fmt.Sprintf(numberFormat, c)
	return cell1, cell2
}

func main() {
	drawTable("°C", "°F", -40, 100, ctof)
	fmt.Println()
	drawTable("°F", "°C", 40, 200, ftoc)
}
