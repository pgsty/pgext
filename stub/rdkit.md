


## Usage

> [rdkit: Cheminformatics and molecule toolkit PostgreSQL cartridge](https://github.com/rdkit/rdkit)

The RDKit PostgreSQL cartridge provides the `mol` datatype for molecules, `fp` datatype for fingerprints, substructure and similarity search operations, and GiST index support.

```sql
CREATE EXTENSION rdkit;
```

### Data Types

| Type | Description |
|------|-------------|
| `mol` | Molecular structure (from SMILES, SMARTS, etc.) |
| `bfp` | Bit vector fingerprint |
| `sfp` | Sparse (count) fingerprint |

### Molecule Input/Output

```sql
-- Create molecule from SMILES
SELECT 'c1ccccc1'::mol;

-- Check if SMILES is valid
SELECT is_valid_smiles('c1ccccc1');

-- Convert molecule to SMILES
SELECT mol_to_smiles('c1ccccc1'::mol);
```

### Substructure Search

```sql
-- Substructure match operator
SELECT 'c1ccccc1O'::mol @> 'c1ccccc1'::mol;   -- true (phenol contains benzene)
SELECT 'c1ccccc1'::mol <@ 'c1ccccc1O'::mol;    -- true

-- Using SMARTS patterns
SELECT 'c1ccccc1O'::mol @> 'c1ccc(O)cc1'::mol;
```

### Similarity Search

```sql
-- Tanimoto similarity (returns value between 0 and 1)
SELECT tanimoto_sml(morganbv_fp('c1ccccc1'::mol), morganbv_fp('c1ccccc1O'::mol));

-- Dice similarity
SELECT dice_sml(morganbv_fp('c1ccccc1'::mol), morganbv_fp('c1ccccc1O'::mol));
```

### Fingerprint Functions

```sql
-- Morgan fingerprint (bit vector)
SELECT morganbv_fp('c1ccccc1'::mol);
SELECT morganbv_fp('c1ccccc1'::mol, 2);  -- radius=2

-- RDKit fingerprint
SELECT rdkit_fp('c1ccccc1'::mol);

-- Topological torsion fingerprint
SELECT torsionbv_fp('c1ccccc1'::mol);

-- Atom pair fingerprint
SELECT atompairbv_fp('c1ccccc1'::mol);
```

### Descriptor Calculations

```sql
SELECT mol_amw('c1ccccc1'::mol);          -- average molecular weight
SELECT mol_logp('c1ccccc1'::mol);         -- LogP
SELECT mol_hba('c1ccccc1O'::mol);         -- H-bond acceptors
SELECT mol_hbd('c1ccccc1O'::mol);         -- H-bond donors
SELECT mol_numrotatablebonds('c1ccccc1'::mol); -- rotatable bonds
SELECT mol_numatoms('c1ccccc1'::mol);     -- number of atoms
SELECT mol_numheavyatoms('c1ccccc1'::mol);    -- heavy atoms
SELECT mol_numrings('c1ccccc1'::mol);     -- number of rings
```

### GiST Index Support

Create indexes for fast substructure and similarity searches:

```sql
-- Substructure search index
CREATE INDEX idx_mol ON molecules USING gist(molecule);

-- Fingerprint similarity index
CREATE INDEX idx_fp ON molecules USING gist(morganbv_fp(molecule));
```

Query with index support:

```sql
-- Substructure search
SELECT * FROM molecules WHERE molecule @> 'c1ccccc1'::mol;

-- Similarity search (with threshold)
SET rdkit.dice_threshold = 0.5;
SELECT * FROM molecules WHERE morganbv_fp(molecule) % morganbv_fp('c1ccccc1O'::mol);
```

### GUC Parameters

| Parameter | Default | Description |
|-----------|---------|-------------|
| `rdkit.tanimoto_threshold` | 0.5 | Threshold for Tanimoto similarity operator `<%>` |
| `rdkit.dice_threshold` | 0.5 | Threshold for Dice similarity operator `%` |
