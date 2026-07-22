## Usage

Sources:

- [Official README](https://github.com/tatut/pgprolog/blob/e3629cc2f5a56da98c22939436d9d803e2f9a6d5/README.md)
- [Rust language-handler implementation](https://github.com/tatut/pgprolog/blob/e3629cc2f5a56da98c22939436d9d803e2f9a6d5/src/lib.rs)
- [Extension control file](https://github.com/tatut/pgprolog/blob/e3629cc2f5a56da98c22939436d9d803e2f9a6d5/pgprolog.control)

`pgprolog` is a proof-of-concept procedural-language handler that embeds Scryer Prolog in PostgreSQL. It lets a function body contain Prolog facts plus a `handle/2` predicate, but upstream explicitly says it is not ready for actual use.

### Minimal Workflow

Extension installation is superuser-only. The language is not created automatically, so create `plprolog` with the installed handler before defining a function.

```sql
CREATE EXTENSION pgprolog;
CREATE LANGUAGE plprolog HANDLER plprolog_call_handler;

CREATE FUNCTION parent_of(text)
RETURNS text
AS $$
parent(alice,bob).
handle(In,Out) :- parent(In,Out).
$$ LANGUAGE plprolog;

SELECT parent_of('alice');
```

The published example demonstrates text input and text output. The project does not document a production type-mapping matrix, trigger support, set-returning functions, transaction/SPI access, modules, persistent Prolog state, or resource controls.

### Trust and Prototype Boundaries

`CREATE LANGUAGE` without `TRUSTED` creates an untrusted language, so only superusers may create functions in it. Keep it untrusted: arbitrary Prolog programs execute a young embedded runtime inside a PostgreSQL backend, and parser, allocation, recursion, or panic defects can terminate or exhaust that backend.

Version `0.0.0` pins pgrx `0.11.3` and Scryer Prolog `0.9.4`; advertised Cargo features stop at PostgreSQL 16. Use only in an isolated lab with statement and process limits. Before any further adoption, the handler needs explicit type/error semantics, cancellation tests, memory and recursion bounds, privilege review, fuzzing, crash isolation, and a maintained compatibility policy.
