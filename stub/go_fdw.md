## Usage

Sources:

- [go_fdw README at the reviewed commit](https://github.com/dennwc/go-fdw/blob/93f15338153f6e07f1673e8733863cc17258a181/README.md)
- [go_fdw install SQL at the reviewed commit](https://github.com/dennwc/go-fdw/blob/93f15338153f6e07f1673e8733863cc17258a181/go_fdw--1.0.sql)
- [go_fdw.control at the reviewed commit](https://github.com/dennwc/go-fdw/blob/93f15338153f6e07f1673e8733863cc17258a181/go_fdw.control)

`go_fdw` is an experimental Go/cgo template for developing a PostgreSQL foreign data wrapper, not a general connector to Go data. The bundled example supports table scans, EXPLAIN, and table options; developers replace the example's `SetTable` implementation with their own data source logic.

### Exercise the Bundled Example

```sql
CREATE EXTENSION go_fdw;

CREATE FOREIGN TABLE public.gotest (
  id integer NOT NULL,
  name text NOT NULL
)
SERVER "go-fdw"
OPTIONS (foo 'bar');

SELECT * FROM public.gotest;
```

The extension install script creates the `go_fdw` wrapper and a server named `go-fdw` automatically.

### Caveats

- Upstream calls the project experimental and documents testing only with PostgreSQL 9.6 and Go 1.8.1. Validate and port it before considering use with a current server.
- This is a development template with a basic working implementation. Its table options and returned rows depend on the Go code compiled into the library.
- Replacing the compiled shared library during development requires a PostgreSQL restart before retesting, according to the README; this is not a `shared_preload_libraries` requirement.
