# moneta

[Русский](README.md) | **English**

moneta is a CLI utility for converting foreign currencies to belarusian rubles.

Currency rates are fetched from the National Bank of the Republic of Belarus.

## Usage

You can get the usage with the `-h`/`--help` flag.

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

To build moneta, you can use a script from the `scripts` directory, depending on your platform:

- `build.sh` on Unix-like systems,
- `build.bat` on Windows.

Alternatively, you can use Make:

```shell
make
```

## License

MIT
