## Usage

Sources:

- [Official README](https://github.com/zvectorlabs/pgauthz/blob/a24a3bae33d1bac320c233e99ceb9153a4d0667d/README.md)
- [Official quick start](https://github.com/zvectorlabs/pgauthz/blob/a24a3bae33d1bac320c233e99ceb9153a4d0667d/docs/quickstart.md)
- [API reference](https://github.com/zvectorlabs/pgauthz/blob/a24a3bae33d1bac320c233e99ceb9153a4d0667d/docs/api-reference.md)
- [Configuration guide](https://github.com/zvectorlabs/pgauthz/blob/a24a3bae33d1bac320c233e99ceb9153a4d0667d/docs/configuration.md)

pgauthz implements Zanzibar-style relationship-based authorization inside PostgreSQL. Policies define object and subject types, relations, computed permissions, conditions, and wildcards; relationship tuples then drive checks. It targets PostgreSQL 16 through 18 and keeps policy, tuple, revision, changelog, and assertion data in the `authz` schema.

### Core Workflow

Define a policy, write a relationship, and check it on the same database transaction boundary as application data:

```sql
CREATE EXTENSION pgauthz;

SELECT pgauthz_define_policy('
  type user {}
  type document {
    relations
      define viewer: [user]
      define editor: [user]
    permissions
      define view = viewer + editor
  }
');

SELECT pgauthz_add_relation(
  'document', 'doc1', 'viewer', 'user', 'alice'
);

SELECT pgauthz_check(
  'document', 'doc1', 'view', 'user', 'alice'
);
```

Use `pgauthz_check_with_context` when a declared condition requires runtime JSON context, and enforce the boolean result in the same statement or security-definer boundary that protects the resource.

### Important Objects

- `pgauthz_define_policy` parses and versions the active authorization model.
- `pgauthz_add_relation` and `pgauthz_write_relationships` create or change relationship tuples.
- `pgauthz_check` and `pgauthz_check_with_context` evaluate a relation or computed permission.
- `pgauthz_list_objects` and `pgauthz_list_subjects` enumerate reachable resources or principals.
- `pgauthz_expand` returns a permission tree for debugging.
- `pgauthz_read_relationships`, `pgauthz_read_latest_policy`, and `pgauthz_read_changes` expose stored state and changes.
- `authz.*` GUCs control check strategy, cache TTL and capacity, revision quantization, tracing, and OpenTelemetry export.

### Security and Consistency Notes

An authorization function does not enforce access unless every data path actually calls it or a correct RLS/policy wrapper. Revoke direct access to protected tables and to the underlying `authz` relations from application roles. Caches are disabled by default; enabling TTLs or revision quantization trades freshness for speed, so test revocation latency explicitly. Treat policy changes and relationship writes as security-sensitive audited operations, validate condition context types, and protect OpenTelemetry endpoints from leaking authorization metadata.
