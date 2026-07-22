## Usage

Sources:

- [Official extension SQL](https://github.com/cahutton/numeric_domains/blob/50496725313939a340f0c8129d9e9f2491009e51/numeric_domains--1.0.sql)
- [Official README](https://github.com/cahutton/numeric_domains/blob/50496725313939a340f0c8129d9e9f2491009e51/README.md)
- [Official extension control file](https://github.com/cahutton/numeric_domains/blob/50496725313939a340f0c8129d9e9f2491009e51/numeric_domains.control)

`numeric_domains` installs reusable sign-constrained domains over `smallint`, `integer`, `bigint`, and `numeric`, plus bounded `latitude` and `longitude` domains. It is a compact way to put simple numeric invariants into PostgreSQL's type system.

### Core Workflow

Create the extension, then use the domains in columns or function signatures:

```sql
CREATE EXTENSION numeric_domains;

CREATE TABLE measurements (
  retry_count nonnegative_integer,
  adjustment nonzero_numeric,
  lat latitude,
  lon longitude
);

INSERT INTO measurements VALUES (0, -1.25, 39.9042, 116.4074);
```

An out-of-range non-null value fails its domain check:

```sql
INSERT INTO measurements(lat, lon) VALUES (91, 0);
```

### Installed Domains

Each of `smallint`, `integer`, `bigint`, and `numeric` has `negative_*`, `nonpositive_*`, `nonzero_*`, `nonnegative_*`, and `positive_*` variants. `latitude` accepts values from `-90` through `90`, and `longitude` accepts `-180` through `180`.

### Semantic Caveats

- None of the domains is declared `NOT NULL`. A domain `CHECK` evaluates to unknown for `NULL`, which passes; add a column-level `NOT NULL` constraint when absence is invalid.
- These are domains, not distinct arithmetic types. Expressions generally use the underlying numeric type, and assignments back to a domain recheck its constraint.
- Bounds are inclusive for `latitude` and `longitude`; the extension does not define coordinate reference systems, axis order, normalization, or geospatial operators.
- The install SQL contains copy/paste errors in several `COMMENT ON DOMAIN` targets, overwriting comments on `negative_*` domains. The checks remain correctly attached, but catalog descriptions may be misleading.
- Before adopting the global unqualified names, check for collisions with existing domains and decide which schema should own them; the extension is relocatable at installation time.
