## Usage

Sources:

- [README](https://github.com/apjapj/pg_bloat_cause/blob/b5de883b65784eccf6f9d0c8e65d68d7252349d2/README.md)
- [Control](https://github.com/apjapj/pg_bloat_cause/blob/b5de883b65784eccf6f9d0c8e65d68d7252349d2/pg_bloat_cause.control)
- [Install and grants](https://github.com/apjapj/pg_bloat_cause/blob/b5de883b65784eccf6f9d0c8e65d68d7252349d2/INSTALL.md)
- [Install SQL](https://github.com/apjapj/pg_bloat_cause/blob/b5de883b65784eccf6f9d0c8e65d68d7252349d2/pg_bloat_cause--0.7.0.sql)

`pg_bloat_cause` 0.7.0 diagnoses why dead rows are not being removed, using PostgreSQL catalog and statistics evidence. It is pure SQL, tested on PostgreSQL 14–19rc1, with superuser installation and no preload or restart.

### Read a Diagnosis

```sql
CREATE EXTENSION pg_bloat_cause;
SELECT * FROM bloat_cause_text();
```

Findings distinguish pinned cleanup horizons, ineffective autovacuum and non-HOT update pressure. They are estimates and observations, not an exact physical bloat measurement. A logical slot's catalog_xmin holds catalog cleanup rather than automatically pinning ordinary user-table rows. Action helpers present commands for operator review and do not execute repairs.

### Permissions, History and Alerts

Readers need explicit view/function grants and pg_read_all_stats to inspect other sessions. Sampling, pruning and NOTIFY functions also need write permissions on their history and alert tables. An external scheduler must call them; no worker is installed. Capturing holder query text is opt-in and may retain secrets. Restore the relevant roles and grants with the extension's application history.
