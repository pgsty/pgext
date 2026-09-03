## Usage

Sources:

- [Official documentation](https://github.com/yansheng836/pg_toastinspect/blob/65571a9267e3123eb0d9d477dff6c5a5246e948a/README.md)
- [Extension control file](https://github.com/yansheng836/pg_toastinspect/blob/65571a9267e3123eb0d9d477dff6c5a5246e948a/pg_toastinspect.control)
- [Official repository](https://github.com/yansheng836/pg_toastinspect)

`pg_toastinspect` Low-level inspection of TOAST chunk identifiers, sizes, and stored values.

### Enablement

Install the files for the intended server, then create `pg_toastinspect` in the target database:

```sql
CREATE EXTENSION pg_toastinspect;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
-- Get the chunk ID of a TOASTed value
SELECT ctid, id, get_toast_chunk_id(response_content), response_content
FROM ai_call_log_copy1
WHERE id = 668310181431480320;

-- Get complete TOAST information
SELECT ctid, id, get_toast_info(response_content), response_content
FROM ai_call_log_copy1
WHERE id = 668310181431480320;

-- Access individual fields from the composite type
SELECT
    ctid,
    id,
    (get_toast_info(response_content)).raw_size,
    (get_toast_info(response_content)).ext_size,
    (get_toast_info(response_content)).chunk_id,
    (get_toast_info(response_content)).reltoastrelid
FROM ai_call_log_copy1
WHERE id = 668310181431480320;
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `get_toast_info` | FUNCTION | Callable function from the reviewed install surface. |
| `get_toast_chunk_id` | FUNCTION | Callable function from the reviewed install surface. |
| `toast_pointer_info` | TYPE | User-facing data type created by the extension. |

### Operations and Boundaries

- Treat extension upgrades as database changes: review the upstream upgrade path, privileges, locks, and backup/restore behavior first.
