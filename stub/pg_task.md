## Usage

> Sources: [pg_task upstream README](https://github.com/RekGRpth/pg_task), [PGXN pg_task](https://pgxn.org/dist/pg_task/), [local metadata](../db/extension.csv).

`pg_task` is a background-worker scheduler for running SQL asynchronously at a planned time. Upstream documents PostgreSQL, Greenplum, and Greengage support.

Enable the worker at server start, then create the extension in the database that will own the task table:

```conf
shared_preload_libraries = 'pg_task'
```

```sql
CREATE EXTENSION pg_task;
```

### Schedule Tasks

Schedule work by inserting rows into the configured task table, which defaults to `public.task` in database `postgres` unless changed with GUCs.

```sql
-- Run SQL immediately
INSERT INTO task (input) VALUES ('SELECT now()');

-- Run SQL after 5 minutes
INSERT INTO task (plan, input) VALUES (now() + '5 min'::interval, 'SELECT now()');

-- Run SQL at a specific time
INSERT INTO task (plan, input) VALUES ('2029-07-01 12:51:00', 'SELECT now()');

-- Repeat SQL every 5 minutes
INSERT INTO task (repeat, input) VALUES ('5 min', 'SELECT now()');

-- Exceptions are caught and written to the error column
INSERT INTO task (input) VALUES ('SELECT 1/0');

-- Limit concurrent tasks in a group.
-- max = 1 means one task at a time for this group.
INSERT INTO task ("group", max, input) VALUES ('billing', 1, 'SELECT refresh_billing_cache()');

-- Run SQL on a remote database
INSERT INTO task (input, remote) VALUES ('SELECT now()', 'user=user host=host');
```

### Task Table

The task table is meant to be user-visible. Upstream notes that users may add columns or partition it.

Key columns:

| Name | Type | Default | Description |
|------|------|---------|-------------|
| id | bigserial | autoincrement | Primary key |
| parent | bigint | pg_task.id | Parent task id |
| plan | timestamptz | statement_timestamp() | Planned start time |
| start | timestamptz | | Actual start time |
| stop | timestamptz | | Actual stop time |
| active | interval | 1 hour | Period after plan time when task is active |
| live | interval | 0 sec | Max lifetime of background worker |
| repeat | interval | 0 sec | Auto repeat interval |
| timeout | interval | 0 sec | Allowed time for task run |
| count | int | 0 | Max task count before worker exit |
| max | int | 0 | Max concurrent tasks in group; negative values mean pause between tasks in milliseconds |
| pid | int | | Process id executing task |
| state | enum | PLAN | PLAN, TAKE, WORK, DONE, STOP |
| delete | bool | true | Auto delete when output and error are null |
| drift | bool | false | Compute next repeat by stop time |
| header | bool | true | Show column headers in output |
| group | text | 'group' | Task grouping name |
| input | text | | SQL command(s) to execute |
| output | text | | Received result(s) |
| error | text | | Caught error |
| remote | text | | Remote database connection string |

### Configuration

Key settings:

| Name | Type | Default | Description |
|------|------|---------|-------------|
| pg_task.delete | bool | true | Auto delete completed tasks |
| pg_task.drift | bool | false | Compute repeat by stop time |
| pg_task.header | bool | true | Show column headers in task output |
| pg_task.string | bool | true | Quote only strings in output |
| pg_task.count | int | 0 | Default maximum number of tasks per worker before exit |
| pg_task.fetch | int | 100 | Number of task rows fetched at once |
| pg_task.limit | int | 1000 | Limit task rows at once |
| pg_task.max | int | 0 | Default max concurrent tasks in group |
| pg_task.run | int | 2147483647 | Maximum concurrently executing tasks in work |
| pg_task.sleep | int | 1000 | Check tasks every N milliseconds |
| pg_task.active | interval | 1 hour | Period after plan time when a task remains active for execution |
| pg_task.live | interval | 0 sec | Maximum lifetime of a worker process |
| pg_task.repeat | interval | 0 sec | Default repeat interval |
| pg_task.timeout | interval | 0 sec | Default task timeout |
| pg_task.data | text | postgres | Database name for tasks table |
| pg_task.user | text | postgres | User name for tasks table |
| pg_task.schema | text | public | Schema name for tasks table |
| pg_task.table | text | task | Table name for tasks table |
| pg_task.json | json | `[{"data":"postgres"}]` | Multi-database configuration |
| pg_task.id | bigint | 0 | Current task id, read-only session setting |

### Multi-Database Configuration

To run tasks on multiple databases, configure via JSON:

```conf
pg_task.json = '[{"data":"database1"},{"data":"database2","user":"username2"},{"data":"database3","schema":"schema3"}]'
```

If the specified database, user, schema or table does not exist, `pg_task` will create them.

The local metadata marks this package as `headless`, so it needs `shared_preload_libraries` rather than a user-facing SQL-only install path. Validate background scheduling behavior on the target PostgreSQL version before relying on it for critical jobs.
