## Usage

Sources:

- [Official upstream README](https://github.com/pyramation/totp/blob/8f724aa2ba53abe49272cbdc90c2004df848b8ea/extensions/@launchql/base32/readme.md)
- [Official extension control file (launchql-base32.control)](https://github.com/pyramation/totp/blob/8f724aa2ba53abe49272cbdc90c2004df848b8ea/extensions/@launchql/base32/launchql-base32.control)
- [Official extension SQL (launchql-base32--0.0.3.sql)](https://github.com/pyramation/totp/blob/8f724aa2ba53abe49272cbdc90c2004df848b8ea/extensions/@launchql/base32/sql/launchql-base32--0.0.3.sql)

`launchql-base32` — First you'll want to start the postgres docker (you can also just use docker-compose up -d):. Use it when implementing the corresponding security, audit, or access-control workflow. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION "launchql-base32";

select base32.encode('foo');
-- MZXW6===


select base32.decode('MZXW6===');
-- foo
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

- The reviewed control file declares default version `0.0.3`.
- Install the confirmed extension dependencies first: `pgcrypto`, `plpgsql`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
