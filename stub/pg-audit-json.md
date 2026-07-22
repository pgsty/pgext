## Usage

Sources:

- [Official README](https://github.com/m-martinez/pg-audit-json/blob/68bbc4417c44dac145084ccb703a24e2a199977c/README.md)
- [Version 1.0.2 extension SQL](https://github.com/m-martinez/pg-audit-json/blob/68bbc4417c44dac145084ccb703a24e2a199977c/sql/pg-audit-json--1.0.2.sql)
- [Official audit tests](https://github.com/m-martinez/pg-audit-json/blob/68bbc4417c44dac145084ccb703a24e2a199977c/test/sql/audit_test.sql)

`pg-audit-json` 1.0.2 is a trigger-based table audit system. It records INSERT, UPDATE, DELETE, and TRUNCATE events in `audit.log`, storing old rows and recursive `jsonb` diffs together with session, query, transaction, client, and timestamp metadata.

### Enable Auditing

Install as a PostgreSQL administrator, then have the owner of `audit.log` attach triggers to each target table:

```sql
CREATE EXTENSION "pg-audit-json";

CREATE TABLE app.customer (
  id bigint PRIMARY KEY,
  email text,
  profile jsonb,
  updated_at timestamptz
);

SELECT audit.audit_table(
  'app.customer'::regclass,
  true,
  true,
  ARRAY['updated_at']::text[]
);

SELECT action, row_data, changed_fields, session_user_name
FROM audit.log
WHERE relid = 'app.customer'::regclass
ORDER BY id DESC;
```

The four-argument `audit.audit_table` selects row-level logging, query-text logging, and ignored columns. The one-argument wrapper enables row values and query text with no ignored columns. Statement-level events are still used for TRUNCATE; if row auditing is disabled, INSERT/UPDATE/DELETE are recorded without row values.

Applications may set `audit.application_name` and `audit.application_user_name` with `SET LOCAL` to attach application identity. INSERT stores the new row in `changed_fields`; UPDATE stores the old row in `row_data` and only changed values in `changed_fields`; DELETE stores the old row; ignored-only updates are omitted.

### Security and Coverage Boundaries

Audit rows are written in the same transaction as the change, so rollback removes both. Table owners, superusers, or code able to disable/drop triggers can bypass it. It does not audit SELECT, general DDL, role changes, failed statements, or actions on tables without its triggers; use server audit logging for those requirements.

`client_query` and full row JSON can contain credentials or personal data. Revoke direct access to `audit.log`, define retention and redaction rules, and account for storage/index/write amplification. The extension configures `audit.log` as dumpable extension data, so backups may contain the audit payload. Schema changes can alter row shape, and dropping/recreating a table changes `relid`; query by both identity metadata and retention epoch where necessary.
