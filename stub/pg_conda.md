## Usage

Sources:

- [Official README](https://github.com/JeanChristopheMorinPerso/pg_conda/blob/bb2ee4f348de550a76599ab85bfcb65ebad2601c/README.md)
- [Type and function implementation](https://github.com/JeanChristopheMorinPerso/pg_conda/blob/bb2ee4f348de550a76599ab85bfcb65ebad2601c/src/lib.rs)
- [Extension control file](https://github.com/JeanChristopheMorinPerso/pg_conda/blob/bb2ee4f348de550a76599ab85bfcb65ebad2601c/pg_conda.control)
- [Build compatibility declaration](https://github.com/JeanChristopheMorinPerso/pg_conda/blob/bb2ee4f348de550a76599ab85bfcb65ebad2601c/Cargo.toml)

`pg_conda` provides the `condaversion` type for storing and comparing Conda ecosystem version strings with Rattler's Conda ordering rules. Use it for package metadata and dependency analysis after testing its comparison semantics against the spellings present in your data.

### Core Workflow

```sql
CREATE EXTENSION pg_conda;

CREATE TABLE package_release (
    package text NOT NULL,
    version condaversion NOT NULL
);

INSERT INTO package_release VALUES
    ('demo', '1.5.0'),
    ('demo', '2.0.0'),
    ('demo', '1.0.0.post1');

SELECT package, min(version), max(version)
FROM package_release
GROUP BY package;

SELECT conda_is_post('1.0.0.post1'::condaversion),
       conda_segments('1.2.3.alpha1'::condaversion);
```

Input is parsed with Rattler's Conda version parser, while output preserves the original source spelling.

### Objects

- `condaversion` supports equality, ordering, hashing, B-tree operators, and casts to and from `text`.
- `min(condaversion)` and `max(condaversion)` use the extension's ordering implementation.
- `conda_is_dev(condaversion)` and `conda_is_post(condaversion)` identify development and post-release components.
- `conda_has_epoch(condaversion)` and `conda_has_local(condaversion)` identify epoch and local components.
- `conda_segments(condaversion) → text[]` returns parsed version components.
- `conda_major(condaversion) → bigint` and `conda_minor(condaversion) → bigint` return numeric components when available, otherwise `NULL` for unsupported component shapes.

### Comparison Caveat

In the reviewed `1.0.1` source, equality and hashing compare the normalized Rattler version, but total ordering compares that version and then the preserved source text. Consequently, equivalent spellings such as `1.1.0` and `1.01.0` can compare equal while sorting differently. This violates the consistency expected by PostgreSQL B-tree comparison and can affect B-tree or unique indexes and the choice returned by `min(condaversion)` or `max(condaversion)`.

Canonicalize stored spellings or avoid relying on those index and aggregate semantics until the behavior is fixed upstream. Test equality, ordering, uniqueness, and upgrade behavior with your real version corpus.

### Compatibility

The control file is not trusted and the extension is not relocatable, so installation normally requires a privileged role and objects remain in the installation schema. The current Cargo feature set declares PostgreSQL 13 through 17 and defaults to PostgreSQL 17; verify a package built specifically for the target major version.

The reviewed tree uses `pgrx` 0.18.1 and Rattler 0.36.0. Invalid input raises a PostgreSQL error. Pin the extension and dependency versions because changes in Rattler's parser or ordering can change database-visible comparison behavior.
