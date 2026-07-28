## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/graph_component/graph_component-1.0.1/README.md)
- [Official extension control file (graph_component.control)](https://api.pgxn.org/src/graph_component/graph_component-1.0.1/graph_component.control)
- [Official extension SQL (graph_component--1.0.0.sql)](https://api.pgxn.org/src/graph_component/graph_component-1.0.1/graph_component--1.0.0.sql)

`graph_component` — It is very hard to compute graph components on pure PostgreSQL. With this extension you can do this very efficiently. Extension does this on pointers using a minimal amount of RAM. This allows you to build components of hundreds of thousands of vertices based on many millions of pairs in seconds. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION graph_component;

SELECT
  get_component(graph_components(array[1,2,3,4,5]))
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `get_component(state graph_component_hashmap)` is an extension function and returns `SETOF int4`.
- `get_component_id(state graph_component_hashmap)` is an extension function and returns `TABLE`.
- `graph_components_final(state internal)` is an extension function and returns `graph_component_hashmap`.
- `graph_components_step_arr(state internal, vertex int[])` is an extension function and returns `internal`.
- `graph_components` is an aggregate exposed by the extension.
- `graph_component_hashmap` is an extension-defined type.

### Requirements and Caveats

- The reviewed control file declares default version `1.0.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
