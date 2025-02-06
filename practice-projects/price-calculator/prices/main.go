package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate     float64
	InputPrices []float64
	//TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]float64)

	for _, price := range job.InputPrices {
		calculation := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = calculation
	}

	fmt.Println(result)
}

func (job *TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("prices.txt")

	if err != nil {
		fmt.Println("Could not open file!")
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Could not scan file!")
		fmt.Println(err)
		file.Close()
		return
	}

	var prices []float64 = make([]float64, len(lines))
	for index, line := range lines {
		price, err := strconv.ParseFloat(line, 64)

		if err != nil {
			fmt.Println("Converting price to float failed.")
			fmt.Println(err)
			file.Close()
			return
		}

		prices[index] = price
	}

	job.InputPrices = prices
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
	}
}
