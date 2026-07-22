## Usage

Sources:

- [Official README](https://github.com/arunchaganty/pg-span/blob/0ff0b6f1070cf6410299701b7fcfe6368513eaa2/README.md)
- [Official usage documentation](https://github.com/arunchaganty/pg-span/blob/0ff0b6f1070cf6410299701b7fcfe6368513eaa2/doc/span.mmd)
- [Extension SQL](https://github.com/arunchaganty/pg-span/blob/0ff0b6f1070cf6410299701b7fcfe6368513eaa2/sql/span.sql)

`span` adds a compact data type for identifying a text interval inside a document. A value has the form `document:begin-end`, such as `article-42:10-25`, and can be stored, compared, grouped, and indexed.

### Core Workflow

```sql
CREATE EXTENSION span;

CREATE TABLE annotations (
    annotation_id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    location span NOT NULL,
    label text NOT NULL
);

INSERT INTO annotations(location, label) VALUES
    ('article-42:10-25', 'person'),
    ('article-42:30-44', 'place');

CREATE INDEX annotations_location_idx
ON annotations(location);

SELECT get_span_doc(location),
       get_span_begin(location),
       get_span_end(location),
       label
FROM annotations
ORDER BY location;
```

The default B-tree operator class supports `=`, `<>`, `<`, `<=`, `>`, and `>=`. A hash operator class supports equality and hash aggregation. The extension also defines `min(span)` and `max(span)` aggregates.

### Constructors and Accessors

- `span(text)` converts validated text to `span`; `text(span)` converts back.
- `is_span(text)` checks whether text has a valid document and ordered offsets.
- `get_span_doc(span)` returns the document identifier.
- `get_span_begin(span)` and `get_span_end(span)` return integer offsets.

```sql
SELECT is_span('article-42:10-25'),
       span('article-42:10-25'),
       text('article-42:10-25'::span);
```

### Caveats

This type models an identifier plus two integer offsets; it does not provide overlap, containment, union, or range-index operators. Ordering is useful for deterministic sorting and B-tree lookup, not interval-search semantics. Invalid forms, missing endpoints, and a begin value greater than the end are rejected. Confirm whether application offsets are bytes, characters, or another unit and keep that convention external to the type. Version `0.1.0` is old and fixed-length, so validate document-identifier length and PostgreSQL-version compatibility before migration. No preload or restart is required.
