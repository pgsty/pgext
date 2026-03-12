

## Usage

> [autoinc: auto-increment trigger function](https://www.postgresql.org/docs/current/contrib-spi.html)

Provides a trigger function that auto-increments specified columns using sequences.

```sql
CREATE EXTENSION autoinc;
```

### Trigger Function

| Function | Description |
|---|---|
| `autoinc()` | Auto-increment columns from sequences on insert (and optionally update) |

Parameters are pairs of (column name, sequence name). The value is replaced only if initially zero or NULL.

### Examples

```sql
CREATE SEQUENCE next_id;
CREATE TABLE ids (id int4, idesc text);

CREATE TRIGGER ids_autoinc
  BEFORE INSERT ON ids
  FOR EACH ROW
  EXECUTE FUNCTION autoinc('id', 'next_id');

INSERT INTO ids (idesc) VALUES ('first');  -- id is auto-assigned from next_id

-- Multiple auto-increment columns
CREATE TRIGGER multi_autoinc
  BEFORE INSERT ON my_table
  FOR EACH ROW
  EXECUTE FUNCTION autoinc('col1', 'seq1', 'col2', 'seq2');
```
