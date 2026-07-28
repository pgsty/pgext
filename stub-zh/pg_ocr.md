## 用法

来源：

- [官方上游 README](https://github.com/z-xiao-m/pg_ocr/blob/7069c5ecfcec4d5eea16d3a843d577ff61bdcb5e/README.md)
- [官方扩展控制文件 (pg_ocr.control)](https://github.com/z-xiao-m/pg_ocr/blob/7069c5ecfcec4d5eea16d3a843d577ff61bdcb5e/pg_ocr.control)
- [官方扩展 SQL (pg_ocr--0.0.1.sql)](https://github.com/z-xiao-m/pg_ocr/blob/7069c5ecfcec4d5eea16d3a843d577ff61bdcb5e/pg_ocr--0.0.1.sql)

`pg_ocr` — 一个用于光学字符识别（OCR）的 PostgreSQL 扩展 —— 可直接从 SQL 运行文本识别。无需外部服务，只需一个 SQL 查询。可用于相应的向量、模型或检索工作流。在安装扩展依赖项并验证之前，请勿将其集成到应用程序 SQL 中。

### 核心工作流

```sql
CREATE EXTENSION pg_ocr;

-- Recognize text from a file
SELECT pg_ocr_text_from_file('/path/to/data/plate.png');

-- Recognize text from bytea
SELECT pg_ocr_text(pg_read_binary_file('/path/to/data/plate.png'));

-- Get structured result with confidence and bounding box
SELECT r.* FROM pg_ocr(pg_read_binary_file('/path/to/data/plate.png')) AS r;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小查询，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `pg_ocr(image bytea)` 是一个扩展函数，返回 `jsonb`。
- `pg_ocr_from_file(image_path text)` 是一个扩展函数，返回 `jsonb`。
- `pg_ocr_get_engine()` 是一个扩展函数，返回 `void`。
- `pg_ocr_set_model_path(path text)` 是一个扩展函数，返回 `void`。
- `pg_ocr_setup()` 是一个扩展函数，返回 `void`。
- `pg_ocr_text(image bytea)` 是一个扩展函数，返回 `text`。
- `pg_ocr_text_from_file(image_path text)` 是一个扩展函数，返回 `text`。
- `pg_ocr_version()` 是一个扩展函数，返回 `text`。

### 要求与注意事项

- 控制文件声明默认版本为 `0.0.1`。
- 请先安装确认的扩展依赖项：`plpython3u`。
- 控制文件将扩展标记为可重定位。
- 控制文件将扩展标记为可信。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源的一致性。
