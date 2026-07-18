## Usage

Sources:

- [Upstream pre-alpha documentation](https://github.com/NAlexPear/pg_branch/blob/d400e373139ed550db0ccbeba569f224007bee7b/README.md)
- [ProcessUtility hook implementation](https://github.com/NAlexPear/pg_branch/blob/d400e373139ed550db0ccbeba569f224007bee7b/src/hooks.rs)
- [Extension control file](https://github.com/NAlexPear/pg_branch/blob/d400e373139ed550db0ccbeba569f224007bee7b/pg_branch.control)

`pg_branch` is a pre-alpha proof of concept that accelerates `CREATE DATABASE ... WITH TEMPLATE` by taking BTRFS copy-on-write snapshots. It targets Linux clusters whose data directory has first been converted into the required BTRFS subvolume layout.

```sql
CREATE EXTENSION pg_branch;
CREATE DATABASE branch_db WITH TEMPLATE source_db;
```

Creating the extension loads a backend-local utility hook; upstream documents no shared-preload setup. The tool changes database-file creation semantics and invokes filesystem operations, so a bug or layout mismatch can damage a cluster. Use only on disposable data after verifying backups, subvolumes, mount options, concurrent sessions, crash cleanup, tablespaces, and PostgreSQL-major compatibility. It is not production-ready.
