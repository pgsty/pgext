## Usage

Sources:

- [Official README](https://github.com/postgrespro/jsonb_schema/blob/master/README.md)
- [Extension control file](https://github.com/postgrespro/jsonb_schema/blob/master/jsonb_schema.control)
- [Version 1.0 extension SQL](https://github.com/postgrespro/jsonb_schema/blob/master/jsonb_schema--1.0.sql)
- [Upstream regression example](https://github.com/postgrespro/jsonb_schema/blob/master/sql/test.sql)

jsonb_schema separates the structural schema of a JSONB value from its scalar data and stores each distinct schema once in a database table. It can reduce storage when many documents share the same shape, but the packed representation is not self-contained and is not directly queryable like the original document.

### Core Workflow

Pack documents before storage and unpack them when the original JSONB value is needed.

```sql
CREATE EXTENSION jsonb_schema;

CREATE TABLE compact_events (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    packed jsonb NOT NULL
);

INSERT INTO compact_events (packed)
VALUES
    (jsonb_pack('{"kind":"click","user":10,"tags":["a","b"]}'::jsonb)),
    (jsonb_pack('{"kind":"view","user":11,"tags":["c"]}'::jsonb));

SELECT id, jsonb_unpack(packed) AS document
FROM compact_events;
```

The packed value is a two-element JSONB array containing a schema identifier and encoded data. The shared schema bytes live in the extension-owned dictionary table.

### Installed Objects

Version 1.0 creates the schema dictionary table, unique indexes for schema bytes and identifiers, two C-level schema/data conversion functions, and the pack and unpack wrappers. Packing inserts a new schema row on first use and reuses the identifier on conflict.

### Storage and Lifecycle Boundaries

Measure with representative data before adopting the format: documents with repeated shapes may benefit, while highly variable shapes pay dictionary and encoding overhead. Indexes or expressions aimed at original JSON fields must unpack first, so ordinary JSONB query and index behavior does not transparently carry over to the packed column.

Packed rows depend on the exact identifier mapping in the dictionary table. Dump, restore, replicate, or move that table together with all packed values; copying only the JSONB column to another database can decode against the wrong schema or return no value. Do not delete dictionary rows still referenced by packed documents. The extension does not provide reference counting or schema garbage collection.

Packing writes shared metadata even though it is called as a function, so treat it as a write operation for transaction, permission, and concurrency planning. The extension is relocatable and declares no preload or restart requirement.
