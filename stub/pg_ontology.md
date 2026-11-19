## Usage

Sources:

- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_ontology/pg_ontology.control)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_ontology/pg_ontology--1.1.sql)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_ontology/meson.build)

`pg_ontology` `1.1` provides ontology triples, hierarchy traversal and history. This preview module is scoped to the PostgreSQL 18 fork in AlohaDB 21.1.0; stock PostgreSQL compatibility is not established.

### Workflow

```sql
CREATE EXTENSION pg_ontology;
SELECT ontology_assert('Honolulu', 'capital_of', 'Hawaii');
SELECT * FROM ontology_query('Honolulu', NULL, NULL);
```

Installation requires a superuser; function use also requires the appropriate object privileges.

### Operational boundaries

ontology_subclass manages a cycle-checked type hierarchy; ontology_related supports bounded graph traversal; ontology_history records assertions and retractions. This is a relational triple store, not a general OWL inference engine. Control writes to ontology tables and review history retention. Version 1.1 supplies a matching install/upgrade script with bidirectional traversal and change tracking.

This module corresponds to the former upstream name `alohadb_ontology`. Follow AlohaDB’s cluster/object conversion procedure and do not substitute an unrelated implementation with a similar name. Rehearse dependencies, privileges and backup restoration before upgrades.
