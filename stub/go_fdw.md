## Usage

Sources:

- [Official README](https://github.com/dennwc/go-fdw/blob/93f15338153f6e07f1673e8733863cc17258a181/README.md)
- [Extension SQL](https://github.com/dennwc/go-fdw/blob/93f15338153f6e07f1673e8733863cc17258a181/go_fdw--1.0.sql)
- [Bundled Go example](https://github.com/dennwc/go-fdw/blob/93f15338153f6e07f1673e8733863cc17258a181/fdw.go)

`go_fdw` 1.0 is an experimental template for writing a PostgreSQL foreign data wrapper in Go. The distributed example does not connect to an external system: it registers a hard-coded generator that returns ten synthetic rows. Developers are expected to replace `fdw.go` and rebuild the shared library for their own source.

### Bundled Example

```sql
CREATE EXTENSION go_fdw;

CREATE FOREIGN TABLE public.gotest (
  id integer NOT NULL,
  name text NOT NULL
)
SERVER "go-fdw"
OPTIONS (foo 'bar');

SELECT * FROM gotest;
```

The installation creates foreign data wrapper `go_fdw` and server `go-fdw`. In the bundled implementation, `id` is derived from the row counter and `name` becomes text such as a row/column label. Nullable columns are left NULL; the demonstration explicitly handles only integer and text types. Table options are passed to the Go callbacks but `foo='bar'` has no meaning in the example.

### Development Boundary

The framework implements table estimates, scans, rescan/close behavior, table options, and `EXPLAIN` properties. It has no generic remote protocol, write callbacks, transaction integration, qualifier pushdown, join pushdown, authentication, or retry behavior. Those semantics belong entirely to the custom Go code supplied to `SetTable`.

Upstream reports testing only with PostgreSQL 9.6 and Go 1.8.1 on Ubuntu x64. The module crosses PostgreSQL C and Go runtime boundaries and uses old server APIs; it must be rebuilt and stress-tested for the exact PostgreSQL and Go versions, especially around memory, signals, concurrency, and backend shutdown. Replacing the library requires new PostgreSQL backend processes because an already loaded shared object is not reloaded in place.
