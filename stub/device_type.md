## Usage

Sources:

- [Official README](https://github.com/adjust/pg-device_type/blob/master/README.md)
- [Version 0.0.2 SQL](https://github.com/adjust/pg-device_type/blob/master/device_type--0.0.2.sql)
- [Current C implementation](https://github.com/adjust/pg-device_type/blob/master/src/device_type.c)
- [Current control file](https://github.com/adjust/pg-device_type/blob/master/device_type.control)

`device_type` defines a one-byte, enum-like base type for a fixed set of device classes. It is smaller than text and supports comparison plus B-tree and hash indexes, but its ordering follows the implementation's internal ordinal values rather than lexical order.

### Core Workflow

Create the extension and use the type for controlled device classifications:

```sql
CREATE EXTENSION device_type;

CREATE TABLE session_device (
  session_id bigint PRIMARY KEY,
  device device_type NOT NULL
);

INSERT INTO session_device VALUES (1, 'phone'), (2, 'tablet');
CREATE INDEX ON session_device (device);
SELECT * FROM session_device WHERE device = 'phone';
```

Input is case-folded and must resolve to one of the supported tokens.

### Values and Ordering

The accepted values, in comparison order, are `bot`, `console`, `ipod`, `mac`, `pc`, `phone`, `server`, `simulator`, `tablet`, `tv`, and `unknown`. The type is a C base type, not a PostgreSQL enum, and the extension does not define general casts to text or integers.

### Version and Operational Notes

The catalog tracks 0.0.2, while the current upstream control file declares 0.0.3. The upstream 0.0.2-to-0.0.3 update script contains no user-facing SQL object change; the inspected value set is unchanged, while current builds mark functions parallel-safe on supported servers. Because on-disk bytes and sort order depend on the C ordinal layout, any future value insertion or reordering requires careful upgrade testing. Build and regression-test the exact source version selected for deployment.
