## Usage

Sources:

- [Official upstream README](https://github.com/seanwevans/pg_os/blob/b822d1f3a83657eac385d2004421236d5e8b1d4d/README.md)
- [Official extension control file (pg_os.control)](https://github.com/seanwevans/pg_os/blob/b822d1f3a83657eac385d2004421236d5e8b1d4d/pg_os.control)
- [Official extension SQL (pg_os--1.0.sql)](https://github.com/seanwevans/pg_os/blob/b822d1f3a83657eac385d2004421236d5e8b1d4d/pg_os--1.0.sql)

`pg_os` — pg_os is a PostgreSQL extension that models operating-system concepts entirely inside the database. Use it when an application needs this specific database capability. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION pg_os;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `acquire_semaphore(process_id INTEGER, sem_name TEXT)` is an extension function and returns `VOID`.
- `allocate_memory(user_id INTEGER, process_id INTEGER, segment_size INTEGER)` is an extension function and returns `VOID`.
- `allocate_page(thread_id INTEGER)` is an extension function and returns `BIGINT`.
- `assign_role_to_user(user_id INTEGER, role_id INTEGER)` is an extension function and returns `VOID`.
- `change_file_permissions(user_id INTEGER, file_id INTEGER, new_perms TEXT)` is an extension function and returns `VOID`.
- `check_mail(user_id INTEGER)` is an extension function and returns `SETOF`.
- `check_permission(p_user_id INTEGER, p_resource_type TEXT, p_action TEXT)` is an extension function and returns `BOOLEAN`.
- `cleanup_terminated_processes(timeout_interval INTERVAL DEFAULT '1 hour')` is an extension function and returns `VOID`.
- `create_file(user_id INTEGER, filename TEXT, parent_id INTEGER, is_dir BOOLEAN DEFAULT FALSE)` is an extension function and returns `INTEGER`.
- `create_mutex(mutex_name TEXT)` is an extension function and returns `VOID`.
- `create_role(role_name TEXT)` is an extension function and returns `INTEGER`.
- `create_semaphore(sem_name TEXT, initial_count INTEGER, max_val INTEGER)` is an extension function and returns `VOID`.
- `create_user(name TEXT)` is an extension function and returns `INTEGER`.
- `enqueue_io_request(device_name TEXT, request_type TEXT, data TEXT)` is an extension function and returns `VOID`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- Install the confirmed extension dependencies first: `plpgsql`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
