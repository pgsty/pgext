## Usage

Sources:

- [PGXN 1.0.0 README](https://pgxn.org/dist/gsscode/1.0.0/README.html)
- [gsscode control file](https://api.pgxn.org/src/gsscode/gsscode-1.0.0/gsscode.control)
- [gsscode 1.0.0 SQL definitions](https://api.pgxn.org/src/gsscode/gsscode-1.0.0/gsscode--1.0.0.sql)
- [PostgreSQL license](https://api.pgxn.org/src/gsscode/gsscode-1.0.0/LICENSE)

`gsscode` stores a nine-character UK Office for National Statistics/Government Statistical Service geography code in 32 bits. Use it when canonical GSS codes need compact storage, btree ordering, exact comparison, or indexed country/type-prefix matching; it validates the code shape, not whether a particular code exists in the ONS register.

### Core Workflow

```sql
CREATE EXTENSION gsscode;

CREATE TABLE areas (
    code gsscode PRIMARY KEY,
    label text NOT NULL
);

INSERT INTO areas VALUES ('E01000001', 'Example LSOA');

SELECT code, country(code), gss_type(code), area(code)
FROM areas
WHERE code % 'E01';
```

The `%` operator accepts a country letter, a three-character country/type prefix, a complete code, or an array of prefixes. It belongs to the `gsscode_ops` btree family, so a btree index can answer a supported prefix predicate directly. `!%` is its negation.

### Objects and Registry Data

- `gsscode` is the packed type; its text form is always the uppercase nine-character representation.
- `is_valid(text)` checks the accepted lexical form without raising an input error.
- `country(gsscode)`, `gss_type(gsscode)`, and `area(gsscode)` return the packed components.
- `description(...)` and `type_info(...)` look up the three-character geography type in `gsscode_types`.
- `isnan(gsscode)` identifies ONS reserved codes whose six-digit area component is `999999`.
- `left(gsscode, integer)` plus `~`, `~*`, `!~`, and `!~*` preserve common text-query syntax after a column conversion.

`gsscode_types` contains the small shipped registry of geography types, not the hundreds of thousands of individual area names. PostgreSQL includes this table in extension configuration dumps so locally refreshed rows survive dump and restore.

### Boundaries

Input must be exactly one letter followed by eight digits. Arithmetic packing deliberately accepts future country/type combinations, so successful parsing does not prove that ONS assigned the code. General regular expressions and `left(...)` render the value as text and do not use the packed prefix operator class; prefer `%` for indexable literal-prefix searches or create an expression index for a recurring `left(...)` predicate.

The 1.0.0 release does not publish a PostgreSQL-major support matrix. Validate the build against the target server before deployment. Refreshing `gsscode_types` is optional and belongs to the separate `gsscode_ons_refresh` extension.
