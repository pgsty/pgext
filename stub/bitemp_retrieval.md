## Usage

Sources:

- [Official upstream README](https://github.com/aditditto/bitemp_retrieval/blob/8bc2fc4f9c7a84f92f0c1224ea557bd7f53a76b6/README.md)
- [Official extension control file (bitemp_retrieval.control)](https://github.com/aditditto/bitemp_retrieval/blob/8bc2fc4f9c7a84f92f0c1224ea557bd7f53a76b6/bitemp_retrieval.control)
- [Official extension SQL (bitemp_retrieval--0.0.1.sql)](https://github.com/aditditto/bitemp_retrieval/blob/8bc2fc4f9c7a84f92f0c1224ea557bd7f53a76b6/bitemp_retrieval--0.0.1.sql)

`bitemp_retrieval` — Postgresql extension for retrieval functionality on bitemporal tables created using pg_bitemporal extension by Henrietta Dombrovskaya. Use it for the corresponding scheduling, temporal, or time-series workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION bitemp_retrieval;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `bitemp_contains_now(effective_range temporal_relationships.timeperiod, asserted_range temporal_relationships.timeperiod)` is an extension function and returns `BOOLEAN`.
- `bitemp_contains_ts(effective_range temporal_relationships.timeperiod, asserted_range temporal_relationships.timeperiod, effective_ts timestamptz, asserted_ts timestamptz)` is an extension function and returns `BOOLEAN`.
- `bitemporal_internal.ll_register_temporal_attribute_property(p_schema TEXT, p_table TEXT, p_attr_name TEXT, p_attr_property bitemporal_internal.temporal_attribute_property_enum)` is an extension function and returns `INTEGER`.
- `get_interval_overlap(a temporal_relationships.timeperiod, b temporal_relationships.timeperiod)` is an extension function and returns `temporal_relationships`.
- `get_sum(int, int)` is an extension function and returns `int`.
- `interval_contains_now(interv temporal_relationships.timeperiod)` is an extension function and returns `BOOLEAN`.
- `interval_contains_ts(interv temporal_relationships.timeperiod, ts timestamptz)` is an extension function and returns `BOOLEAN`.
- `interval_join(a temporal_relationships.timeperiod, b temporal_relationships.timeperiod)` is an extension function and returns `temporal_relationships`.
- `interval_joinable(a temporal_relationships.timeperiod, b temporal_relationships.timeperiod)` is an extension function and returns `BOOLEAN`.
- `interval_len(interv temporal_relationships.timeperiod)` is an extension function and returns `INTERVAL`.
- `intervals_contains_now(intervs temporal_relationships.timeperiod[])` is an extension function and returns `BOOLEAN`.
- `intervals_contains_ts(intervs temporal_relationships.timeperiod[], ts timestamptz)` is an extension function and returns `BOOLEAN`.
- `ita_now(p_schema TEXT, p_table TEXT, p_group_by TEXT[], p_aggr_funcs TEXT[], p_aggr_target TEXT[], p_aggr_fieldnames TEXT[])` is an extension function and returns `SETOF`.
- `mwta_now(p_schema TEXT, p_table TEXT, p_group_by TEXT[], p_aggr_funcs TEXT[], p_aggr_target TEXT[], p_aggr_fieldnames TEXT[], p_window_size INTERVAL)` is an extension function and returns `SETOF`.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
