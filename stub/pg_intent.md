## Usage

Sources:

- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_intent/pg_intent.control)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_intent/pg_intent--1.1.sql)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_intent/meson.build)

`pg_intent` `1.1` provides rule-based document intent classification. This preview module is scoped to the PostgreSQL 18 fork in AlohaDB 21.1.0; stock PostgreSQL compatibility is not established.

### Workflow

```sql
CREATE EXTENSION pg_intent;
SELECT intent_register('refund', 'Refund request', ARRAY['refund','return']);
SELECT * FROM intent_infer('Please refund this order');
```

Installation requires a superuser; function use also requires the appropriate object privileges.

### Operational boundaries

An intent registry, parent hierarchy, document tags and keyword phrase matching drive classification. intent_tag, intent_auto_tag and intent_find manage tagged documents; cycle checks protect the hierarchy. This is keyword/rule classification rather than probabilistic inference. Review dictionary edits and grants on the extension-owned public tables.

This module corresponds to the former upstream name `alohadb_intent`. Follow AlohaDB’s cluster/object conversion procedure and do not substitute an unrelated implementation with a similar name. Rehearse dependencies, privileges and backup restoration before upgrades.
