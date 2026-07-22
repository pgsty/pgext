## Usage

Sources:

- [Official README](https://github.com/Edwards-R/nonoms/blob/4ef6c5fcf05f0cf03a6ed852a20d2f791f539c5c/readme.md)
- [Version 1.2.0 SQL implementation](https://github.com/Edwards-R/nonoms/blob/4ef6c5fcf05f0cf03a6ed852a20d2f791f539c5c/nonoms--1.2.0.sql)
- [Initialization documentation](https://github.com/Edwards-R/nonoms/blob/4ef6c5fcf05f0cf03a6ed852a20d2f791f539c5c/docs/procedure/init_nonoms.md)
- [Rank-creation documentation](https://github.com/Edwards-R/nonoms/blob/4ef6c5fcf05f0cf03a6ed852a20d2f791f539c5c/docs/procedure/insert_rank.md)

`nonoms` 1.2.0 implements a normalized biological-nomenclature model following NoNomS protocols. It creates a configurable rank hierarchy and records current understandings, synonyms, compositions, merges, and splits. Upstream labels the project beta; validate its model against the nomenclatural rules of the consuming application.

### Initialize a New Nomenclature

Use a dedicated empty schema named `nomenclature`. The 1.2.0 split implementation contains a literal `nomenclature` schema reference, so another extension schema is not fully relocatable in practice.

```sql
CREATE SCHEMA nomenclature;
CREATE EXTENSION nonoms WITH SCHEMA nomenclature;

CALL nomenclature.init_nonoms('Example scheme', 2026);
CALL nomenclature.insert_rank('kingdom', 1, 'Kingdom');
CALL nomenclature.insert_current_understanding(
  2, 1, 'Animalia', 'Linnaeus', 1758
);

SELECT * FROM nomenclature.rank ORDER BY id;
SELECT * FROM nomenclature.kingdom ORDER BY id;
```

`init_nonoms` creates the `rank` catalog and the capstone rank. Add every rank in parent-first order before loading names; upstream states that inserting ranks in the middle or adding ranks after population is not supported.

### Main Procedures and Types

- `insert_rank`: creates one rank table, its composition table, indexes, foreign keys, and trigger.
- `insert_current_understanding` and `insert_synonym_understanding`: add current names and synonyms.
- `merge_understandings`: replaces multiple current understandings with a merged understanding and propagates children.
- `split_understanding`: creates destinations and composition records using `split_result[]`.
- `create_binomial_view`: creates the cross-rank binomial view after the rank structure is ready.
- `rank`, per-rank tables, and per-rank `_composition` tables: persistent nomenclature state.

### Migration and Integrity Boundaries

Initialization refuses a populated schema unless `override => true`; do not use the override without a full object review and recoverable backup. Rank creation and merge/split operations execute dynamic DDL or recursive data updates and create constraints, triggers, indexes, and temporary tables. Run them as controlled migrations, not inside untrusted application requests.

The 1.2.0 source contains no current PostgreSQL compatibility matrix and no declared license file in the reviewed repository. It also has hard-coded schema assumptions in `split_understanding`. Pin the source, test every procedure on a copy of real hierarchy data, and verify current/synonym pointers, compositions, child propagation, rollback, dump, and restore before production adoption.
