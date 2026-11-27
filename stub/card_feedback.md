## Usage

Sources:

- [Official source](https://github.com/sanchitachaurasia/CS-631/blob/79c0152b5fcd67a2e847c3b128d0dbdef503f642/README.md)
- [Official source](https://github.com/sanchitachaurasia/CS-631/blob/79c0152b5fcd67a2e847c3b128d0dbdef503f642/project/extension/card_feedback.control)
- [Official source](https://github.com/sanchitachaurasia/CS-631/blob/79c0152b5fcd67a2e847c3b128d0dbdef503f642/project/extension/card_feedback--1.0.sql)
- [Official source](https://github.com/sanchitachaurasia/CS-631/blob/79c0152b5fcd67a2e847c3b128d0dbdef503f642/project/extension/card_feedback.c)
- [Official source](https://github.com/sanchitachaurasia/CS-631/blob/79c0152b5fcd67a2e847c3b128d0dbdef503f642/project/extension/Makefile)

`card_feedback` 1.0 exposes inspection and reset functions for a research implementation of cardinality feedback in PostgreSQL 16.13. It depends on this project's modified planner and feedback storage APIs. The module cannot provide those features when installed on an otherwise stock server.

### Inspect and Reset

Install the matching patched server and module in an isolated research environment. Create the extension as a superuser:

```sql
CREATE EXTENSION card_feedback;
SELECT * FROM card_feedback_entries();
SELECT card_feedback_reset();
```

`card_feedback_entries()` returns the signature, original row estimate, latest observed row count, exponentially smoothed row count, observation count and update time. `card_feedback_reset()` clears the feedback store and returns how many entries were removed; call it only when resetting the experiment is intended.

### Privileges and Interpretation

Both SQL functions are volatile and parallel-unsafe. The installation script does not revoke default PUBLIC execution, and the C wrappers contain no explicit superuser check. Restrict function execution before allowing unrelated users to connect, especially the reset function.

The SQL extension does not itself enable the planner changes or create the underlying feedback store; those belong to the patched core. An empty result is not proof that a workload was observed. Compare estimates and actual execution statistics under the project's documented benchmark configuration. No shared preload requirement is declared for this inspection wrapper.

### Maintenance

The extension is relocatable, but its ABI depends on the accompanying core patch. Rebuild and qualify both together when changing PostgreSQL versions. The reviewed tree supplies version 1.0 without an SQL upgrade chain; preserve experiment inputs and results separately. A license for the project-specific additions was not identified in the reviewed files.
