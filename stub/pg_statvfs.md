## Usage

Sources:

- [Official documentation](https://github.com/michaelpq/pg_plugins/blob/ae57c1f3df697fb942f4576f873a655187193ace/pg_statvfs/README)
- [Control file](https://github.com/michaelpq/pg_plugins/blob/ae57c1f3df697fb942f4576f873a655187193ace/pg_statvfs/pg_statvfs.control)
- [Version 1.0 SQL](https://github.com/michaelpq/pg_plugins/blob/ae57c1f3df697fb942f4576f873a655187193ace/pg_statvfs/pg_statvfs--1.0.sql)
- [Privilege and path checks](https://github.com/michaelpq/pg_plugins/blob/ae57c1f3df697fb942f4576f873a655187193ace/pg_statvfs/pg_statvfs.c)

`pg_statvfs` 1.0 exposes filesystem capacity, inode availability, and mount flags through the server's statvfs call. It inspects the filesystem containing a server-side path, not a client directory or the size of an individual relation.

### Core Workflow

After installing the extension files, use a superuser session:

```sql
CREATE EXTENSION pg_statvfs;
SELECT * FROM pg_statvfs(current_setting('data_directory'));
SELECT f_frsize * f_blocks AS total_bytes,
       f_frsize * f_bavail AS available_bytes
FROM pg_statvfs(current_setting('data_directory'));
```

### Result and Access

`pg_statvfs(path text)` returns one record. `f_bsize` is the preferred block size; `f_frsize` is the unit for block counts. `f_blocks`, `f_bfree`, and `f_bavail` report total, free, and unprivileged-available blocks. `f_files`, `f_ffree`, and `f_favail` describe inode counts; `f_fsid`, `f_namemax`, and `flags` expose filesystem identity, name limits, and supported mount flags.

The C function explicitly requires superuser privileges and validates the path before asking the operating system. Missing or inaccessible paths produce an error. The extension needs no preload or restart of its own. Its control file permits relocation, but does not mark it trusted. The upstream subdirectory does not declare a complete PostgreSQL-major compatibility matrix; operating-system support and mount-flag coverage also vary.
