## Usage

Sources:

- [Official repository](https://github.com/mhagander/jsonb_delete_array)
- [Extension control file](https://github.com/mhagander/jsonb_delete_array/blob/master/jsonb_delete_array.control)
- [Version 1.0 extension SQL](https://github.com/mhagander/jsonb_delete_array/blob/master/jsonb_delete_array--1.0.sql)
- [Version 1.0 C implementation](https://github.com/mhagander/jsonb_delete_array/blob/master/jsonb_delete_array.c)
- [Current PostgreSQL JSON operators](https://www.postgresql.org/docs/current/functions-json.html)

jsonb_delete_array adds a function and subtraction operator that remove multiple top-level object keys or matching top-level string elements from a JSONB value in one call. It does not recursively descend into nested objects or arrays.

### Core Workflow

Pass the input document and a one-dimensional text array of names or string values to delete.

```sql
CREATE EXTENSION jsonb_delete_array;

SELECT jsonb_delete_array(
    '{"keep":1,"drop_a":2,"drop_b":3}'::jsonb,
    ARRAY['drop_a', 'drop_b']
);

SELECT '["keep","drop","other","drop"]'::jsonb
       - ARRAY['drop']::text[];
```

For an object, a matching top-level key and its value are removed. For an array, matching string elements are removed; non-string elements are retained. Null entries in the deletion array are ignored, and an empty deletion array leaves the input unchanged.

### Boundaries and Errors

The function rejects scalar JSONB input and a multidimensional deletion array. Matching is exact and case-sensitive. Nested values are copied without traversal, so deleting a name does not remove the same name below the top level.

### PostgreSQL Version Compatibility

Current PostgreSQL already documents a built-in subtraction overload accepting a text array. Because this extension creates the same operator signature, installation can conflict on PostgreSQL releases that provide it in core. Check the target server's operator catalog and prefer the supported built-in behavior when it is available.

The extension is relocatable and declares no preload or restart requirement. Review application queries before removal or replacement, because a same-looking operator supplied by another PostgreSQL version must be tested for the exact object, array, scalar, and nesting behavior the application expects.
