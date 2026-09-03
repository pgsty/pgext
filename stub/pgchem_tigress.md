## Usage

Sources:

- [Extension control file](https://gitlab.com/api/v4/projects/telukir%2Fpgchem_tigress/repository/files/extension%2Fpgchem_tigress.control/raw?ref=a0e83cc13bf964861eb931592bb0c0f7574a0c5b)
- [Versioned install SQL](https://gitlab.com/api/v4/projects/telukir%2Fpgchem_tigress/repository/files/extension%2Fpgchem_tigress--3.2.sql/raw?ref=a0e83cc13bf964861eb931592bb0c0f7574a0c5b)
- [Official user guide source](https://gitlab.com/api/v4/projects/telukir%2Fpgchem_tigress/repository/files/doc%2Fuserguide%2Fpgchem_usersguide%2Fpgchem_usersguide.tex/raw?ref=a0e83cc13bf964861eb931592bb0c0f7574a0c5b)

`pgchem_tigress` is a legacy chemoinformatics extension that provides `molecule` and fingerprint types, chemical-property functions, exact/substructure operators, and a GiST operator class.

### Enablement

Install the 3.2 library and its OpenBabel, Barsoi, and InChI runtime dependencies, then create the relocatable extension as a superuser:

```sql
CREATE EXTENSION pgchem_tigress;
SELECT pgchem_version();
```

No preload is declared. Upstream is inactive and its source-era installation notes target PostgreSQL greater than 8 with hard-coded auxiliary-data paths; modern majors, library ABIs, dump/restore, and upgrade behavior require a dedicated compatibility test.

### Store and Index Molecules

The `molecule` input function accepts formats including SMILES, InChI, and MDL molfile text. A GiST index accelerates equality and substructure containment.

```sql
CREATE TABLE compounds (
  id  bigint PRIMARY KEY,
  mol molecule NOT NULL
);

CREATE INDEX compounds_mol_gist
ON compounds USING gist (mol);

INSERT INTO compounds VALUES
  (1, 'c1ccccc1'::molecule);
```

### Search Operators

`=` tests exact molecular equality. `A <= B` asks whether A is a substructure of B, and `A >= B` asks whether A contains B. `@` returns a Tanimoto coefficient and is not GiST-accelerated.

```sql
SELECT id
FROM compounds
WHERE 'c1ccccc1'::molecule <= mol;

SELECT id,
       'c1ccccc1'::molecule @ mol AS similarity
FROM compounds
ORDER BY similarity DESC;
```

Molecule serialization and fingerprints changed across historical releases. Rebuild affected indexes and migrate stored values according to the exact release notes when changing binaries; never swap the shared library underneath existing `molecule` columns without an explicit migration test.
