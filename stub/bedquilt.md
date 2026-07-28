## Usage

Sources:

- [Official extension control file (bedquilt.control)](https://api.pgxn.org/src/bedquilt/bedquilt-2.1.0/bedquilt.control)
- [Official extension SQL (bedquilt--0.2.0.sql)](https://api.pgxn.org/src/bedquilt/bedquilt-2.1.0/sql/bedquilt--0.2.0.sql)

`bedquilt` — A json object store. Use it when an application needs this specific database capability. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION bedquilt;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `bq_add_constraints(i_coll text, i_jdoc json)` is an extension function and returns `boolean`.
- `bq_check_id_type(i_jdoc json)` is an extension function and returns `VOID`.
- `bq_collection_exists(i_coll text)` is an extension function and returns `boolean`.
- `bq_constraint_name_exists(i_coll text, i_name text)` is an extension function and returns `boolean`.
- `bq_count(i_coll text, i_doc json)` is an extension function and returns `integer`.
- `bq_create_collection(i_coll text)` is an extension function and returns `BOOLEAN`.
- `bq_delete_collection(i_coll text)` is an extension function and returns `BOOLEAN`.
- `bq_doc_set_key(i_jdoc json, i_key text, i_val anyelement)` is an extension function and returns `json`.
- `bq_find(i_coll text, i_json_query json)` is an extension function and returns `table`.
- `bq_find_one(i_coll text, i_json_query json)` is an extension function and returns `table`.
- `bq_find_one_by_id(i_coll text, i_id text)` is an extension function and returns `table`.
- `bq_generate_id()` is an extension function and returns `char`.
- `bq_insert(i_coll text, i_jdoc json)` is an extension function and returns `text`.
- `bq_list_collections()` is an extension function and returns `table`.

### Requirements and Caveats

- The reviewed control file declares default version `2.1.0`.
- Install the confirmed extension dependencies first: `plpgsql`, `pgcrypto`, `plpython3u`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
