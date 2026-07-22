## Usage

Sources:

- [Official extension control file](https://github.com/sroeschus/loginfo/blob/e78fb0312856ff0c4fc4000fa49f75be24ad3d67/loginfo.control)
- [Official upstream documentation](https://github.com/sroeschus/loginfo/blob/e78fb0312856ff0c4fc4000fa49f75be24ad3d67/README.md)

`loginfo` 1.0 exposes PostgreSQL logging configuration, log-file inventory, and log contents through SQL views. It is intended for database-side diagnostics and monitoring, but it can reveal sensitive statements and server paths.

### Core Workflow

```sql
CREATE EXTENSION loginfo;

SELECT * FROM log_info;
SELECT * FROM ls_log_files;
SELECT * FROM cat_log_file;
```

`log_info` reports the active logging-related settings and current log target. `ls_log_files` lists `.log` files under `log_directory`, while `cat_log_file` reads the current text log. When CSV logging is enabled, `ls_csv_files` lists CSV logs and `cat_csv_file` exposes their 23 fields as text columns.

### Configuration and Caveats

To use the CSV views, include `csvlog` in `log_destination`; the upstream README says this change requires a restart. The file-list views select by suffix rather than expanding the `log_filename` pattern. Log scans can be large and may expose credentials, personal data, query text, filesystem locations, or error details, so grant access only to trusted operators and apply restrictive predicates and retention controls.
