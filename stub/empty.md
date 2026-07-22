## Usage

Sources:

- [Official README](https://github.com/tvondra/empty/blob/ac8d6da58df3763a462a5fdd1f781269cff7d67d/README.md)
- [Official extension control file](https://github.com/tvondra/empty/blob/ac8d6da58df3763a462a5fdd1f781269cff7d67d/empty.control)
- [Official extension SQL](https://github.com/tvondra/empty/blob/ac8d6da58df3763a462a5fdd1f781269cff7d67d/empty--1.0.sql)
- [Official C implementation](https://github.com/tvondra/empty/blob/ac8d6da58df3763a462a5fdd1f781269cff7d67d/empty.c)

`empty` 1.0 is an archived PostgreSQL extension-development sandbox, not one cohesive user feature. It collects demonstration code for C functions, a fixed 2×2 `matrix` type, a sample FDW, logical decoding callbacks, table scanning, shared memory, and an error-log hook. Use it only for source study or controlled experiments.

### Minimal Experiment

The library requests shared memory and an LWLock tranche in `_PG_init`, so load it through `shared_preload_libraries` and restart PostgreSQL before creating the extension. A small self-contained example exercises the matrix code:

```conf
shared_preload_libraries = 'empty'
```

```sql
CREATE EXTENSION empty;

SELECT '(1 2 3 4)'::matrix + '(5 6 7 8)'::matrix;
SELECT '(1 2 3 4)'::matrix * '(5 6 7 8)'::matrix;

SELECT *
FROM matrix_powers('(1 1 1 0)'::matrix, 4);
```

The parser accepts exactly four integers in the displayed parenthesized format. `matrix_powers(matrix, count)` emits successive powers with a zero-based ordinal.

### Demonstration Surface

- Arithmetic examples: `int4_plus`, `numeric_plus`, and `random_string`.
- Matrix example: `matrix`, operators `+` and `*`, and `matrix_powers`.
- Inspection example: `read_table(regclass)` scans a table and writes every value to server warnings rather than returning rows.
- FDW example: `empty_fdw` contains experimental file/scan code with no stable documented contract.
- The same shared library exports logical-decoding callbacks and counts selected server message levels through a global emit-log hook.

These pieces are unrelated teaching experiments; the terse README does not define supported behavior or a compatibility policy.

### Safety and Maintenance Boundaries

`read_table` can disclose complete row contents through server logs, and the global hooks affect the whole instance once preloaded. Restrict installation and execution, never point the sample at sensitive data, and test hook interaction with other preloaded modules.

The project was archived after its 2019 commit and uses server internals that change across PostgreSQL releases. Some declarations and implementation choices are intentionally illustrative rather than production-grade. Do not treat `empty` as a supported FDW, monitoring extension, random generator, matrix library, or decoding plugin; copy and review only the specific pattern needed.
