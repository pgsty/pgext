## Usage

Sources:

- [Official README](https://github.com/shestero/pglas/blob/main/README.md)
- [Official Rust implementation](https://github.com/shestero/pglas/blob/main/src/lib.rs)
- [Official control file](https://github.com/shestero/pglas/blob/main/pglas.control)
- [Official LAS 2.0 specification](https://www.cwls.org/wp-content/uploads/2017/02/Las2_Update_Feb2017.pdf)

`pglas` reads Log ASCII Standard (LAS) 2.0 well-log files from the PostgreSQL server's filesystem. It reports the file's null marker, lists curve names, and returns depth/value pairs for one curve. These are well-logging LAS files, not LiDAR point-cloud files that share the `.las` suffix.

### Core Workflow

Create the extension as a superuser and query a server-local LAS 2.0 file:

```sql
CREATE EXTENSION pglas;

SELECT las_na('/srv/well-logs/A10.las');

SELECT *
FROM las_curves('/srv/well-logs/A10.las');

SELECT *
FROM las_curve('/srv/well-logs/A10.las', 'Gamma')
ORDER BY "DEPT";
```

`las_curve` pairs the selected curve with the `DEPT` curve, which the implementation expects to be first in the file. LAS null-marker values are returned as SQL `NULL` where the parser recognizes them.

### API

- `las_na(file text) RETURNS double precision`: the file's declared N/A marker.
- `las_curves(file text)`: rows `(IDX bigint, CURVE text)` enumerating curve headers.
- `las_curve(file text, curve text)`: rows `(DEPT double precision, VAL double precision)` for a selected curve.

### Caveats

The functions accept arbitrary server-side paths and open files with the PostgreSQL server process's privileges. Restrict `EXECUTE` to trusted roles and place readable files in a dedicated directory; otherwise the API can disclose server-local data or filesystem error details.

Only LAS 2.0 is supported. The implementation assumes the first curve is depth and relies on the `lasrs` parser, so validate headers, null markers, curve lengths, case-sensitive names, and malformed records before loading results into permanent tables. The packaged pgrx build must match the PostgreSQL major version.
