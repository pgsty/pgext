## Usage

Sources:

- [Official documentation](https://github.com/brunoenten/pg_fsm/blob/f781e13a985330b45eb78195cc1df18eacd84dd3/README.md)
- [Extension control file](https://github.com/brunoenten/pg_fsm/blob/f781e13a985330b45eb78195cc1df18eacd84dd3/fsm.control)
- [Official repository](https://github.com/brunoenten/pg_fsm)

`fsm` Trigger-enforced finite state machines for application tables.

### Enablement

Install the files for the intended server, then create `fsm` in the target database:

```sql
CREATE EXTENSION fsm;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
-- Example business table
CREATE TABLE public.orders (
  id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  customer_id bigint NOT NULL,
  total numeric(12,2) NOT NULL
);

-- Attach FSM columns/triggers
SELECT fsm.add_to_table('public.orders'::regclass);

-- Define transitions from start -> pending -> paid|cancelled
SELECT fsm.add_transition('public.orders'::regclass, 'start',   'create',  'pending');
SELECT fsm.add_transition('public.orders'::regclass, 'pending', 'pay',     'paid');
SELECT fsm.add_transition('public.orders'::regclass, 'pending', 'cancel',  'cancelled');

-- Insert row (FSM columns must remain default on INSERT)
INSERT INTO public.orders (customer_id, total) VALUES (10, 99.99);

-- Append one event to move start -> pending
UPDATE public.orders SET new_event = 'create' WHERE id = 1;

-- Append another event to move pending -> paid
UPDATE public.orders SET new_event = 'pay' WHERE id = 1;

SELECT id, fsm_previous_state, fsm_current_state
FROM public.orders
WHERE id = 1;
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `fsm.event` | TYPE | User-facing data type created by the extension. |
| `fsm.add_transition` | FUNCTION | Callable function from the reviewed install surface. |
| `fsm.add_to_table` | FUNCTION | Callable function from the reviewed install surface. |
| `fsm.machines` | TABLE | Extension-owned table; account for its data in backup and upgrades. |
| `fsm.run_machine` | FUNCTION | Callable function from the reviewed install surface. |
| `fsm.add_callback` | FUNCTION | Callable function from the reviewed install surface. |
| `fsm.append_event` | FUNCTION | Callable function from the reviewed install surface. |
| `fsm.events_callbacks` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- The verified PostgreSQL-major evidence covers 9.6, 10, 11, 12, 13, 14, 15, 16, 17, 18; do not infer unlisted majors.
- The extension fixes or creates schema objects under `fsm`; include them in privilege and backup review.
