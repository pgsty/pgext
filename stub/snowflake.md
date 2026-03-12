

## Usage

> [snowflake: Snowflake ID sequences for PostgreSQL](https://github.com/pgEdge/snowflake)

Provides `int8` and `sequence` based unique ID generation using the Snowflake format, suitable for distributed systems.

```sql
CREATE EXTENSION snowflake;
```

### Configuration

Set the node identifier in `postgresql.conf` (required, values 1-1023):

```ini
snowflake.node = 1
```

### Functions

| Function | Description |
|---|---|
| `snowflake.nextval([sequence regclass])` | Generate the next Snowflake ID (uses internal sequence if none specified) |
| `snowflake.currval([sequence regclass])` | Return the current value of the sequence |
| `snowflake.get_epoch(snowflake int8)` | Extract the timestamp as epoch (seconds since 2023-01-01) |
| `snowflake.get_count(snowflake int8)` | Extract the count part (resets per millisecond) |
| `snowflake.get_node(snowflake int8)` | Extract the node identifier |
| `snowflake.format(snowflake int8)` | Return a JSONB with `node`, `ts`, and `count` fields |

### Examples

```sql
-- Generate a snowflake ID
SELECT snowflake.nextval();
-- 136169504773242881

-- Use with a named sequence
CREATE SEQUENCE orders_id_seq;
SELECT snowflake.nextval('orders_id_seq'::regclass);

-- Extract components
SELECT snowflake.get_epoch(136169504773242881);
-- 1704996539.845

SELECT to_timestamp(snowflake.get_epoch(136169504773242881));
-- 2024-01-11 13:08:59.845-05

SELECT snowflake.get_node(136169504773242881);
-- 1

SELECT snowflake.format(136169504773242881);
-- {"id": 1, "ts": "2024-01-11 13:08:59.845-05", "count": 0}

-- Use as default column
CREATE TABLE orders (
  id int8 DEFAULT snowflake.nextval() PRIMARY KEY,
  data text
);
```
