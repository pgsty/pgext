## Usage

Sources:

- [Official repository README](https://github.com/xaneets/pg-debyte/blob/e0b6e4d4c98781acc0cfbc458fc95de9e0fbb1af/README.md)
- [Official extension control file](https://github.com/xaneets/pg-debyte/blob/e0b6e4d4c98781acc0cfbc458fc95de9e0fbb1af/pg_debyte_ext/pg_debyte_ext.control)
- [Official example-extension implementation](https://github.com/xaneets/pg-debyte/blob/e0b6e4d4c98781acc0cfbc458fc95de9e0fbb1af/pg_debyte_ext/src/lib.rs)

`pg_debyte_ext` 0.1.0 is the example PostgreSQL extension in the `pg-debyte` Rust workspace. It decodes registered Bincode `bytea` payloads into `jsonb`, either from an explicit type identifier and schema version or from a self-describing envelope. Upstream intends developers to build their own registry-backed extension rather than treat this demo schema as a universal decoder.

### Core Workflow

The pinned example registers one demo type under UUID `11111111-1111-1111-1111-111111111111`, plus a second known-schema decoder:

```sql
CREATE EXTENSION pg_debyte_ext;

SELECT bytea_to_json_by_id(
  decode('010464656d6f', 'hex'),
  '11111111-1111-1111-1111-111111111111'::uuid,
  1::smallint
);

SELECT bytea_to_json_demo_record_second(
  decode('01067365636f6e6401', 'hex')
);
```

The first call returns a JSON object with an integer ID and text label. The second decodes the example's fixed `DemoRecordSecond` schema. `bytea_to_json_auto(bytea)` accepts only the repository's envelope format and dispatches using the type, schema version, codec, and action identifiers encoded in that envelope.

### Important Objects

- `bytea_to_json_by_id(bytea, uuid, smallint)` routes raw bytes through an explicitly selected registered decoder.
- `bytea_to_json_auto(bytea)` parses the envelope and dispatches automatically.
- `bytea_to_json_demo_record_second(bytea)` is the generated known-schema function for the second demo record.
- `pg_debyte.max_input_bytes`, `pg_debyte.max_output_bytes`, and `pg_debyte.max_json_bytes` are user-settable safety limits; the pinned source defaults each to 950 MiB.

### Extension-Builder Boundary

The workspace's reusable crates are version 0.2.1 at the pinned commit, while the example extension control file remains 0.1.0. Custom extensions define serializable Rust types, codecs, UUID/schema-version keys, optional actions such as bounded Zstandard decompression, and a static registry. A raw payload must match the selected Bincode schema exactly; an envelope must reference registered codec, action, and decoder IDs.

### Operational Notes

The example supplies pgrx build features for PostgreSQL 15 and 17 and requires a build matching the target major version. It has no preload or restart requirement: library initialization registers the GUCs and demo registry. Decode errors, unknown types, bad envelopes, oversized input/output, and Rust panics are converted to PostgreSQL errors, but limits should be reduced to values appropriate for the workload. Never decode untrusted large payloads with the 950 MiB defaults without resource testing.
