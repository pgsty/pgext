

## Usage

> [pgtt: Extension to add Global Temporary Tables feature to PostgreSQL](https://github.com/darold/pgtt)

### Creating a Global Temporary Table

```sql
CREATE EXTENSION pgtt;

-- ON COMMIT PRESERVE ROWS: data persists across transactions within a session
CREATE GLOBAL TEMPORARY TABLE test_gtt (
    id integer,
    lbl text
) ON COMMIT PRESERVE ROWS;

-- ON COMMIT DELETE ROWS: data is deleted at transaction commit
CREATE GLOBAL TEMPORARY TABLE session_data (
    id integer,
    value text
) ON COMMIT DELETE ROWS;
```

The `GLOBAL` keyword can also be used as a comment to avoid warnings:

```sql
CREATE /*GLOBAL*/ TEMPORARY TABLE test_gtt (
    LIKE other_table INCLUDING DEFAULTS INCLUDING CONSTRAINTS INCLUDING INDEXES
) ON COMMIT PRESERVE ROWS;
```

### CREATE AS Form

```sql
CREATE /*GLOBAL*/ TEMPORARY TABLE gtt_copy
AS SELECT * FROM source_table WITH DATA;
```

### Using Global Temporary Tables

```sql
INSERT INTO test_gtt VALUES (1, 'one'), (2, 'two');
SELECT * FROM test_gtt;  -- visible only in current session
```

### Creating Indexes

```sql
CREATE INDEX ON test_gtt (id);
```

### Constraints

All constraints except FOREIGN KEYS are supported:

```sql
CREATE GLOBAL TEMPORARY TABLE t2 (
    c1 serial PRIMARY KEY,
    c2 VARCHAR(50) UNIQUE NOT NULL,
    c3 boolean DEFAULT false
);
```

### Dropping

```sql
DROP TABLE test_gtt;  -- can be dropped even while used by other sessions
```

### Configuration

```sql
SET pgtt.enabled TO off;   -- disable GTT rerouting
SET pgtt.enabled TO on;    -- re-enable GTT rerouting
```

### Key Behaviors

- GTT content is session-local; other sessions cannot see your data
- The table structure is persistent (visible to all users), but data is per-session
- Requires loading via `session_preload_libraries = 'pgtt'`
- Partitioning is not supported on GTTs
