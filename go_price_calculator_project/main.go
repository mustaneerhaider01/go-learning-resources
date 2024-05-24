package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errChans := make([]chan error, len(taxRates))

	for i, taxRate := range taxRates {
		doneChans[i] = make(chan bool)
		errChans[i] = make(chan error)
		// Swappable structs funtionality can be achieved via an interface.
		// So every struct that adheres to that interface can be used as the value.
		// Both fm and cmdm struct values, adheres to IOManager interface which is what required by
		// the NewTaxIncludedPriceJob() method as the first parameter.
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()

		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[i], errChans[i])

		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// }
	}

	for i := range taxRates {
		select {
		case err := <-errChans[i]:
			if err != nil {
				fmt.Println("Could not process job")
				fmt.Println(err)
			}
		case <-doneChans[i]:
			fmt.Println("Process done!")
		}
	}
}
