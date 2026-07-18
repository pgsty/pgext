## Usage

Sources:

- [pg_py_plan_forwarding README at the reviewed commit](https://github.com/eduard-gr/pg_py_plan_forwarding/blob/74205a8f864d9cfbefaa7085887ce26debabe059/README.md)
- [pg_py_plan_forwarding.control at the reviewed commit](https://github.com/eduard-gr/pg_py_plan_forwarding/blob/74205a8f864d9cfbefaa7085887ce26debabe059/pg_py_plan_forwarding.control)
- [Empty version 0.1 install SQL](https://github.com/eduard-gr/pg_py_plan_forwarding/blob/74205a8f864d9cfbefaa7085887ce26debabe059/pg_py_plan_forwarding--0.1.sql)
- [Planner and executor hook implementation](https://github.com/eduard-gr/pg_py_plan_forwarding/blob/74205a8f864d9cfbefaa7085887ce26debabe059/pg_py_plan_forwarding.c)
- [Example Python callbacks](https://github.com/eduard-gr/pg_py_plan_forwarding/blob/74205a8f864d9cfbefaa7085887ce26debabe059/pg_py_plan_forwarding.py)
- [Build configuration](https://github.com/eduard-gr/pg_py_plan_forwarding/blob/74205a8f864d9cfbefaa7085887ce26debabe059/Makefile)

`pg_py_plan_forwarding` is optimizer-learning code that installs `planner`, `executor_start`, and `executor_end` hooks backed by embedded Python. It serializes PostgreSQL internal nodes with `nodeToString`, passes the strings to Python, and reconstructs the returned plan with `stringToNode`.

The version `0.1` install SQL is empty, so creating the extension only records it. In a disposable backend built for the source's exact Python and PostgreSQL versions, the library can be activated explicitly for a planner-only experiment:

```sql
CREATE EXTENSION pg_py_plan_forwarding;
LOAD 'pg_py_plan_forwarding';
EXPLAIN SELECT 1;
```

The example Python callback logs the parse tree and plan and returns the plan unchanged.

### Caveats

- The build hard-codes Python 3.9 headers, while initialization hard-codes `/home/vagrant/pg_py_plan_forwarding` as the module path. It is not a portable Python environment integration.
- A Python callback can return malformed or version-incompatible PostgreSQL internal-node text. Reconstructing that as a plan can crash a backend or corrupt execution semantics.
- The reviewed executor-hook fallback does not call the standard executor when no earlier hook exists, and its Python reference management contains unsafe paths. Do not run ordinary statements after activation without repairing and auditing the C code.
- This is not a stable SQL API or a production plan-management feature. PostgreSQL internal plan representations and hook signatures change between releases.
