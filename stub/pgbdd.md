## Usage

Sources:

- [Official upstream README](https://github.com/utwente-db/dubio/blob/b5210d373e24752a5d7e48cb2abdeb1a132df210/pgbdd/README.txt)
- [Official extension control file (pgbdd.control)](https://github.com/utwente-db/dubio/blob/b5210d373e24752a5d7e48cb2abdeb1a132df210/pgbdd/pgbdd.control)
- [Official extension SQL (pgbdd--0.0.1.sql)](https://github.com/utwente-db/dubio/blob/b5210d373e24752a5d7e48cb2abdeb1a132df210/pgbdd/pgbdd--0.0.1.sql)

`pgbdd` — See file(s) example-dictionary.sql(.log) for an example using the 'dictionary' type. Use it when application data needs this type, domain, or its operators. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pgbdd;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `add(d dictionary, vardefs text)` is an extension function and returns `dictionary`.
- `alg_bdd(alg cstring, expression cstring)` is an extension function and returns `bdd`.
- `alternatives(dict dictionary, var cstring)` is an extension function and returns `text`.
- `and_accum(internal bdd, next bdd)` is an extension function and returns `bdd`.
- `bdd_bytea_in(expression bytea)` is an extension function and returns `bdd`.
- `bdd_equal(lhs_bdd bdd,rhs_bdd bdd)` is an extension function and returns `BOOLEAN`.
- `bdd_equiv(lhs_bdd bdd,rhs_bdd bdd)` is an extension function and returns `BOOLEAN`.
- `bdd_fast_equiv(lhs_bdd bdd,rhs_bdd bdd)` is an extension function and returns `BOOLEAN`.
- `bdd_in(expression cstring)` is an extension function and returns `bdd`.
- `bdd_out(dict bdd)` is an extension function and returns `cstring`.
- `contains(bdd bdd, var cstring, val integer)` is an extension function and returns `BOOLEAN`.
- `debug(dict dictionary)` is an extension function and returns `text`.
- `del(d dictionary, vardefs text)` is an extension function and returns `dictionary`.
- `dictionary_in(dictname cstring)` is an extension function and returns `dictionary`.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
