## Usage

Sources:

- [Project README](https://github.com/no0p/postbis/blob/ce454ebfbc27e0b6c8357ef6bfc8da1c4b2967c8/README.txt)
- [Extension control file](https://github.com/no0p/postbis/blob/ce454ebfbc27e0b6c8357ef6bfc8da1c4b2967c8/postbis.control)
- [Version 1.0 SQL API](https://github.com/no0p/postbis/blob/ce454ebfbc27e0b6c8357ef6bfc8da1c4b2967c8/sql/postbis--1.0.sql)
- [Sequence regression tests](https://github.com/no0p/postbis/tree/ce454ebfbc27e0b6c8357ef6bfc8da1c4b2967c8/test/sql)

`postbis` 1.0 provides compact native types for DNA, RNA, amino-acid, and aligned sequences. It also provides configurable alphabets and type modifiers, casts, sequence operations, biological transformations, comparison operators, and B-tree and hash operator classes.

### Store typed sequences

```sql
CREATE EXTENSION postbis;

CREATE TABLE specimen (
  specimen_id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  dna dna_sequence(SHORT, FLC, CASE_SENSITIVE) NOT NULL,
  rna rna_sequence(IUPAC, CASE_SENSITIVE),
  protein aa_sequence(IUPAC, CASE_SENSITIVE)
);

INSERT INTO specimen (dna, rna, protein)
VALUES ('AACCGGTT', 'AACGUU', 'ACDEFG');

SELECT specimen_id,
       char_length(dna) AS bases,
       substr(dna, 3, 4)::text AS fragment
FROM specimen;
```

Input validation depends on the selected alphabet, case-sensitivity, and type modifiers. Verify that casts reject symbols outside the required biological convention and that aligned and unaligned types are not mixed accidentally.

### Transform and translate sequences

```sql
SELECT complement('ACGTN'::dna_sequence)::text;
-- TGCAN

SELECT reverse_complement('ACGTN'::dna_sequence)::text;
-- NACGT

SELECT transcribe('AACGTT'::dna_sequence)::text;
-- AACGUU

SELECT translate('AUGGCCUAA'::rna_sequence)::text;
-- MA
```

The extension also exposes `reverse_transcribe()`, `six_frame()`, `get_alphabet()`, `entropy()`, `gc_content()`, and sequence generators. The translation functions accept explicit translation tables when the standard genetic code is not appropriate.

### Inspect compression and add indexes

```sql
SELECT char_length(sequence) AS symbols,
       octet_length(sequence) AS storage_bytes,
       compression_ratio(sequence) AS storage_ratio
FROM (
  SELECT repeat('ACGT', 256)::dna_sequence AS sequence
) AS sample;

CREATE INDEX specimen_dna_btree ON specimen USING btree (dna);
CREATE INDEX specimen_dna_hash  ON specimen USING hash  (dna);
```

Equality, ordering, concatenation, substring, search, and length functions are available for the sequence types. Check plans and realistic data distributions before relying on an index for a production workload.

### Packaging and durability risk

Pigsty applies a downstream compatibility patch and packages PostBIS 1.0 for PostgreSQL 14–18. That packaging result does not change the upstream lifecycle: the project is inactive and has no extension upgrade path beyond 1.0.

The custom types use native compressed on-disk representations. Treat stored values and indexes as tied to an exact tested build. Before adoption or migration, prove dump and restore, binary and logical upgrades, replication, driver decoding, index rebuilds, malformed input handling, and large-sequence memory behavior.

Functions such as `reverse()`, `char_length()`, and `substr()` overload familiar names, so schema qualification and controlled `search_path` settings matter. For new durable datasets, prefer maintained sequence tooling or plain PostgreSQL types unless the extension has been locally audited, packaged, and assigned an explicit long-term migration owner.
