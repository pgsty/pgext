## Usage

Sources:

- [Pinned control file](https://github.com/NguyenLe1605/pgmolji/blob/42b9486e29319026df7c957b484eaf47fceabecd/pgmolji.control)
- [Pinned executor-hook implementation](https://github.com/NguyenLe1605/pgmolji/blob/42b9486e29319026df7c957b484eaf47fceabecd/src/lib.rs)

`pgmolji` is an experimental pgrx executor-hook extension. When its library is loaded in a backend, it logs a simplified representation of supported sequential-scan plans and replaces every returned `text` or `varchar` byte with a randomly selected emoji. It intentionally changes query results and is suitable only for demonstrations or isolated experiments.

### Demonstration

Use a disposable database and session:

```sql
CREATE EXTENSION pgmolji;

CREATE TEMP TABLE demo (id integer, message text);
INSERT INTO demo VALUES (1, 'hello');

SELECT id, message FROM demo WHERE id = 1;
```

The integer is passed through, while `message` is replaced by random emoji output. The module also emits a `[pgmolji]` log message for the simple sequential-scan plan.

### Exposed and Hooked Behavior

- `hello_pgmolji()` returns a text greeting before the destination receiver transforms it.
- `_PG_init()` installs an `ExecutorRun` hook when the shared library is loaded in a backend.
- The custom destination receiver transforms output attributes whose type OID is `text` or `varchar`; other types pass through.
- Plan rendering understands a top-level sequential scan and a small set of expression nodes/operators. Unsupported plans or expressions are logged as unsupported/unknown.

### Safety Boundaries

There is no GUC, role filter, relation filter, or supported opt-out once the hook is active in a backend. Output is nondeterministic, and transformation is based on the stored byte length, not Unicode characters. Applications, tests, exports, monitoring queries, and administrative tools can therefore receive corrupted text values.

The implementation is version `0.0.0`, requires superuser installation, and is non-relocatable. It includes low-level executor and tuple-destination code with a limited test surface; nulls, toasted values, nested plans, non-sequential scans, and changing PostgreSQL internals are significant risk areas. Do not preload it cluster-wide or install it in a production database. No restart is required for a backend-local load.
