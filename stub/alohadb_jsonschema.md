## Usage

Sources:

- [Official alohadb_jsonschema.control](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_jsonschema/alohadb_jsonschema.control)
- [Official alohadb_jsonschema--1.0.sql](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_jsonschema/alohadb_jsonschema--1.0.sql)
- [Official jsonschema_validator.c](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_jsonschema/jsonschema_validator.c)

`alohadb_jsonschema` 1.0 validates JSONB documents against a subset of JSON Schema Draft-07 and offers an optional named-schema registry. These sources belong to the AlohaDB kernel; stock PostgreSQL compatibility is not established.

### Validate Documents

```sql
CREATE EXTENSION alohadb_jsonschema;
SELECT jsonschema_is_valid('{"age":30}'::jsonb,
  '{"type":"object","required":["age"],"properties":{"age":{"type":"integer","minimum":0}}}'::jsonb);
SELECT * FROM jsonschema_validate('{"age":-1}'::jsonb,
  '{"properties":{"age":{"minimum":0}}}'::jsonb);
SELECT jsonschema_register('person', '{"type":"object","required":["age"]}'::jsonb);
SELECT jsonschema_validate_named('{"age":30}'::jsonb, 'person');
```

### Objects and Constraints

`jsonschema_is_valid()` returns a boolean and is immutable, strict and parallel-safe, suitable for a CHECK with a fixed schema. `jsonschema_validate()` returns validity plus error path and message. `alohadb_jsonschema_registry` stores names, JSONB schemas and creation times. Named validation is stable rather than immutable because it reads this registry. Control registry writes through database privileges; changing a registered schema does not automatically revalidate stored application data.

### Support Boundary

Supported checks include types, properties, required fields, numeric/string/array/object bounds, pattern, enum, const and logical schema composition. `$ref` resolves internal JSON pointers only. Do not assume remote references or every Draft-07 keyword is implemented. Installation requires administrator privileges; the control marks the extension non-relocatable and not trusted. No shared-preload requirement is present in this module.
