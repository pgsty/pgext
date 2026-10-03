## Usage

Sources:

- [Official PGXN1.1.0 distribution](https://pgxn.org/dist/fsm_core/1.1.0/)
- [1.1.0 README](https://api.pgxn.org/src/fsm_core/fsm_core-1.1.0/README.md)
- [1.1.0 control](https://api.pgxn.org/src/fsm_core/fsm_core-1.1.0/fsm_core.control)
- [1.1.0 SQL](https://api.pgxn.org/src/fsm_core/fsm_core-1.1.0/fsm_core--1.1.0.sql)
- [1.1.0 metadata](https://api.pgxn.org/src/fsm_core/fsm_core-1.1.0/META.json)

`fsm_core` is a finite-state-machine toolkit that stores FSM definitions, instances, transitions, and event logs inside PostgreSQL. A machine definition is loaded from JSON, instances are created by name and version, and events are sent through SQL functions with optional `pgmq` queues.

This document describes the PGXN 1.1.0 distribution, which requires PostgreSQL 15 or later, `ltree` 1.2 or later and `pgmq` 1.4.4 or later. The fixed schema is `fsm_core`; installation requires a superuser. No preload or restart is required by this SQL extension. Current repository migrations are a separate source and may expose a different API.

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

`load_fsm_from_json_v2()` checks JSON against `fsm_core.fsm_json_schema()`, expands states and transitions, then caches the raw definition in `fsm_json`. In 1.1.0, JSON-schema violations produce a NOTICE and loading continues; the schema-validation exception is commented out. Validate definitions before loading rather than relying on this check to reject every invalid definition. Keep deployed definitions and version identifiers stable so existing instances continue against their original definition. Set the psql variable fsm_json to your validated machine JSON before running the example.

### Create an Instance

```sql
SELECT fsm_core.create_fsm_instance_from_name_v2(
  input_fsm_name     := 'creditCheck',
  input_fsm_version  := 'v01',
  input_fsm_context  := '{"applicant_id":"a-42"}'::jsonb,
  create_pgmq_queue  := true
) AS creation_result
\gset
SELECT :'creation_result'::jsonb AS creation_status;
SELECT :'creation_result'::jsonb ->> 'fsm_instance_id' AS fsm_instance_id
\gset
```

The PGXN 1.1.0 function checks that the named FSM exists, inserts an `fsm_instance`, and copies transition authorization rows. When `create_pgmq_queue` is true, it attempts to create a `pgmq` queue named by the instance UUID and send `initialTransition_event`. Queue creation and initial-event failures are caught and reported in the returned JSON. Before sending further events, require `queue_created` to be true and inspect `send_event_result`, `message` and `extra_message` to confirm the initial event succeeded. The psql example retains the returned `fsm_instance_id` for the next call.

### Send Events

```sql
SELECT fsm_core.send_event_to_fsm_queue_with_event_logs_v2(
  input_fsm_instance_id                 := :'fsm_instance_id'::uuid,
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

### Dependencies and Operation

Enable `ltree` and `pgmq` before installing `fsm_core`. The release README also requires `pg_jsonschema` 0.3.3 or later, although the control and META dependency lists omit it. The JSON loader calls `fsm_core.jsonschema_validation_errors`, which is not defined by the distribution's SQL script; verify that the documented JSON-schema helper is available in that schema before loading definitions. Installing a dependency into a different schema alone does not provide that qualified function.

Queued events persist as application data. Supplying `create_pgmq_queue => true` requests the per-instance queue and its initial event; check the returned status before proceeding. A consumer still needs to process queued work. Sending an event does not by itself guarantee that an asynchronous worker has executed the transition. Review queue retention, consumers and permissions together.

Inspect function and table grants before exposing machine creation or arbitrary event submission to application roles. The SQL surface also contains legacy v1 and lower-level helpers; use the verified v2 entry points for this release.
