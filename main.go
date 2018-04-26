package main

import (
	"fmt"
	"log"

	"github.com/0xdevalias/poc-go-filehippo/api"
)

func main() {
	// TODO: Need to figure how the program is generating it's accessToken.. Seems to change each time I scan
	c := api.DefaultClient("TODO")

	// TODO: I think we may need to get the program IDs we want to scan from c.ProgramDefinitions() first
	//progDefs, err := c.ProgramDefinitions()
	//if err != nil {
	//	log.Fatalf("error fetching program definitions: %v", err)
	//}
	//fmt.Sprintf("Program Definitions: %#v\n", progDefs)

	// TODO: Figure out this properly
	programs := []api.Program{
		{
			nil,
			1337,
			[]api.Files{
				{
					File:   "",
					VerPv:  "",
					VerPvr: "",
					VerFv:  "",
					VerFvr: "",
					Len:    0,
					MD5:    "",
				},
			},
		},
	}

	results, err := c.ScanResults(api.DefaultScanRequest(programs))
	if err != nil {
		log.Fatalf("error fetching scan results: %v", err)
	}

	fmt.Sprintf("Scan Results: %#v\n", results)
}
