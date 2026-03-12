

## Usage

> [wal2mongo: PostgreSQL logical decoding output plugin for MongoDB](https://github.com/HighgoSoftware/wal2mongo)

A logical decoding output plugin that formats PostgreSQL WAL changes as MongoDB-compatible commands, enabling data replication from PostgreSQL to MongoDB.

### Configuration

In `postgresql.conf`:

```ini
wal_level = logical
max_replication_slots = 10
```

### Using with SQL Functions

```sql
-- Create a logical replication slot
SELECT * FROM pg_create_logical_replication_slot('w2m_slot', 'wal2mongo');

-- Perform DML operations
CREATE TABLE books (id SERIAL PRIMARY KEY, title VARCHAR(100), author VARCHAR(100));
INSERT INTO books (id, title, author) VALUES (123, 'My Book', 'Author');

-- Peek at changes (MongoDB format)
SELECT * FROM pg_logical_slot_peek_changes('w2m_slot', NULL, NULL);
-- Output: db.books.insertOne( { id:123, title:"My Book", author:"Author" } )

-- Consume changes
SELECT * FROM pg_logical_slot_get_changes('w2m_slot', NULL, NULL);

-- Drop the slot
SELECT pg_drop_replication_slot('w2m_slot');
```

### Using with pg_recvlogical

```bash
pg_recvlogical -d postgres --slot w2m_slot --create-slot -P wal2mongo
pg_recvlogical -d postgres --slot w2m_slot --start -f -
```

### Replicating to MongoDB

The output can be applied directly in the MongoDB shell:

```javascript
// Copy the output from pg_logical_slot_get_changes
db.books.insertOne( { id:123, title:"My Book", author:"Author" } )
```

Or save to a `.js` file and import:

```bash
mongo < changes.js
```

### Output Format

- **INSERT**: `db.<table>.insertOne( { <columns> } )`
- **UPDATE**: `db.<table>.updateOne( { <key> }, { $set: { <changes> } } )`
- **DELETE**: `db.<table>.deleteOne( { <key> } )`

Tables need a primary key or replica identity for UPDATE/DELETE operations to be captured.
