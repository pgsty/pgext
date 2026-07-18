## Usage

Sources:

- [Project README](https://github.com/no0p/postbis/blob/ce454ebfbc27e0b6c8357ef6bfc8da1c4b2967c8/README.txt)
- [Extension control file](https://github.com/no0p/postbis/blob/ce454ebfbc27e0b6c8357ef6bfc8da1c4b2967c8/postbis.control)
- [Version 1.0 SQL API](https://github.com/no0p/postbis/blob/ce454ebfbc27e0b6c8357ef6bfc8da1c4b2967c8/sql/postbis--1.0.sql)
- [Sequence regression tests](https://github.com/no0p/postbis/tree/ce454ebfbc27e0b6c8357ef6bfc8da1c4b2967c8/test/sql)

`postbis` 1.0 is an abandoned native extension for compressed biological sequences. It defines `dna_sequence`, `rna_sequence`, `aa_sequence`, aligned variants, configurable alphabets and type modifiers, plus casts, substring and length operations, transformations, comparison operators, and B-tree and hash support.

### Store and transform a DNA sequence

```sql
CREATE EXTENSION postbis;

CREATE TABLE specimen (
  specimen_id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  sequence dna_sequence NOT NULL
);

INSERT INTO specimen (sequence) VALUES ('ACGTACGT'::dna_sequence);

SELECT specimen_id,
       char_length(sequence) AS bases,
       reverse_complement(sequence)::text AS reverse_complement
FROM specimen;
```

Input validation depends on the selected alphabet, case-sensitivity, and type modifiers. Verify that casts reject symbols outside the required biological convention and that aligned and unaligned types are not mixed accidentally.

### Durability and compatibility risk

The custom types use native compressed on-disk representations. The repository contains old imported Subversion metadata, no current release or supported-PostgreSQL matrix, and no upgrade path beyond 1.0. Treat stored values and indexes as tied to an exact tested build. Before adoption or migration, prove dump and restore, binary and logical upgrades, replication, driver decoding, index rebuilds, malformed input handling, and large-sequence memory behavior.

Functions such as `reverse()`, `char_length()`, and `substr()` overload familiar names, so schema qualification and controlled `search_path` settings matter. For new durable datasets, prefer maintained sequence tooling or plain PostgreSQL types unless the extension has been locally audited, packaged, and assigned an explicit long-term migration owner.
