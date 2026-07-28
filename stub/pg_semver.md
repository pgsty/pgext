## Usage

Sources:

- [Official upstream README](https://github.com/eendroroy/pg_semver/blob/42cddda9d5f36161841cd1fba5cd9fdf9a704cdb/README.md)
- [Official extension control file (pg_semver.control)](https://github.com/eendroroy/pg_semver/blob/42cddda9d5f36161841cd1fba5cd9fdf9a704cdb/pg_semver.control)
- [Official extension SQL (pg_semver--0.0.1.sql)](https://github.com/eendroroy/pg_semver/blob/42cddda9d5f36161841cd1fba5cd9fdf9a704cdb/pg_semver--0.0.1.sql)

`pg_semver` — **Version** Data type (SEMVER) for postgresql. Use it when application data needs this type, domain, or its operators. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_semver;

SELECT PG_SEMVER_CMP('1.0.0-alpha.1', '1.0.0-alpha.2');
 pg_semver_cmp
---------------
            -1
(1 row)

SELECT PG_SEMVER_CMP('0.0.1', '0.0.1');
 pg_semver_cmp
---------------
             0
(1 row)

SELECT PG_SEMVER_CMP('0.0.2', '0.0.1');
 pg_semver_cmp
---------------
             1
(1 row)
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `hash_ver(semver)` is an extension function and returns `int`.
- `pg_semver_bump(semver, int)` is an extension function and returns `semver`.
- `pg_semver_car(semver, semver)` is an extension function and returns `boolean`.
- `pg_semver_cmp(semver, semver)` is an extension function and returns `int`.
- `pg_semver_eq(semver, semver)` is an extension function and returns `boolean`.
- `pg_semver_ge(semver, semver)` is an extension function and returns `boolean`.
- `pg_semver_gt(semver, semver)` is an extension function and returns `boolean`.
- `pg_semver_in(cstring)` is an extension function and returns `semver`.
- `pg_semver_le(semver, semver)` is an extension function and returns `boolean`.
- `pg_semver_lt(semver, semver)` is an extension function and returns `boolean`.
- `pg_semver_ncar(semver, semver)` is an extension function and returns `boolean`.
- `pg_semver_ne(semver, semver)` is an extension function and returns `boolean`.
- `pg_semver_nsat(semver, semver)` is an extension function and returns `boolean`.
- `pg_semver_out(semver)` is an extension function and returns `cstring`.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
