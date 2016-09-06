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

    // SCHRITT 1 
    // Standardeingabe lesen, als JSON interpretieren und damit eine
    // Struktur fuellen, die als Grundlage fuer die nachfolgende Berechnung dient.

    data, err := ioutil.ReadAll(os.Stdin)	// alles von der Standardeingabe lesen
    if err != nil {				// und in einen String packen
	fmt.Fprintf(os.Stderr,"%v",err)
	os.Exit(1)
    }
    err = json.Unmarshal(data,&addin)		// String als JSON interpretieren und damit
    if err != nil {				// die Struktur addin fuellen
	fmt.Fprintf(os.Stderr,"%v",err)
	os.Exit(1)
    }

    // SCHRITT 2
    // Die eigentliche Berechnung durchfuehren.

    res = addin.A + addin.B			// Hier natuerlich total simpel, i. A.
						// findet hier die eigentliche Arbeit statt.

    // SCHRITT 3
    // Das Ergebnis in einen JSON-Wert umwandeln und auf die
    // Standardausgabe schreiben.

    output, err := json.Marshal(&res)		// Wir machen aus dem Resultat der Berechnung
    if err != nil {				// einen JSON-Wert als String
	fmt.Fprintf(os.Stderr,"%v",err)
	os.Exit(1)
    }
    fmt.Print(string(output))			// und schreiben diesen String auf die
                                                // Standardausgabe.
}
