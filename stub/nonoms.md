## Usage

Sources:

- [nonoms README at the reviewed commit](https://github.com/Edwards-R/nonoms/blob/4ef6c5fcf05f0cf03a6ed852a20d2f791f539c5c/readme.md)
- [nonoms 1.2.0 install SQL at the reviewed commit](https://github.com/Edwards-R/nonoms/blob/4ef6c5fcf05f0cf03a6ed852a20d2f791f539c5c/nonoms--1.2.0.sql)
- [nonoms initialization procedure at the reviewed commit](https://github.com/Edwards-R/nonoms/blob/4ef6c5fcf05f0cf03a6ed852a20d2f791f539c5c/procedure/init_nonoms.sql)

`nonoms` 1.2.0 implements a normalized biological-nomenclature model following NoNomS protocols. It creates rank-specific tables and tracks current names, synonyms, compositions, merges, and splits. Install it in a dedicated empty schema, then initialize the capstone and add ranks in parent-first order.

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

`init_nonoms` creates the `rank` table and the initial capstone rank. Subsequent procedures dynamically create one nomenclature table and one composition table per rank.

### Caveats

- Upstream explicitly describes the project as beta. Validate its nomenclatural rules and migration behavior against the application's domain model.
- Initialization is designed for a fresh extension schema and refuses a populated schema unless its override flag is used. Do not enable that override without reviewing every existing object.
- Rank operations execute dynamic DDL and create triggers, foreign keys, and indexes. Run initialization and structural changes as controlled migrations.
- The reviewed repository contains no declared license or compatibility matrix for current PostgreSQL majors.
