# JTL Parse
> Stream decode JMeter JTL files and output them as XML, CSV

Parse JMeter JTL files, supporting:

- Nested samples
- JMeter custom variables
- Responses
- Assertions
- Cookies
- More...

The parser is a **stream decoder**, meaning it's safe to use
for very large files

## Usage

```bash
❯ ./jtlparse -h
Usage: jtlparse [--output OUTPUT] FILENAME

Positional arguments:
  FILENAME

Options:
  --output OUTPUT, -o OUTPUT
                         specify the output type, valid options: csv,xml [default: csv]
  --help, -h             display this help and exit
```

## Parse Validating

Output/Input XML is not in canonical format, in order to diff them,
use the included `compare.sh` script:

```bash
$ ./compare.sh data.jtl
```

## Output

Currently CSV output is not configurable, and outputs only:
- label
- timestamp
- latency

## License

MIT
