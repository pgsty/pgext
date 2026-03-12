

## Usage

> [insert_username: track who modified a table row](https://www.postgresql.org/docs/current/contrib-spi.html)

Provides a trigger function that stores the current user's name into a specified text column.

```sql
CREATE EXTENSION insert_username;
```

### Trigger Function

| Function | Description |
|---|---|
| `insert_username()` | Store current username in the specified column |

Parameter: name of the text column to store the username.

### Examples

```sql
CREATE TABLE audit_log (
  id serial PRIMARY KEY,
  data text,
  modified_by text
);

-- Track who inserts
CREATE TRIGGER set_insert_user
  BEFORE INSERT ON audit_log
  FOR EACH ROW
  EXECUTE FUNCTION insert_username('modified_by');

-- Track who updates
CREATE TRIGGER set_update_user
  BEFORE UPDATE ON audit_log
  FOR EACH ROW
  EXECUTE FUNCTION insert_username('modified_by');

INSERT INTO audit_log (data) VALUES ('test');
SELECT modified_by FROM audit_log;  -- returns current user
```
