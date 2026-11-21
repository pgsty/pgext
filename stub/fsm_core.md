## Usage

Sources:

- [fsm_core-v1.2.0](https://github.com/pgfsm/fsm/releases/tag/fsm_core-v1.2.0)
- [README.md](https://github.com/pgfsm/fsm/blob/fsm_core-v1.2.0/packages/database-src/README.md)
- [fsm_core-1.2.0.zip](https://github.com/pgfsm/fsm/releases/download/fsm_core-v1.2.0/fsm_core-1.2.0.zip)


`fsm_core` 1.2.0 stores finite-state-machine definitions, instances, transitions and event history in PostgreSQL. Its SQL interface loads validated JSON, creates instances by name and version, and submits events through `pgmq` queues.

### Install and Configure

The release requires PostgreSQL 15 or later and was tested on 15–18. Required extensions are `ltree` 1.2+, `pgmq` 1.4.4+ and `pg_cron` 1.5+. Add `pg_cron` to the existing `shared_preload_libraries`, restart PostgreSQL if that setting changes, and install into the database selected by `cron.database_name`.

After installing the matching extension files, a superuser can enable the fixed, nonrelocatable `fsm_core` schema:

```sql
CREATE EXTENSION fsm_core CASCADE;
SELECT fsm_core.version();
```

Installation creates `fsm_*` roles and the `fsm_schedule_all_pending` scheduler job, which runs every five seconds. Inspect role memberships, queue permissions and the scheduler before granting application access. `pg_jsonschema` is optional: if its validation helper exists in `fsm_core`, the loader reports schema problems but does not reject a definition on that check alone.

### Load a Definition

Validate the machine with the upstream compiler and set the psql variable fsm_json to that JSON before this call. The 1.2.0 signature has no input_fsm_type argument:

```sql
SELECT fsm_core.load_fsm_from_json_v2(
  json_input        := :'fsm_json'::jsonb,
  root_node_text    := 'root',
  input_fsm_name    := 'creditCheck',
  input_fsm_version := 'v01'
);
```

Concurrent loads of the same name and version are serialized by an advisory transaction lock. Identical JSON returns the cached definition; different JSON under an existing name and version raises an error. Use a new version for a changed machine. The optional input_dependent_children argument records child-machine dependencies.

### Create an Instance and Send Events

```sql
SELECT fsm_core.create_fsm_instance_from_name_v2(
  input_fsm_name    := 'creditCheck',
  input_fsm_version := 'v01',
  input_fsm_context := '{"applicant_id":"a-42"}'::jsonb,
  create_pgmq_queue := true
) AS creation_result
\gset
SELECT :'creation_result'::jsonb AS creation_status;
SELECT :'creation_result'::jsonb ->> 'fsm_instance_id' AS fsm_instance_id
\gset
```

The returned JSON includes `fsm_instance_id`, `queue_created`, `send_event_result`, `message` and `extra_message`. Check these before submitting work: queue creation and initial-event failures can be reported in the result. Initial context is shallow-merged with the definition's static context, with supplied object keys taking precedence.

For this top-level machine, submit an event declared by its definition:

```sql
SELECT fsm_core.send_event_to_fsm_queue_with_event_logs_v2(
  input_fsm_instance_id            := :'fsm_instance_id'::uuid,
  input_fsm_instance_id_fsm_type   := 'fsm',
  input_fsm_instance_id_fsm_version := 'v01',
  input_event_name                 := 'APPROVE',
  input_event_action_type          := 'user',
  input_event_data                 := '{"approved_by":"manager"}'::jsonb,
  input_event_delay                := 0
);
```

The helper queues the event and writes event history. The old parent-queue parameters are absent from this release's signature. Queuing an event is not proof that a worker has processed it; monitor scheduler activity, queue backlog and instance state.

### Objects and Upgrade Boundaries

`fsm_json`, `fsm_states` and `fsm_transitions` hold definitions; `fsm_instance` holds instances; `fsm_instance_queue_event_logs` records queue events. `resolve_state_value_v2` and `macrostep_v2` expose lower-level state resolution and transition logic. Prefer the verified v2 entry points and restrict access to security-definer functions and underlying tables.

The release notes describe sequential upgrade scripts guarded by `require_version`, but the reviewed 1.2.0 ZIP contains the full install script and no version-to-version upgrade scripts. Preserve populated databases and obtain a verified upgrade artifact before upgrading an existing installation. Do not rerun the full install or assume `ALTER EXTENSION ... UPDATE` has an available upgrade path.
