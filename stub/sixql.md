## Usage

Sources:

- [Official README](https://github.com/wrnrlr/sixql/blob/4b4231ca0d3ef779c9ccb4bcd0f9f3c4e41ab969/readme.md)
- [Extension implementation](https://github.com/wrnrlr/sixql/blob/4b4231ca0d3ef779c9ccb4bcd0f9f3c4e41ab969/src/lib.rs)
- [Cargo and PostgreSQL compatibility](https://github.com/wrnrlr/sixql/blob/4b4231ca0d3ef779c9ccb4bcd0f9f3c4e41ab969/Cargo.toml)

`sixql` 0.0.0 is an abandoned pgrx experiment for rendering SIXEL terminal graphics from PostgreSQL. The pinned implementation is a prototype, not a plotting API: its functions return NULL, and generated escape data is written through server logging or process standard output rather than returned over the PostgreSQL wire protocol.

### Available Experiment

A superuser can install the native module on a test server with `libsixel`, then call its two generated SQL functions:

```sql
CREATE EXTENSION sixql;

SELECT hello_sixql();

SELECT line_plot(ARRAY[0.0, 1.0, 0.5]::real[]);
```

`hello_sixql()` constructs a fixed SIXEL sample and returns NULL. `line_plot(real[])` renders a fixed 100×100 sine curve and also returns NULL; the input points are ignored in this revision. Neither function returns an image, byte array, file path, or table value that a normal SQL client can consume.

### Terminal and Server Boundaries

SIXEL requires a compatible terminal and native `libsixel`, but PostgreSQL's wire protocol and ordinary logging paths do not form a reliable binary terminal channel. Server logs may escape, split, prefix, suppress, or redirect the control sequence, and emitting it from a server process does not target the user's `psql` terminal. Do not enable raw control-sequence rendering in shared logs.

The Cargo features target PostgreSQL 11–15 through pgrx 0.8.3. The code uses unsafe FFI callbacks and experimental SIXEL crates and has no functional PostgreSQL regression test. Treat it only as source material for a future design that returns an explicit image payload; it is not suitable for production or untrusted callers.
