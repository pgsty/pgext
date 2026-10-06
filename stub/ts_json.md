## Usage

Sources:

- [README.md](https://github.com/za-arthur/apod_fts/blob/635ddf0dae7dde6dc2e594ef1b2a7b4f6fd21d10/README.md)
- [modules/ts_json/ts_json--1.0.sql](https://github.com/za-arthur/apod_fts/blob/635ddf0dae7dde6dc2e594ef1b2a7b4f6fd21d10/modules/ts_json/ts_json--1.0.sql)
- [modules/ts_json/ts_json.c](https://github.com/za-arthur/apod_fts/blob/635ddf0dae7dde6dc2e594ef1b2a7b4f6fd21d10/modules/ts_json/ts_json.c)
- [modules/ts_json/ts_json.control](https://github.com/za-arthur/apod_fts/blob/635ddf0dae7dde6dc2e594ef1b2a7b4f6fd21d10/modules/ts_json/ts_json.control)

`ts_json` is a small helper from the APOD full-text-search demonstration. It traverses JSON/JSONB scalar values and joins their text with a chosen delimiter for a text-search pipeline.

### Core Workflow

```sql
CREATE EXTENSION ts_json;
SELECT jsonb_values('{"title":"PostgreSQL","body":"database"}'::jsonb, ' ');
```

### Operational Boundaries

`json_values(json, text)` and `jsonb_values(jsonb, text)` return text; the second argument is a delimiter, not a field name. Keys are omitted, while string, number, boolean and null values are included. The extension does not install the demonstration website, RUM indexes or text-search dictionaries.

It is a historical C extension with no current major-version matrix. Installation needs a superuser. No preload is required; validate behavior on the target PostgreSQL version before relying on the old source.
