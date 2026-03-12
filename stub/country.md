

## Usage

> [country: ISO 3166-1 alpha-2 country code type](https://github.com/adjust/pg-country)

The `country` extension provides an enumerated data type for ISO 3166-1 alpha-2 country codes with minimal storage overhead.

```sql
CREATE EXTENSION country;

CREATE TABLE locations (
    id    serial PRIMARY KEY,
    code  country
);

INSERT INTO locations VALUES (1, 'US'), (2, 'GB'), (3, 'FR');

SELECT * FROM locations ORDER BY code;
```

### Data Type

The `country` type restricts values to valid ISO 3166-1 alpha-2 codes (two-letter country designations like `US`, `GB`, `FR`, `DE`, etc.). Invalid codes are rejected on input.

### Operators

Standard comparison operators are supported: `=`, `<>`, `<`, `>`, `<=`, `>=`.

### Index Support

Both B-tree and hash index operator classes are provided for efficient querying and sorting on `country` columns.
