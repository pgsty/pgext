## Usage

Sources:

- [Official README](https://github.com/adjust/pg-ajversion/blob/502776cf0c17d7ad034a2b5bb36ba742f3e8334d/README.md)
- [Version 0.0.1 SQL API](https://github.com/adjust/pg-ajversion/blob/502776cf0c17d7ad034a2b5bb36ba742f3e8334d/ajversion--0.0.1.sql)
- [C input, output, and comparison implementation](https://github.com/adjust/pg-ajversion/blob/502776cf0c17d7ad034a2b5bb36ba742f3e8334d/src/ajversion.c)

`ajversion` 0.0.1 defines a compact, sortable version type for values with up to three numeric components. It is useful for storing and indexing simple `major.minor.patch` values, but it is not a complete Semantic Versioning implementation: prerelease identifiers, build metadata, and strict input validation are not supported.

### Core Workflow

```sql
CREATE EXTENSION ajversion;

CREATE TABLE software_release (
  package text NOT NULL,
  version ajversion NOT NULL
);

INSERT INTO software_release VALUES
  ('demo', '6.7.4'),
  ('demo', '12.5.7');

SELECT version, major(version), minor(version), patch(version)
FROM software_release
WHERE version >= '7.0.0'::ajversion
ORDER BY version;
```

The default B-tree and hash operator classes allow ordinary comparison indexes:

```sql
CREATE INDEX software_release_version_idx
ON software_release (version);
```

### Main Objects

- `ajversion`: a four-byte version value.
- `major(ajversion)`, `minor(ajversion)`, `patch(ajversion)`: extract components.
- `=`, `<>`, `<`, `<=`, `>`, `>=`: compare packed values numerically.
- `btree_ajversion_ops`, `hash_ajversion_ops`: default B-tree and hash operator classes.
- Casts from `integer`, `numeric`, `real`, and `double precision`: convert numeric input to `ajversion`.

### Parsing and Range Boundaries

The representation reserves 12 bits for the major component and 10 bits each for minor and patch. The valid ranges are therefore major `0..4095` and minor/patch `0..1023`; larger components raise an error.

Input parsing is intentionally permissive. It scans for at most three digit sequences among arbitrary nonnumeric characters, fills missing components with zero, and normalizes output to three components. For example, `'15'` becomes `15.0.0`, while text with no digits becomes `0.0.0`. Validate input separately when malformed text must be rejected, and do not use `ajversion` for SemVer precedence involving prerelease or build fields.
