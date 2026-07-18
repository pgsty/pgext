## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/storm-db/pgxc_dns/blob/441b94223abd693620b5125ae473dfd40a36e790/README)
- [Version 1.0 SQL objects](https://github.com/storm-db/pgxc_dns/blob/441b94223abd693620b5125ae473dfd40a36e790/pgxc_dns--1.0.sql)
- [C implementation](https://github.com/storm-db/pgxc_dns/blob/441b94223abd693620b5125ae473dfd40a36e790/pgxc_dns.c)
- [Extension control file](https://github.com/storm-db/pgxc_dns/blob/441b94223abd693620b5125ae473dfd40a36e790/pgxc_dns.control)

`pgxc_dns` is historical Postgres-XC 1.0 code that generates BIND SDB zone rows according to relative coordinator usage. It was intended to expose coordinator A records through the `pgxc_dns_zone` view for DNS-based load distribution.

```sql
CREATE EXTENSION pgxc_dns;
SELECT name, ttl, rdtype, rdata
FROM pgxc_dns_zone
ORDER BY name, rdtype, rdata;
```

Configuration uses `pgxc_dns.zone`, `pgxc_dns.name`, `pgxc_dns.host`, `pgxc_dns.zone_ttl`, and `pgxc_dns.ttl`. Each coordinator must report its own address while sharing the zone settings.

This is not an extension for current PostgreSQL, Postgres-XL, or modern distributed PostgreSQL products. Upstream testing was on CentOS 6.3 with Postgres-XC based on PostgreSQL 9.2, and the documented BIND configuration embeds database credentials. Treat it as archival source only. Use supported service discovery, health checks, and authenticated DNS integration instead of deploying it.
