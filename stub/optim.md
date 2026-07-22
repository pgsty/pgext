## Usage

Sources:

- [Monorepo README](https://github.com/ChristoferBerruz/hyper_db/blob/07c3b41e90ef434f29b3a97ac701bfce62e73363/README.md)
- [Extension implementation](https://github.com/ChristoferBerruz/hyper_db/blob/07c3b41e90ef434f29b3a97ac701bfce62e73363/optim/src/lib.rs)
- [Extension control file](https://github.com/ChristoferBerruz/hyper_db/blob/07c3b41e90ef434f29b3a97ac701bfce62e73363/optim/optim.control)

`optim` is an early pgrx proof of concept from a database course project. It exposes optimization-related PostgreSQL types plus a single demonstration function; it does not yet implement an optimizer or a documented hyperparameter-search workflow.

### Core Workflow

The control file requires a superuser to install the untrusted C-compatible module. After installation, the only documented callable operation is the generated `hello_optim()` function:

```sql
CREATE EXTENSION optim;
SELECT hello_optim();
```

The expected result is `Hello, optim`. Treat this as a smoke test of the pgrx module, not an optimization result.

### Source-Defined Objects

- `OptimizationDirection`: pgrx `PostgresEnum` with `Maximize` and `Minimize` variants.
- `Hyperparameter`: pgrx `PostgresType` containing one `f64` value.
- `Trial`: pgrx `PostgresType` containing a vector of `Hyperparameter` values and an `f64` score.
- `hello_optim()`: returns the constant text `Hello, optim`.

The exact SQL spelling and I/O syntax of the pgrx-generated types should be inspected on the installed build before use; upstream provides no SQL examples for them.

### Operational Notes

The crate version is `0.0.0`, the control file is generated from the Cargo version, and the README describes the repository as a student database project. The control file declares `relocatable = false`, `superuser = true`, and `trusted = false`; no preload setting is defined. There is no documented persistence model, planner hook workflow, release contract, or compatibility guarantee. Use `optim` only for source evaluation or isolated experimentation until upstream defines a stable SQL API.
