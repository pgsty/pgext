## Usage

Sources:

- [Official README](https://github.com/ntkme/postgresql-cidr/blob/0a79fac02992d1863b6f98a533e51abc25c3d69c/README.md)
- [Version 1.0 extension SQL](https://github.com/ntkme/postgresql-cidr/blob/0a79fac02992d1863b6f98a533e51abc25c3d69c/cidr--1.0.sql)
- [Upstream regression examples](https://github.com/ntkme/postgresql-cidr/blob/0a79fac02992d1863b6f98a533e51abc25c3d69c/sql/cidr.sql)

`cidr` adds a two-argument `pg_catalog.cidr(inet, inet)` function that returns the smallest set of CIDR networks covering an inclusive IPv4 or IPv6 address range. Unlike built-in `inet_merge`, it does not widen the result to one network that may contain addresses outside the requested range.

### Core Workflow

```sql
CREATE EXTENSION cidr;

SELECT network
FROM cidr('192.0.2.5'::inet, '192.0.2.20'::inet) AS r(network);
```

Use it to materialize an exact network cover for range-oriented access rules or reporting:

```sql
INSERT INTO allowed_networks(network)
SELECT network
FROM cidr('2001:db8::10'::inet, '2001:db8::ff'::inet) AS r(network);
```

### Function Contract

- Both arguments must belong to the same address family.
- Both arguments must be host addresses with a full-width mask, not input networks.
- The first address must be less than or equal to the second.
- The set-returning result consists of non-overlapping `cidr` values that exactly cover the inclusive range.

### Operational Notes

Version `1.0` is pure PL/pgSQL, relocatable at the extension level, and declares no preload dependency. The function itself is deliberately created in `pg_catalog`, so it is visible without changing `search_path`; account for that global name when auditing or removing the extension.

Large ranges are represented compactly where alignment permits, but highly unaligned boundaries can return many rows. Materialize or bound results when using user-supplied ranges. Mixed IPv4/IPv6 input, network-masked input, and descending ranges raise errors rather than being normalized automatically.
