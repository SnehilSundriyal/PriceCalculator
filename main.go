package main

import (
	"fmt"  
	//"example.com/price-calculator/cmdmanager"
	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)


func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))
	
	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		//cmdm := cmdmanager.New()
		fm := filemanager.New("prices.txt", fmt.Sprintf("results_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index], errorChans[index])

		// if err != nil {
		// 	fmt.Println("could not process job")
		// 	fmt.Println(err)
		// }
	}

	for index := range taxRates {
		select {
		case err := <- errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <- doneChans[index]:
			fmt.Println("Done")
		}
	}
}

