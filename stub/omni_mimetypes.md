


## Usage

> [omni_mimetypes: MIME types](https://docs.omnigres.org/omni_mimetypes/reference/)

The `omni_mimetypes` extension provides MIME type data sourced from IANA, Apache, and Nginx (via mime-db).

### Tables

- **`omni_mimetypes.mime_types`** -- `id`, `name` (unique), `source` (iana/apache/nginx), `compressible`, `charset`
- **`omni_mimetypes.file_extensions`** -- `id`, `extension` (without dot prefix)
- **`omni_mimetypes.mime_types_file_extensions`** -- Junction table linking MIME types to file extensions

### Find MIME Type by Extension

```sql
SELECT mime_types.name
FROM omni_mimetypes.mime_types
JOIN omni_mimetypes.mime_types_file_extensions mtfe ON mtfe.mime_type_id = mime_types.id
JOIN omni_mimetypes.file_extensions ON mtfe.file_extension_id = file_extensions.id
WHERE file_extensions.extension = 'js';
-- Result: application/javascript
```

### Update MIME Database

```sql
CREATE EXTENSION omni_httpc CASCADE;
WITH db AS (
    SELECT * FROM omni_httpc.http_execute(
        omni_httpc.http_request(
            'https://cdn.jsdelivr.net/gh/jshttp/mime-db@1.52.0/db.json'))
)
SELECT omni_mimetypes.import_mime_db(convert_from(body, 'utf8')::jsonb) FROM db;
```
