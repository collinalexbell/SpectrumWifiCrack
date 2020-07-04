# SpectrumWifiCrack

## Description 

The spectrum wifi router's default password looks pretty predictable.
```
// <5letterWord><5letterWord><n [0:1000]>
ex := "spacesaver911"
```

This produces ~30000000000 combinations. 

I just got a new GPU that can hash PBKDF2 at ~215 KH/s, which would render the password in ~43 hours.

## Dependencies
- hashcat
- golang

## How to run
```
$ chmod +x ./run
$ ./run
```
