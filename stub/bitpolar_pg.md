## Usage

Sources:

- [Official upstream README](https://github.com/mmgehlot/bitpolar/blob/97f1885b472eb9b973713c46842c69a653fc3afd/README.md)
- [Official extension control file (bitpolar_pg.control)](https://github.com/mmgehlot/bitpolar/blob/97f1885b472eb9b973713c46842c69a653fc3afd/bitpolar-pg/bitpolar_pg.control)
- [Official implementation source](https://github.com/mmgehlot/bitpolar/blob/97f1885b472eb9b973713c46842c69a653fc3afd/bitpolar-pg/src/lib.rs)

`bitpolar_pg` — BitPolar: near-optimal vector quantization — 3-8 bit compression with zero training. 58 integrations across every major AI framework. Use it for the corresponding vector, model, or retrieval workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION bitpolar_pg;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `bitpolar_compress` is an extension function.
- `bitpolar_compression_ratio` is an extension function.
- `bitpolar_decompress` is an extension function.
- `bitpolar_inner_product` is an extension function.
- `bitpolar_version()` is an extension function.

### Requirements and Caveats

- The reviewed control file declares default version `0.3.3`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
