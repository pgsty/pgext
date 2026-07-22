## Usage

Sources:

- [Official control file](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/cas_codes/cas_codes.control)
- [Official CAS Codes README](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/cas_codes/README.md)
- [Generated PostgreSQL 17 extension SQL](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/cas_codes/extension/usr/share/postgresql/17/extension/cas_codes--0.1.0.sql)

`cas_codes` adds the `CAS` type for Chemical Abstracts Service Registry Numbers. Input is parsed with format and check-digit validation, making the type useful for rejecting malformed chemical identifiers at the database boundary.

### Core Workflow

```sql
CREATE EXTENSION cas_codes;

CREATE TABLE reagents (
    name text NOT NULL,
    cas_number CAS NOT NULL UNIQUE
);

INSERT INTO reagents VALUES ('water', '7732-18-5');
SELECT * FROM reagents WHERE cas_number = '7732-18-5'::CAS;
```

An invalid format or checksum is rejected by the type input function. The value is displayed in the conventional hyphenated form.

### SQL Objects

- `CAS` is a variable-length custom type in the current PGRX build; the README notes that fixed-size custom types were not yet supported by its PGRX version.
- Comparison operators `=`, `<>`, `<`, `<=`, `>`, and `>=` are available.
- `CAS_btree_ops` supports B-tree ordering and uniqueness, while `CAS_hash_ops` supports hashing.

### Operational Notes

Version 0.1.0 is non-relocatable. Its control file sets `superuser = true` and `trusted = false`, so a superuser must create it; it declares no prerequisite extension or preload requirement. The generated SQL is tied to the PostgreSQL/PGRX build, so install matching control, SQL, and shared-library artifacts together before restoring schemas that use `CAS`.
