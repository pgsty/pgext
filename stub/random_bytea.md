## Usage

Sources:

- [Official upstream README](https://github.com/ringerc/scrapcode/blob/1d713af60d0b3957771448a685706d35779c4491/postgresql/random_bytea/README)
- [Official extension control file (random_bytea.control)](https://github.com/ringerc/scrapcode/blob/1d713af60d0b3957771448a685706d35779c4491/postgresql/random_bytea/random_bytea.control)
- [Official extension SQL (random_bytea--1.0.sql)](https://github.com/ringerc/scrapcode/blob/1d713af60d0b3957771448a685706d35779c4491/postgresql/random_bytea/random_bytea--1.0.sql)

`random_bytea` — This is a toy extension I wrote while playing with random bytea string generation in PostgreSQL. It was prompted by this Stack Overflow post:. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION random_bytea;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `random_bytea(integer)` is an extension function and returns `bytea`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
