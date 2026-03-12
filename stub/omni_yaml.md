


## Usage

> [omni_yaml: YAML toolkit](https://docs.omnigres.org/omni_yaml/yaml/)

The `omni_yaml` extension provides YAML-to-JSON and JSON-to-YAML conversion functions.

### Functions

**`omni_yaml.to_json(text)`** -- Converts YAML text to JSON:

```sql
SELECT omni_yaml.to_json('key: value');
-- Returns: {"key": "value"}
```

**`omni_yaml.to_yaml(json)`** -- Converts JSON to YAML text:

```sql
SELECT omni_yaml.to_yaml('{"key": "value"}'::json);
-- Returns: key: value
```

### JSONB Usage

Direct JSONB support is not provided. Cast the JSON result to `jsonb` as needed:

```sql
SELECT omni_yaml.to_json('key: value')::jsonb;
```
