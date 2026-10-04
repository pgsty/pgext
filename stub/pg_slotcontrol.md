## Usage

Sources:

- [pg_slotcontrol/pg_slotcontrol.control](https://github.com/mhagander/pg_slotsync/blob/655f5836dc6736e3aa7bcdb521b6395645da263e/pg_slotcontrol/pg_slotcontrol.control)
- [pg_slotcontrol/README.md](https://github.com/mhagander/pg_slotsync/blob/655f5836dc6736e3aa7bcdb521b6395645da263e/pg_slotcontrol/README.md)
- [pg_slotcontrol/pg_slotcontrol--1.0.sql](https://github.com/mhagander/pg_slotsync/blob/655f5836dc6736e3aa7bcdb521b6395645da263e/pg_slotcontrol/pg_slotcontrol--1.0.sql)
- [pg_slotcontrol/pg_slotcontrol.c](https://github.com/mhagander/pg_slotsync/blob/655f5836dc6736e3aa7bcdb521b6395645da263e/pg_slotcontrol/pg_slotcontrol.c)
- [LICENSE](https://github.com/mhagander/pg_slotsync/blob/655f5836dc6736e3aa7bcdb521b6395645da263e/LICENSE)

`pg_slotcontrol` exposes `pg_slotmove(text, pg_lsn)` to move a physical replication slot retention position forward. This is a 2017 source snapshot; current PostgreSQL compatibility is not established.

### Core Workflow

```sql
CREATE EXTENSION pg_slotcontrol;
SELECT slot_name, slot_type, restart_lsn FROM pg_replication_slots;
```

```sql
SELECT pg_slotmove('physical_standby', '1/86000100'::pg_lsn);
```

### Operational Boundaries

Only move a slot after ensuring required WAL exists elsewhere. The C implementation rejects logical slots, caps the target at the current WAL write position and never moves backwards. The boolean result indicates whether the position changed. Installation requires a superuser; PUBLIC execution is revoked, so delegate the function deliberately. No preload is specified. The separate `pg_slotsync` program is not installed by creating this extension.
