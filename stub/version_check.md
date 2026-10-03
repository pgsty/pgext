## Usage

Sources:

- [version_check.control](https://github.com/VladlenPopolitov/version_control/blob/844aa0cd66391a9848f1242c1287a243d9d44baa/version_check.control)
- [README.md](https://github.com/VladlenPopolitov/version_control/blob/844aa0cd66391a9848f1242c1287a243d9d44baa/README.md)
- [version_check.c](https://github.com/VladlenPopolitov/version_control/blob/844aa0cd66391a9848f1242c1287a243d9d44baa/version_check.c)
- [version_check--5.0.sql](https://github.com/VladlenPopolitov/version_control/blob/844aa0cd66391a9848f1242c1287a243d9d44baa/version_check--5.0.sql)
- [Makefile](https://github.com/VladlenPopolitov/version_control/blob/844aa0cd66391a9848f1242c1287a243d9d44baa/Makefile)

`version_check` supplies boolean predicates for PostgreSQL major versions. The reviewed 5.0 SQL installs predicates for versions 13 through 17, while the current C source explicitly builds only against PostgreSQL 17.

### Core Workflow

```sql
CREATE EXTENSION version_check;
SELECT getversion13(), getversion14(), getversion15(), getversion16(), getversion17();
```

### Operational Boundaries

Install as a superuser; no preload or restart is required by this implementation. In the PG17 build only the version-17 predicate is true. These are compiled predicates, not a general runtime compatibility probe. The README contains older filenames and comments; the control, SQL and compiler guard establish this exact version. No explicit license was found.
