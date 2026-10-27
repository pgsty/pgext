## Usage

Sources:

- [README](https://github.com/thomasdarimont/pgauthz/blob/64ffad245fdfed11cae42ccc3849082340b9b420/extensions/pg-cel/README.md)
- [Control](https://github.com/thomasdarimont/pgauthz/blob/64ffad245fdfed11cae42ccc3849082340b9b420/extensions/pg-cel/pg_cel.control)
- [Implementation](https://github.com/thomasdarimont/pgauthz/blob/64ffad245fdfed11cae42ccc3849082340b9b420/extensions/pg-cel/src/lib.rs)
- [Cargo](https://github.com/thomasdarimont/pgauthz/blob/64ffad245fdfed11cae42ccc3849082340b9b420/extensions/pg-cel/Cargo.toml)

`pg_cel` 0.2.0 evaluates Common Expression Language conditions for pgauthz. It is a Rust/pgrx shared library; PostgreSQL 18 is the default build, with Cargo features for 13–18. Core use does not require preload or restart.

### Evaluate a Condition

```sql
CREATE SCHEMA IF NOT EXISTS authz;
CREATE EXTENSION pg_cel SCHEMA authz;
SELECT authz.cel_compile_check('request.age >= 18');
SELECT authz.cel_eval_bool('request.age >= 18',
                           '{"request":{"age":21}}');
```

### Integration and Failure Boundaries

`cel_compile_check` rejects malformed expressions; `cel_eval_bool` evaluates against JSON input and returns a boolean or NULL for a non-boolean result. Compile or evaluation errors raise SQL errors. Policy callers must treat errors, NULL and non-true results as denial. The pgauthz engine does this and expects the functions in `authz`; the extension itself is relocatable.

The expression evaluator has no external side effects. Input variable names and types must match the JSON context. This is the optional CEL evaluator in pgauthz release 0.19.0, not a replacement for the engine's authorization model or its migrations.
