## Usage

Sources:

- [Official documentation](https://github.com/fukanchik/probackup_ctl/blob/e03cb514d05c82830379d152924bfeaf81e42ef4/README.md)
- [Extension control file](https://github.com/fukanchik/probackup_ctl/blob/e03cb514d05c82830379d152924bfeaf81e42ef4/probackup_ctl.control)
- [Official repository](https://github.com/fukanchik/probackup_ctl)

`probackup_ctl` Inspect and operate pg_probackup catalogs, including filesystem, SFTP, and S3 storage.

### Enablement

Install the files for the intended server, then create `probackup_ctl` in the target database:

```sql
CREATE EXTENSION probackup_ctl;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
select * from probackup.log(catalog_id=>2,pg_instance=>'dba1',backup_id=>(select backup_id from probackup.show(catalog_id=>2)));
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `probackup.backup` | FUNCTION | Callable function from the reviewed install surface. |
| `probackup.show` | FUNCTION | Callable function from the reviewed install surface. |
| `probackup.catalogs` | TABLE | Extension-owned table; account for its data in backup and upgrades. |
| `probackup.register_catalog` | FUNCTION | Callable function from the reviewed install surface. |
| `probackup.delete` | FUNCTION | Callable function from the reviewed install surface. |
| `probackup.log` | FUNCTION | Callable function from the reviewed install surface. |
| `probackup.s3_config` | TABLE | Extension-owned table; account for its data in backup and upgrades. |
| `probackup.sftp_config` | TABLE | Extension-owned table; account for its data in backup and upgrades. |

### Operations and Boundaries

- The extension fixes or creates schema objects under `probackup`; include them in privilege and backup review.
