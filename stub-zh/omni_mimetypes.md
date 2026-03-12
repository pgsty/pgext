


## 用法

> [omni_mimetypes: MIME 类型](https://docs.omnigres.org/omni_mimetypes/reference/)

`omni_mimetypes` 扩展提供来自 IANA、Apache 和 Nginx（通过 mime-db）的 MIME 类型数据。

### 表

- **`omni_mimetypes.mime_types`** -- `id`、`name`（唯一）、`source`（iana/apache/nginx）、`compressible`、`charset`
- **`omni_mimetypes.file_extensions`** -- `id`、`extension`（不带点前缀）
- **`omni_mimetypes.mime_types_file_extensions`** -- 连接 MIME 类型与文件扩展名的关联表

### 通过扩展名查找 MIME 类型

```sql
SELECT mime_types.name
FROM omni_mimetypes.mime_types
JOIN omni_mimetypes.mime_types_file_extensions mtfe ON mtfe.mime_type_id = mime_types.id
JOIN omni_mimetypes.file_extensions ON mtfe.file_extension_id = file_extensions.id
WHERE file_extensions.extension = 'js';
-- 结果: application/javascript
```

### 更新 MIME 数据库

```sql
CREATE EXTENSION omni_httpc CASCADE;
WITH db AS (
    SELECT * FROM omni_httpc.http_execute(
        omni_httpc.http_request(
            'https://cdn.jsdelivr.net/gh/jshttp/mime-db@1.52.0/db.json'))
)
SELECT omni_mimetypes.import_mime_db(convert_from(body, 'utf8')::jsonb) FROM db;
```
