## Usage

Sources:

- [Official control file](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/pgrx_validation/pgrx_validation.control)
- [Official PGRX Validation README](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/pgrx_validation/README.md)
- [Generated PostgreSQL 17 extension SQL](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/pgrx_validation/extension/usr/share/postgresql/17/extension/pgrx_validation--0.0.0.sql)

`pgrx_validation` exposes the EMI project's shared validation rules as PostgreSQL functions. They are intended for database constraints that apply the same string, numeric, UUID, timestamp, and icon checks used by other application targets.

### Core Workflow

```sql
CREATE EXTENSION pgrx_validation;

CREATE TABLE contacts (
    email text NOT NULL CHECK (must_be_email(email)),
    score double precision NOT NULL CHECK (must_be_strictly_positive_f64(score))
);
```

The generated functions are `STRICT`, so a null input yields null. A PostgreSQL `CHECK` accepts null; use `NOT NULL` separately when absence must be rejected.

### Validation Functions

- Strings: `must_be_email(text)`, `must_not_be_empty(text)`, `must_not_be_padded(text)`, `must_not_contain_consecutive_whitespace(text)`, `must_not_contain_control_characters(text)`, `must_be_paragraph(text)`, and `must_be_distinct(text, text)`.
- Integers and UUIDs: `must_be_positive_i16(smallint)`, `must_be_positive_i32(int)`, `must_be_strictly_positive_i16(smallint)`, `must_be_strictly_positive_i32(int)`, `must_be_distinct_i16(smallint, smallint)`, `must_be_distinct_i32(int, int)`, and `must_be_distinct_uuid(uuid, uuid)`.
- Floating point: `must_be_strictly_positive_f32(real)`, `must_be_strictly_positive_f64(double precision)`, plus the `must_be_greater_than_f32`, `must_be_smaller_than_f32`, `must_be_strictly_greater_than_f32`, `must_be_strictly_smaller_than_f32`, `must_be_strictly_greater_than_f64`, and `must_be_strictly_smaller_than_f64` bound checks.
- Other checks: `must_be_smaller_than_utc(timestamptz, timestamptz)`, `must_be_strictly_smaller_than_utc(timestamptz, timestamptz)`, and `must_be_font_awesome_class(text)`.

### Operational Notes

The reviewed package version is 0.0.0 and the API may still change. The extension is non-relocatable; its control file sets `superuser = true` and `trusted = false`, so a superuser must create it. It declares no prerequisite extension or preload requirement. Treat these predicates as application validation, not security sanitization, and regression-test constraints when upgrading the shared rules.
