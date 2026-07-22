## Usage

Sources:

- [Official README](https://github.com/theforeman/postgresql-evr/blob/ff3ad993a913e403f2966a86ba9c0a751fb24357/README.md)
- [Extension control file](https://github.com/theforeman/postgresql-evr/blob/ff3ad993a913e403f2966a86ba9c0a751fb24357/evr.control)
- [Version 0.0.2 extension SQL](https://github.com/theforeman/postgresql-evr/blob/ff3ad993a913e403f2966a86ba9c0a751fb24357/evr--0.0.2.sql)

`evr` is an archived, pure-PL/pgSQL helper for decomposing RPM epoch/version/release strings into composite values. Despite the project description, it does not implement RPM version comparison, ordering operators, or an indexable EVR data type.

### Core Workflow

Install the extension and inspect how a version string is split into numeric and alphabetic segments:

```sql
CREATE EXTENSION evr;

SELECT rpmver_array('1.20beta-3');
```

`evr_trigger()` expects the target table to have columns named exactly `epoch`, `version`, `release`, and `evr`:

```sql
CREATE TABLE package_version (
  package text,
  epoch integer,
  version text,
  release text,
  evr evr_t
);

CREATE TRIGGER package_version_evr
BEFORE INSERT OR UPDATE OF epoch, version, release ON package_version
FOR EACH ROW EXECUTE FUNCTION evr_trigger();

INSERT INTO package_version(package, epoch, version, release)
VALUES ('demo', 1, '2.4beta', '3.el9');
```

### Objects

- `evr_array_item` is a composite `(n numeric, s text)` segment: numeric runs populate `n`, alphabetic runs populate `s`.
- `evr_t` stores `epoch integer`, `version evr_array_item[]`, and `release evr_array_item[]`.
- `rpmver_array(varchar)` removes non-alphanumeric separators and returns the segment array.
- `evr_trigger()` populates `NEW.evr` from the four fixed column names.
- `empty(text)`, `isalpha(char)`, `isalphanum(char)`, and `isdigit(char)` are public helper functions used by the parser.

### Caveats

Parsing is ASCII-only and discards punctuation, so it does not preserve the original version string. It also does not implement RPM's comparison rules, including special version semantics; never sort or compare packages by the textual rendering of `evr_t`.

When `version` or `release` is NULL, the trigger parses the literal string `empty` rather than producing an empty array. The generic helper names are installed in the target schema and can collide with application functions. The repository is archived, so treat this as a historical parser and implement comparison separately if exact RPM EVR semantics are required.
