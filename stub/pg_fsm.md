## Usage

Sources:

- [extensions/pg_fsm/pg_fsm.control](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_fsm/pg_fsm.control)
- [README.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/README.md)
- [extensions/pg_fsm/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_fsm/Cargo.toml)
- [extensions/pg_fsm/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_fsm/src/lib.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/LICENSE)
- [extensions/pg_fsm/src/machine.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_fsm/src/machine.rs)
- [extensions/pg_fsm/src/binding.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_fsm/src/binding.rs)
- [extensions/pg_fsm/src/transition.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_fsm/src/transition.rs)

`pg_fsm` 0.3.0 defines finite state machines in `pgfsm`, with transition guards, actions, table bindings and history.

### Core Workflow

```sql
CREATE EXTENSION pg_fsm;
SELECT pgfsm.create_machine('order_flow', 'draft', 'Order workflow');
SELECT pgfsm.add_state('order_flow', 'submitted');
SELECT pgfsm.add_transition('order_flow', 'draft', 'submitted', 'submit');
SELECT * FROM pgfsm.list_machines();
```

### Operational Boundaries

The control does not require superuser-only installation, but object-creation privileges still apply. No preload is required. Define states and transitions, then bind an application state column or call the transition API. Guards/actions execute database expressions, so restrict their authors. `pg_fsm.enabled` disables enforcement when false; `pg_fsm.disable_default_notify` suppresses the implicit notification. Version 0.3.0 adds idempotent seeding with `if_not_exists`. This project is distinct from the extension whose canonical name is `fsm`. This is an unsupported proof of concept under Matroid Source Available License 1.0. APIs may change. Version 0.3.0 is the new upgrade baseline: earlier 0.2.0 installations require a rehearsed data migration/recreation, not ordinary ALTER EXTENSION UPDATE. Back up data and dependencies before following that destructive upstream path.
