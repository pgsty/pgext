## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/frectonz/pg-ethopic/blob/6748b47634034085b503723ec4b4007a7cda7b8b/README.md)
- [Rust implementation](https://github.com/frectonz/pg-ethopic/blob/6748b47634034085b503723ec4b4007a7cda7b8b/src/lib.rs)
- [Cargo feature matrix](https://github.com/frectonz/pg-ethopic/blob/6748b47634034085b503723ec4b4007a7cda7b8b/Cargo.toml)
- [Extension control file](https://github.com/frectonz/pg-ethopic/blob/6748b47634034085b503723ec4b4007a7cda7b8b/pg_ethopic.control)

`pg_ethopic` formats PostgreSQL dates in the Ethiopic calendar and integers with Ethiopic numerals.

```sql
CREATE EXTENSION pg_ethopic;
SELECT ethopic_date(CURRENT_DATE);
SELECT ethopic_date(DATE '2026-01-01', '{year}/{month}/{day}');
SELECT ethopic_number(420);
```

The custom date template supports the fields implemented by the upstream Rust library; validate templates and expected calendar boundaries with domain examples. Results are formatted text, not a distinct calendar-aware date type, so retain the original PostgreSQL `date` for sorting and arithmetic.

Version 0.0.1 is an early pgrx extension. The README provides container images for PostgreSQL 13–17, while Cargo also defines a PostgreSQL 12 feature. Test conversion results, negative and large numbers, locale-independent output, and upgrades before using formatted values in exported records or legal documents.
