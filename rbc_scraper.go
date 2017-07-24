package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	months = []string{"JAN ", "FEB ", "MAR ", "APR ", "MAY ", "JUN ", "JUL ", "AUG ", "SEP ", "OCT ", "NOV ", "DEC "}
)

func ScrapeRbcStatements() {
	files, _ := ioutil.ReadDir("./")
	csv_file, err := os.Create("transactions.csv")
	if err != nil {
		log.Fatal("Failed to create file:", err.Error())
	}
	defer csv_file.Close()

	csv_writer := csv.NewWriter(csv_file)
	defer csv_writer.Flush()

	err = csv_writer.Write([]string{"Date", "Details", "Amount"})
	if err != nil {
		log.Println("Cannot write to file:", err.Error())
	}

	for _, f := range files {
		if strings.ToLower(filepath.Ext(f.Name())) == ".pdf" {
			err := exec.Command("pdf2htmlEX", f.Name(), "tmp.html").Run()
			if err != nil {
				log.Fatal("Failed on:", f.Name(), err.Error())
			}

			html_file, err := os.Open("tmp.html")
			if err != nil {
				log.Fatal(err)
			}
			defer html_file.Close()
			doc, err := goquery.NewDocumentFromReader(html_file)
			if err != nil {
				log.Fatal(err)
			}

			// Find the review items
			doc.Find(".x1").Each(func(i int, s *goquery.Selection) {
				text := s.Text()
				if in_strings(text[:4], months) { // its a transaction
					//fmt.Printf("%s", text)
					last_space := strings.LastIndex(text, " ")
					date := text[:6]
					details := strings.TrimSpace(text[13:last_space])
					var amount string
					if strings.Contains(text[last_space:], "$") { // is full transaction
						amount = strings.TrimSpace(text[last_space:])
					} else { // get additional transaction details
						e := s
						for {
							e = e.Next()
							if e.HasClass("xd") {
								details = fmt.Sprintf("%s | %s", details, strings.TrimSpace(e.Text()))
							} else {
								amount = strings.TrimSpace(e.Text())
								break
							}
						}
					}
					err := csv_writer.Write([]string{date, details, amount})
					if err != nil {
						log.Println("Cannot write to file:", err.Error())
					}
					// log.Printf("%s,%s,%s", date, details, amount)
				}
			})

			os.Remove("tmp.html")
		}
	}
}

func in_strings(query string, strings []string) bool {
	for _, x := range strings {
		if x == query {
			return true
		}
	}
	return false
}

func main() {
	ScrapeRbcStatements()
}
