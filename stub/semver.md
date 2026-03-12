

## Usage

> [semver: semantic versioning data type](https://github.com/theory/pg-semver)

The `semver` extension provides a data type implementing [Semantic Versioning 2.0.0](https://semver.org/spec/v2.0.0.html).

```sql
CREATE EXTENSION semver;

SELECT '1.2.1'::semver;
SELECT '1.2.0'::semver > '1.2.0-b1'::semver;  -- true (prerelease < release)
```

### Operators

| Operator | Description | Example | Result |
|----------|-------------|---------|--------|
| `=` | Equal | `'1.2.0'::semver = '1.2.00'::semver` | `t` |
| `<>` | Not equal | `'1.2.0'::semver <> '1.2.00'::semver` | `f` |
| `<` | Less than | `'3.4.0-b1'::semver < '3.4.0'::semver` | `t` |
| `<=` | Less or equal | `'3.4.0-b1'::semver <= '3.4.0'::semver` | `t` |
| `>` | Greater than | `'3.4.0-b1'::semver > '3.4.0'::semver` | `f` |
| `>=` | Greater or equal | `'3.4.0-b1'::semver >= '3.4.0'::semver` | `f` |

### Functions

| Function | Description | Example | Result |
|----------|-------------|---------|--------|
| `to_semver(text)` | Permissive parse | `to_semver('1.0')` | `1.0.0` |
| `is_semver(text)` | Validate format | `is_semver('1.2.0')` | `true` |
| `semver(text)` | Strict cast | `semver('1.2.1')` | `1.2.1` |
| `get_semver_major(semver)` | Major version | `get_semver_major('4.2.0')` | `4` |
| `get_semver_minor(semver)` | Minor version | `get_semver_minor('4.2.0')` | `2` |
| `get_semver_patch(semver)` | Patch version | `get_semver_patch('4.2.0')` | `0` |
| `get_semver_prerelease(semver)` | Prerelease part | `get_semver_prerelease('2.1.0-b2+bfb13')` | `b2` |

Supports casts from `text`, `numeric`, `real`, `double precision`, `integer`, `bigint`, `smallint`.

### Range Type

The `semverrange` type supports standard range operators:

```sql
SELECT '1.0.5'::semver <@ '[1.0.0, 2.0.0)'::semverrange;  -- true
```

### Aggregate Functions

`MIN(semver)` and `MAX(semver)` are supported. Btree and hash indexes are available.
