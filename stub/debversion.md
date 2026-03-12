

## Usage

> [debversion: Debian version number type for PostgreSQL](https://github.com/ATIX-AG/postgresql-debversion-evr)

The `debversion` extension provides Debian package version comparison functionality, implementing the same sorting logic used by `dpkg`.

```sql
CREATE EXTENSION debversion;
```

### Data Type

The `debversion` type stores Debian package version strings and compares them according to the Debian versioning specification (epoch:upstream-revision format).

```sql
CREATE TABLE packages (
    name    text,
    version debversion
);

INSERT INTO packages VALUES ('foo', '1.0-1'), ('foo', '2.0-1'), ('foo', '1.0-2');

SELECT * FROM packages ORDER BY version;
```

### Version Comparison

```sql
SELECT '1.60-26+b1'::debversion < '1.60+git20161116.90da8a0-1'::debversion;
```

### Operators

Standard comparison operators are supported: `=`, `<>`, `<`, `>`, `<=`, `>=`.

The comparison algorithm mirrors `dpkg --compare-versions`, ensuring results identical to standard Debian package management utilities.
