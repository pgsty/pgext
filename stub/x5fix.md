## Usage

Sources:

- [x5fix control file at the reviewed commit](https://github.com/darthunix/x5fix/blob/3c28e84d98ec17a0e2cf0da7be5abb3b571334e1/x5fix.control)
- [x5fix 0.1 install SQL at the reviewed commit](https://github.com/darthunix/x5fix/blob/3c28e84d98ec17a0e2cf0da7be5abb3b571334e1/x5fix--0.1.sql)
- [x5fix C implementation at the reviewed commit](https://github.com/darthunix/x5fix/blob/3c28e84d98ec17a0e2cf0da7be5abb3b571334e1/x5fix.c)

`x5fix` 0.1 is abandoned, undocumented repair code for legacy Greenplum persistent-filespace internals. It exposes one function, `add_entry_gp_persistent_filespace_node(oid, smallint, text, smallint, text)`. There is no safe general-purpose invocation; after installation in a matching forensic environment, inspect only the registered signature:

```sql
CREATE EXTENSION x5fix;

SELECT 'add_entry_gp_persistent_filespace_node(oid,smallint,text,smallint,text)'::regprocedure;
```

### Caveats

- Do not run the repair function from an ordinary maintenance playbook. It takes Greenplum internal locks and mutates persistent catalog and shared-memory state.
- The install SQL declares the mutating function `IMMUTABLE`, which is incorrect and can permit unsafe planner assumptions.
- The C source disables WAL flushing for its tuple replacement. It also writes the primary path into the mirror-path field instead of the supplied mirror path, indicating a serious implementation defect.
- The repository has no README, tests, license, releases, or supported-version matrix. It depends on private legacy Greenplum headers and cannot be treated as a PostgreSQL-compatible extension.
- Any recovery use requires a build matched to the exact Greenplum server, verified backups, an offline reproduction, and vendor-qualified review of the source.
