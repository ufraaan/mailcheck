package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
)

func main() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.Style().Options.SeparateRows = true

	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 4, WidthMax: 40},
		{Number: 6, WidthMax: 40},
	})

	t.AppendHeader(table.Row{"Domain", "MX", "SPF", "SPF Record", "DMARC", "DMARC Record"})

	check := func(domain string) {
		r := checkDomain(domain)
		if r != nil {
			t.AppendRow(*r)
		}
	}

	if len(os.Args) > 1 {
		for _, d := range os.Args[1:] {
			check(d)
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			check(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	}

	t.Render()
}

func checkDomain(domain string) *table.Row {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s (MX): %v\n", domain, err)
		return nil
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s (TXT): %v\n", domain, err)
		return nil
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s (DMARC): %v\n", domain, err)
		return nil
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	return &table.Row{domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord}
}
