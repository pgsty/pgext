## Usage

Sources:

- [README at the reviewed commit](https://github.com/sangli00/pg_checkdestroy/blob/f78520270a5a411d2604056a8c0dd2473556e505/README.md)
- [Extension control file](https://github.com/sangli00/pg_checkdestroy/blob/f78520270a5a411d2604056a8c0dd2473556e505/pg_checkdestroy.control)
- [Empty extension install SQL](https://github.com/sangli00/pg_checkdestroy/blob/f78520270a5a411d2604056a8c0dd2473556e505/pg_checkdestroy--1.0.sql)
- [Hook implementation](https://github.com/sangli00/pg_checkdestroy/blob/f78520270a5a411d2604056a8c0dd2473556e505/pg_checkdestroy.c)

`pg_checkdestroy` is an experimental hook module that attempts to suppress table `DROP`, `TRUNCATE`, and unqualified `DELETE` statements. It installs parser, executor, and utility hooks only when loaded through `shared_preload_libraries`; the extension SQL itself is empty.

### Configuration

```conf
shared_preload_libraries = 'pg_checkdestroy'
```

After restarting PostgreSQL, register the extension and confirm that its session switch is enabled:

```sql
CREATE EXTENSION pg_checkdestroy;
SHOW pg_checkdestroy.work;
```

`pg_checkdestroy.work` defaults to `on`, but it is a `PGC_USERSET` setting and can be disabled in any session with sufficient access to connect.

### Caveats

- Upstream explicitly says the code is not fully tested and advises against online/production use. The source dates from 2018 and depends on old PostgreSQL internal hook APIs.
- The hook suppresses selected commands by skipping normal execution rather than raising a clear policy error. A client can therefore receive misleading completion behavior.
- This is not a security boundary: `pg_checkdestroy.work` is user-settable, coverage is based on particular parse/plan shapes, and privileged users can bypass or unload the module.
- The implementation chains several global hooks and assumes per-query state that may interact badly with errors, nested statements, prepared plans, or other hook extensions. Use native privileges, ownership controls, and tested auditing instead of relying on this module to protect data.
