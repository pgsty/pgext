

## Usage

> [numeral: text numeral data types (English, German, Roman)](https://github.com/df7cb/postgresql-numeral)

The `numeral` extension provides three custom numeric data types that use textual numerals instead of digits.

```sql
CREATE EXTENSION numeral;
```

### Data Types

- **`numeral`**: English numerals using short scale (10^9 = billion)
- **`zahl`**: German numerals using long scale (10^9 = Milliarde)
- **`roman`**: Roman numerals

All three are internally binary-compatible with `bigint` and implicitly cast to it.

### Examples

```sql
-- English numerals
SELECT 'thirty'::numeral + 'twelve'::numeral;
-- forty-two

-- German numerals
SELECT 'siebzehn'::zahl * 'dreiundzwanzig'::zahl;
-- dreihunderteinundneunzig

-- Roman numerals
SELECT 'MCMLV'::roman + 'II'::roman * 'XXX'::roman;
-- MMXV
```

### Operators

Standard arithmetic operators (`+`, `-`, `*`, `/`) work through the implicit `bigint` cast. All existing bigint operators and functions are available.
