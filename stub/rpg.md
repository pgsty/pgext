## Usage

Sources:

- [Official README](https://github.com/ryapric/rpg/blob/master/README.md)
- [Extension control file](https://github.com/ryapric/rpg/blob/master/rpg.control)
- [Version 0.1 extension SQL](https://github.com/ryapric/rpg/blob/master/rpg--0.1.sql)

rpg is a small pure-SQL collection of R-inspired PL/pgSQL helpers. Version 0.1 provides date formatting, English letter arrays, and an experimental dynamic table-binding function; the upstream README is explicitly unfinished, so the SQL file is the authoritative API surface.

### Safe Helpers

The date helper returns a year and zero-padded month separated by an optional delimiter. The letter helper returns the lower-case alphabet by default or the upper-case alphabet for the recognized alternative.

```sql
CREATE EXTENSION rpg;

SELECT year_month(DATE '2026-07-22');
SELECT year_month(DATE '2026-07-22', '/');
SELECT letters();
SELECT letters('upper');
```

An unrecognized letter-case argument returns null because version 0.1 has no fallback branch. Null input also propagates through the function logic rather than being rejected by a strict declaration.

### Experimental Table Binding

The table-binding helper accepts three text arguments naming an output table and two input tables, then constructs a create-table-as union. Its implementation concatenates unquoted identifiers, searches information_schema by table name without schema qualification, sorts each input's columns alphabetically, and does not add missing columns despite the source comment claiming support for different column counts.

Do not pass untrusted values to that function and do not use it as a general substitute for an explicit union. It is vulnerable to identifier ambiguity and SQL injection, and inputs with different column counts or incompatible alphabetic column order can fail or produce incorrect alignment. Prefer a reviewed statement that quotes schema-qualified identifiers and names every output column.

### Operational Boundary

The extension depends on PL/pgSQL being installed, is relocatable, and declares no preload or restart requirement. Because there are no meaningful upstream regression tests and the project documentation is incomplete, treat version 0.1 as a convenience prototype: pin behavior with local tests before putting even the simple helpers into durable application logic.
