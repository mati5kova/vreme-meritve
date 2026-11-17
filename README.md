# Vreme meritve

Go modul za beleženje temperaturnih meritev vremenskih postaj.

## Paketi

- `vreme`: upravljanje vremenskih postaj, meritev in statistike
- `cmd/vreme-meritve`: aplikacija z ukaznim vmesnikom

## Zagon

```bash
go run ./cmd/vreme-meritve
```

## Parametri

- `--minTemp`: najnižja dovoljena temperatura (privzeto -50)
- `--maxTemp`: najvišja dovoljena temperatura (privzeto 50)
- `--stMeritev`: najmanjše število meritev (privzeto 3)

## Primer

```bash
go run ./cmd/vreme-meritve --minTemp=-20 --maxTemp=40 --stMeritev=2
```

## Namestitev modula

```bash
go get github.com/mati5kova/vreme-meritve@v1.0.0
```
