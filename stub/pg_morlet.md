## Usage

Sources:

- [Official source](https://github.com/merkhet-io/morletdb/blob/paper-v0.31/extensions/pg_morlet/pg_morlet.control)
- [Official source](https://github.com/merkhet-io/morletdb/blob/paper-v0.31/docs/compatibility.md)
- [Official source](https://github.com/merkhet-io/morletdb/blob/paper-v0.31/extensions/pg_morlet/sql/pg_morlet--0.3.11--0.3.12.sql)
- [Official source](https://github.com/merkhet-io/morletdb/blob/paper-v0.31/extensions/pg_morlet/sql/pg_morlet--0.3.0--0.3.1.sql)
- [Official source](https://github.com/merkhet-io/morletdb/blob/paper-v0.31/extensions/pg_morlet/README.md)
- [Official source](https://github.com/merkhet-io/morletdb/blob/paper-v0.31/docs/recording-workflow.md)

`pg_morlet` 0.3.12 stores waveforms and compares compatible windows with versioned Morlet features. The exact paper-v0.31 source tag contains this control/SQL version. PostgreSQL 17 is the reference platform; 18 has functional CI coverage, while 19 was tested as a release candidate. This remains an experimental extension using pgrx 0.19.3.

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

### Continuous recordings and maintenance

The current workflow uses `morlet_start_recording`, `morlet_append_chunk` / `morlet_append_chunks`, `morlet_read_interval` and `morlet_search_interval`. Interval reads reconstruct chunks under caller privileges; incompatible or partly hidden intervals fail explicitly. Optional envelope probes require `vector`; ordinary storage and exact search do not.

The 0.3.11-to-0.3.12 SQL update corrects `morlet_build_envelopes` accounting for compacted episodes and `morlet_collection_status` backlog accounting. Install matching library and SQL files before updating:

```sql
ALTER EXTENSION pg_morlet UPDATE TO '0.3.12';
```

The current tag supplies a versioned upgrade chain. Upstream tests cover older alpha upgrades on PostgreSQL 17; 18/19 use fresh installations. Some README paragraphs still describe an obsolete 0.1.0 scaffold with no upgrade path; use the actual control, SQL chain and compatibility document for this revision. Preserve recordings and backups; do not drop a populated extension to follow the old walkthrough.
