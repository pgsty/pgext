## Usage

Sources:

- [README](https://github.com/RahmanTektas/postgresql-dna-extension/blob/8b5d85bcc5f3f75d2ca8f8b4b33f3cf265bdab2d/README.md)
- [Control](https://github.com/RahmanTektas/postgresql-dna-extension/blob/8b5d85bcc5f3f75d2ca8f8b4b33f3cf265bdab2d/dna_sequence/dna_sequence.control)
- [SQL](https://github.com/RahmanTektas/postgresql-dna-extension/blob/8b5d85bcc5f3f75d2ca8f8b4b33f3cf265bdab2d/dna_sequence/dna_sequence--1.0.sql)
- [Source](https://github.com/RahmanTektas/postgresql-dna-extension/blob/8b5d85bcc5f3f75d2ca8f8b4b33f3cf265bdab2d/dna_sequence/dna_sequence.c)
- [SP-GiST tests](https://github.com/RahmanTektas/postgresql-dna-extension/blob/8b5d85bcc5f3f75d2ca8f8b4b33f3cf265bdab2d/tests/test_spgist.sql)

`dna_sequence` is an academic C extension for genomic sequence values and indexed k-mer searches. This source snapshot declares extension version 1.0; upstream does not publish a PostgreSQL-major compatibility matrix or an explicit license. Validate it on a disposable database before adopting its on-disk types.

### Core Workflow

After installing the shared library and SQL files, a superuser can create the extension. Its control file permits schema relocation; no preload setting or separate extension dependency is declared.

```sql
CREATE EXTENSION dna_sequence;
SELECT length('ACGTACGT'::dna);
SELECT * FROM generate_kmers('ACGTACGT'::dna, 3);

CREATE TABLE dna_kmers (value kmer);
INSERT INTO dna_kmers
SELECT * FROM generate_kmers('ACGTACGT'::dna, 3);
CREATE INDEX dna_kmers_spgist ON dna_kmers USING spgist (value);
SELECT value FROM dna_kmers WHERE value ^@ 'AC'::kmer;
SELECT value FROM dna_kmers WHERE value <@ 'ANG'::qkmer;
```

`dna` stores sequences of the four concrete DNA bases; `kmer` accepts 1–32 concrete bases. `qkmer` is a query pattern with IUPAC ambiguity codes. `generate_kmers(dna, integer)` returns the overlapping k-mers as a set, and `length` has overloads for all three types.

### Query Surface

| Object | Meaning |
| --- | --- |
| `equals(kmer,kmer)`, `=` and `<>` | Exact k-mer equality and inequality |
| `starts_with(kmer,kmer)`, `^@` | Test whether a k-mer starts with a prefix |
| `contains(qkmer,kmer)`, `contained(kmer,qkmer)`, `@>` and `<@` | Test a concrete k-mer against an ambiguity-code query pattern |
| `kmer_hash_ops` | Hash indexing and grouping support |
| `kmer_btree_ops` | Ordered comparisons and B-tree indexing |
| `kmer_spgist_ops` | SP-GiST support for equality, prefix and pattern predicates |

### Boundaries

Types reject invalid bases and invalid k-mer lengths. The extension supplies native type and index code, so logical dumps and a tested restore path matter when changing builds. Its unqualified type names can collide with other genomics extensions in the same schema; use the actual canonical extension name `dna_sequence`, not the separate catalog entry `dna`.

The upstream test/build helpers include commands that drop the extension with dependent objects. Do not run that development reset workflow against stored application data. This documentation describes the declared SQL interface, not a production-support or runtime-test certification.
