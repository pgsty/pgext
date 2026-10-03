## Usage

Sources:

- [pg_bio.control](https://github.com/jonatas/pg_bio/blob/93b50d613ba0fd9555df386f359ebf9bb2f9cd91/pg_bio.control)
- [README.md](https://github.com/jonatas/pg_bio/blob/93b50d613ba0fd9555df386f359ebf9bb2f9cd91/README.md)
- [Cargo.toml](https://github.com/jonatas/pg_bio/blob/93b50d613ba0fd9555df386f359ebf9bb2f9cd91/Cargo.toml)
- [src/lib.rs](https://github.com/jonatas/pg_bio/blob/93b50d613ba0fd9555df386f359ebf9bb2f9cd91/src/lib.rs)

`pg_bio` provides experimental Rust SQL primitives for molecular coordinates, sequence comparison, fingerprints and biological data ingestion. Its control uses the Cargo version 0.0.0; this is an early source snapshot.

### Core Workflow

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pg_bio;
SELECT distance_angstroms(
  create_residue_coord(0, 0, 0, 'A'),
  create_residue_coord(3, 4, 0, 'B')
);
SELECT sequence_alignment_score('ACGT', 'ACGA');
```

### Operational Boundaries

The control requires `vector` and superuser installation; no preload is specified. Coordinate helpers include `create_residue_coord`, `distance_angstroms` and `z_order_encode`; the latter clamps coordinates to the source-defined range. Other interfaces include sequence alignment, SMILES/fingerprint helpers and UniProt retrieval. Network/file ingestion has server-side effects and must be restricted. Model-named embedding helpers are not evidence of trained-model inference; do not infer clinical, predictive or performance guarantees. PostgreSQL major compatibility and licensing are left unspecified where source evidence is insufficient.
