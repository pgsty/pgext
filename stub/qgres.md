## Usage

Sources:

- [Official extension control file (qgres.control)](https://api.pgxn.org/src/qgres/qgres-0.1.2/qgres.control)
- [Official extension SQL (qgres.sql)](https://api.pgxn.org/src/qgres/qgres-0.1.2/sql/qgres.sql)

`qgres` — Simple queue system. Use it when an application needs this specific database capability. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION qgres;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `consume(queue_id _queue.queue_id%TYPE , consumer_name _sp_consumer.consumer_name%TYPE , row_limit int DEFAULT 2^31-1)` is an extension function and returns `TABLE`.
- `consume(queue_name _queue.queue_name%TYPE , consumer_name _sp_consumer.consumer_name%TYPE , row_limit int DEFAULT 2^31-1)` is an extension function and returns `TABLE`.
- `consumer__drop(queue_name _queue.queue_name%TYPE , consumer_name _sp_consumer.consumer_name%TYPE)` is an extension function and returns `void`.
- `consumer__register(queue_name _queue.queue_name%TYPE , consumer_name _sp_consumer.consumer_name%TYPE)` is an extension function and returns `void`.
- `qgres_temp.build_add(first_arg text , call text , data_type regtype)` is an extension function and returns `void`.
- `qgres_temp.build_publish(first_arg text , call text , data_type regtype)` is an extension function and returns `void`.
- `qgres_temp.role__create(role_name name)` is an extension function and returns `void`.
- `queue__drop(queue_name _queue.queue_name%TYPE , force boolean DEFAULT false)` is an extension function and returns `void`.
- `queue__get(queue_id _queue.queue_id%TYPE)` is an extension function and returns `queue`.
- `queue__get(queue_name _queue.queue_name%TYPE)` is an extension function and returns `queue`.
- `queue__get_id(queue_name _queue.queue_name%TYPE)` is an extension function and returns `int`.
- `queue_entry(bytea bytea DEFAULT NULL , jsonb jsonb DEFAULT NULL , text text DEFAULT NULL)` is an extension function and returns `queue_entry`.
- `queue_entry` is an extension-defined type.
- `queue_type` is an extension-defined type.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.2`.
- Install the confirmed extension dependencies first: `plpgsql`, `citext`.
- The control file marks the extension as relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
