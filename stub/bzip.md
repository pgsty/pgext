


## Usage

> [bzip: Bzip compression and decompression](https://github.com/steve-chavez/pg_bzip)

### Functions

- `bzcat(data bytea) returns bytea` -- Decompress bzip2 data, similar to the `bzcat` command.
- `bzip2(data bytea, compression_level int default 9) returns bytea` -- Compress data using bzip2.

### Examples

Decompress a bzip2-compressed file:

```sql
SELECT convert_from(bzcat(pg_read_binary_file('/path/to/data.csv.bz2')), 'utf8') AS contents;
```

Compress data with bzip2:

```sql
SELECT bzip2(repeat('my text to be compressed', 1000)::bytea) AS compressed;
```

Compress with a custom compression level (1-9):

```sql
SELECT bzip2('some data'::bytea, 5);
```
