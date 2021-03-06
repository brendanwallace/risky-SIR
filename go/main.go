package main

import (
	//"errors"
	"github.com/brendanwallace/riskySIR/simulate"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"golang.org/x/exp/rand"
	"time"
)

var NFlag = flag.Int("N", 1000, "number of people in simulation")
var trialsFlag = flag.Int("T", 1000, "times to run the simulation")
var alphaRFlag = flag.Float64("ar", 0.0016,
	"infectiousness parameter of risky behavior (default 0.0008)")
var alphaCFlag = flag.Float64("ac", 0.0001,
	"infectiousness parameter of community spread (default 0.0002)")

var AFlag = flag.Float64("A", 1, "beta distribution A parameter")
var BFlag = flag.Float64("B", 1, "beta distribution B parameter")

func main() {
	flag.Parse()
	rand.Seed(uint64(time.Now().UnixNano()))

	// Set up the parameters of the simulation
	params := simulate.Parameters{
		//RiskynessDistribution: distribution,
		A: *AFlag,
		B: *BFlag,
		AlphaC: *alphaCFlag,
		AlphaR: *alphaRFlag,
		DiseaseLength: 10,
		// these get added by computeR0 function:
		R0c: -1,
		R0r: -1,
		R0: -1,
		N: *NFlag,
		Trials: *trialsFlag,
	}
	params = simulate.ComputeR0(params)
	filename := fmt.Sprintf("%v.json", params.FileDescriptionLong())
	fmt.Println("starting simulation. will save output as:")
	fmt.Println(filename)

	
	/////////////////////////////////////
	// Run the simulation
	/////////////////////////////////////
	var results simulate.Results = simulate.Run(params)

	// Output to appropriately named file
	file, jsonErr := json.MarshalIndent(results, "", "\t")
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}


	writeFileErr := ioutil.WriteFile("data/"+filename, file, 0644)
	if writeFileErr != nil {
		log.Fatal(writeFileErr)
	}
}