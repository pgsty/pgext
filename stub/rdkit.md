## Usage

Sources:

- [Cartridge documentation](https://github.com/rdkit/rdkit/blob/Release_2026_03_6/Docs/Book/Cartridge.md)
- [Control](https://github.com/rdkit/rdkit/blob/Release_2026_03_6/Code/PgSQL/rdkit/rdkit.control)
- [SQL template](https://github.com/rdkit/rdkit/blob/Release_2026_03_6/Code/PgSQL/rdkit/rdkit.sql.in)
- [Upgrade script](https://github.com/rdkit/rdkit/blob/Release_2026_03_6/Code/PgSQL/rdkit/update_sql/rdkit--4.7.0--4.8.0.sql.in)

`rdkit` provides molecular types, substructure search, fingerprints and chemical descriptors. This page follows the PostgreSQL cartridge shipped in RDKit 2026.03.6.

### Create The Extension

```sql
CREATE EXTENSION rdkit;
```

The cartridge adds chemistry-specific types including `mol`, `bfp`, and `sfp`.

### Core Search Operators

The cartridge documentation covers:

- `@>` and `<@` for substructure matching.
- `@=` for exact molecular equality.
- `%`, `<%>`, and `<#>` style fingerprint similarity and KNN operators for similarity search.

These are typically combined with GiST indexes over fingerprint columns.

### Fingerprints And Similarity

Common fingerprint functions documented for SQL usage include `morgan_fp`, `morganbv_fp`, `featmorgan_fp`, `rdkit_fp`, `atompair_fp`, `torsion_fp`, `layered_fp`, and `maccs_fp`.

Example from the cartridge docs:

```sql
SELECT tanimoto_sml(
  morganbv_fp('c1ccccc1'::mol),
  morganbv_fp('c1ccccc1O'::mol)
);
```

### Descriptors And Validation

The cartridge docs also expose validation and descriptor helpers such as:

- `is_valid_smiles()`
- `is_valid_ctab()`
- `is_valid_smarts()`
- `mol_amw()`
- `mol_hba()`
- `mol_numrings()`

These functions are the main user-facing surface for SQL analytics on molecular structures.

### Version and Upgrade Boundary

The RDKit toolkit release 2026.03.6 contains SQL extension version `4.8.0`; these are different version namespaces. The upstream `4.7.0` to `4.8.0` SQL template changes function costs, but its `fmcs_smiles` statement has a line comment that swallows the cost clause and terminator. Do not assume an unpatched `ALTER EXTENSION rdkit UPDATE` path works: verify the installed upgrade script and test the update on a restored copy first. A packaging patch is distinct from the upstream source.
