## Usage

Sources:

- [Extension control file](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_ledger/pg_ledger.control)
- [Ledger SQL API and transaction checks](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_ledger/src/lib.rs)
- [Account-rule implementation](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_ledger/src/rules.rs)

`pg_ledger` version `0.2.0` is a double-entry accounting engine in the fixed `pgledger` schema. It records immutable journal headers and lines, checks transaction-level debit/credit balance, manages currencies and fiscal periods, and can generate entries from rules attached to application tables.

### Core Workflow

Post balanced entries in one transaction and query the resulting account totals:

```sql
CREATE EXTENSION pg_ledger;

BEGIN;
SELECT pgledger.debit('cash', 100.00, 'sale received');
SELECT pgledger.credit('revenue', 100.00, 'sale received');
SELECT pgledger.check_balance();
COMMIT;

SELECT pgledger.account_balance('cash');
SELECT * FROM pgledger.trial_balance();
SELECT * FROM pgledger.journal_entries();
```

The central type is `pgledger.ledgeramount`, with arithmetic, comparisons, B-tree indexing, and `sum`, `min`, and `max` aggregates. Core entry functions are `debit`, `credit`, `entry`, `reverse`, and `check_balance`. Administration functions include `create_rule`, `enable_rule`, `disable_rule`, `drop_rule`, `set_exchange_rate`, `convert`, `open_period`, `close_period`, and `period_status`.

### Operational Notes

The extension rejects unbalanced ledger activity before commit when `pg_ledger.enabled` is on. `pg_ledger.strict_mode` controls whether mutations of ledger-amount columns without rules raise an error; both are superuser-settable session GUCs. Journal tables are protected by immutability triggers, so corrections use `reverse` rather than UPDATE or DELETE. The control file is non-relocatable and superuser-only. Validate rule-generated accounts, rounding, savepoints, fiscal closing, backup/restore, and application authorization before treating the journal as an accounting control.
