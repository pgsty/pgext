## Usage

Sources:

- [Upstream README](https://github.com/danolivo/anotherSelfJoin/blob/b4461d76e07fb83e1a18b7e6cee8953b6d0c27c4/README.md)
- [Version 0.1 install SQL](https://github.com/danolivo/anotherSelfJoin/blob/b4461d76e07fb83e1a18b7e6cee8953b6d0c27c4/optpaths--0.1.sql)
- [Planner-hook implementation](https://github.com/danolivo/anotherSelfJoin/blob/b4461d76e07fb83e1a18b7e6cee8953b6d0c27c4/optpaths.c)

optpaths 0.1 is an experimental planner hook that attempts to remove a narrow class of redundant self joins when uniqueness conditions allow it.

### Isolated evaluation

Load the library at server start, restart, and test only in a disposable database:

```conf
shared_preload_libraries = 'optpaths'
```

```sql
CREATE EXTENSION optpaths;
EXPLAIN SELECT p.* FROM test AS p
JOIN (SELECT * FROM test WHERE b = 0) AS q ON p.a = q.a;
SET remove_self_joins = off;
EXPLAIN SELECT p.* FROM test AS p
JOIN (SELECT * FROM test WHERE b = 0) AS q ON p.a = q.a;
```

Compare plans and results with the hook enabled and disabled before considering a workload.

### Caveats

- The install SQL creates a generic table named test and inserts two rows. Installation can collide with an existing table and modifies the chosen extension schema.
- CREATE EXTENSION does not load the planner library. Preloading and a server restart are required for predictable activation.
- The hook mutates internal planner structures and covers only limited join shapes. There is no current major-version compatibility matrix or regression suite proving equivalent results.
- remove_self_joins and log_level are user-settable settings. Verbose logging can flood clients or server logs.
- Test exact query semantics, prepared plans, parallel plans, partitioning, row security, and upgrades in isolation. Do not enable this prototype cluster-wide without review.
