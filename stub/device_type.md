## Usage

Sources:

- [Official extension control file](https://github.com/adjust/pg-device_type/blob/2d19dcc0ab9df0808579aff22bd412e44c64741b/device_type.control)
- [Official PGXN distribution page](https://pgxn.org/dist/device_type/)

`device_type` — Fixed-byte enum-like data type for mobile device types with comparison operators and btree/hash operator classes.

The reviewed catalog snapshot records version `0.0.2`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "device_type";
SELECT extversion
FROM pg_extension
WHERE extname = 'device_type';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
