## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/siilike/postgresql-vpack/blob/88148fa952875c6fc3ef49c3add0eec4a48418e0/README.md)
- [Version 0.1 SQL API](https://github.com/siilike/postgresql-vpack/blob/88148fa952875c6fc3ef49c3add0eec4a48418e0/vpack--0.1.sql)
- [Extension control file](https://github.com/siilike/postgresql-vpack/blob/88148fa952875c6fc3ef49c3add0eec4a48418e0/vpack.control)

`vpack` adds ArangoDB VelocyPack document storage to PostgreSQL. The `vpack` type is broadly analogous to `jsonb` but supports additional binary/logical types; `vpackpath` provides a JSONPath-like query language. The extension also offers JSON/JSONB/BSON conversion, document operators, B-tree/hash/GIN operator classes, builders, and path functions.

```sql
CREATE EXTENSION vpack;
CREATE TABLE events (id bigint PRIMARY KEY, body vpack);
INSERT INTO events VALUES (1, '{"kind":"login","ok":true}'::vpack);
SELECT body ->> 'kind' FROM events WHERE body @> '{"ok":true}'::vpack;
CREATE INDEX ON events USING gin (body);
```

The README says some pieces and tests are unfinished, tags are not fully supported by `vpackpath`, type-to-tag mappings are not fully specified, and several `jsonb`-style functions are missing. It also attributes some benchmark speed to omitted input validation. Validate every untrusted payload and do not assume JSON-equivalent canonicalization or comparison semantics.

VelocyPack, BSON, JSON, and PostgreSQL types do not have lossless one-to-one mappings. Test numeric ranges, duplicate keys, binary values, tags, path errors, ordering, index rechecks, and round trips. Pin the upstream VelocyPack and Mongo C-driver ABIs, and verify binary dump/restore and logical replication before durable use.
