## Usage

Sources:

- [Official documentation](https://github.com/sindrip/pg_punycode/blob/67792c9a5e2fc17627532e4696c0861096f6efba/README.md)
- [Extension control file](https://github.com/sindrip/pg_punycode/blob/67792c9a5e2fc17627532e4696c0861096f6efba/crates/pg_punycode/pg_punycode.control)
- [Build manifest](https://github.com/sindrip/pg_punycode/blob/67792c9a5e2fc17627532e4696c0861096f6efba/crates/pg_punycode/Cargo.toml)

`pg_punycode` Punycode and IDNA/UTS-46 domain-name conversion functions.

### Enablement

Install the files for the intended server, then create `pg_punycode` in the target database:

```sql
CREATE EXTENSION pg_punycode;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
CREATE EXTENSION pg_punycode;

SELECT domain_to_ascii('Bücher.DE');                  -- xn--bcher-kva.de
SELECT domain_to_unicode('xn--fiqs8s');               -- 中国
SELECT punycode_encode('bücher');                     -- bcher-kva
SELECT punycode_decode('MajiKoi5-783gue6qz075azm5e'); -- MajiでKoiする5秒前
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `domain_to_ascii` | FUNCTION | Callable function from the reviewed install surface. |
| `punycode_encode` | FUNCTION | Callable function from the reviewed install surface. |
| `domain_to_unicode` | FUNCTION | Callable function from the reviewed install surface. |
| `punycode_decode` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- The verified PostgreSQL-major evidence covers 13, 14, 15, 16, 17, 18, 19; do not infer unlisted majors.
- Catalog lifecycle is preview; test upgrades, dump/restore, and server compatibility before production use.
