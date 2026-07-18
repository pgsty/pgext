## Usage

Sources:

- [Upstream README](https://github.com/patchsoft/postcode/blob/13ccded329c2e12081987d197cfe44413c35bd4e/README.md)
- [Extension control file](https://github.com/patchsoft/postcode/blob/13ccded329c2e12081987d197cfe44413c35bd4e/postcode.control)
- [SQL installation script](https://github.com/patchsoft/postcode/blob/13ccded329c2e12081987d197cfe44413c35bd4e/postcode--1.3.0.sql)
- [PGXN release page](https://pgxn.org/dist/postcode/)

`postcode` version `1.3.0` provides compact `postcode` and `dps` types for UK postcodes and delivery-point suffixes. It normalizes input, supplies B-tree comparison support, supports partial matching with `%`, and exposes validation and formatting helpers.

### Example

```sql
CREATE EXTENSION postcode;
CREATE TABLE addresses (code postcode);
INSERT INTO addresses VALUES ('SW1A 1AA'), ('LS2 4AA');
SELECT postcode_validate('SW1A 1AA');
SELECT * FROM addresses WHERE code % 'SW1A';
SELECT to_char(code, 'AD SW') FROM addresses;
```

Syntactic validation does not prove that a postcode is currently allocated or that an address exists; verify such claims against an authoritative current address dataset. Partial forms can be ambiguous (for example, district versus sector), so define accepted input levels in application code. The upstream release is from 2015, has no current compatibility matrix, and its embedded UK-format assumptions may be stale.
