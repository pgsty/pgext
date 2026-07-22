## Usage

Sources:

- [Pinned upstream README](https://bitbucket.org/neadwerx/event_manager/src/760428da0d8e7a33c24e202ca83a3e1789ad552d/README.md)
- [Pinned usage guide](https://bitbucket.org/neadwerx/event_manager/src/760428da0d8e7a33c24e202ca83a3e1789ad552d/docs/usage.md)
- [Pinned setup guide](https://bitbucket.org/neadwerx/event_manager/src/760428da0d8e7a33c24e202ca83a3e1789ad552d/docs/setup.md)
- [Pinned version 0.1 installation SQL](https://bitbucket.org/neadwerx/event_manager/src/760428da0d8e7a33c24e202ca83a3e1789ad552d/sql/event_manager--0.1.sql)

`event_manager` version `0.1` implements loosely coupled DML-trigger processing with an event queue followed by a work queue. A work-item query turns an `INSERT`, `UPDATE`, or `DELETE` event into JSON parameters; an action then runs local SQL synchronously or is processed asynchronously by the separate `event_manager` daemon.

### Core Workflow

The extension is installed by a superuser in its fixed `event_manager` schema. The following psql-oriented skeleton registers a synchronous local action; use the identifiers returned by `\gset` in the final insert:

```sql
CREATE EXTENSION event_manager;

CREATE TABLE public.orders (id integer PRIMARY KEY, note text);
CREATE TABLE public.order_events (order_id integer, event_op char(1));

INSERT INTO event_manager.tb_event_table(schema_name, table_name)
VALUES ('public', 'orders')
RETURNING event_table \gset

INSERT INTO event_manager.tb_action(query, static_parameters)
VALUES (
  'INSERT INTO public.order_events(order_id, event_op) VALUES (?order_id?::integer, ?event_op?::char(1))',
  '{}'::jsonb
)
RETURNING action \gset

INSERT INTO event_manager.tb_event_table_work_item(
  source_event_table, action, work_item_query, op, execute_asynchronously
)
VALUES (
  :'event_table',
  :'action',
  'SELECT jsonb_build_object(''order_id'', ?pk_value?, ''event_op'', ?op?) AS parameters',
  ARRAY['I']::char(1)[],
  false
);

INSERT INTO public.orders VALUES (1, 'first order');
SELECT * FROM public.order_events;
```

For asynchronous processing, start the upstream service with at least one event worker and one work worker (`-E 1 -W 1`). Remote HTTP actions cannot run synchronously.

### Important Objects

- `event_manager.tb_event_table`: registers source tables watched by generated row triggers.
- `event_manager.tb_event_table_work_item`: associates operations, an optional when-function, a parameter-producing query, an action, and synchronous/asynchronous mode.
- `event_manager.tb_action`: stores local DML or `GET`, `POST`, and `PUT` HTTP actions plus static parameters.
- `event_manager.tb_event_queue` and `event_manager.tb_work_queue`: hold event snapshots and executable action parameters.
- `event_manager.tb_setting`: persists extension GUC values and applies them to the current database.
- `?pk_value?`, `?op?`, `?OLD.column?`, `?NEW.column?`, and named session-GUC bind tokens: inject event state into work-item queries; those queries must return `jsonb` aliased as `parameters`.

### Operational Notes

The version `0.1` schema stores source primary-key values as `integer`, so it is not a generic queue for UUID, composite, or bigint keys. Synchronous actions execute in the triggering transaction; asynchronous actions run later and therefore have different rollback and visibility semantics. Failed queue entries are marked failed and are not retried automatically. If unlogged queues are selected during installation, crash recovery truncates them and they are not replicated.

Writes to `tb_action`, `tb_event_table_work_item`, and `tb_setting` can cause arbitrary SQL, outbound HTTP, trigger creation, or database-level configuration changes. Restrict those tables to trusted administrators and validate bind-token substitution because the official guide explicitly warns of SQL-injection risk. The source tree contains version `0.2` documentation and upgrade SQL, but the reviewed control default is `0.1`; do not assume `0.2` configuration-manager or worker-scaling behavior until that upgrade is installed and tested.
