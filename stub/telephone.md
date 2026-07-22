## Usage

Sources:

- [Official telephone type documentation at the reviewed commit](https://github.com/pgstuff/telephone/blob/abdf084b6460f2dbfe747c0c19092a845ad67e02/doc/telephone.md)
- [Official README at the reviewed commit](https://github.com/pgstuff/telephone/blob/abdf084b6460f2dbfe747c0c19092a845ad67e02/README.md)
- [SQL objects and operator classes](https://github.com/pgstuff/telephone/blob/abdf084b6460f2dbfe747c0c19092a845ad67e02/sql/telephone.sql)
- [Parser and representation implementation](https://github.com/pgstuff/telephone/blob/abdf084b6460f2dbfe747c0c19092a845ad67e02/src/telephone.c)
- [Extension control file](https://github.com/pgstuff/telephone/blob/abdf084b6460f2dbfe747c0c19092a845ad67e02/telephone.control)

`telephone` adds a normalized telephone-number type with B-tree and hash indexing. It preserves presentation metadata while deriving an identity that ignores punctuation and maps letters to telephone keypad digits. It also represents pause, confirmation, extension, and partial-number directives. Version 0.0.1 is marked unstable and its dialing-plan data is historical.

### Core Workflow

```sql
CREATE EXTENSION telephone;

CREATE TABLE contact_phone (
  phone telephone PRIMARY KEY,
  label text NOT NULL
);

INSERT INTO contact_phone VALUES ('1 (800) 555-01AZ', 'support');

SELECT phone, telephone_to_format(phone, 'E123')
FROM contact_phone
WHERE phone = '18005550129'::telephone;
```

Values beginning with `+` enter calling-code mode and are checked against the extension's embedded numbering-plan data. Other values use digits mode. For strict calling-code validation, cast the `+` form directly; wrapper behavior can catch errors and fall back to digits mode, hiding the difference between invalid and unsupported input.

### Main Objects

`telephone_mode_get`, `telephone_calling_code_get`, `telephone_service_get`, `telephone_fictitious_get`, `telephone_fictitious_is`, `telephone_geo_is`, `telephone_geo_parts_get`, `telephone_domestic_numbers_get`, `telephone_extension_numbers_get`, and `telephone_domestic_prefer_get` inspect normalized values. `telephone_domestic_assume_set` and `telephone_set` construct or reinterpret them. Comparison operators, B-tree and hash operator classes, and aggregates such as `min`/`max` work on telephone identity rather than original formatting.

### Data and Safety Boundaries

International support is incomplete; the North American Numbering Plan is the most developed path, while many calling codes are partial or unsupported. Numbering plans change over time, so the embedded data must not be treated as current carrier, routing, emergency, or regulatory validation.

The SQL defines a binary-compatible assignment cast from `bytea` to `telephone` without a validation function. Untrusted raw bytes can bypass the text parser and violate internal assumptions; revoke or avoid that path. Test equality, formatting, fictitious ranges, geographic helpers, dumps, and every accepted calling plan before using the type in durable constraints.
