## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/MasaoFujii/postgres_fdw_plus/blob/e04bf0a9948333e415e3a5ea68a0af861b19e918/README.md)
- [Version 1.0 SQL objects](https://github.com/MasaoFujii/postgres_fdw_plus/blob/e04bf0a9948333e415e3a5ea68a0af861b19e918/postgres_fdw_plus--1.0.sql)
- [FDW implementation](https://github.com/MasaoFujii/postgres_fdw_plus/blob/e04bf0a9948333e415e3a5ea68a0af861b19e918/postgres_fdw_plus.c)

`postgres_fdw_plus` forks PostgreSQL's `postgres_fdw` and adds optional two-phase commit for foreign transactions. It creates the wrapper named `postgres_fdw`, so core `postgres_fdw` and this extension cannot coexist in one database.

```sql
CREATE EXTENSION postgres_fdw_plus;
CREATE SERVER remote_pg FOREIGN DATA WRAPPER postgres_fdw
  OPTIONS (host 'db.internal', dbname 'app');
SET postgres_fdw.two_phase_commit = on;
```

Every remote server must support prepared transactions and be configured for the resulting prepared-state load. Track and monitor `pg_prepared_xacts` on all participants. The extension records committed local transactions so resolver functions can decide whether to commit or roll back stranded remote transactions.

Several critical GUCs are changeable by ordinary users, including a debug option that deliberately stops after prepare and a tracking switch that can make resolution uncertain. Enforce safe role/database defaults and deny untrusted sessions. Never use forced rollback unless tracking completeness is proven. Test crash points across prepare, local commit, remote commit, recovery, failover, timeouts, user mappings, and cleanup; 2PC can retain locks and WAL indefinitely. Use only with PostgreSQL 16 or later versions explicitly tested upstream.
