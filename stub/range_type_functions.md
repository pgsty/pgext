## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/range_type_functions/range_type_functions-0.0.4/README.md)
- [Official extension control file (range_type_functions.control)](https://api.pgxn.org/src/range_type_functions/range_type_functions-0.0.4/range_type_functions.control)
- [Official extension SQL (range_type_functions.sql)](https://api.pgxn.org/src/range_type_functions/range_type_functions-0.0.4/sql/range_type_functions.sql)
- [Current official source repository](https://github.com/decibel/range_type_functions)

`range_type_functions` — This extension serves two purposes: 1. Extend the capabilities of range functions, with an eye toward moving the most useful of those functions into the core. 2. Facilitate back-porting of functions from newer versions of PostgreSQL. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION range_type_functions;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `element_range_comp(element anyelement, range anyrange)` is an extension function and returns `smallint`.
- `get_bound_expr(range anyrange, literal anyelement)` is an extension function and returns `text`.
- `get_bounds_condition_expr(range anyrange, text default 'x')` is an extension function and returns `text`.
- `get_collation_expr(range anyrange)` is an extension function and returns `text`.
- `get_lower_bound_condition_expr(range anyrange, text default 'x')` is an extension function and returns `text`.
- `get_subtype_element_expr(range anyrange, text default 'x')` is an extension function and returns `text`.
- `get_upper_bound_condition_expr(range anyrange, text default 'x')` is an extension function and returns `text`.
- `is_singleton(range anyrange)` is an extension function and returns `boolean`.
- `to_range(elem anyelement, range anyrange)` is an extension function and returns `anyrange`.
- `to_range(low anyelement, high anyelement, bounds text, range anyrange)` is an extension function and returns `anyrange`.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.4`.
- The control file marks the extension as relocatable.
- The former `moat` GitHub URL is unavailable; the surviving upstream source repository is linked above.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
