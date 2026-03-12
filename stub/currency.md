

## Usage

> [currency: ISO 4217 currency code type](https://github.com/adjust/pg-currency)

The `currency` extension provides a data type for ISO 4217 currency codes using only a single byte of storage per value.

```sql
CREATE EXTENSION currency;

CREATE TABLE transactions (
    id                serial,
    payment_currency  currency
);

INSERT INTO transactions VALUES (1, 'USD'), (2, 'EUR'), (3, 'USD');

SELECT * FROM transactions ORDER BY payment_currency;
 id | payment_currency
----+------------------
  2 | EUR
  1 | USD
  3 | USD
```

### Operators

Standard comparison operators are supported: `=`, `<>`, `<`, `>`, `<=`, `>=`.

B-tree index support is included for efficient ordering and lookups.

### Functions

```sql
-- List all supported currency codes
SELECT * FROM supported_currencies() currency ORDER BY currency;
```
