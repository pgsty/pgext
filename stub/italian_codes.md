## Usage

Sources:

- [Upstream English documentation](https://github.com/dvarrazzo/italian_codes/blob/3d4d0604bfe12addf55d19683ed28d9a36187fb6/doc/italian_codes.rst)
- [Version 1.0 SQL implementation](https://github.com/dvarrazzo/italian_codes/blob/3d4d0604bfe12addf55d19683ed28d9a36187fb6/sql/italian_codes.sql)
- [Extension control file](https://github.com/dvarrazzo/italian_codes/blob/3d4d0604bfe12addf55d19683ed28d9a36187fb6/italian_codes.control)

`italian_codes` provides text domains and validation functions for Italian fiscal codes: `codice_fiscale` accepts 16-character natural-person or 11-digit legal-person codes, and `partita_iva` accepts Italian VAT numbers. It validates supplied identifiers; it does not derive a fiscal code from personal data.

```sql
CREATE EXTENSION italian_codes;

SELECT codice_fiscale('mss trs 53b19 h892p'::text);
SELECT partita_iva('0575346 048 3'::text);
```

The constructor functions normalize whitespace, and `codice_fiscale` also uppercases input. Pass an explicitly typed `text` expression: PostgreSQL can interpret a bare literal such as `codice_fiscale('X')` as a domain literal and bypass the normalizing function.

Use `codice_fiscale_error(text)` or `partita_iva_error(text)` when validation errors should be returned as text instead of raised as `check_violation`; these helpers expect already-normalized input. The lower-level normalizers do not validate content, and validation messages are in Italian.

Fiscal-code generation is intentionally absent because omocodia can require a Revenue Service-issued alternative. Preserve the authoritative identifier, distinguish format validation from identity verification, and avoid treating either domain as proof that a person or business exists.
