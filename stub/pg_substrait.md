## Usage

Sources:

- [Official README](https://github.com/PrimeDataConversion/substrait-postgres/blob/d9ead4261c32d86e3a19bfcf72ce591c8359d6ed/README.md)
- [Official control file](https://github.com/PrimeDataConversion/substrait-postgres/blob/d9ead4261c32d86e3a19bfcf72ce591c8359d6ed/pg_substrait.control)
- [Official build manifest](https://github.com/PrimeDataConversion/substrait-postgres/blob/d9ead4261c32d86e3a19bfcf72ce591c8359d6ed/Cargo.toml)

`pg_substrait` executes Substrait query plans through PostgreSQL's native executor. It accepts either protobuf bytes or JSON and returns a dynamically inferred record set, making it useful for interoperability experiments with systems that produce Substrait plans; it is version `0.1.0`, so supported relations, expressions, and type mappings must be tested rather than assumed.

### Core Workflow

```sql
CREATE EXTENSION pg_substrait;

SELECT *
FROM from_substrait_json('{
  "version": {"minorNumber": 54},
  "relations": [{
    "root": {
      "names": ["result"],
      "input": {
        "project": {
          "expressions": [{"literal": {"i32": 42}}]
        }
      }
    }
  }]
}');
```

The extension infers output names and types, so an `AS` clause is optional. Supply one when the caller needs an explicit stable row contract. For binary plans, pass protobuf data to `from_substrait(bytea)`; `from_substrait_json(text)` handles JSON.

### Important Objects

- `from_substrait(plan bytea)` decodes and executes a protobuf plan.
- `from_substrait_json(json_plan text)` decodes and executes a JSON plan.
- Both return `SETOF RECORD` with extension-provided schema inference.

### Caveats

A Substrait plan is executable input, not passive data. Restrict function execution, apply statement and resource limits, and never accept an untrusted plan merely because it is valid JSON or protobuf. Verify that PostgreSQL permissions still constrain every referenced relation and function in the packaged build.

The README declares PostgreSQL 13 through 17, while the manifest defaults to PostgreSQL 15 and pins pgrx/Substrait library versions. Cross-version producers may emit newer extensions or semantics that this release does not understand. Maintain producer/consumer compatibility tests covering types, nulls, functions, joins, errors, transaction behavior, and malformed or oversized plans.
