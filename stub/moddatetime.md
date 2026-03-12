

## Usage

> [moddatetime: track modification timestamp](https://www.postgresql.org/docs/current/contrib-spi.html)

Provides a trigger function that stores the current timestamp when a row is modified.

```sql
CREATE EXTENSION moddatetime;
```

### Trigger Function

| Function | Description |
|---|---|
| `moddatetime()` | Store current timestamp in the specified column on UPDATE |

Parameter: name of the `timestamp` or `timestamp with time zone` column to update.

### Examples

```sql
CREATE TABLE documents (
  id serial PRIMARY KEY,
  content text,
  modified_at timestamp with time zone
);

CREATE TRIGGER set_modified
  BEFORE UPDATE ON documents
  FOR EACH ROW
  EXECUTE FUNCTION moddatetime('modified_at');

UPDATE documents SET content = 'new content' WHERE id = 1;
-- modified_at is automatically set to current timestamp
```
