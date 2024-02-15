package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	equazione := os.Args[1] //+c+bx+ax^2=0,
	a := 0
	b := 0
	c := 0

	abc := strings.FieldsFunc(equazione, func(singolaRuna rune) bool {
		return singolaRuna == 'x' || singolaRuna == '^'
	})

	a, _ = strconv.Atoi(abc[1]) //a già definito.
	//+-c+-b +-a 2=0, divido c e b.

	cnosegno, bnosegno, found := strings.Cut(abc[0][1:], "+") //escludo il segno di c, sennò verrebbe effettuato un taglio li.

	if !found { //se tra c e b c'era un - e non un più
		cnosegno, bnosegno, _ = strings.Cut(abc[0][1:], "-")
		b, _ = strconv.Atoi("-" + bnosegno) //b è negativo.
	} else { //se tra c e b c'è un +
		b, _ = strconv.Atoi(bnosegno)
	}

	c, _ = strconv.Atoi(string(equazione[0]) + cnosegno) //considero il segno di c, che sta nel primo carattere dell'equazione.

	//calcolo soluzioni.
	delta := float64((b * b) - (4 * a * c))

	if delta < 0 {
		fmt.Println("Non ci sono soluzioni reali.")
	} else if delta == 0 {
		fmt.Println("Esiste un'unica soluzione reale: ", -b/2*a)
	} else {
		fmt.Printf("Ci sono due soluzioni reali: %.2f e %.2f\n", (float64(-b)+math.Sqrt(delta))/(2.0*float64(a)), (float64(-b)-math.Sqrt(delta))/(2.0*float64(a)))
	}

}
