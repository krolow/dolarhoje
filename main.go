package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/leekchan/accounting"
)

func main() {
	doc, err := goquery.NewDocument("http://www.valor.com.br/valor-data")

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	result := doc.Find("#block-valor_data_blocks-ticker-moedas span")
	value, err := strconv.ParseFloat(strings.Replace(result.First().Text(), ",", ".", -1), 64)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	ac := accounting.Accounting{Symbol: "R$ ", Precision: 2, Thousand: ".", Decimal: ","}

	if len(os.Args) > 1 {
		reais, err := strconv.ParseFloat(os.Args[1], 64)

		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		value = value * reais
	}

	fmt.Println(ac.FormatMoney(value))
}
