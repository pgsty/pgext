## Usage

Sources:

- [Upstream README](https://github.com/mweber26/timestampandtz/blob/fc3597ec859201d8f1e122c56bbf355d4c79b04f/README.md)
- [Install SQL](https://github.com/mweber26/timestampandtz/blob/fc3597ec859201d8f1e122c56bbf355d4c79b04f/timestampandtz--1.0.0.sql)
- [C type implementation](https://github.com/mweber26/timestampandtz/blob/fc3597ec859201d8f1e122c56bbf355d4c79b04f/timestampandtz.c)
- [Fixed timezone identifier table](https://github.com/mweber26/timestampandtz/blob/fc3597ec859201d8f1e122c56bbf355d4c79b04f/zones.c)
- [PGXN release metadata](https://pgxn.org/dist/timestampandtz/)

timestampandtz stores an instant together with the named source timezone, preserving local-time display while comparisons and B-tree ordering use the UTC instant. The extension's SQL default version is 1.0.0; PGXN labels the same install script as distribution release 1.0.2 from 2018.

### Preserve a source timezone

```sql
CREATE EXTENSION timestampandtz;
SELECT '2014-09-01 22:15:00 @ US/Eastern'::timestampandtz;
SELECT tzmove(
  '2014-09-01 22:15:00 @ US/Eastern'::timestampandtz,
  'US/Pacific'
);
SELECT '2014-09-01 22:15:00 @ US/Eastern'::timestampandtz
     = '2014-09-01 19:15:00 @ US/Pacific'::timestampandtz;
```

The final comparison is true because the two values represent the same instant. Interval arithmetic preserves the stored timezone and applies wall-clock daylight-saving rules.

### Caveats

- The two-byte timezone identifier is an extension-owned fixed table, not PostgreSQL's live timezone catalog. New or renamed IANA zones may be unavailable, although offset rules still come from the server's timezone data.
- Preserve the exact identifier table when moving binary data between builds. The binary receive function does not validate identifier bounds before later table lookup.
- Malformed textual values containing an empty side around the @ separator can reach unsafe buffer-boundary logic in the published C parser. Reject untrusted input before casting and test repaired code for exposed workloads.
- Installation adds a timezone overload inside pg_catalog and therefore requires elevated trust. Review object ownership and conflicts during install and removal.
- Upstream's last release is from 2018 and targets PostgreSQL 9.2-era internals. There is no current major-version compatibility statement; rebuild and run regression, daylight-saving, binary-copy, dump, restore, and index tests before adoption.
