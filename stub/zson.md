## Usage

Sources:

- [Pinned upstream README](https://github.com/postgrespro/zson/blob/e214c79711b9eb92786280756a549b9e4cb2a215/README.md)
- [Pinned version 1.1 installation SQL](https://github.com/postgrespro/zson/blob/e214c79711b9eb92786280756a549b9e4cb2a215/zson--1.1.sql)
- [Pinned extension control file](https://github.com/postgrespro/zson/blob/e214c79711b9eb92786280756a549b9e4cb2a215/zson.control)
- [Pinned official benchmark notes](https://github.com/postgrespro/zson/blob/e214c79711b9eb92786280756a549b9e4cb2a215/docs/benchmark.md)

`zson` version `1.1` is a dictionary-compressed, JSONB-compatible storage type. It learns frequently repeated strings—including keys, string values, and array elements—from representative JSONB documents, then uses that shared dictionary when encoding new `zson` values.

### Core Workflow

Train a dictionary before inserting compressed values. The training argument is a two-dimensional array of table and JSONB-column names:

```sql
CREATE EXTENSION zson;

CREATE TABLE zson_training (doc jsonb);
INSERT INTO zson_training VALUES
  ('{"kind":"event","service":"billing","status":"ok"}'::jsonb),
  ('{"kind":"event","service":"billing","status":"failed"}'::jsonb);

SELECT zson_learn('{{"zson_training","doc"}}'::text[][]);
SELECT * FROM zson_dict ORDER BY dict_id, word_id;

CREATE TABLE zson_events (id bigint GENERATED ALWAYS AS IDENTITY, payload zson);
INSERT INTO zson_events(payload)
SELECT doc::zson FROM zson_training;

SELECT payload ->> 'service' FROM zson_events;
SELECT zson_info(payload) FROM zson_events LIMIT 1;
```

Run `zson_learn()` again after the document vocabulary changes materially. New values use the new dictionary, while existing values retain their original dictionary identifier.

### Important Objects

- `zson`: variable-length compressed type with assignment cast from `jsonb` and implicit cast to `jsonb`, allowing JSONB operators after decompression.
- `zson_learn(tables_and_columns, max_examples, min_length, max_length, min_count)`: samples JSONB columns and creates the next dictionary with up to 65,534 frequent strings.
- `zson_dict(dict_id, word_id, word)`: extension-managed dictionary table included in extension configuration dumps.
- `zson_info(zson)`: reports format and dictionary-version information for a stored value.
- `zson_extract_strings(jsonb)`: recursively extracts object keys and string values used during training.

### Operational Notes

The extension supports installation only in schema `public`. `zson_learn()` constructs dynamic SQL from table and column names and scans sample data, so allow only trusted administrators to run it, bound `max_examples`, and train on a representative distribution. Compression and throughput gains are data- and workload-dependent; reproduce the official benchmark method with the actual dataset instead of treating published numbers as guarantees.

Dictionaries are cached, and the upstream guide says a new dictionary may take about a minute to become active in all backends. Never delete an old `dict_id` while any stored value still references it, or those documents can become unreadable; keeping a few kilobytes of old dictionary data is safer. Test backup/restore with `zson_dict`, mixed dictionary versions, expression indexes that cast to `jsonb`, and the CPU cost of decompression before replacing JSONB columns with `zson`.
