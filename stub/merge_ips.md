## Usage

Sources:

- [Official upstream documentation](https://github.com/cybertec-postgresql/merge_ips/blob/3cf654cef7c7ce7fafebc9b0799f15da8725da52/README)
- [Official extension SQL](https://github.com/cybertec-postgresql/merge_ips/blob/3cf654cef7c7ce7fafebc9b0799f15da8725da52/merge_ips--1.0.sql)
- [Official extension control file](https://github.com/cybertec-postgresql/merge_ips/blob/3cf654cef7c7ce7fafebc9b0799f15da8725da52/merge_ips.control)

`merge_ips` 1.0 combines adjacent IPv4 or IPv6 addresses and networks into the smallest exact set of enclosing networks without filling gaps. It is a pure-PL/pgSQL extension installed in the fixed `merge_ips` schema. The upstream repository is archived, so pin and test the implementation before using it in persistent workflows.

### Merge Networks

Pass an `inet[]` array to the set-returning function or its array wrapper:

```sql
CREATE EXTENSION merge_ips;

SELECT *
FROM merge_ips.merge_ips(ARRAY[
  '10.0.52.0/30',
  '10.0.52.4/30',
  '10.0.52.16/30',
  '10.0.52.20/30'
]::inet[]);

SELECT merge_ips.merge_ips_array(ARRAY[
  '192.0.2.0/25',
  '192.0.2.128/25'
]::inet[]);
```

The first query returns `10.0.52.0/29` and `10.0.52.16/29`; it does not invent a network covering the gap between them. The second returns an array containing `192.0.2.0/24`.

### Function Behavior

- `merge_ips.merge_ips(inet[])` returns a set of `inet` values, removes duplicates and networks contained by another result, and repeatedly combines exact sibling networks.
- `merge_ips.merge_ips_array(inet[])` collects the same result into an `inet[]`.

The output order is not part of the SQL contract; add `ORDER BY` when deterministic presentation matters.

### Operational Notes

The implementation creates, truncates, renames, and drops fixed-name temporary work tables in the current session. It is therefore not a side-effect-free immutable calculation and can conflict with nested or reentrant calls in the same session. Avoid calling it from immutable expressions, indexes, or parallelized application code sharing one session. Keep IPv4 and IPv6 inputs separate, test host bits and containment cases, and sort results explicitly. For large arrays, benchmark temporary-table I/O and loop cost; a maintained application library may be a better fit when this archived implementation is too slow or operationally intrusive.
