## Usage

Sources:

- [2.0.0 extension README](https://github.com/monumental-archive/edtf/blob/v2.0.0/crates/edtf-postgres/README.md)
- [2.0.0 control](https://github.com/monumental-archive/edtf/blob/v2.0.0/crates/edtf-postgres/edtf_postgres.control)
- [2.0.0 release](https://github.com/monumental-archive/edtf/releases/tag/v2.0.0)

`edtf_postgres` 2.0.0 validates and normalizes Extended Date/Time Format expressions, calculates date bounds and compares temporal relations for ISO 8601-2:2019 Annex A.

### Enable and Query

Release 2.0.0 supports PostgreSQL 15–19. The extension is trusted and relocatable: a role with database `CREATE` privilege can install it. No preload is required.

```sql
CREATE EXTENSION edtf_postgres;
SELECT edtf_valid('1985-04-12');
SELECT edtf_level('1985-04-12/..');
SELECT edtf_canonical('1985-04-12');
SELECT edtf_min('1985-04'), edtf_max('1985-04');
SELECT edtf_relation('1985', '1986');
```

The SQL interface operates on text and derived dates. It does not introduce a stored EDTF base type. Preserve the original input when qualifiers or uncertainty matter, and use derived bounds only for the intended query semantics.

### Version 2.0 Boundaries

Version 1.3.5 is the final release with PostgreSQL 14 support. In 2.0.0, `edtf_min` and `edtf_max` return negative or positive infinity for bounds outside PostgreSQL's finite date range, replacing the earlier NULL result. Bounds in years at or before zero are corrected. Review application predicates and any stored derived values when adopting the release.

Y-prefixed years with 19 or more digits are rejected. Normalization and interpretation of ambiguous ranges also change in this major release; revalidate representative historical inputs before replacing existing normalized output.
