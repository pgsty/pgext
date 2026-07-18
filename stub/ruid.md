## Usage

Sources:

- [README at the reviewed commit](https://github.com/sbertrang/pg-ruid/blob/cbf18924308bb1771c9f79bd2d1b9b0fafaba58a/README.md)
- [Extension control file](https://github.com/sbertrang/pg-ruid/blob/cbf18924308bb1771c9f79bd2d1b9b0fafaba58a/ruid.control)
- [Type, operators, indexes, and helper functions](https://github.com/sbertrang/pg-ruid/blob/cbf18924308bb1771c9f79bd2d1b9b0fafaba58a/sql/ruid.sql)
- [Repository PGXN metadata](https://github.com/sbertrang/pg-ruid/blob/cbf18924308bb1771c9f79bd2d1b9b0fafaba58a/META.json)
- [PGXN distribution page](https://pgxn.org/dist/ruid/)

`ruid` defines a Readable and Usable Identifier type with the same 16-byte internal value as PostgreSQL `uuid` but a compact, modified Base64 text representation. A typical value is 22 characters rather than the canonical 36-character UUID form. Binary assignment casts connect the two types without conversion functions, and the extension supplies default B-tree and hash operator classes.

### Generate and Convert Identifiers

The extension depends on `uuid-ossp` for UUID generation and namespace helpers.

```sql
CREATE EXTENSION "uuid-ossp";
CREATE EXTENSION ruid;

SELECT ruid_v4();
SELECT '5b0ba39c-6597-11e2-9d36-001b211595d1'::uuid::ruid;
```

Helper functions include `ruid_nil`, `ruid_v1`, `ruid_v1mc`, `ruid_v4`, `ruid_v5`, and namespace helpers `ruid_dns`, `ruid_oid`, `ruid_url`, and `ruid_x500`.

### Caveats

- The compact representation changes text interchange only; the database value still occupies 16 bytes like `uuid`. External systems must explicitly understand this modified Base64 alphabet.
- The control file reports version `0.0.4`, while the repository's PGXN `META.json` still reports `0.0.3`. The install SQL and control file are the relevant API/version evidence for this catalog entry.
- The README lists `ruid_v3`, but the reviewed install SQL does not create that function. Do not rely on it without inspecting the installed extension.
- The install script explicitly sets `search_path` to `public` even though the control file says the extension is relocatable. Test non-public-schema installation and relocation rather than assuming extension-schema isolation.
- The source and documentation date from 2013 and provide no modern PostgreSQL compatibility matrix. Test dump/restore, casts, indexes, and client decoding on the exact target release.
