## Usage

Sources:

- [Extension control file](https://github.com/zhongjn/mtree/blob/40773bc60f4516015feda337dc0082d1e4466657/mtree.control)
- [Version 0.0.1 install SQL](https://github.com/zhongjn/mtree/blob/40773bc60f4516015feda337dc0082d1e4466657/mtree--0.0.1.sql)
- [Access-method implementation](https://github.com/zhongjn/mtree/blob/40773bc60f4516015feda337dc0082d1e4466657/mtree.c)

mtree 0.0.1 installs an index access-method handler named mtree. It does not install any data type, operator, or operator class, so it cannot create a useful user index by itself.

### Installation inspection

Use only on a disposable server built for the target PostgreSQL major:

```sql
CREATE EXTENSION mtree;
SELECT amname, amtype FROM pg_am WHERE amname = 'mtree';
```

The query confirms registration of the access method; it is not an end-to-end index example.

### Caveats

- No compatible operator class is supplied. A normal CREATE INDEX command has no supported type and operator semantics to use.
- The implementation is a 2020 fork of PostgreSQL GiST internals and calls private backend, buffer, locking, and write-ahead-log interfaces. Such code is tied closely to a PostgreSQL major version.
- Upstream provides no usage documentation, regression tests, crash-recovery evidence, compatibility matrix, or license.
- Treat mtree as incomplete index infrastructure. Do not place production data behind it without an independent code audit plus concurrency, crash, recovery, and upgrade testing.
