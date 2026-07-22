## Usage

Sources:

- [Official README](https://github.com/fusiongyro/damm/blob/0829d6bc4e51f724a9186a55e26208bcf2152836/README.md)
- [Official version 1.0 SQL](https://github.com/fusiongyro/damm/blob/0829d6bc4e51f724a9186a55e26208bcf2152836/damm--1.0.sql)
- [Official control file](https://github.com/fusiongyro/damm/blob/0829d6bc4e51f724a9186a55e26208bcf2152836/damm.control)

`damm` implements the Damm decimal check-digit algorithm in pure SQL. It can generate human-visible identifiers that detect every single-digit error and adjacent digit transposition, validate existing codes, or derive sequential checked IDs from a PostgreSQL sequence.

### Core Workflow

Create a sequence and use `nextdamm` as the default for the `damm_code` domain:

```sql
CREATE EXTENSION damm;
CREATE SEQUENCE thing_id_seq;

CREATE TABLE things (
  id damm_code PRIMARY KEY DEFAULT nextdamm('thing_id_seq'),
  label text NOT NULL
);

INSERT INTO things (label) VALUES ('sample') RETURNING id;

SELECT generate_damm(20140426);
SELECT valid_damm_code(201404265);
```

`generate_damm` appends one decimal check digit. `valid_damm_code` recomputes that digit and returns a boolean.

### API

- `damm_check_digit(bigint)`: computes the check digit without appending it.
- `generate_damm(bigint)`: returns the input multiplied by ten plus its check digit.
- `valid_damm_code(bigint)`: validates the final digit.
- `damm_code`: `bigint` domain constrained by `valid_damm_code`.
- `nextdamm(varchar)`: calls `nextval` on the named sequence and returns a checked value.

### Caveats

A Damm digit detects common transcription errors; it is not a cryptographic checksum, uniqueness mechanism, or authorization control. The sequence supplies uniqueness, while the domain only validates the decimal representation.

The extension owns a lookup table and is not dynamically relocatable after installation, although it can be installed into a chosen schema. Schema-qualify the domain/functions when they are outside the application's search path, and test upper bounds because appending a digit can overflow `bigint` for sufficiently large inputs.
