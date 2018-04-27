package main

import (
	"fmt"
	"log"

	"github.com/0xdevalias/poc-go-filehippo/api"
)

func main() {
	c := api.DefaultClient("TODO").WithDebug()

	// TODO: I think we may need to get the program IDs we want to scan from c.ProgramDefinitions() first
	progDefs, err := c.ProgramDefinitions()
	if err != nil {
		log.Fatalf("error fetching program definitions: %v", err)
	}
	fmt.Sprintf("Program Definitions: %#v\n", progDefs)

	//// TODO: Figure out this properly
	////programs := []api.Program{}
	//programs := []api.Program{
	//	{
	//		[]string{},
	//		319,
	//		[]api.Files{
	//			{
	//				File:   "C:\\Program Files\\Java\\jre1.8.0_161\\bin\\java.exe",
	//				VerPv:  "8.0.1610.12",
	//				VerPvr: "8.0.1610.12",
	//				VerFv:  "8.0.1610.12",
	//				VerFvr: "8.0.1610.12",
	//				Len:    206912,
	//				MD5:    "A85792C102EE127B0B9845355D734F59",
	//			},
	//		},
	//	},
	//}
	//
	//results, err := c.ScanResults(api.DefaultScanRequest(programs))
	//if err != nil {
	//	log.Fatalf("error fetching scan results: %v", err)
	//}
	//
	//fmt.Sprintf("Scan Results: %#v\n", results)
}
