## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/fuzzy_logic/fuzzy_logic-1.1.0/README.md)
- [Official extension control file (lukasiewicz_logic.control)](https://api.pgxn.org/src/fuzzy_logic/fuzzy_logic-1.1.0/lukasiewicz_logic.control)
- [Official extension SQL (lukasiewicz_logic--1.1.0.sql)](https://api.pgxn.org/src/fuzzy_logic/fuzzy_logic-1.1.0/sql/lukasiewicz_logic--1.1.0.sql)

`lukasiewicz_logic` — Fuzzy logic =========== This extension provides basic logical operators (conjunction, disjunction, implication and negation) for three basic fuzzy logics - Łukasiewicz, Gödel and product. For the Łukasiewicz logic, there are also operators for weak conjunction and disjunction. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION lukasiewicz_logic;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `lukasiewicz_conjunction(a fuzzy_boolean, b fuzzy_boolean)` is an extension function and returns `fuzzy_boolean`.
- `lukasiewicz_disjunction(a fuzzy_boolean, b fuzzy_boolean)` is an extension function and returns `fuzzy_boolean`.
- `lukasiewicz_negation(a fuzzy_boolean)` is an extension function and returns `fuzzy_boolean`.
- `lukasiewicz_residuum(a fuzzy_boolean, b fuzzy_boolean)` is an extension function and returns `fuzzy_boolean`.
- `lukasiewicz_weak_conjunction(a fuzzy_boolean, b fuzzy_boolean)` is an extension function and returns `fuzzy_boolean`.
- `lukasiewicz_weak_disjunction(a fuzzy_boolean, b fuzzy_boolean)` is an extension function and returns `fuzzy_boolean`.
- `fuzzy_boolean` is an extension-defined domain.

### Requirements and Caveats

- The reviewed control file declares default version `1.1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
