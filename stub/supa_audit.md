## Usage

Sources:

- [Archived project README](https://github.com/supabase/supa_audit/blob/4f527f062e85ec0549c45009c3860d50ddf7f282/README.md)
- [Extension control file](https://github.com/supabase/supa_audit/blob/4f527f062e85ec0549c45009c3860d50ddf7f282/supa_audit.control)
- [Version 0.3.1 SQL implementation](https://github.com/supabase/supa_audit/blob/4f527f062e85ec0549c45009c3860d50ddf7f282/supa_audit--0.3.1.sql)
- [Regression tests](https://github.com/supabase/supa_audit/tree/4f527f062e85ec0549c45009c3860d50ddf7f282/test/sql)

`supa_audit` 0.3.1 is an archived pure-SQL trigger audit extension for PostgreSQL 12 and later. It stores insert, update, delete, and truncate events in `audit.record_version`. For row events it records JSONB current and previous rows and derives UUID record identifiers from the table OID and primary-key values.

### Enable tracking on a table

```sql
CREATE EXTENSION supa_audit CASCADE;

CREATE TABLE public.account (
  id bigint PRIMARY KEY,
  display_name text NOT NULL
);

SELECT audit.enable_tracking('public.account'::regclass);

SELECT id, record_id, old_record_id, op, ts, record, old_record
FROM audit.record_version
WHERE table_oid = 'public.account'::regclass::oid
ORDER BY id;
```

`CASCADE` installs the required `uuid-ossp` extension if necessary. Tracking requires a primary key. Query by indexed `table_oid`, and define separate retention and archive jobs because the audit table grows with writes.

### Audit is transactional, sensitive data

Trigger rows commit and roll back with the source transaction. They are not an independent tamper-evident forensic log: table owners or privileged roles can disable triggers, alter history, or avoid the path. A truncate event records only a marker, not every deleted row. Export important events to separately controlled storage.

Every tracked column, including credentials, tokens, personal data, and large values, can be copied into current and old JSONB. Restrict `audit` schema access, encrypt and minimize retained data, and include it in privacy deletion and backup policies. Upstream warns that auditing reduces write throughput and does not recommend it above 3,000 peak operations per second. Table OIDs contribute to record identifiers and are not portable identities, so prove continuity across dump and restore. The project is archived; its security-definer functions and nonfixed UUID-function search path require review before any continued deployment.
