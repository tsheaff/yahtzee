# yahtzee

Emperical Yahtzee probability calculator, written in golang

Run with `go run ./main/yahtzee.go` from terminal. This will output something like the following:

```
$ go run ./main/yahtzee.go
with 3 tries to get a yahtzee with 5 dice, we have a yahtzee rate of 4.6118% (sample size 1000000)
```

Modify these constants at the top of `yahtzee.go` to get different scenarios:

```
const NumbersPerDie = 6
const NumDiePerSet = 5
const NumTurnsForYahtzee = 3
```

See [here](http://datagenetics.com/blog/january42012/index.html) for a nice article going over the deeper math behind this.
