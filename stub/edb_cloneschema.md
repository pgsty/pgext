## Usage

Sources:

- [Official EDB Clone Schema documentation](https://www.enterprisedb.com/docs/epas/latest/database_administration/14_edb_clone_schema/)
- [Official setup guide](https://www.enterprisedb.com/docs/epas/latest/database_administration/14_edb_clone_schema/setting_up_edb_clone_schema/)
- [Official local-copy guide](https://www.enterprisedb.com/docs/epas/latest/database_administration/14_edb_clone_schema/copying_a_schema/)

`edb_cloneschema` copies a schema and its database objects locally or between EPAS databases, using foreign servers, scheduler jobs, and parallel clone workers.

### Prerequisites and Enablement

Install and preload the EPAS worker packages, restart, then create all required extensions in each participating database:

```ini
shared_preload_libraries = 'parallel_clone,edb_job_scheduler'
```

```sql
CREATE EXTENSION postgres_fdw SCHEMA public;
CREATE EXTENSION dblink SCHEMA public;
CREATE EXTENSION edb_job_scheduler;
CREATE EXTENSION dbms_job;
CREATE EXTENSION parallel_clone;
CREATE EXTENSION edb_cloneschema;
CREATE TRUSTED LANGUAGE plperl;
SELECT edb_util.create_clone_log_dir();
```

The source/target user mappings must authenticate roles with the privileges required to read and create every cloned object.

### Copy a Schema

Define the foreign server and mapping documented for local or remote mode, then call the appropriate `localcopyschema`/`localcopyschema_nb` or `remotecopyschema`/`remotecopyschema_nb` function. The nonblocking variants schedule work through EDB Job Scheduler; use `process_status_from_log` to follow the status file.

### Operational Boundaries

Schema copy can create large transactions, WAL volume, locks, indexes, foreign keys, and background-worker demand. Tune `work_mem`, `maintenance_work_mem`, `max_worker_processes`, checkpoint settings, WAL size, and `max_locks_per_transaction` for the planned object set. Protect credentials and log files, test unsupported object types, and verify post-copy ownership, privileges, dependencies, and data consistency before cutover.

