## Usage

Sources:

- [Official README](https://github.com/mediait/lpathtree/blob/6b1fd0acd1d4a698d9f4b04d59eb4e144a1eda15/README.md)
- [Official extension SQL](https://github.com/mediait/lpathtree/blob/6b1fd0acd1d4a698d9f4b04d59eb4e144a1eda15/lpathtree--1.0.sql)

`lpathtree` is a fork of PostgreSQL's ltree module that uses `/` between labels. It is useful for materialized paths that need slash syntax or label characters such as `{`, `}`, `!`, and `*`, at the cost of removing the original `ltxtquery` feature.

### Core Workflow

Store paths in the `lpathtree` type and add a GiST index for ancestor, descendant, and path-query operators.

```sql
CREATE EXTENSION lpathtree;

CREATE TABLE taxonomy (path lpathtree PRIMARY KEY);
CREATE INDEX taxonomy_path_gist ON taxonomy USING gist (path);

INSERT INTO taxonomy(path) VALUES
  ('Top/Science'),
  ('Top/Science/Astronomy'),
  ('Top/Hobbies');

SELECT path
FROM taxonomy
WHERE 'Top/Science'::lpathtree @> path;

SELECT path
FROM taxonomy
WHERE path ~ 'Top/*/Astronomy'::lpathquery;
```

### Main Objects

- `lpathtree` and `lpathquery` are the path and pattern types.
- `@>` and `<@` test ancestor and descendant relationships; `~` matches an `lpathquery`; `||` concatenates paths.
- `nlevel`, `subpath`, `index`, and `lca` expose depth, slicing, subpath search, and lowest-common-ancestor operations.
- The extension provides B-tree ordering and GiST operator classes for scalar paths and path arrays.

### Compatibility Boundary

This code was forked from an older PostgreSQL ltree implementation. Slash-separated values are not interchangeable with dot-separated `ltree` values, and `ltxtquery` is deliberately absent. Treat migration from `ltree` as a data and query rewrite, and validate this old C extension on the exact PostgreSQL major version before production use.
