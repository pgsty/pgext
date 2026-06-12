Source: [fsm_core PGXN README](https://github.com/Nirajkashyap/fsm/blob/main/packages/database-src/pgxn-dist/README.md), [control file](https://github.com/Nirajkashyap/fsm/blob/main/packages/database-src/pgxn-dist/fsm_core.control), [1.1.0 SQL definition](https://github.com/Nirajkashyap/fsm/blob/main/packages/database-src/pgxn-dist/fsm_core--1.1.0.sql), [example definitions README](https://github.com/Nirajkashyap/fsm/blob/main/apps/fsm-core-example/README.md).

## Usage

`fsm_core` is a finite-state-machine toolkit that stores FSM definitions, instances, transitions, queue events, and event logs inside PostgreSQL. A machine definition is loaded from JSON, instances are created by name and version, and events are sent through SQL functions with optional `pgmq` queues.

The extension is PG15+ in the local catalog. The control file requires `ltree` and `pgmq`; the upstream README also lists `pg_jsonschema` as a requirement, and the Pigsty package dependency includes `postgresql-$v-pg-jsonschema`.

### Core Tables and Types

`fsm_core` creates an enum `fsm_state_type` with `atomic`, `compound`, `parallel`, `final`, and `history`, plus tables including:

- `fsm_core.fsm_json` for loaded FSM definitions.
- `fsm_core.fsm_states` for expanded state nodes and ltree paths.
- `fsm_core.fsm_transitions` for transition rules.
- `fsm_core.fsm_instance` for running instances.
- `fsm_core.fsm_instance_lock` for advisory/concurrency state.
- `fsm_core.fsm_instance_queue_event_logs` and `fsm_core.fsm_promise_queue_event_logs` for queued event history.

### Load a Machine Definition

```sql
SELECT fsm_core.load_fsm_from_json_v2(
  json_input        := :'fsm_json'::jsonb,
  root_node_text    := 'root',
  input_fsm_type    := 'workflow',
  input_fsm_name    := 'creditCheck',
  input_fsm_version := 'v01'
);
```

`load_fsm_from_json_v2()` validates the JSON against `fsm_core.fsm_json_schema()`, caches the raw definition in `fsm_json`, then expands states and transitions. Upstream examples keep immutable version folders such as `fsm/creditCheck/v01/xstate-fsm.json` and `fsm/creditCheck/v02/xstate-fsm.json`; keep deployed versions immutable so existing instances continue against their original definition.

### Create an Instance

```sql
SELECT fsm_core.create_fsm_instance_from_name_v2(
  input_fsm_name     := 'creditCheck',
  input_fsm_version  := 'v01',
  input_fsm_context  := '{"applicant_id":"a-42"}'::jsonb,
  create_pgmq_queue  := true
);
```

The function checks that the named FSM exists, inserts an `fsm_instance`, copies transition authorization rows for that instance, and, when `create_pgmq_queue` is true, creates a `pgmq` queue named by the instance UUID and sends `initialTransition_event`.

The returned JSON includes `queue_created`, `fsm_name`, `fsm_version`, `fsm_instance_id`, `fsm_instance_context`, `send_event_result`, `message`, and `extra_message`.

### Send Events

```sql
SELECT fsm_core.send_event_to_fsm_queue_with_event_logs_v2(
  input_fsm_instance_id                 := '00000000-0000-0000-0000-000000000000'::uuid,
  input_fsm_instance_id_fsm_type         := 'workflow',
  input_fsm_instance_id_fsm_version      := 'v01',
  input_send_to_parent_queue_id          := fsm_core.pg_system_queue_uuid(),
  input_send_to_parent_queue_type        := fsm_core.pg_system_queue_type(),
  input_send_to_parent_queue_id_event_name := fsm_core.pg_system_event_name(),
  input_event_name                       := 'APPROVE',
  input_event_action_type                := 'user',
  input_event_data                       := '{"approved_by":"manager"}'::jsonb,
  input_event_delay                      := 0
);
```

This helper writes to the instance queue with `pgmq.send()` and records the event in `fsm_instance_queue_event_logs`. For nested FSM and promise flows, `send_event_to_queue_from_fsm_instance_id_v2()` dispatches to the child-FSM or promise queue helper based on `fsmtype`.

### Resolve and Step State

```sql
SELECT fsm_core.resolve_state_value_v2(
  input_json        := '{"value":"pending"}'::jsonb,
  input_fsm_name    := 'creditCheck',
  input_fsm_version := 'v01'
);

SELECT fsm_core.macrostep_v2(
  event_name        := 'APPROVE',
  input_state_value := ARRAY['pending']::text[],
  fsm_name_param    := 'creditCheck',
  fsm_version_param := 'v01'
);
```

The SQL surface also includes lower-level `microstep_v2()`, `fsm_worker_v2()`, lock helpers, archive helpers, and v1 compatibility functions. Prefer v2 entry points for new usage when both versions are present.

### Dependency and Source Caveats

- Enable `ltree` and `pgmq` before `fsm_core`; account for `pg_jsonschema` because upstream requirements and Pigsty package dependencies list it.
- `pgmq` queues are part of the async execution model, so queue names and retention should be operated like application data.
- The upstream repo has no release tag for `1.1.0`; the authoritative packaged source for this stub is the `packages/database-src/pgxn-dist` SQL/control set plus the PGXN README.
- Public docs are sparse compared with the SQL surface. Treat unlisted helper functions as internal unless the SQL definition or example app demonstrates their role.
