## Usage

Sources:

- [Official pre-alpha documentation](https://github.com/NAlexPear/pg_branch/blob/d400e373139ed550db0ccbeba569f224007bee7b/README.md)
- [ProcessUtility hook implementation](https://github.com/NAlexPear/pg_branch/blob/d400e373139ed550db0ccbeba569f224007bee7b/src/hooks.rs)
- [Extension control file](https://github.com/NAlexPear/pg_branch/blob/d400e373139ed550db0ccbeba569f224007bee7b/pg_branch.control)

`pg_branch` is a pre-alpha proof of concept that accelerates `CREATE DATABASE ... WITH TEMPLATE` by replacing PostgreSQL's normal file copy with BTRFS copy-on-write snapshots. It targets disposable Linux test clusters, not production databases.

### Core Workflow

Before installing the extension, place the cluster on BTRFS and use the upstream initialization procedure to convert the required data-directory subdirectories into BTRFS subvolumes. Then load the hook in the administrative session and create a database from a template:

```sql
CREATE EXTENSION pg_branch;

CREATE DATABASE branch_db
WITH TEMPLATE source_db;
```

Creating the extension loads a backend-local `ProcessUtility` hook; upstream does not document `shared_preload_libraries`. Run the command from the session where the extension library has been loaded and verify the snapshot layout independently.

### Safety Boundaries

Version `0.0.1` invokes BTRFS filesystem operations and changes database-file creation semantics. The README explicitly says the project is a proof of concept and not ready for production, and lists broader option and filesystem support as future work.

A layout mismatch, unsupported `CREATE DATABASE` option, concurrent session, crash, tablespace, backup tool, or PostgreSQL-major change can invalidate assumptions and damage the cluster. Use only disposable data after verifying recoverable backups, mount options, subvolumes, ownership, crash cleanup, tablespaces, and template-connection behavior. Do not infer support for ZFS, XFS, cluster branching, or production failover from the project name.
