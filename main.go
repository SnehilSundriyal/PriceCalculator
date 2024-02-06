package main

import (
	"fmt"  
	//"example.com/price-calculator/cmdmanager"
	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)


func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	
	for _, taxRate := range taxRates {
		//cmdm := cmdmanager.New()
		fm := filemanager.New("prices.txt", fmt.Sprintf("results_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("could not process job")
			fmt.Println(err)
		}
	}
}
