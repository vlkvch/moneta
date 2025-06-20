# moneta

**English** | [Русский](README.ru.md)

moneta is a CLI utility for converting foreign currencies to BYN.

Currency rates are fetched from the National Bank of the Republic of Belarus.

## Usage

You can get the usage using the `-h`/`--help` flag.

```
Usage: moneta [option...]

Options:
  -amount float
      Set amount to convert.
  -from string
      Set currency to convert from. (default "RUB")
  -quiet
      Display less output.
```

On the first run, data is fetched from the network and then cached on your disk. The cache is considered invalid the next day.

## Building

To build moneta, you can use a build script in the `scripts` directory depending on your platform:

- `build.sh` for Unix;
- `build.bat` for Windows.

Alternatively, you can use Make:

```shell
make
```

## License

MIT
