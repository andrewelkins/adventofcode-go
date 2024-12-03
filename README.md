## Running Locally
### Requirements
Go 1.16+ is required because [embed][embed] is used for input files.

`brew install golang`

Use `go run main.go -part <1 or 2>` will be usable to run the actual inputs for that day.

Use `go test -run RegExpToMatchFunctionNames .` to run examples and unit tests via the `main_test.go` files.

## Scripts (used for all years but 2019)
Makefile should be fairly self-documenting. Alternatively you can run the binaries yourself via `go run` or `go build`.

`make help` prints a help message.

### Make skeleton files
```sh
for ((i=1; i<26; i++)); do
make skeleton DAY=$i YEAR=2023
done
```

Note that skeletons use [embed][embed] and __will not compile__ without an `input.txt` file located in the same folder. Input files can be made via `make input`.
```sh
make skeleton DAY=5 YEAR=2023
make input DAY=5 YEAR=2023 AOC_SESSION_COOKIE=your_cookie
```

### Fetch inputs and write to input.txt files
Requires passing your cookie from AOC from either `-cookie` flag, or `AOC_SESSION_COOKIE` env variable.
```sh
make input DAY=1 YEAR=2023
```

[embed]: https://golang.org/pkg/embed/

### Original Source
Learning Go with some help here. Alex Chao has a fanastic repo from many years of use. I've forked it to give me some assistance. [GH](https://github.com/alexchao26/advent-of-code-go/tree/main).

