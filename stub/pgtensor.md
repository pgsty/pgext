## Usage

Sources:

- [Official upstream README](https://github.com/sam-harri/pgtensor/blob/969352abd2d425943c46aaac3508c41d0c731a49/README.md)
- [Official extension control file (pgtensor.control)](https://github.com/sam-harri/pgtensor/blob/969352abd2d425943c46aaac3508c41d0c731a49/pgtensor.control)
- [Official implementation source](https://github.com/sam-harri/pgtensor/blob/969352abd2d425943c46aaac3508c41d0c731a49/src/lib.rs)

`pgtensor` — Open-source Postgres extension that adds a tensor type, and an ONNX inference engine using dynamic worker processes. Built with pgrx for Rust. Use it for the corresponding vector, model, or retrieval workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pgtensor;

CREATE TABLE t (x tensor(2,3));
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The catalog records version `0.0.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- The control file marks the extension as not trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
