## Usage

Sources:

- [Upstream README](https://github.com/diffix/pg_diffix/blob/94b4edff1e697de73f0649d7fab64890bdb52bc7/README.md)
- [Extension control file](https://github.com/diffix/pg_diffix/blob/94b4edff1e697de73f0649d7fab64890bdb52bc7/pg_diffix.control)
- [Administrator guide](https://github.com/diffix/pg_diffix/blob/94b4edff1e697de73f0649d7fab64890bdb52bc7/docs/admin_guide.md)
- [Analyst guide](https://github.com/diffix/pg_diffix/blob/94b4edff1e697de73f0649d7fab64890bdb52bc7/docs/analyst_guide.md)
- [SQL installation script](https://github.com/diffix/pg_diffix/blob/94b4edff1e697de73f0649d7fab64890bdb52bc7/pg_diffix--fir.sql)

`pg_diffix` implements dynamic query anonymization for PostgreSQL. Administrators label tables as public or personal, identify one or more anonymizing-ID columns, and assign roles the `direct`, `anonymized_trusted`, or `anonymized_untrusted` access level. Queries against personal tables are then restricted and aggregate results receive suppression and noise.

### Minimal administration outline

```sql
CREATE EXTENSION pg_diffix;

CALL diffix.mark_personal('public.customers', 'customer_id');
CALL diffix.mark_role('analyst', 'anonymized_untrusted');

ALTER DATABASE appdb
  SET session_preload_libraries TO 'pg_diffix';
```

Reconnect after setting the preload library, then inspect the active access level and issue an allowed grouped query:

```sql
SELECT diffix.access_level();

SELECT city, count(*)
FROM public.customers
GROUP BY city;
```

The anonymized modes intentionally allow only a limited SQL subset. Correct privacy depends on complete table/column labels, sound anonymizing IDs, appropriate suppression settings, role isolation, and review of RLS policies, user-defined functions, and every other installed extension. A role labeled `direct` bypasses anonymization.

The administrator guide explicitly calls the `Fir` generation a work in progress, and PostgreSQL 13 or later is required. The control file and catalog use the nonnumeric version `fir`, while `META.json` at the same commit reports distribution version `1.2.0`; do not treat those labels as an upgrade sequence. Deploy only after a privacy review and adversarial testing against the actual schema and workload.
