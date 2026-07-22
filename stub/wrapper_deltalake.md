## Usage

Sources:

- [Official README](https://github.com/Giangblackk/wrapper_deltalake/blob/0459b9f308ee790df21dafdd9f366b5b54fb33b0/README.md)
- [Extension control file](https://github.com/Giangblackk/wrapper_deltalake/blob/0459b9f308ee790df21dafdd9f366b5b54fb33b0/wrapper_deltalake.control)
- [Rust implementation](https://github.com/Giangblackk/wrapper_deltalake/blob/0459b9f308ee790df21dafdd9f366b5b54fb33b0/src/lib.rs)

wrapper_deltalake is an unfinished proof of concept that planned to connect Supabase Wrappers and delta-rs. The checked version does not implement a foreign data wrapper, Delta Lake reader, server options, or table mapping; it exposes only a greeting function.

### Core Workflow

The complete user-visible behavior is the demonstration call:

```sql
CREATE EXTENSION wrapper_deltalake;

SELECT hello_wrapper_deltalake();
```

It returns a fixed greeting. There is no supported way to define a foreign server or query a Delta table.

### Important Objects

- `hello_wrapper_deltalake` returns the fixed text `Hello, wrapper_deltalake`.
- `wrapper_deltalake` is the extension name, but no FDW handler or validator is registered.

### Operational Notes

The catalog and control file use version 0.0.0, and the README describes future intent rather than implemented behavior. Do not invent configuration from Supabase Wrappers or delta-rs documentation: those projects' capabilities are not present here. Use a maintained, actually implemented Delta Lake integration for real data access.
