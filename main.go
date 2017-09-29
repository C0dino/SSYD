package main

import (
	"fmt"
	"math"
)

// func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	fmt.Fprint(w, "Welcome!\n")
// }

// func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
// }

func rBC(i float64) float64 { // Reader Board Required Calculation
	return math.Ceil(i / 2)
}
func main() {

	// router := httprouter.New()
	// router.GET("/", Index)
	// router.GET("/hello/:name", Hello)

	// log.Fatal(http.ListenAndServe(":8080", router))

	fmt.Println(`
	========================================================================================
	Welcome to the Security System Design (SSYD) Program.  Enter Readers required to begin
	========================================================================================
	`)

	var tR, tI, tO, rK4, rK8, v100, v200, v300, ep1502, mr52, mr16IN, mr16OUT, mEL, mES, mEP, syCL, hEL, hES, hEP float64
	var auxI, auxO bool

	fmt.Print("Enter Total Readers: ")
	fmt.Scan(&tR)              // Prompts user for Total Readers
	if (tR > 0) && (tR <= 4) { //Determine Access Control Hardware Needed - Readers (0-4 HID)
		rK4 = 1
		hEL = 0
		hES = 0
		if (tR > 0) && (tR <= 2) { //Determine Access Control Hardware Needed - Readers (1-2 Mercury)
			ep1502 = 1
			syCL = 1
			mES = 1

		} else if (tR > 2) && (tR <= 4) { //Determine Access Control Hardware Needed - Readers (2-4 Mercury)
			syCL = 1
			ep1502 = 1
			mr52 = 1
			mES = 1
		}
	} else if (tR > 4) && (tR <= 8) { //Determine Access Control Hardware Needed - Readers (5-8)
		rK8 = 1
		syCL = 1
		ep1502 = 1
		mr52 = rBC(tR - 2)
	} else if tR >= 9 { //Determine Access Control Hardware Needed - Readers (9+)
		rK8 = 1
		v100 = rBC(tR - 8)
		syCL = 1
		ep1502 = 1
		mr52 = rBC(tR - 2)
	}
	tHRB := (rK4 * 2) + (rK8 * 4) + v100 //Total HID Reader Boards
	tMRB := ep1502 + mr52                //Total Mercury Reader Boards

	fmt.Print("Enter Total Inputs: ")
	fmt.Scan(&tI) // Prompts user for Total inputs
	fmt.Print("Utilize Auxiliary Inputs? (True/False): ")
	fmt.Scan(&auxI) // Prompts user for True/False to Use Aux Inputs on Reader Boards

	//Input Boards Required Calculations
	if (tI > 0) && (!auxI) {
		v200 = math.Ceil(tI / 16)
		mr16IN = math.Ceil(tI / 16)
	} else if (tI > 0) && (auxI) {
		v200 = math.Ceil((tI - (tHRB * 2)) / 16)
		mr16IN = math.Ceil((tI - (tMRB * 2)) / 16)
	}

	fmt.Print("Enter Total Outputs: ")
	fmt.Scan(&tO) // Prompts user for Total inputs
	fmt.Print("Utilize Auxiliary Outputs? (True/False): ")
	fmt.Scan(&auxO) // Prompts user for True/False to Use Aux Inputs on Reader Boards

	//Output Boards Required Calculations
	if (tO > 0) && (!auxO) {
		v300 = math.Ceil(tO / 12)
		mr16OUT = math.Ceil(tO / 16)
	} else if (tO > 0) && (auxO) {
		v300 = math.Ceil((tO - (tHRB * 2)) / 12)
		mr16OUT = math.Ceil((tO - (tMRB * 2)) / 16)
	}
	// Enclosures Required Calculations - Break out into Function(?)

	tHB := v100 + v200 // Total HID Boards
	if tHB >= 1 {      // Determine Boards Remaining that need Enclosures - HID
		hBR := (tHB / 5) - math.Floor(tHB/5)
		if (hBR <= 0) || (hBR >= .7) {
			hEL = math.Ceil(tHB / 5)
			hES = 0
			hEP = hEL
		} else if (hBR <= .6) && (hBR > 0) {
			hEL = math.Floor(tHB / 5)
			hES = 1
			hEP = hES + hEL
		}
	}

	tMB := ep1502 + mr52 + mr16IN + syCL //Total Mercury Boards & Cloud Links
	if tMB > 0 {                         // Determine Boards Remaining that need enclosures - Mercury
		mBR := ((tMB) / 5) - math.Floor((tMB)/5)
		if (mBR <= 0) || (mBR >= .7) {
			mEL = math.Ceil((tMB) / 5)
			mES = 0
			mEP = mEL
		} else if (mBR <= .6) && (mBR > 0) {
			mEL = math.Floor((tMB) / 5)
			mES = 1
			mEP = mES + mEL
		}
	}

	fmt.Println(`
	========================================================================================
	HID Equipment Required
	========================================================================================
	`)
	fmt.Println("4 Reader Kits Required: ", rK4)
	fmt.Println("8 Reader Kits Required: ", rK8)
	fmt.Println("V100's Required: ", v100)
	fmt.Println("V200's Required: ", v200)
	fmt.Println("V300's Required: ", v300)
	fmt.Println("Large Enclosures Needed: ", hEL, "Small Enclosures Needed: ", hES, "Enclosure PSU's Needed: ", hEP)

	fmt.Println(`
	========================================================================================
	Mercury Equipment Required
	========================================================================================
	`)
	fmt.Println("Synergis Cloudlinks Needed: ", syCL)
	fmt.Println("EP1502 Required: ", ep1502)
	fmt.Println("mr52 Required: ", mr52)
	fmt.Println("mr16IN Required: ", mr16IN)
	fmt.Println("mr16OUT Required: ", mr16OUT)
	fmt.Println("Large Enclosures Needed: ", mEL, "Small Enclosures Needed: ", mES, "Enclosure PSU's Needed: ", mEP)
}
