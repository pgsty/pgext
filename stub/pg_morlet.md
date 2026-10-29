## Usage

Sources:

- [MorletDB README at v0.3.0-alpha.1](https://github.com/merkhet-io/morletdb/blob/e85a24d1775cc971b2c4b444fbf3863910f8582c/README.md)
- [pg_morlet README](https://github.com/merkhet-io/morletdb/blob/e85a24d1775cc971b2c4b444fbf3863910f8582c/extensions/pg_morlet/README.md)
- [Control file](https://github.com/merkhet-io/morletdb/blob/e85a24d1775cc971b2c4b444fbf3863910f8582c/extensions/pg_morlet/pg_morlet.control)
- [SQL upgrade](https://github.com/merkhet-io/morletdb/blob/e85a24d1775cc971b2c4b444fbf3863910f8582c/extensions/pg_morlet/sql/pg_morlet--0.2.0--0.3.0.sql)

`pg_morlet` 0.3.0 stores waveform recordings and compares compatible windows using versioned Morlet features. Release 0.3.0-alpha.1 is a developer alpha for PostgreSQL 17 with pgrx 0.16.1; it adds acquisition timelines and context-aware comparisons while keeping the synchronous, transactional storage model.

### Record and search

```sql
CREATE EXTENSION pg_morlet;
SELECT morlet_create_stream_collection('public', 'signals');
SELECT morlet_enable_acquisitions('signals'::regclass);
SELECT morlet_create_acquisition('signals'::regclass, 'run-001', 'motor-1');
```

Acquisitions align channels on an integer nanosecond timeline with declared offsets, clock source and uncertainty. `morlet_build_acquisition_segments`, `morlet_read_acquisition` and `morlet_compare_acquisition` expose the new aligned workflow. The versioned SQL script defines the acquisition setup and signatures; comparisons additionally require populated recordings and built context segments. It does not resample channels or estimate their clock offsets.

### Consistency and access

Ingestion and feature publication remain transactional. Caller privileges and row-security policies apply to collection tables; direct edits can break cross-table invariants. Searches compare only compatible rates, channels, units, windows and extractor versions. Similarity scores are investigation signals, not calibrated probabilities.

### Upgrade boundary

The SQL version is 0.3.0. The tested update script accepts SQL 0.2.0 from 0.2.0-alpha.1 and the chained 0.1.0 SQL from 0.1.0-alpha.3. Older 0.1.0-alpha.1 installations need a reinstall. This remains a developer alpha with no stable schema or production-capacity guarantee.
