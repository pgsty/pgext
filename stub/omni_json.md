


## Usage

> [omni_json: JSON toolkit](https://docs.omnigres.org/omni_json/table_mapping/)

The `omni_json` extension provides table-to-JSON mapping with custom transformations when converting rows to JSON.

### Define a Table Mapping

```sql
SELECT omni_json.define_table_mapping(example, '{}');
```

### Column Renaming

```sql
SELECT omni_json.define_table_mapping(example, '{
    "columns": { "dob": { "path": "date_of_birth" } }
}');
```

### Nested Keys

```sql
SELECT omni_json.define_table_mapping(example, '{
    "columns": {
        "first_name": { "path": ["name", "first"] },
        "last_name":  { "path": ["name", "last"] }
    }
}');
```

### Column Exclusion

```sql
SELECT omni_json.define_table_mapping(example, '{
    "columns": { "dob": { "exclude": true } }
}');
```

### Column Transformation

Input (JSON to record) and output (record to JSON) transforms:

```json
{
    "columns": {
        "password": {
            "transform": {
                "input":  { "type": "text", "function": "encrypt_password" },
                "output": { "type": "text", "function": "mask_password" }
            }
        }
    }
}
```

### Querying as JSON

```sql
SELECT to_jsonb(products.*) FROM products;
```

### Inserting from JSON

```sql
INSERT INTO people (dob, first_name, last_name)
    (SELECT dob, first_name, last_name
     FROM jsonb_populate_record(null::people,
         '{"first_name": "Jane", "last_name": "Doe", "dob": "1981-12-12"}'));
```
