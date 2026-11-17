// Package vreme omogoča beleženje temperaturnih meritev
// vremenskih postaj in izračun osnovne statistike.
package vreme

import (
	"fmt"
	"math"
	"sort"
)

// Postaja predstavlja vremensko postajo z njenimi meritvami.
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

var postaje = make(map[string]Postaja)

var (
	minTemp   = -50.0
	maxTemp   = 50.0
	stMeritev = 3
)

// NastaviParametre nastavi dovoljeno temperaturno območje
// in najmanjše število meritev za izračun statistike.
// Ob neveljavnih parametrih vrne napako.
func NastaviParametre(min, max float64, st int) error {
	if math.IsNaN(min) || math.IsNaN(max) ||
		math.IsInf(min, 0) || math.IsInf(max, 0) ||
		min > max {
		return fmt.Errorf("neveljavno temperaturno območje")
	}
	if st < 1 {
		return fmt.Errorf("stMeritev mora biti vsaj 1")
	}

	minTemp = min
	maxTemp = max
	stMeritev = st
	return nil
}

// DodajPostajo doda novo postajo z dano oznako, imenom in krajem.
// Če oznaka že obstaja, obstoječo postajo nadomesti z novo.
func DodajPostajo(oznaka, ime, kraj string) {
	postaje[oznaka] = Postaja{
		ime:  ime,
		kraj: kraj,
	}
}

// DodajMeritev doda temperaturo postaji z dano oznako.
// Neveljavno meritev ali neobstoječo postajo izpiše kot opozorilo.
func DodajMeritev(oznaka string, temperatura float64) {
	if math.IsNaN(temperatura) ||
		math.IsInf(temperatura, 0) ||
		temperatura < minTemp || temperatura > maxTemp {
		fmt.Printf("Temperatura %v ni na intervalu [%v, %v] °C\n",
			temperatura, minTemp, maxTemp)
		return
	}

	postaja, ok := postaje[oznaka]
	if !ok {
		fmt.Printf("Postaja z oznako %s ne obstaja\n", oznaka)
		return
	}

	postaja.meritve = append(postaja.meritve, temperatura)
	postaje[oznaka] = postaja
}

// oznakePostaj vrne urejen seznam oznak vseh postaj.
func oznakePostaj() []string {
	oznake := make([]string, 0, len(postaje))
	for oznaka := range postaje {
		oznake = append(oznake, oznaka)
	}
	sort.Strings(oznake)
	return oznake
}

// IzpisMeritev izpiše podatke postaj in njihove meritve.
func IzpisMeritev() {
	fmt.Println("MERITVE:")
	for _, oznaka := range oznakePostaj() {
		p := postaje[oznaka]
		fmt.Printf("%s - %s (%s): %v\n",
			oznaka, p.ime, p.kraj, p.meritve)
	}
}

// statistika izračuna najnižjo, najvišjo in povprečno
// temperaturo postaje. Če meritev ni dovolj, vrne ok=false.
func statistika(oznaka string) (
	najnizja, najvisja, povprecna float64, ok bool,
) {
	postaja, obstaja := postaje[oznaka]
	if !obstaja || len(postaja.meritve) < stMeritev {
		return 0, 0, 0, false
	}

	najnizja = postaja.meritve[0]
	najvisja = postaja.meritve[0]

	for i, meritev := range postaja.meritve {
		if meritev < najnizja {
			najnizja = meritev
		}
		if meritev > najvisja {
			najvisja = meritev
		}
		povprecna += (meritev - povprecna) / float64(i+1)
	}

	return najnizja, najvisja, povprecna, true
}

// IzpisPorocila izpiše statistiko vseh postaj in opis
// temperature glede na izračunano povprečje.
func IzpisPorocila() {
	for _, oznaka := range oznakePostaj() {
		postaja := postaje[oznaka]
		min, max, povprecje, ok := statistika(oznaka)

		if !ok {
			fmt.Printf("%s: premalo meritev\n", postaja.ime)
			continue
		}

		var opis string
		switch {
		case povprecje >= 25:
			opis = "Vroče"
		case povprecje >= 15:
			opis = "Toplo"
		case povprecje >= 5:
			opis = "Hladno"
		default:
			opis = "Mrzlo"
		}

		fmt.Printf(
			"%s: min %.1f °C, max %.1f °C, povprečje %.1f °C -> %s\n",
			postaja.ime, min, max, povprecje, opis,
		)
	}
}
