# moneta

moneta is a CLI utility for converting foreign currencies to BYN.

Currency rates are fetched from the National Bank of the Republic of Belarus.

## Usage

You can get the usage using the `-h`/`--help` flag.

## Building

The simplest way to build moneta is to use either [Task](https://taskfile.dev/):

```
$ task
```

or GNU Make:

```
$ make
```

Alternatively, you can build the app using the native Go tooling:

```
$ go build -o moneta ./cmd/cli
```

## License

MIT
