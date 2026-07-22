## Usage

Sources:

- [Official README](https://github.com/marconeperes/pgBRTypes/blob/5491ecb720e6dab34ba4211509c2b75ca9f05c29/README.md)
- [Official extension SQL](https://github.com/marconeperes/pgBRTypes/blob/5491ecb720e6dab34ba4211509c2b75ca9f05c29/sql/pgbrtypes--1.0.sql)
- [Official GUC implementation](https://github.com/marconeperes/pgBRTypes/blob/5491ecb720e6dab34ba4211509c2b75ca9f05c29/pgBRTypes.c)

`pgbrtypes` adds validated `cpf` and `cnpj` types for Brazilian individual and company taxpayer registration numbers. It normalizes formatted or unformatted input, preserves leading zeros in the typed value, and can present a conventional display mask.

### Core Workflow

```sql
CREATE EXTENSION pgbrtypes;

CREATE TABLE counterparties (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  person_cpf cpf,
  company_cnpj cnpj
);

INSERT INTO counterparties (person_cpf, company_cnpj)
VALUES ('021.516.186-68', '14.320.135/0001-49');

CREATE INDEX counterparties_cpf_idx
ON counterparties (person_cpf);

SELECT person_cpf, company_cnpj
FROM counterparties;
```

Invalid document numbers are rejected by the type input functions. Use text input when leading zeros matter; integer literals cannot represent the original textual width even though implicit integer conversions are available.

### Types and Helpers

- `cpf` and `cnpj` are fixed-length eight-byte types with binary send/receive functions.
- `is_cpf(text)` and `is_cnpj(text)` validate candidate values without storing them; overloads also accept bigint and C-string inputs.
- Comparison operators and the `cpf_op_class` and `cnpj_op_class` B-tree operator classes support equality/range predicates and indexes.
- Implicit casts accept integer/bigint input, while assignment casts convert typed values back to bigint.

Validation checks the identifier's syntactic/check-digit rules; it does not prove that a person or company exists or owns the number.

### Output Formatting

The user-settable GUCs change display only, not stored identity or comparison behavior:

```sql
SHOW cpf.enable_format;
SHOW cnpj.enable_format;

SET cpf.enable_format = off;
SET cnpj.enable_format = off;
```

Because text output changes by session, integrations should choose a stable setting or cast to bigint deliberately. Treat CPF/CNPJ values as sensitive identifiers and apply access control, logging redaction, and retention policies independently of formatting.

### Compatibility

The README requires PostgreSQL 9.1 or later but reports testing only on 9.6; the reviewed source was last changed in 2017 and does not document Windows support. Build and test the exact PostgreSQL/OS target, including indexes, binary dump/restore, upgrade, malformed input, and locale-independent output, before production use.
