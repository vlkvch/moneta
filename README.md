# moneta

moneta is a CLI utility for converting foreign currencies to BYN.

Currency rates are fetched from the National Bank of the Republic of Belarus.

## Usage

You can get the usage using the `-h`/`--help` flag.

```
Usage: moneta [option...]

Options:
  -amount Set the amount to convert
  -from   Set the currency to convert from (default RUB)
  -quiet  Display less output
```

On the first run, data is fetched from the network and then cached on your disk. The cache is considered invalid the next day.

## Building

The simplest way to build moneta is to use either GNU Make:

```shell
make
```

or [Task](https://taskfile.dev/):

```shell
task
```

Alternatively, you can build the app using the native Go tooling:

```shell
go build -o moneta ./cmd/cli
```

## License

MIT
