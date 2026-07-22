## Usage

Sources:

- [Official README at the catalog source revision](https://github.com/furstenheim/is_jsonb_valid/blob/3afa28d70500791a233dc17c919d6136b83cb2ef/README.md)
- [Extension control file at the catalog source revision](https://github.com/furstenheim/is_jsonb_valid/blob/3afa28d70500791a233dc17c919d6136b83cb2ef/is_jsonb_valid.control)
- [Version 0.1.4 extension SQL](https://github.com/furstenheim/is_jsonb_valid/blob/3afa28d70500791a233dc17c919d6136b83cb2ef/is_jsonb_valid--0.1.4.sql)
- [PGXN distribution page](https://pgxn.org/dist/is_jsonb_valid/)

is_jsonb_valid provides native boolean validation of JSONB values against JSON Schema Draft 4 or Draft 7. It is suitable for checks and predicates that need only pass/fail output; it does not return a path or diagnostic explaining a failed keyword.

### Core Workflow

Choose the function that matches the schema draft and pass the schema first, followed by the value being validated.

```sql
CREATE EXTENSION is_jsonb_valid;

SELECT is_jsonb_valid(
    '{"type":"object","required":["id"],"properties":{"id":{"type":"integer"}}}'::jsonb,
    '{"id":42}'::jsonb
);

SELECT is_jsonb_valid_draft_v7(
    '{"if":{"exclusiveMaximum":0},"else":{"multipleOf":2}}'::jsonb,
    '4'::jsonb
);
```

The two functions are strict and immutable in version 0.1.4. This makes them usable in constraints and indexes when the schema is stable.

```sql
CREATE TABLE events (
    payload jsonb NOT NULL,
    CONSTRAINT events_payload_valid CHECK (
        is_jsonb_valid(
            '{"type":"object","required":["id"]}'::jsonb,
            payload
        )
    )
);
```

### Draft Boundary and Limitations

The base validator follows Draft 4; the explicitly named alternative follows Draft 7. Select deliberately, because keyword meaning differs between drafts and the extension does not infer the draft from a schema declaration.

At the pinned revision, remote references are unsupported. Local references are limited to definitions reachable from the schema root, and identifier changes along a reference chain are not resolved. Format validation is also unsupported. These gaps mean a true result is only a claim about the implemented subset, not every behavior in the complete JSON Schema specifications.

No preload or restart is declared. Keep schema literals and application expectations under version control, and add representative valid and invalid cases whenever a constraint depends on this extension.
