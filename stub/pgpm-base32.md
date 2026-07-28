## Usage

Sources:

- [Official upstream README](https://github.com/constructive-io/pgpm-modules/blob/45d62108f508af564b72bd6512e86a5ec7cdc923/packages/base32/README.md)
- [Official extension control file (pgpm-base32.control)](https://github.com/constructive-io/pgpm-modules/blob/45d62108f508af564b72bd6512e86a5ec7cdc923/packages/base32/pgpm-base32.control)
- [Official extension SQL (pgpm-base32--0.15.5.sql)](https://github.com/constructive-io/pgpm-modules/blob/45d62108f508af564b72bd6512e86a5ec7cdc923/packages/base32/sql/pgpm-base32--0.15.5.sql)

`pgpm-base32` — @pgpm/base32 implements Base32 encoding and decoding entirely in PostgreSQL using plpgsql. Use it for the corresponding SQL or database utility workflow. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION "pgpm-base32";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `base32.base32_alphabet(input int)` is an extension function and returns `char`.
- `base32.base32_alphabet_to_decimal(input text)` is an extension function and returns `text`.
- `base32.base32_alphabet_to_decimal_int(input text)` is an extension function and returns `int`.
- `base32.base32_to_decimal(input text)` is an extension function and returns `text[]`.
- `base32.binary_to_int(input text)` is an extension function and returns `int`.
- `base32.decimal_to_chunks(input text[])` is an extension function and returns `text[]`.
- `base32.decode(input text)` is an extension function and returns `text`.
- `base32.encode(input text)` is an extension function and returns `text`.
- `base32.fill_chunks(input text[])` is an extension function and returns `text[]`.
- `base32.string_nchars(text, int)` is an extension function and returns `text[]`.
- `base32.to_ascii(input text)` is an extension function and returns `int[]`.
- `base32.to_base32(input text[])` is an extension function and returns `text`.
- `base32.to_binary(input int)` is an extension function and returns `text`.
- `base32.to_binary(input int[])` is an extension function and returns `text[]`.

### Requirements and Caveats

- The reviewed control file declares default version `0.15.5`.
- Install the confirmed extension dependencies first: `pgcrypto`, `plpgsql`, `pgpm-verify`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
