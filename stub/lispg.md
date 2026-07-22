## Usage

Sources:

- [Extension control file](https://github.com/niquola/lispg/blob/e43d7961477259e910761218a5541211d2cf8860/lispg.control)
- [Extension SQL for 1.0](https://github.com/niquola/lispg/blob/e43d7961477259e910761218a5541211d2cf8860/lispg--1.0.sql)
- [Evaluator implementation](https://github.com/niquola/lispg/blob/e43d7961477259e910761218a5541211d2cf8860/lispg_operations.c)

`lispg` version 1.0 is an experimental Lisp-like value type and evaluator for PostgreSQL. Its implemented evaluator is a very small numeric calculator, so it should not be treated as a general Lisp runtime or a procedural language.

### Core Workflow

```sql
CREATE EXTENSION lispg;

SELECT '(+ 2 3)'::lispg;
SELECT lispg_debug('(+ 2 3)');
```

The custom type parses and serializes parenthesized expressions. The debug evaluator parses a C string, evaluates the expression, and returns its textual representation.

### SQL Objects

- `lispg` is a variable-length, extended-storage type.
- `lispg_in(cstring)` and `lispg_out(lispg)` are its input and output functions.
- `lispg_debug(cstring) RETURNS text` evaluates an expression for inspection.

The pinned implementation registers only binary `+` and `-` numeric operations. It has no documented persistence, sandboxing, SQL access, or stable language API. Use it only for experimentation against the exact build you have installed.
