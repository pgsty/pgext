## Usage

Sources:

- [Official extension README](https://github.com/monumental-archive/edtf/blob/edtf-postgres-v1.2.3/crates/edtf-postgres/README.md)
- [Extension control file](https://github.com/monumental-archive/edtf/blob/edtf-postgres-v1.2.3/crates/edtf-postgres/edtf_postgres.control)
- [pgrx manifest](https://github.com/monumental-archive/edtf/blob/edtf-postgres-v1.2.3/crates/edtf-postgres/Cargo.toml)

`edtf_postgres` exposes Extended Date/Time Format validation, normalization, bounds, and temporal relations for ISO 8601-2:2019 Annex A strings.

### Enablement

Release 1.2.3 publishes artifacts for PostgreSQL 14–18 on amd64 and arm64. Install the matching artifact, then run:

```sql
CREATE EXTENSION edtf_postgres;
```

The extension is relocatable and trusted, so a role with database `CREATE` privilege can install it. It does not require preloading.

### Validation and Canonical Form

Use `edtf_valid` before accepting external text, `edtf_level` to inspect the supported EDTF level, and `edtf_canonical` to normalize equivalent forms.

```sql
SELECT edtf_valid('1985-04-12');
SELECT edtf_level('1985-04-12/..');
SELECT edtf_canonical('1985-04-12');
```

`edtf_min` and `edtf_max` return index-friendly date bounds. `edtf_relation` compares two expressions and returns a three-valued temporal relation.

```sql
SELECT edtf_min('1985-04'), edtf_max('1985-04');
SELECT edtf_relation('1985', '1986');
```

### Data Modeling Boundary

The extension operates on text and derived dates; it does not replace application validation or introduce a stored EDTF base type. Preserve the original expression when uncertainty and qualifiers are significant, and index derived bounds only for the query semantics your application intends.

