## Usage

> Sources: [GitHub README](https://raw.githubusercontent.com/rdkit/rdkit/master/README.md), [Cartridge docs](https://www.rdkit.org/docs/Cartridge.html), [project site](https://www.rdkit.org)
> RDKit provides a PostgreSQL cartridge for cheminformatics storage, search, fingerprints, and descriptors.

```sql
CREATE EXTENSION rdkit;
```

The cartridge adds molecular types and operators for substructure search, exact match, similarity search, fingerprint generation, and many descriptor functions.

### Core Types

- `mol` for molecule and query molecule values.
- `bfp` for bit-vector fingerprints.
- `sfp` for sparse count fingerprints.

### Search Operators

- `@>` checks whether the right-hand molecule is a substructure of the left-hand molecule.
- `<@` performs the reverse substructure check.
- `@=` checks exact molecular equality.
- `#` is the Dice similarity threshold operator.
- `<%>` supports Tanimoto KNN searches.
- `<#>` supports Dice KNN searches.

### Fingerprints

- `morgan_fp` and `morganbv_fp` for Morgan fingerprints.
- `featmorgan_fp` and `featmorganbv_fp` for feature-based Morgan fingerprints.
- `rdkit_fp` for hashed RDKit fingerprints.
- `atompair_fp`, `atompairbv_fp`, `torsion_fp`, and `torsionbv_fp`.
- `layered_fp` and `maccs_fp`.

### Descriptor Examples

```sql
SELECT is_valid_smiles('c1ccccc1');
SELECT tanimoto_sml(morganbv_fp('c1ccccc1'::mol), morganbv_fp('c1ccccc1O'::mol));
SELECT mol_amw('c1ccccc1'::mol);
SELECT mol_hba('c1ccccc1O'::mol);
SELECT mol_numrings('c1ccccc1'::mol);
```

### Notes

- The docs describe additional molecule-validation helpers such as `is_valid_smiles`, `is_valid_ctab`, and `is_valid_smarts`.
- The cartridge also exposes many additional descriptor and fingerprint size GUCs, but the stub keeps only the core user-facing entry points.
