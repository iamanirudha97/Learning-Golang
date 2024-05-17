package prices

import (
	"fmt"

	conversion "example.com/app/conversions"
	"example.com/app/filemanager"
)

type TaxInclusivePriceJob struct {
	TaxRate            float64
	InputPrices        []float64
	TaxInclusivePrices map[string]float64
}

func (job *TaxInclusivePriceJob) Process() {
	job.LoadData()
	result := make(map[string]float64)

	for _, price := range job.InputPrices {
		// formattedVal, err := strconv.ParseFloat(fmt.Sprintf("%.2f", price*(1+job.TaxRate)), 64)
		// if err != nil {
		// 	fmt.Printf("Error while formatting")
		// 	return
		// }
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}
	job.TaxInclusivePrices = result
	filemanager.WriteToJson(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), job)
}

func NewTaxInclusicePriceJob(taxRate float64) *TaxInclusivePriceJob {
	return &TaxInclusivePriceJob{
		// InputPrices: []float64{10, 20, 30},
		TaxRate: taxRate,
	}
}

func (job *TaxInclusivePriceJob) LoadData() {
	lines, err := filemanager.ReadLines("prices.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloat(lines)

	if err != nil {
		fmt.Println(err)
	}
	job.InputPrices = prices
}
