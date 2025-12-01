package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/kwegl/PS_naloga5/redovalnica"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "redovalnica1",
		Usage: "redovalnica1 beleži ocene in jih izpisuje",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "stOcen",
				Usage: "Minimalno število ocena za povprečje",
				Value: 6,
			},
			&cli.IntFlag{
				Name:  "minOcena",
				Usage: "Minimalna ocena, ki se jo lahko vpiše",
				Value: 1,
			},
			&cli.IntFlag{
				Name:  "maxOcena",
				Usage: "Maksimalna ocena, ki se jo lahko vpiše",
				Value: 10,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			stOcen := cmd.Int("stOcen")
			minOcena := cmd.Int("minOcena")
			maxOcena := cmd.Int("maxOcena")
			return run(stOcen, minOcena, maxOcena)
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(stOcen int, minOcena int, maxOcena int) error {
	var studenti = make(map[string]redovalnica.Student)
	studenti["63210001"] = *redovalnica.NewStudent("Ana", "Novak", []int{10, 9, 8}, stOcen, minOcena, maxOcena)
	studenti["63210002"] = *redovalnica.NewStudent("Boris", "Kralj", []int{6, 7, 5, 8}, stOcen, minOcena, maxOcena)
	studenti["63210003"] = *redovalnica.NewStudent("Janez", "Novak", []int{8, 5, 3, 7}, stOcen, minOcena, maxOcena)
	redovalnica.IzpisVsehOcen(studenti)
	redovalnica.DodajOceno(studenti, "63210001", 8)
	redovalnica.DodajOceno(studenti, "63210001", 10)
	redovalnica.DodajOceno(studenti, "63210001", 9)
	redovalnica.DodajOceno(studenti, "63210003", 5)
	redovalnica.DodajOceno(studenti, "63210003", 9)
	redovalnica.DodajOceno(studenti, "63210003", 8)
	fmt.Println("----------------------------")
	redovalnica.IzpisVsehOcen(studenti)
	fmt.Println("----------------------------")
	redovalnica.IzpisiKoncniUspeh(studenti)
	return nil
}
