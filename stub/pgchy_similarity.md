## Usage

Sources:

- [Version 1.0 SQL objects](https://github.com/zandmitrij/pgchy_similarity/blob/76e4b9ba0530ba16a577141b0375268261626a37/sql/pgchy_similarity--1.0.sql)
- [Tanimoto implementation](https://github.com/zandmitrij/pgchy_similarity/blob/76e4b9ba0530ba16a577141b0375268261626a37/src/pgchy_similarity.c)
- [Benchmark notes](https://github.com/zandmitrij/pgchy_similarity/blob/76e4b9ba0530ba16a577141b0375268261626a37/benchmark/BENCHMARK.md)

`pgchy_similarity` computes thresholded Tanimoto similarity for molecular fingerprints. The SQL helper `smiles_to_fingerprint(text)` uses untrusted PL/Python, Chython, and NumPy to build a 2048-bit Morgan fingerprint; `is_tanimoto_similar(bytea, bytea, float4)` compares two packed fingerprints.

```sql
CREATE EXTENSION plpython3u;
CREATE EXTENSION pgchy_similarity;
WITH q AS (SELECT smiles_to_fingerprint('CCO') AS fp)
SELECT molecule_id
FROM molecule, q
WHERE is_tanimoto_similar(q.fp, molecule.fingerprint, 0.7);
```

Fingerprint meaning depends on exact Chython/NumPy versions, radius, length, bit order, and packed layout. Pin the Python environment and validate known molecule pairs. A change requires recomputing every stored fingerprint.

Untrusted PL/Python is superuser-controlled and can access the server operating system. The C function interprets a packed `bytea` layout directly, so reject malformed or wrong-length inputs before native code. Tanimoto thresholding is a screening method, not chemical identity or safety validation. Benchmark false positives/negatives and memory/CPU, and restrict both fingerprint generation and comparison functions to reviewed roles.
