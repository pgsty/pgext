

## Usage

> [lolor: Logical-replication-friendly replacement for PostgreSQL large objects](https://github.com/pgEdge/lolor)

Makes PostgreSQL large objects compatible with logical replication by storing them in non-catalog tables.

### Enabling

```sql
CREATE EXTENSION lolor;
```

Configure the node identifier in `postgresql.conf`:

```ini
lolor.node = 1  -- unique node ID (1 to 2^28)
```

Optionally adjust the search path:

```sql
SET search_path = lolor, "$user", public, pg_catalog;
```

### Large Object Operations

Once installed, the standard `lo_*` functions are redirected to use lolor's tables:

```sql
-- Create a large object
SELECT lo_create(0);

-- Import a file into a large object
SELECT lo_import('/path/to/file.bin');

-- Export a large object to a file
SELECT lo_export(oid, '/path/to/output.bin');

-- Open, read, write, seek, close
SELECT lo_open(oid, x'40000'::int);  -- INV_WRITE
SELECT lowrite(fd, 'data'::bytea);
SELECT loread(fd, 1024);
SELECT lo_close(fd);

-- Delete a large object
SELECT lo_unlink(oid);
```

### Replication Setup

Add lolor tables to your replication set:

```sql
-- For spock/pgedge replication
SELECT spock.repset_add_table('default', 'lolor.pg_largeobject');
SELECT spock.repset_add_table('default', 'lolor.pg_largeobject_metadata');
```

### Internal Tables

The extension manages large objects in:

- `lolor.pg_largeobject` - stores object data chunks
- `lolor.pg_largeobject_metadata` - stores object metadata

### Limitations

- Native PostgreSQL large object functionality cannot be used while lolor is active
- Migration of existing native large objects to lolor is not supported
- `ALTER LARGE OBJECT`, `GRANT ON LARGE OBJECT`, `COMMENT ON LARGE OBJECT`, and `REVOKE ON LARGE OBJECT` are not supported
- Requires PostgreSQL 16 or newer
