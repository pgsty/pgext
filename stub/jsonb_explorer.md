## Usage

Sources:

- [Official extension SQL](https://github.com/erthalion/jsonb_explorer/blob/515b434872ded15b64c88bd4101c061afb15698e/jsonb_explorer--1.0.sql)
- [Official C implementation](https://github.com/erthalion/jsonb_explorer/blob/515b434872ded15b64c88bd4101c061afb15698e/jsonb_explorer.c)
- [Official tree formatter implementation](https://github.com/erthalion/jsonb_explorer/blob/515b434872ded15b64c88bd4101c061afb15698e/jsonb_explorer_utils.c)

`jsonb_explorer` is a small C extension that renders a JSONB document as an indented text tree. Version `1.0` exposes only two functions, and one has a source-level return-type defect; use the working tree renderer only in disposable inspection sessions.

### Core Workflow

Install the extension and render a document with `jsonb_tree(jsonb)`:

```sql
CREATE EXTENSION jsonb_explorer;

SELECT jsonb_tree(
  '{"service":{"name":"api","ports":[8080,8081]}}'::jsonb
);
```

The result is presentation text containing tree-drawing characters, keys, scalar values, and array-size annotations. It is intended for human inspection, not for stable machine parsing.

### Important Objects

- `jsonb_tree(jsonb)` returns the formatted tree as `text` and rejects NULL input because the function is declared strict.
- `jsonb_paths(jsonb)` is declared to return `text[][]`, but the C implementation constructs and returns a single `text` datum. That ABI/type mismatch makes the function unsafe to call.

### Operational Caveats

Do not call `jsonb_paths(jsonb)` unless a maintained downstream build has corrected both the SQL declaration and implementation. The repository contains no substantive API documentation or compatibility matrix, and its last reviewed commit is from 2018. The formatter also uses PostgreSQL internal JSONB iterator APIs, so a binary must be rebuilt and tested for the exact PostgreSQL major version. Prefer built-in JSONB operators, `jsonb_each`, `jsonb_array_elements`, or a maintained client-side formatter for application logic.
