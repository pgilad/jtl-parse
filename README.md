# JTL Parse
> Stream decode JMeter JTL files and output them as XML, CSV

Parse JMeter JTL files, supporting:

- Nested samples
- JMeter custom variables
- Responses
- Assertions
- Cookies
- More...

The parser is a **stream decoder**, meaning it's safe to use for very large files

## Usage

```bash
❯ ./jtl-parse -h
Usage: jtl-parse [--output OUTPUT] FILENAME

Positional arguments:
  FILENAME

Options:
  --output OUTPUT, -o OUTPUT
                         specify the output type, valid options: csv,xml,json [default: json]
  --help, -h             display this help and exit
```

## Parse Validating

Output/Input XML is not in canonical format, in order to diff them,
use the included `compare.sh` script:

```bash
$ ./compare.sh data.jtl
```

## Output Formats

- XML

- JSON

- CSV

Sample CSV output:
```csv
Label,Timestamp,Response Time,Latency,Users
Label1,1519028940050,1205,1205,1
Label2,1519028941268,93,93,1
```

## License

MIT © [Gilad Peleg](https://www.giladpeleg.com)
