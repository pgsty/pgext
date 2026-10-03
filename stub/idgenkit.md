## Usage

Sources:

- [README.md](https://github.com/rahulbsw/idgenkit/blob/37126f591680b741305b54463d87336537e93d05/README.md)
- [postgres/idgenkit.control](https://github.com/rahulbsw/idgenkit/blob/37126f591680b741305b54463d87336537e93d05/postgres/idgenkit.control)
- [postgres/idgenkit--0.1.0.sql](https://github.com/rahulbsw/idgenkit/blob/37126f591680b741305b54463d87336537e93d05/postgres/idgenkit--0.1.0.sql)

`idgenkit` 0.1.0 generates several ID formats on PostgreSQL 14–18. Basic random generators work without preloading; relative IDs and some Snowflake configurations have additional requirements.

### Core Workflow

```sql
CREATE EXTENSION idgenkit;
SELECT ulid_generate(), uuidv7_generate(), nanoid_generate();
CREATE TABLE events (id uuid PRIMARY KEY DEFAULT ulid_generate_uuid(), body text);
```

### Functions

`ulid_generate_monotonic()` and `uuidv7_generate_monotonic()` increase within one backend, not globally. `ulid_to_uuid`, `ulid_from_uuid`, `ulid_timestamp` and `uuidv7_timestamp` convert or inspect IDs. `snowflake_generate`, `snowflake_timestamp`, `snowflake_machine_id`, `snowflake_sequence` and `snowflake_from_timestamp` operate on 64-bit IDs. `relid_generate`, `relid_generate_monotonic`, `relid_tag` and `relid_timestamp` provide keyed tags and timestamp extraction.

### Configuration and Boundaries

The control marks the extension trusted and relocatable. Keyed relative IDs require `idgenkit` in `shared_preload_libraries`, a restart and a superuser-managed `idgenkit.relid_secret` of at least 16 bytes. Protect that secret outside version control; anyone allowed to execute the tag function can test guessed keys. The 30-bit tags can collide and are not an authorization boundary.

Snowflake uses shared state. PostgreSQL 14–16 requires preload and restart; 17+ can initialize it on first use without preload. Assign distinct `idgenkit.machine_id` values to separate servers and configure `idgenkit.snowflake_epoch_ms` consistently. Store textual ULIDs with `COLLATE "C"` for chronological ordering. IDs expose timing information and are not authentication secrets.
