## Usage

Sources:

- [Official upstream README](https://github.com/opmdg/opm-core/blob/ae89f025407ab144e1e30abd7d6580f258945d61/README.md)
- [Official extension control file (opm_core.control)](https://github.com/opmdg/opm-core/blob/ae89f025407ab144e1e30abd7d6580f258945d61/pg/opm_core.control)
- [Official extension SQL (opm_core--2.5--2.6.sql)](https://github.com/opmdg/opm-core/blob/ae89f025407ab144e1e30abd7d6580f258945d61/pg/opm_core--2.5--2.6.sql)

`opm_core` — Central module of the OPM suite. Use it when collecting or interpreting the corresponding PostgreSQL statistics. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION opm_core;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `public.clone_graph(p_id_graph bigint)` is an extension function and returns `bigint`.
- `public.drop_account(IN p_account text)` is an extension function and returns `TABLE`.
- `public.get_graph(p_id_graph bigint)` is an extension function and returns `TABLE`.
- `public.get_sampled_metric_data(p_id_metric bigint, p_timet_begin timestamp with time zone, p_timet_end timestamp with time zone, p_sample_num integer)` is an extension function and returns `TABLE`.
- `public.get_server(IN p_id bigint)` is an extension function and returns `TABLE`.
- `public.get_service(IN p_id bigint)` is an extension function and returns `TABLE`.
- `public.grant_appli(IN p_role name)` is an extension function and returns `TABLE`.
- `public.grant_dispatcher(IN p_whname name, IN p_rolname name)` is an extension function and returns `TABLE`.
- `public.js_time(timestamptz)` is an extension function and returns `bigint`.
- `public.js_timetz(timestamptz)` is an extension function and returns `bigint`.
- `public.list_accounts()` is an extension function and returns `TABLE`.
- `public.list_graphs()` is an extension function and returns `TABLE`.
- `public.list_graphs(p_id_server bigint)` is an extension function and returns `TABLE`.
- `public.list_graphs_templates(IN p_id bigint)` is an extension function and returns `TABLE`.

### Requirements and Caveats

- The reviewed control file declares default version `2.6`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
