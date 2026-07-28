## 用法

来源：

- [官方扩展控制文件 (bedquilt.control)](https://api.pgxn.org/src/bedquilt/bedquilt-2.1.0/bedquilt.control)
- [官方扩展 SQL (bedquilt--0.2.0.sql)](https://api.pgxn.org/src/bedquilt/bedquilt-2.1.0/sql/bedquilt--0.2.0.sql)

`bedquilt` — 一个 JSON 对象存储。当应用程序需要此特定数据库功能时使用它。在集成到应用程序 SQL 之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION bedquilt;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `bq_add_constraints(i_coll text, i_jdoc json)` 是一个扩展函数，返回 `boolean`。
- `bq_check_id_type(i_jdoc json)` 是一个扩展函数，返回 `VOID`。
- `bq_collection_exists(i_coll text)` 是一个扩展函数，返回 `boolean`。
- `bq_constraint_name_exists(i_coll text, i_name text)` 是一个扩展函数，返回 `boolean`。
- `bq_count(i_coll text, i_doc json)` 是一个扩展函数，返回 `integer`。
- `bq_create_collection(i_coll text)` 是一个扩展函数，返回 `BOOLEAN`。
- `bq_delete_collection(i_coll text)` 是一个扩展函数，返回 `BOOLEAN`。
- `bq_doc_set_key(i_jdoc json, i_key text, i_val anyelement)` 是一个扩展函数，返回 `json`。
- `bq_find(i_coll text, i_json_query json)` 是一个扩展函数，返回 `table`。
- `bq_find_one(i_coll text, i_json_query json)` 是一个扩展函数，返回 `table`。
- `bq_find_one_by_id(i_coll text, i_id text)` 是一个扩展函数，返回 `table`。
- `bq_generate_id()` 是一个扩展函数，返回 `char`。
- `bq_insert(i_coll text, i_jdoc json)` 是一个扩展函数，返回 `text`。
- `bq_list_collections()` 是一个扩展函数，返回 `table`。

### 要求与注意事项

- 审核的控制文件声明默认版本为 `2.1.0`。
- 先安装并验证确认的扩展依赖项：`plpgsql`, `pgcrypto`, `plpython3u`。
- 控制文件标记该扩展为可重定位。
- 在生产使用前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
