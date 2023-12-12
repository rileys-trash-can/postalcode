package main

import (
	gq "github.com/PuerkitoBio/goquery"
	"github.com/rileys-trash-can/postalcode"
	"gopkg.in/resty.v1"

	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: dl <address>")
	}

	url := os.Args[1]
	r, err := resty.R().Get(url)
	if err != nil {
		log.Fatalf("Failed to GET '%s': %s", url, err)
	}

	_, outfile := filepath.Split(r.Request.RawRequest.URL.Path)

	doc, err := gq.NewDocumentFromReader(bytes.NewReader(r.Body()))
	if err != nil {
		log.Fatalf("Failed to decode body: %s", err)
	}

	sel := doc.Find(".unit")
	plzes := make([]plz.PLZ, sel.Length())

	var cbuf int64
	sel.Each(func(i int, s *gq.Selection) {
		place := s.Find(".place").Text()
		scode := s.Find(".code")

		scode = scode.Find("span")
		codes := make([]int, scode.Length())
		scode.Each(func(i int, s *gq.Selection) {
			cbuf, err = strconv.ParseInt(s.Nodes[0].FirstChild.Data,
				10, 64)
			if err != nil {
				log.Fatalf("Failed to decode code '%s': %s",
					s.Nodes[0].FirstChild.Data, err)
			}

			codes[i] = int(cbuf)
		})

		plzes[i] = plz.PLZ{
			Name: place,
			Code: codes,
		}
	})

	// json part:
	encodeJSON("data/"+outfile+".json", plzes)
	encodeGO(outfile, outfile+".go", plzes)
}

func encodeJSON(path string, plz []plz.PLZ) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		log.Fatalf("Failed to open outfile '%s': %s", path, err)
	}

	defer f.Close()

	enc := json.NewEncoder(f)
	enc.Encode(plz)
}

func encodeGO(name, path string, plz []plz.PLZ) {
	name = strings.ReplaceAll(name, "-", "_")

	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		log.Fatalf("Failed to open outfile '%s': %s", path, err)
	}

	defer f.Close()

	fmt.Fprintf(f, `package plz

var %s_map = map[string]PLZ{
`, name)
	for _, plz := range plz {
		fmt.Fprintf(f, "\"%s\": %s,\n", plz.Name, goenc(plz))
	}

	fmt.Fprint(f, "}\n")
	fmt.Fprintf(f, "var %s_slice = []PLZ{\n", name)
	for _, plz := range plz {
		fmt.Fprintf(f, "%s,\n", goenc(plz))
	}

	fmt.Fprint(f, "}\n")
}

func goenc(plz plz.PLZ) string {
	return fmt.Sprintf("%#v", plz)[4:]
}
