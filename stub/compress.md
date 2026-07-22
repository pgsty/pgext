## Usage

Sources:

- [Official README](https://github.com/snaiffer/pg_compress/blob/4d254039d5bd36f4a8c8bc9add225cf4f0c9b92e/README.md)
- [Extension SQL](https://github.com/snaiffer/pg_compress/blob/4d254039d5bd36f4a8c8bc9add225cf4f0c9b92e/compress--0.0.2.sql)
- [Archive implementation](https://github.com/snaiffer/pg_compress/blob/4d254039d5bd36f4a8c8bc9add225cf4f0c9b92e/compress_cpp.cpp)

`compress` version `0.0.2` converts a one-file ZIP archive between Base64 text and plain text. It is useful only when a text-oriented SQL interface is appropriate; it is not a general archive-management API.

### Core Workflow

```sql
CREATE EXTENSION compress;

SELECT zip_base64('message.txt', 'hello from PostgreSQL');

SELECT unzip_base64(
  'UEsDB...base64-encoded-zip...'
);
```

`zip_base64(filename, data)` creates a ZIP containing one named file and returns its Base64 representation. `unzip_base64(archive)` decodes the Base64 input, opens the ZIP, and returns the contents of entry zero as text.

### Objects

- `zip_base64(text, text) returns text`: create a single-file archive; the first argument is the stored filename and the second is its contents.
- `unzip_base64(text) returns text`: extract the first entry from a Base64-encoded ZIP archive.

Both functions are declared `IMMUTABLE` and `STRICT`.

### Operational Notes

The interface uses `text`, not `bytea`, and the implementation processes archives in backend memory. It neither exposes entry selection nor extracts multiple files. Apply input-size limits before accepting untrusted compressed data, because decompressed output may be much larger than the Base64 input. No preload or restart is required after the library is installed.
