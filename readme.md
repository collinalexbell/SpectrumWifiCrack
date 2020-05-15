# SpectrumWifiCrack
The spectrum wifi router's default password looks pretty predictable.
```
// <5letterWord><5letterWord><n [0:1000]>
ex := "spacesaver911"
```

This produces 33489369000 combinations. On my macbook air with 4000k/s hashrate, this would take 95 days to crack.
On a GPU, it would take significantly less time.

If I were to write all the enumerated passwords out to a file, each on a new line, the file would be 435 gigabytes in size. ~34 of those gigabytes would be newlines.


