package price

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadData() {
	var lines []string
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("unable to open file with ERROR: ", err)
		return
	}
	Scanner := bufio.NewScanner(file)
	for Scanner.Scan() {
		lines = append(lines, Scanner.Text())
	}
	err = Scanner.Err()
	if err != nil {
		fmt.Println("error reading file content")
		file.Close()
		return
	}
	priceArr := make([]float64, len(lines))
	for lineIndex, line := range lines {
		price, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("unable to parse float from")
			return
		}
		priceArr[lineIndex] = price
	}
	job.InputPrices = priceArr
}

func WriteJson(data interface{}) error {
	file, err := os.Create("price.json")
	if err != nil {
		return errors.New("error creating price.json file")
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("failed to convert data to json")
	}
	file.Close()
	return nil

}

func (job TaxIncludedPriceJob) Process() {
	result := make(map[string]float64)
	job.LoadData()
	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}
	job.TaxIncludedPrices = result
	err := WriteJson(job)
	if err != nil {
		return
	}
}
func NewTaxIncludedPriceJob(TaxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{

		TaxRate: TaxRate,
	}

}
