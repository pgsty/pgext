## Usage

Sources:

- [Official README](https://github.com/grzm/pgcrockford/blob/774cfee2b77086c096547d1b6767cd1e89076f88/README.markdown)
- [Extension SQL](https://github.com/grzm/pgcrockford/blob/774cfee2b77086c096547d1b6767cd1e89076f88/crockford.sql)
- [Input/output implementation](https://github.com/grzm/pgcrockford/blob/774cfee2b77086c096547d1b6767cd1e89076f88/inout.c)

`crockford` 0.8.34 provides compact unsigned integer types whose textual form uses Crockford Base32. It is useful for human-facing identifiers that need case-insensitive, visually tolerant text, while retaining numeric comparison, arithmetic, and index support.

### Core Workflow

Choose the storage width, insert Base32 text, and cast integers when the input represents a numeric value:

```sql
CREATE EXTENSION crockford;

CREATE TABLE ticket (
  id crockford4 PRIMARY KEY,
  label text NOT NULL
);

INSERT INTO ticket VALUES ('10', 'Base32 text'), (10::crockford4, 'Integer value');
SELECT id, id + 1 FROM ticket ORDER BY id;
```

Text input is decoded as Base32, whereas an integer cast preserves the integer value. Consequently, the text literal `10` represents decimal 32, while integer 10 is rendered as `A`.

### Type and Operator Index

- `crockford2` stores an unsigned 16-bit value, up to `1ZZZ`.
- `crockford4` stores an unsigned 32-bit value, up to `3ZZZZZZ`.
- `crockford8` stores an unsigned 64-bit value, up to `FZZZZZZZZZZZZ`.
- Comparison, arithmetic, and bitwise operators are supplied with B-tree and hash operator classes.

The parser accepts lower- or upper-case input and maps the ambiguous letters I/L and O to 1 and 0; U is excluded by the Crockford alphabet.

### Operational Notes

The types are unsigned, so casts to signed PostgreSQL integers can overflow. Arithmetic also performs range checks. If installed into a schema absent from `search_path`, schema-qualify operators with PostgreSQL's `OPERATOR(schema.operator)` syntax. The reviewed upstream version was tested against PostgreSQL 11; validate compilation, dump/restore, casts, and overflow behavior on newer server majors.
