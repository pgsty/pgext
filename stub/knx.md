## Usage

Sources:

- [Official README](https://github.com/mike-kfed/pgknx/blob/c485a440f7347a0bd840018054466735d8f5f9ce/README.md)
- [Type definitions](https://github.com/mike-kfed/pgknx/blob/c485a440f7347a0bd840018054466735d8f5f9ce/knx.sql)
- [Address parser](https://github.com/mike-kfed/pgknx/blob/c485a440f7347a0bd840018054466735d8f5f9ce/knx.c)
- [Operators and index classes](https://github.com/mike-kfed/pgknx/blob/c485a440f7347a0bd840018054466735d8f5f9ce/knx-operators.sql)

`knx` 1.0 adds compact two-byte types for KNX group and individual addresses. Use them to validate common KNX text formats, preserve the underlying 16-bit address value, and support numeric ordering plus B-tree and hash indexes.

### Core Workflow

Choose the presentation appropriate to the address family:

```sql
CREATE EXTENSION knx;

CREATE TABLE knx_endpoint (
  group_address knx_group_address3 PRIMARY KEY,
  device_address knx_individual_address
);

INSERT INTO knx_endpoint VALUES ('17/6/1', '4.3.250');
SELECT group_address, group_address::integer FROM knx_endpoint;
```

Integer casts in both directions expose the packed 16-bit value. Zero is rejected for all three address types.

### Type Index

- `knx_group_address3` renders `main/middle/sub`, with ranges 0–31, 0–7, and 0–255.
- `knx_group_address2` renders `main/sub`, with ranges 0–31 and 0–2047.
- `knx_individual_address` renders `area.line.device`, with ranges 0–15, 0–15, and 0–255.

Each type also accepts a hexadecimal packed address. Comparisons use the packed numeric value, including cross-type comparison operators among the KNX address types.

### Operational Notes

The group-address parser accepts either two-level or three-level slash input for both group types; the declared type only controls the canonical output format. Validate input shape in the application or a CHECK constraint if that distinction matters. The parser rejects the zero/broadcast address and values outside the documented ranges. Before modeling a KNX installation, confirm whether broadcast and reserved addresses need separate representation rather than forcing them into these types.
