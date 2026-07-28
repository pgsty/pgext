## Usage

Sources:

- [Official upstream README](https://github.com/aau-daisy/solvedb/blob/8a15559b5a08747b2217d9146e3dd122379d4de1/README.md)
- [Official extension control file (solverapi.control)](https://github.com/aau-daisy/solvedb/blob/8a15559b5a08747b2217d9146e3dd122379d4de1/SolverAPI/solverapi.control)
- [Official extension SQL (solverapi--1.2.sql)](https://github.com/aau-daisy/solvedb/blob/8a15559b5a08747b2217d9146e3dd122379d4de1/SolverAPI/solverapi--1.2.sql)

`solverapi` — SolveDB: A PostgreSQL-based DBMS for optimization applications. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION solverapi;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `sl_build_dst_ctr(arg sl_solver_arg, vsout sl_viewsql_out, ctr_nr int)` is an extension function and returns `sl_viewsql_dst`.
- `sl_build_dst_ctr_union(arg sl_solver_arg, vsout sl_viewsql_out, ctr_type text)` is an extension function and returns `sl_viewsql_dst`.
- `sl_build_dst_obj(arg sl_solver_arg, vsout sl_viewsql_out)` is an extension function and returns `sl_viewsql_dst`.
- `sl_build_dst_values(arg sl_solver_arg, vsout sl_viewsql_out, cast_to text DEFAULT 'text')` is an extension function and returns `sl_viewsql_dst`.
- `sl_build_out(arg sl_solver_arg)` is an extension function and returns `sl_viewsql_out`.
- `sl_build_out_array1subst(arg sl_solver_arg, par_nr int DEFAULT 1)` is an extension function and returns `sl_viewsql_out`.
- `sl_build_out_arrayNsubst(arg sl_solver_arg, par_pos int[])` is an extension function and returns `sl_viewsql_out`.
- `sl_build_out_defcols(arg sl_solver_arg, colvalues text[][], base sl_viewsql_out DEFAULT NULL)` is an extension function and returns `sl_viewsql_out`.
- `sl_build_out_func1subst(arg sl_solver_arg, func text)` is an extension function and returns `sl_viewsql_out`.
- `sl_build_out_funcNmap(arg sl_solver_arg, base sl_viewsql_out, funcs text[])` is an extension function and returns `sl_viewsql_out`.
- `sl_build_out_funcNsubst(arg sl_solver_arg, funcs text[])` is an extension function and returns `sl_viewsql_out`.
- `sl_build_out_join(arg sl_solver_arg, base sl_viewsql_out, sql text, join_id_col text)` is an extension function and returns `sl_viewsql_out`.
- `sl_build_out_joinvalues(arg sl_solver_arg, sql text, col_varnr text, col_value text)` is an extension function and returns `sl_viewsql_out`.
- `sl_build_out_rename(arg sl_solver_arg, base sl_viewsql_out, col_type sl_attribute_kind, col_alias text)` is an extension function and returns `sl_viewsql_out`.

### Requirements and Caveats

- The reviewed control file declares default version `1.2`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
