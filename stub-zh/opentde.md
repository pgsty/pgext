## 用法

来源：

- [官方扩展控制文件（opentde.control）](https://github.com/ddlifter/cuddly-meme/blob/4740cdde27bf12b026279860af2559db780e1972/opentde.control)
- [官方扩展 SQL（opentde--1.0.sql）](https://github.com/ddlifter/cuddly-meme/blob/4740cdde27bf12b026279860af2559db780e1972/opentde--1.0.sql)

`opentde` — Open Transparent Data Encryption for PostgreSQL。在实现相应的安全、审计或访问控制工作流时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION opentde;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `opentde_blind_index(text)` 是一个扩展函数，返回 `bytea`。
- `opentde_debug_keys()` 是一个扩展函数，返回 `text`。
- `opentde_get_dek_hex(oid)` 是一个扩展函数，返回 `text`。
- `opentde_rotate_master_key(bytea)` 是一个扩展函数，返回 `int4`。
- `opentde_rotate_table_dek(oid)` 是一个扩展函数，返回 `int4`。
- `opentde_set_master_key(bytea)` 是一个扩展函数，返回 `void`。
- `opentde_tableam_handler(internal)` 是一个扩展函数，返回 `table_am_handler`。
- `set_column_level_encryption(enable boolean)` 是一个扩展函数，返回 `void`。
- `opentde` 是一个扩展定义的访问方法。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
