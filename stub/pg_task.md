

## Usage

> [pg_task: PostgreSQL job scheduler](https://github.com/RekGRpth/pg_task)

`pg_task` allows executing any SQL command at any specific time in the background asynchronously. It works with PostgreSQL, Greenplum and Greengage.

First, add to `postgresql.conf`:

```conf
shared_preload_libraries = 'pg_task'
```

Then schedule tasks by inserting into the `task` table:

```sql
-- Run SQL immediately
INSERT INTO task (input) VALUES ('SELECT now()');

-- Run SQL after 5 minutes
INSERT INTO task (plan, input) VALUES (now() + '5 min'::INTERVAL, 'SELECT now()');

-- Run SQL at a specific time
INSERT INTO task (plan, input) VALUES ('2029-07-01 12:51:00', 'SELECT now()');

-- Repeat SQL every 5 minutes
INSERT INTO task (repeat, input) VALUES ('5 min', 'SELECT now()');

-- Exceptions are caught and written to the error column
INSERT INTO task (input) VALUES ('SELECT 1/0');

-- Limit concurrent tasks in a group
INSERT INTO task (group, max, input) VALUES ('group', 1, 'SELECT now()');

-- Run SQL on a remote database
INSERT INTO task (input, remote) VALUES ('SELECT now()', 'user=user host=host');
```


## Task Table Columns

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
| max | int | 0 | Max concurrent tasks in group |
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

You may add any needed columns and/or make partitions on this table.


## Configuration (GUCs)

Key settings:

| Name | Type | Default | Description |
|------|------|---------|-------------|
| pg_task.data | text | postgres | Database name for tasks table |
| pg_task.user | text | postgres | User name for tasks table |
| pg_task.schema | text | public | Schema name for tasks table |
| pg_task.table | text | task | Table name for tasks table |
| pg_task.sleep | int | 1000 | Check tasks every N milliseconds |
| pg_task.delete | bool | true | Auto delete completed tasks |
| pg_task.drift | bool | false | Compute repeat by stop time |
| pg_task.repeat | interval | 0 sec | Default repeat interval |
| pg_task.timeout | interval | 0 sec | Default task timeout |
| pg_task.max | int | 0 | Default max concurrent tasks in group |
| pg_task.run | int | 2147483647 | Max concurrent tasks in work |
| pg_task.json | json | [{"data":"postgres"}] | Multi-database configuration |

### Multi-Database Configuration

To run tasks on multiple databases, configure via JSON:

```conf
pg_task.json = '[{"data":"database1"},{"data":"database2","user":"username2"},{"data":"database3","schema":"schema3"}]'
```

If the specified database, user, schema or table does not exist, `pg_task` will create them.
