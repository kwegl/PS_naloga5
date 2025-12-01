// paket redovalnica se uporablja za beleženje in izpis ocen
//
// Example usage:
//
//	var studenti = make(map[string]Student)
//	studenti["63210001"] = Student{"Ana", "Novak", []int{10, 9, 8}, 6, 1, 10}
//	studenti["63210002"] = Student{"Boris", "Kralj", []int{6, 7, 5, 8}, 6, 1, 10}
//
//	DodajOceno(studenti, "63210001", 8)
//	DodajOceno(studenti, "63210001", 10)
//
//	IzpisVsehOcen(studenti)
//	IzpisiKoncniUspeh(studenti)
package redovalnica

import "fmt"

// Student predstavlja enega študenta
type Student struct {
	ime      string
	priimek  string
	ocene    []int
	stOcen   int
	minOcena int
	maxOcena int
}

// NewStudent ustvari nov Student in ga vrne
func NewStudent(ime string, priimek string, ocene []int, stOcen int, minOcena int, maxOcena int) *Student {
	return &Student{ime, priimek, ocene, stOcen, minOcena, maxOcena}
}

// DodajOceno doda oceno ocena študentu iz seznama studenti, ki ima vpisno številko vpisnaStevilka
func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int) {
	if _, exists := studenti[vpisnaStevilka]; !exists {
		fmt.Printf("Študenta %s ni na seznamu", vpisnaStevilka)
		return
	}
	var s Student = studenti[vpisnaStevilka]
	if ocena > s.maxOcena || ocena < s.minOcena {
		return
	}
	s.ocene = append(s.ocene, ocena)
	studenti[vpisnaStevilka] = s
}

// povprecje izpiše povprečno oceno študenta, če ima ta dovolj ocen
func povprecje(studenti map[string]Student, vpisnaStevilka string) float64 {
	if _, exists := studenti[vpisnaStevilka]; !exists {
		fmt.Printf("Študenta %s ni na seznamu", vpisnaStevilka)
		return -1.0
	}
	var s Student = studenti[vpisnaStevilka]
	if len(s.ocene) < s.stOcen {
		return 0.0
	}
	var vsota int = 0
	for i := 0; i < len(s.ocene); i++ {
		vsota += s.ocene[i]
	}
	return float64(vsota) / float64(len(s.ocene))
}

// IzpisVsehOcen povrsti izpiše vse ocene od vseh študentov iz seznama
func IzpisVsehOcen(studenti map[string]Student) {
	for k, v := range studenti {
		fmt.Printf("%s - %s %s: [", k, v.ime, v.priimek)
		for i, o := range v.ocene {
			fmt.Printf("%d", o)
			if i == len(v.ocene)-1 {
				break
			}
			fmt.Printf(" ")
		}
		fmt.Printf("]\n")
	}
}

// IzpisiKoncniUspeh povrsti izpiše uspeh vsakega študenta iz seznama
func IzpisiKoncniUspeh(studenti map[string]Student) {
	for k, v := range studenti {
		pov := povprecje(studenti, k)
		fmt.Printf("%s %s: povprečna ocena %.1f -> ", v.ime, v.priimek, pov)
		switch {
		case pov >= 9:
			fmt.Println("Odličen študent!")
		case pov >= 6:
			fmt.Println("Povprečen študent")
		case pov < 6:
			fmt.Println("Neuspešen študent")
		}
	}
}
