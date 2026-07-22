## Usage

Sources:

- [Official extension control file](https://github.com/adunstan/file_fixed_length_record_fdw/blob/53f7875d71cd457ae62daf142156af927630c999/file_fixed_length_fdw.control)
- [Official extension implementation](https://github.com/adunstan/file_fixed_length_record_fdw/blob/53f7875d71cd457ae62daf142156af927630c999/file_fixed_length_fdw.c)

`file_fixed_length_fdw` exposes delimiter-free, fixed-width records as a read-only foreign table. Use it when every record has the same byte layout and returning one `text[]` value per record is acceptable.

### Core Workflow

Create the wrapper, a server, and a foreign table with exactly one `text[]` column. The lengths list determines both the number of array elements and the number of bytes consumed by each element.

```sql
CREATE EXTENSION file_fixed_length_fdw;

CREATE SERVER fixed_files
  FOREIGN DATA WRAPPER file_fixed_length_fdw;

CREATE FOREIGN TABLE fixed_records (fields text[])
  SERVER fixed_files
  OPTIONS (
    filename '/srv/import/customer.dat',
    field_lengths '8,30,10',
    trim 'true',
    record_separator 'lf',
    encoding 'UTF8'
  );

SELECT fields FROM fixed_records;
```

### Options and Objects

- `filename` is required and is read by the PostgreSQL server process.
- `field_lengths` is a comma-separated list of byte lengths; the foreign table must still have only one column.
- `trim` removes leading and trailing ASCII spaces from every extracted field.
- `record_separator` accepts `none`, `lf`, `cr`, or `crlf` and contributes zero, one, or two bytes to each physical record.
- `encoding` validates and converts field bytes to the server encoding; without it, the client encoding is used.

The extension installs the `file_fixed_length_fdw` foreign-data wrapper, its handler, and its validator. It implements scanning and rescanning only; there is no write path.

### Operational Notes

Only a superuser can set or change the foreign-table options because `filename` grants server-side file access. The implementation requires every read to return one complete fixed-size record; a short trailing record is an error. Lengths describe bytes rather than characters, so test multibyte encodings carefully. The upstream control-file comment mentions SNMP, but the implementation and repository identify this as a fixed-length file wrapper.
