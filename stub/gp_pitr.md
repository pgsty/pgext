## Usage

Sources:

- [gpdb-doc/markdown/ref_guide/modules/gp_pitr.html.md](https://github.com/open-gpdb/gpdb/blob/c6e63e9a53d01ed6ea3a93e5e280f38e8b9d808b/gpdb-doc/markdown/ref_guide/modules/gp_pitr.html.md)
- [gpcontrib/gp_pitr/gp_pitr--1.1.sql](https://github.com/open-gpdb/gpdb/blob/c6e63e9a53d01ed6ea3a93e5e280f38e8b9d808b/gpcontrib/gp_pitr/gp_pitr--1.1.sql)
- [gpcontrib/gp_pitr/Makefile](https://github.com/open-gpdb/gpdb/blob/c6e63e9a53d01ed6ea3a93e5e280f38e8b9d808b/gpcontrib/gp_pitr/Makefile)
- [LICENSE](https://github.com/open-gpdb/gpdb/blob/c6e63e9a53d01ed6ea3a93e5e280f38e8b9d808b/LICENSE)
- [gpcontrib/gp_pitr/gp_pitr.control](https://github.com/open-gpdb/gpdb/blob/c6e63e9a53d01ed6ea3a93e5e280f38e8b9d808b/gpcontrib/gp_pitr/gp_pitr.control)

`gp_pitr` exposes cluster-wide recovery helpers for Greenplum 6. Version 1.1 provides a combined archiver view and functions used by recovery tooling. Installing the module alone does not configure backups or WAL archiving.

### Core Workflow

```sql
CREATE EXTENSION gp_pitr;
SELECT * FROM gp_stat_archiver ORDER BY gp_segment_id;
```

### Operational Boundaries

`gp_stat_archiver` combines the coordinator and segment archiver statistics, with `gp_segment_id` identifying each source; the coordinator uses -1. The installation script grants public read access to this view.

`gp_create_restore_point(text)` creates a named restore point across the cluster and returns segment identifiers with their recovery LSNs. `gp_switch_wal()` switches WAL files and returns the segment identifier, LSN, and file name. Both functions execute on the coordinator and the installation script revokes their execution privilege from `PUBLIC`; they are privileged interfaces for recovery tooling.

Register the module as a superuser. The C library loads on demand without preload configuration, and the extension is not relocatable. Version 1.1 includes an upgrade script from 1.0. These interfaces depend on Greenplum kernel facilities and do not establish compatibility with vanilla PostgreSQL. A complete recovery procedure still needs the corresponding cluster backup and WAL archive.
