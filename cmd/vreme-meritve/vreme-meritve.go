package main

import (
	"context"
	"log"
	"os"

	"github.com/mati5kova/vreme-meritve/vreme"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "vreme",
		Usage: "Beleženje in statistika temperaturnih meritev",
		Flags: []cli.Flag{
			&cli.Float64Flag{
				Name:  "minTemp",
				Usage: "Najnižja dovoljena temperatura",
				Value: -50,
			},
			&cli.Float64Flag{
				Name:  "maxTemp",
				Usage: "Najvišja dovoljena temperatura",
				Value: 50,
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Usage: "Najmanjše število meritev za statistiko",
				Value: 3,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			err := vreme.NastaviParametre(
				cmd.Float64("minTemp"),
				cmd.Float64("maxTemp"),
				cmd.Int("stMeritev"),
			)
			if err != nil {
				return err
			}

			vreme.DodajPostajo("MEDVODE01", "Postaja Medvode 1 - Glavna", "Medvode")
			vreme.DodajPostajo("MEDVODE02", "Postaja Medvode 2 - Pri Savi", "Medvode")
			vreme.DodajPostajo("SENTVID01", "Postaja Šentvid 1 - Pri gimnaziji Šentvid", "Šentvid")
			vreme.DodajPostajo("SENTVID02", "Postaja Šentvid 2 - Pri škofijski gimnaziji", "Šentvid")

			meritve := map[string][]float64{
				"MEDVODE01": {8.0, 13.5, 16.6},
				"MEDVODE02": {8.5, 14.5, 16.0},
				"SENTVID01": {5.5, 9.5, 14.5},
				"SENTVID02": {6.5, 10.5},
			}

			for oznaka, temperature := range meritve {
				for _, temperatura := range temperature {
					vreme.DodajMeritev(oznaka, temperatura)
				}
			}

			// Neveljavna meritev
			vreme.DodajMeritev("MEDVODE01", 90.0)

			// Neobstoječa oznaka postaje
			vreme.DodajMeritev("NEZNANA_OZNAKA", 10.0)

			// Dodajanje veljavne meritve
			vreme.DodajMeritev("MEDVODE01", 18.0)

			// Izpis meritev
			vreme.IzpisMeritev()

			// Izpis statističnega poročila
			vreme.IzpisPorocila()

			return nil
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
