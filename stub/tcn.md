

## Usage

> [tcn: triggered change notifications via LISTEN/NOTIFY](https://www.postgresql.org/docs/current/tcn.html)

Provides a trigger function that sends NOTIFY events with information about changed rows, enabling asynchronous change tracking.

```sql
CREATE EXTENSION tcn;
```

### Trigger Function

| Function | Description |
|---|---|
| `triggered_change_notification()` | Send NOTIFY on row changes with primary key info |

Optional parameter: custom channel name (defaults to `tcn`).

### Notification Payload Format

```
"table_name",operation,"column"='value',"column"='value'
```

Operations: `I` (INSERT), `U` (UPDATE), `D` (DELETE).

### Examples

```sql
CREATE TABLE tcndata (
  a int NOT NULL,
  b date NOT NULL,
  c text,
  PRIMARY KEY (a, b)
);

-- Attach the trigger
CREATE TRIGGER tcndata_tcn
  AFTER INSERT OR UPDATE OR DELETE ON tcndata
  FOR EACH ROW
  EXECUTE FUNCTION triggered_change_notification();

-- Listen for notifications
LISTEN tcn;

-- Changes trigger notifications:
INSERT INTO tcndata VALUES (1, '2024-01-01', 'test');
-- Notification: "tcndata",I,"a"='1',"b"='2024-01-01'

UPDATE tcndata SET c = 'updated' WHERE a = 1;
-- Notification: "tcndata",U,"a"='1',"b"='2024-01-01'

DELETE FROM tcndata WHERE a = 1;
-- Notification: "tcndata",D,"a"='1',"b"='2024-01-01'

-- Use a custom channel name
CREATE TRIGGER my_trigger
  AFTER INSERT OR UPDATE OR DELETE ON my_table
  FOR EACH ROW
  EXECUTE FUNCTION triggered_change_notification('my_channel');
```
