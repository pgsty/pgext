## Usage

Sources:

- [Official control file](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/iso_codes/iso_codes.control)
- [Official ISO Codes README](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/iso_codes/README.md)
- [Generated PostgreSQL 17 extension SQL](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/iso_codes/extension/usr/share/postgresql/17/extension/iso_codes--0.1.0.sql)

`iso_codes` adds the `CountryCode` PostgreSQL enum containing the extension build's two-letter country-code set. Use it when a column must be constrained to that fixed vocabulary rather than arbitrary text.

### Core Workflow

```sql
CREATE EXTENSION iso_codes;

CREATE TABLE offices (
    name text NOT NULL,
    country CountryCode NOT NULL
);

INSERT INTO offices VALUES ('New York', 'US'), ('Berlin', 'DE');
SELECT * FROM offices WHERE country = 'US'::CountryCode;
```

Values outside the enum are rejected. The generated 0.1.0 SQL includes the country-code labels compiled into that release, including `XK`.

### SQL Objects

- `CountryCode` is the only user-facing SQL object in the reviewed generated SQL.
- Enum ordering follows the declaration order in the generated artifact; do not assume that ordering represents geography or business priority.

### Operational Notes

Version 0.1.0 is non-relocatable. Its control file sets `superuser = true` and `trusted = false`, so creation requires a superuser; it declares no prerequisite extension or preload requirement. Because PostgreSQL enum membership is fixed in the installed SQL, review release diffs before relying on newly assigned or changed codes.
