## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/BrandwatchLtd/trunkfingerprint/blob/a4ed83cd533f233e70b23aac3ebb7d91248c7087/README.md)
- [Version 1.3.0 SQL implementation](https://github.com/BrandwatchLtd/trunkfingerprint/blob/a4ed83cd533f233e70b23aac3ebb7d91248c7087/trunkfingerprint--1.3.0.sql)
- [Extension control file](https://github.com/BrandwatchLtd/trunkfingerprint/blob/a4ed83cd533f233e70b23aac3ebb7d91248c7087/trunkfingerprint.control)

`trunkfingerprint` computes comparable MD5 fingerprints over PostgreSQL catalog structure and, optionally, table data. Use it to detect drift between databases built from the same expected schema—not as a backup verifier or cryptographic integrity system.

```sql
CREATE EXTENSION trunkfingerprint;
SELECT * FROM trunkfingerprint.get_db_fingerprint(
  p_level => 0,
  p_data => false,
  p_exclude_schemas => ARRAY['monitoring']::name[]
);
```

`p_level` 0 returns one fingerprint, 1 returns a value per examined relation, and 2 emits object-level input. `p_table` limits the scan; `p_data` selects structure only, data only, or both; exclusion parameters remove schemas, tables, or columns from data inspection.

The implementation normalizes many OID references and excludes volatile catalog columns, but fingerprints still depend on PostgreSQL version, extension versions, object definitions, collation, and the tool's own normalization rules. Compare like-for-like builds and use level 1 or 2 to diagnose a mismatch.

Data mode scans and serializes table contents and may be expensive, expose sensitive values to the executing role, and change under concurrent writes. Run against a consistent snapshot or quiescent copy, budget I/O and temporary space, and restrict function access. MD5 collision resistance is insufficient for adversarial proof; use authenticated backup manifests for that purpose.
