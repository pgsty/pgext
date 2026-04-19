## Usage

- Sources: [project README](https://github.com/rdkit/rdkit/blob/master/README.md), [cartridge docs](https://www.rdkit.org/docs/Cartridge.html), [2025.03.6 release](https://github.com/rdkit/rdkit/releases/tag/Release_2025.03.6)

RDKit ships a PostgreSQL cartridge for cheminformatics storage, search, fingerprints, and descriptors. The cartridge docs remain the main upstream usage reference; the 2025.03.6 release notes do not call out cartridge-specific user-facing changes.

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
