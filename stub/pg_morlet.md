## Usage

Sources:

- [README](https://github.com/merkhet-io/morletdb/blob/8b9e0af38ec2dffef6c6b5ae69982a29920a352a/extensions/pg_morlet/README.md)
- [Control](https://github.com/merkhet-io/morletdb/blob/8b9e0af38ec2dffef6c6b5ae69982a29920a352a/extensions/pg_morlet/pg_morlet.control)
- [Changelog](https://github.com/merkhet-io/morletdb/blob/8b9e0af38ec2dffef6c6b5ae69982a29920a352a/CHANGELOG.md)
- [Durable ingestion](https://github.com/merkhet-io/morletdb/blob/8b9e0af38ec2dffef6c6b5ae69982a29920a352a/docs/durable-ingestion.md)
- [SQL upgrade](https://github.com/merkhet-io/morletdb/blob/8b9e0af38ec2dffef6c6b5ae69982a29920a352a/extensions/pg_morlet/sql/pg_morlet--0.1.0--0.2.0.sql)

`pg_morlet` 0.2.0 stores uniformly sampled waveform recordings and finds similar intervals using versioned Morlet features. This developer alpha targets PostgreSQL 17 with pgrx 0.16.1. It installs in `public`, needs superuser installation, and requires no shared preload or background worker.

### Record and Search

```sql
CREATE EXTENSION pg_morlet;
SELECT morlet_create_stream_collection('public', 'signals');
SELECT morlet_start_recording(
  'signals', 'motor-001', 4096, 'vibration', 'g', '{"asset":"motor-1"}'
);
SELECT morlet_append_chunk(
  'signals', 'motor-001', 0,
  ARRAY(SELECT sin(i::float8 / 8) FROM generate_series(0,2047) i)
);
SELECT * FROM morlet_search_streams(
  'signals',
  ARRAY(SELECT sin(i::float8 / 8) FROM generate_series(0,511) i),
  4096, 'vibration', 'g', '{"asset":"motor-1"}'
);
```

### Consistency and Access

`morlet_append_chunk` publishes raw chunks, features and extraction checkpoints in the same transaction. Identical retries verify committed input; changed retries, gaps and overlaps fail. `morlet_search_interval` searches from a stored interval, `morlet_read_interval` reconstructs samples, and `morlet_stream_status` reports ingestion progress. Collection tables use the caller's ordinary privileges and row-security policies; direct edits can violate cross-table invariants.

### Limits and Upgrades

Search compares compatible sample rates, channels, units, windows and extractor versions. It scans exact candidates, with a default 4096-segment budget, and returns half-open sample intervals. Chunk input is 1–32768 finite samples; search windows are 512–4096 samples. There is no resampling, asynchronous extraction, ANN index or synchronized multichannel search. Similarity is an investigation clue, not a calibrated confidence score.

SQL version 0.2.0 belongs to release 0.2.0-alpha.1. Its tested update script accepts the 0.1.0 SQL shipped in 0.1.0-alpha.3; 0.1.0-alpha.1 installations need a reinstall. Existing float8[] collections remain unchanged. Re-encoding them into the new checksummed lossless chunk format is a separate data transfer, not part of ALTER EXTENSION.


### 0.2.0 Alpha Storage and Context

New stream collections store losslessly encoded chunks with CRC-32C validation. Raw-only recordings, materialized envelopes and context-conditioned comparison are optional workflows; measured similarity is not a calibrated fault probability. This release remains a developer alpha for PostgreSQL 17, with production capacity and concurrent workloads unproven.
