## Usage

Sources:

- [Upstream README](https://github.com/ozum/boolean_cascaded/blob/88502bb17dd17b0bb590eb29c328343010d64a7a/README.md)
- [Version 2.0 control file](https://github.com/ozum/boolean_cascaded/blob/88502bb17dd17b0bb590eb29c328343010d64a7a/boolean_cascaded.control)
- [Version 2.0 install SQL](https://github.com/ozum/boolean_cascaded/blob/88502bb17dd17b0bb590eb29c328343010d64a7a/sql/boolean_cascaded--2.0.sql)
- [Conversion tests](https://github.com/ozum/boolean_cascaded/blob/88502bb17dd17b0bb590eb29c328343010d64a7a/test/001-conversions.sql)

boolean_cascaded 2.0 is a pure-SQL integer-backed type plus trigger functions for propagating disabled state through parent and child tables. Version 2.0 encodes local state and inherited false counts in one integer; the older composite representation described in parts of the README no longer matches the install SQL.

### Inspect the value semantics

```sql
CREATE EXTENSION boolean_cascaded WITH SCHEMA public;
SELECT 0::public.boolean_cascaded::boolean AS active,
       1::public.boolean_cascaded::boolean AS inactive;
```

Zero casts to true and ordinary nonzero values cast to false. Boolean input first becomes a negative sentinel that the supplied BEFORE trigger is expected to normalize while preserving inherited state.

### Caveats

- The type, several operators, and trigger functions are created in public even when another extension schema is requested. Review global name collisions and ownership before installation.
- Do not update the encoded integer directly. Negative sentinels and decimal-position counts are implementation details, and bypassing the full trigger set can leave parent and child state inconsistent.
- Cascading is synchronous row-trigger work. Large fan-out trees can amplify writes, locks, deadlocks, and replication volume; circular or cross-schema designs require particular care.
- Trigger arguments name tables and columns dynamically, but relation names are not schema-qualified. A caller's search path can select an unintended same-named relation.
- The branch for a non-default status column builds a JSON key from the literal text p_status_col instead of the variable's value, so advertised custom column names may not be updated.
- The project was last changed in 2016 and publishes no current PostgreSQL compatibility matrix. Exercise inserts, moves between parents, concurrent updates, deletes, cycles, dump and restore, and trigger ordering before use.
