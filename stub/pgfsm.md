## Usage

Sources:

- [Official README](https://github.com/michelp/pgfsm/blob/ee79a0e12ef73474a07b34fc36b3b2214e85b42e/README.md)
- [Extension SQL](https://github.com/michelp/pgfsm/blob/ee79a0e12ef73474a07b34fc36b3b2214e85b42e/pgfsm--0.0.1.sql)
- [Official worked example](https://github.com/michelp/pgfsm/blob/ee79a0e12ef73474a07b34fc36b3b2214e85b42e/sql/test.sql)

pgfsm is a compact SQL example for storing finite-state-machine definitions and instances in PostgreSQL. It enforces that each instance uses a declared state and that state updates follow an allowed destination. It is instructional code, not a general workflow engine.

### Core Workflow

Install the extension into a dedicated `fsm` schema, define transitions, create an instance, and advance it by transition label:

```sql
CREATE SCHEMA fsm;
CREATE EXTENSION pgfsm WITH SCHEMA fsm;
SET search_path = fsm, public;

INSERT INTO transition(name, from_state, transition, to_state) VALUES
  ('turnstile', 'locked',   'coin', 'unlocked'),
  ('turnstile', 'unlocked', 'push', 'locked');

INSERT INTO machine(name, state)
VALUES ('turnstile', 'locked');

SELECT * FROM do_transition(1, 'coin');
```

Every usable state must appear as a `from_state` because the machine table has a foreign key to that column.

### Important Objects

- `transition` stores machine name, current state, transition label, and destination state.
- `machine` stores state-machine instances.
- `transitions_for` returns the transition row available from an instance's current state.
- `states_for` returns the declared source states for a machine definition.
- `do_transition` applies a matching label and returns the updated instance.
- `check_valid_state_update` and `check_valid_state_insert` back the validation triggers.

### Design Limits

The primary key on `transition(name, from_state)` permits only one outgoing transition row per state, so it cannot model branching by label. A direct update may choose the permitted destination without presenting the label; revoke direct table updates if transition labels matter. There is no history, callback execution, optimistic version, ownership model, or workflow-definition migration support. Extend the schema deliberately before using the pattern for application state.
