## Usage

Sources:

- [Upstream README](https://github.com/ryapric/rpg/blob/76c6a508ac04fca27ebb0ea3ffc675e7a3fb6442/README.md)
- [Extension SQL](https://github.com/ryapric/rpg/blob/76c6a508ac04fca27ebb0ea3ffc675e7a3fb6442/rpg--0.1.sql)
- [Extension control file](https://github.com/ryapric/rpg/blob/76c6a508ac04fca27ebb0ea3ffc675e7a3fb6442/rpg.control)

`rpg` version `0.1` is a small SQL/PLpgSQL prototype with three utility functions. `year_month` formats a date, and `letters` returns the English alphabet in lower or upper case:

```sql
CREATE EXTENSION rpg;

SELECT year_month(DATE '2026-07-16');
SELECT letters('upper');
```

### Prototype limitations

The third function, `bind_rows(output_table, table_a, table_b)`, attempts to create a table from a `UNION ALL` over two named tables. Its dynamic SQL concatenates table and column identifiers without identifier quoting, and the input relations must still be union-compatible. Never pass untrusted names to `bind_rows`; inspect and patch the implementation before operational use. The upstream README contains no API, compatibility, or support documentation.
