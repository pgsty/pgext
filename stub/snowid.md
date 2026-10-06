## Usage

Sources:

- [README.md](https://github.com/hillac/pg_snowid/blob/2257f3be778d044f62b8138e15e3f6ae330c3068/README.md)
- [snowid/snowid--1.1.sql](https://github.com/hillac/pg_snowid/blob/2257f3be778d044f62b8138e15e3f6ae330c3068/snowid/snowid--1.1.sql)
- [snowid/snowid.control](https://github.com/hillac/pg_snowid/blob/2257f3be778d044f62b8138e15e3f6ae330c3068/snowid/snowid.control)

`snowid` stores a 64-bit identifier while displaying UUID-shaped text. The high 12 bits encode a table identifier and the low 52 bits encode a sequence; this is not a full 128-bit UUID space.

### Core Workflow

```sql
CREATE EXTENSION snowid;
CREATE SEQUENCE demo_snowid_seq MAXVALUE 4503599627370495;
SELECT snowid.gen_snowid(1, 'demo_snowid_seq'::regclass);
```

### Operational Boundaries

The SQL installs the snowid type, comparison operators, B-tree/hash support, casts and sequence helpers under `snowid`. `snowid.pack` combines components; the two-argument `snowid.gen_snowid` consumes a chosen sequence. Keep table IDs unique within the intended domain and prevent sequence overflow; bit masking is not collision prevention.

The installation also includes DDL event-trigger machinery for the single-argument generator. Review its effects on defaults and sequences before enabling it in an existing database. Installation needs a superuser. No preload or current PostgreSQL-major support matrix is declared, and no license was found in the reviewed repository.
