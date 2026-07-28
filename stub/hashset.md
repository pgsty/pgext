## Usage

Sources:

- [Official upstream README](https://github.com/tvondra/hashset/blob/3ef8368d43eb6508de4f2b0e0a5a1d5e497bd44b/README.md)
- [Official extension control file (hashset.control)](https://github.com/tvondra/hashset/blob/3ef8368d43eb6508de4f2b0e0a5a1d5e497bd44b/hashset.control)
- [Official extension SQL (hashset--0.0.1.sql)](https://github.com/tvondra/hashset/blob/3ef8368d43eb6508de4f2b0e0a5a1d5e497bd44b/hashset--0.0.1.sql)

`hashset` — This PostgreSQL extension implements hashset, a data structure (type) providing a collection of unique integer items with fast lookup. Use it when application data needs this type, domain, or its operators. Upstream explicitly says it is not production-ready.

### Core Workflow

```sql
CREATE EXTENSION hashset;

SELECT hashset_add(NULL, 1); -- {1}
SELECT hashset_add('{NULL}', 1); -- {1,NULL}
SELECT hashset_add('{1}', NULL); -- {1,NULL}
SELECT hashset_add('{1}', 1); -- {1}
SELECT hashset_add('{1}', 2); -- {1,2}
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `hashset_add(int4hashset, int)` is an extension function and returns `int4hashset`.
- `hashset_capacity(int4hashset)` is an extension function and returns `bigint`.
- `hashset_cardinality(int4hashset)` is an extension function and returns `bigint`.
- `hashset_cmp(int4hashset, int4hashset)` is an extension function and returns `integer`.
- `hashset_collisions(int4hashset)` is an extension function and returns `bigint`.
- `hashset_contains(int4hashset, int)` is an extension function and returns `boolean`.
- `hashset_difference(int4hashset, int4hashset)` is an extension function and returns `int4hashset`.
- `hashset_eq(int4hashset, int4hashset)` is an extension function and returns `boolean`.
- `hashset_ge(int4hashset, int4hashset)` is an extension function and returns `boolean`.
- `hashset_gt(int4hashset, int4hashset)` is an extension function and returns `boolean`.
- `hashset_hash(int4hashset)` is an extension function and returns `integer`.
- `hashset_intersection(int4hashset, int4hashset)` is an extension function and returns `int4hashset`.
- `hashset_le(int4hashset, int4hashset)` is an extension function and returns `boolean`.
- `hashset_lt(int4hashset, int4hashset)` is an extension function and returns `boolean`.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as relocatable.
- Upstream explicitly says the project is not production-ready.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
