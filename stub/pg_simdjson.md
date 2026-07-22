## Usage

Sources:

- [Upstream README](https://github.com/erthalion/pg_simdjson/blob/4e5a20eab6296d48e41da6f255c70c16bbb3d413/README.md)
- [Version 0.1 SQL definition](https://github.com/erthalion/pg_simdjson/blob/4e5a20eab6296d48e41da6f255c70c16bbb3d413/pg_simdjson--0.1.sql)
- [Parser implementation](https://github.com/erthalion/pg_simdjson/blob/4e5a20eab6296d48e41da6f255c70c16bbb3d413/pg_simdjson.cpp)

`pg_simdjson` version `0.1` is a prototype SIMD JSON parser. It accepts JSON as text and returns PostgreSQL `jsonb`; use it only after benchmarking and validating the packaged build on the target CPU and PostgreSQL version.

### Core Workflow

```sql
CREATE EXTENSION pg_simdjson;

SELECT simdjson_parse('{"id":42,"tags":["fast","json"]}');
SELECT simdjson_parse(payload_text) FROM incoming_events;
```

The extension exposes one function, `simdjson_parse(text) returns jsonb`. It is declared immutable, strict, and parallel safe. Valid JSON objects, arrays, and scalars are converted to normal `jsonb`; invalid input raises an error.

### Semantics and Caveats

Once parsed, PostgreSQL `jsonb` semantics apply: insignificant whitespace and object-key ordering are not preserved, and values are normalized as `jsonb`. Upstream notes that regression results match PostgreSQL's JSONB tests apart from error messages; applications must not depend on identical diagnostics.

This repository embeds an old simdjson source snapshot and describes the project as a prototype. Historical benchmark results are not performance guarantees. There is no preload requirement, but C++ ABI, compiler, CPU-instruction support, malformed-input behavior, and modern PostgreSQL compatibility all need verification before non-lab use.
