## Usage

Sources:

- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_context/pg_context.control)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_context/pg_context--1.1.sql)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_context/meson.build)

`pg_context` `1.1` provides entity disambiguation and context-sensitive ranking. This preview module is scoped to the PostgreSQL 18 fork in AlohaDB 21.1.0; stock PostgreSQL compatibility is not established.

### Workflow

```sql
CREATE EXTENSION pg_ontology;
CREATE EXTENSION pg_context;
SELECT context_register('Apple', 'customer:42', 'customer', NULL, 100, 'sales', ARRAY['invoice']);
SELECT * FROM context_resolve('Apple', 'Apple invoice');
```

Installation requires a superuser; function use also requires the appropriate object privileges.

### Operational boundaries

This extension requires the AlohaDB pg_ontology module. Profiles, aliases and keyword/ontology cues determine resolution scores; context_rerank adjusts candidate rankings. This is rule-based matching, not a trained semantic model. Review dictionary writes and public-schema privileges; 1.1 changes the profile key to support multiple senses.

This module corresponds to the former upstream name `alohadb_context`. Follow AlohaDB’s cluster/object conversion procedure and do not substitute an unrelated implementation with a similar name. Rehearse dependencies, privileges and backup restoration before upgrades.
