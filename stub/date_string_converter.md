## Usage

Sources:

- [Project README](https://github.com/vee-huyvunguyen/postgres_rust_extension/blob/667aaa420473f1b3b89507152dac5e7f4240163e/README.md)
- [date_string_converter implementation](https://github.com/vee-huyvunguyen/postgres_rust_extension/blob/667aaa420473f1b3b89507152dac5e7f4240163e/date_string_converter/src/lib.rs)
- [Extension control file](https://github.com/vee-huyvunguyen/postgres_rust_extension/blob/667aaa420473f1b3b89507152dac5e7f4240163e/date_string_converter/date_string_converter.control)

`date_string_converter` 0.0.0 is a pgrx demonstration extension. Although the repository README describes an intended converter for strings such as `1w 2d 4h 30m`, the pinned extension implementation exposes only a greeting function; it contains no date-duration parser or conversion API.

### Available Workflow

```sql
CREATE EXTENSION date_string_converter;

SELECT hello_date_string_converter();
-- Hello, date_string_converter
```

The only user-facing object in this version is `hello_date_string_converter() RETURNS text`. The extension is non-relocatable and its control file requires superuser installation.

### Limitations

Do not design application logic around the converter described at repository level: there is no function that accepts `1w 2d 4h 30m`, and no result such as `220.5` can be produced by this extension revision. Treat it as a minimal extension scaffold until upstream adds and documents a real conversion function. Its Cargo configuration targets PostgreSQL 11 through 16 via pgrx 0.10.1; newer server compatibility is not established by this source snapshot.
