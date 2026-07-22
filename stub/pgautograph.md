## Usage

Sources:

- [Official Rust implementation](https://github.com/nevindev/pgautograph/blob/a8c59f407f80fb5ed4f1ee8e3ab15745aab03100/src/lib.rs)
- [Official graph-query generator](https://github.com/nevindev/pgautograph/blob/a8c59f407f80fb5ed4f1ee8e3ab15745aab03100/src/queries/mod.rs)
- [Official Rust package manifest](https://github.com/nevindev/pgautograph/blob/a8c59f407f80fb5ed4f1ee8e3ab15745aab03100/Cargo.toml)

`pgautograph` 0.0.0 is an unfinished pgrx experiment that reads foreign keys from the `public` schema and constructs a proposed property-graph DDL string. Its public function returns foreign-key metadata and writes the generated string to the PostgreSQL log; it does not create a graph.

### Inspect Foreign Keys

Create the extension in a disposable database and call its main function:

```sql
CREATE EXTENSION pgautograph;

SELECT * FROM list_foreign_keys();
SELECT hello_pgautograph();
```

`list_foreign_keys()` returns `source_table`, `source_pk`, `source_columns`, `target_table`, and `target_column`. It then logs a string beginning with `CREATE PROPERTY GRAPH graph` at LOG level. Capture and review that log text manually if evaluating the prototype.

### Hard-Coded Model Assumptions

The catalog query only considers foreign keys whose source table is in `public`. It uses unqualified relation names, expects every source table to have a primary key, handles only one column from each foreign-key row, and applies simple English capitalization and singularization. Generated vertex definitions have no properties, while edge generation infers directed or undirected forms from the number of matching references.

### Work-in-Progress Boundary

The function does not return or execute the graph DDL. It unwraps NULL catalog values and SPI results, so a foreign-key table without a primary key can raise a Rust panic. Hash-set iteration also makes generated clause order nondeterministic. The included greeting test expects text different from the function's actual exclamation-mark result, indicating the 0.0.0 tree is not release-gated even by its own test.

Use it only as source material or a disposable prototype. Before any real use, add schema-qualified catalog handling, composite-key support, deterministic output, identifier quoting, NULL/error handling, a return value for generated DDL, and tests against the exact PostgreSQL property-graph syntax. Do not grant it production superuser installation merely to obtain a catalog query that can be written safely in SQL.
