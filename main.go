// A very simple program to show, how to implement
// a solver for Solver4all in the programming language Go.
//
// The solver reads two numbers encoded as JSON object
// with the properties "A" and "B" and 
// returns the sum of the two numbers.
//
// input: {"A":number,"B":number}
//
// output: number
//
package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

type addinT struct {
    A float64 
    B float64 
}

func main() {
    var addin addinT
    var res   float64

    // STEP 1
    // Read standard input, interprete it as JSON and fill a Go struct (or variable)
    // that is used by the following computations.

    data, err := ioutil.ReadAll(os.Stdin)	// Read everything from standard input
    if err != nil {				// and store the input in a string.
	fmt.Fprintf(os.Stderr,"%v",err)		// If something get wrong:
	os.Exit(1)				// exit with status code <> 0.
    }
    err = json.Unmarshal(data,&addin)		// Interprete the input as JSON and fill
    if err != nil {				// a corresponding Go struct.
	fmt.Fprintf(os.Stderr,"%v",err)
	os.Exit(1)
    }

    // STEP 2
    // Do the computation.

    res = addin.A + addin.B			// Extremely simple in this case.

    // STEP 3
    // Represent the result of the computation as JSON value and
    // write the JSON value to standard output.

    output, err := json.Marshal(&res)		// json.Marshal() transforms our result in a
    if err != nil {				// byte array representing the result as
	fmt.Fprintf(os.Stderr,"%v",err)		// JSON value.
	os.Exit(1)
    }
    fmt.Print(string(output))			// We transform the byte array output into a
     						// UTF-8 encoded string and write it to
						// standard output.
}
