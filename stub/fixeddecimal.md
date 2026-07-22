## Usage

Sources:

- [Official README](https://github.com/2ndQuadrant/fixeddecimal/blob/970f56b53a2eef74130527778190597dbd9a6a30/README.md)
- [Extension control file](https://github.com/2ndQuadrant/fixeddecimal/blob/970f56b53a2eef74130527778190597dbd9a6a30/fixeddecimal.control)
- [Extension SQL](https://github.com/2ndQuadrant/fixeddecimal/blob/970f56b53a2eef74130527778190597dbd9a6a30/fixeddecimal--1.1.0.sql)
- [Type implementation](https://github.com/2ndQuadrant/fixeddecimal/blob/970f56b53a2eef74130527778190597dbd9a6a30/fixeddecimal.c)

`fixeddecimal` provides an 8-byte, fixed-scale decimal type for values such as money where a scale of exactly two decimal places and compact storage are useful. It offers exact storage and exact addition or subtraction within its range, but it does not implement arbitrary `numeric` precision or configurable scale.

### Core Workflow

```sql
CREATE EXTENSION fixeddecimal;

CREATE TABLE invoice (
  invoice_id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  amount fixeddecimal(12,2) NOT NULL
);

INSERT INTO invoice(amount) VALUES ('19.95'), ('0.05');
SELECT sum(amount), avg(amount) FROM invoice;
CREATE INDEX invoice_amount_idx ON invoice (amount);
```

The declaration is `fixeddecimal(precision,2)`, with precision from 3 through 17. Values are stored as a signed 64-bit integer scaled by 100, giving a range from `-92233720368547758.08` through `92233720368547758.07`. The extension supplies comparison and arithmetic operators, casts to and from common numeric types, `sum`, `avg`, B-tree and hash operator classes, and a BRIN minmax operator class.

### Precision and Compatibility

Input beyond two fractional digits is truncated rather than rounded, and multiplication or division truncates toward zero after rescaling. Use `round(value::numeric, 2)::fixeddecimal` when business rules require explicit rounding. Overflow raises an error, and the type has no `NaN` representation.

No preload is required. Upstream documents PostgreSQL 9.5 and later but says the code was last tested on PostgreSQL 12; the pinned 1.1.0 source is from 2019. Validate compilation, binary upgrade, dump/restore, casts, aggregates, and indexes on each newer PostgreSQL major before adopting the type. Changing a column later to a different scale requires conversion to another type because only scale 2 is implemented.
