## 用法

来源：

- [官方上游 README](https://github.com/veranemoloko/pg_jewelry_tools/blob/702b3e118c2a0b0a4b941037a8f6d3093f59788d/README.md)
- [官方扩展控制文件 (pg_jewelry_tools.control)](https://github.com/veranemoloko/pg_jewelry_tools/blob/702b3e118c2a0b0a4b941037a8f6d3093f59788d/pg_jewelry_tools.control)
- [官方扩展 SQL (pg_jewelry_tools--0.0.1.sql)](https://github.com/veranemoloko/pg_jewelry_tools/blob/702b3e118c2a0b0a4b941037a8f6d3093f59788d/sql/pg_jewelry_tools--0.0.1.sql)

`pg_jewelry_tools` 是一个为珠宝企业提供实用功能的 PostgreSQL 扩展，目前属于教育用途项目。当 SQL 需要这些专用函数或聚合时可使用它。请以上方链接的固定上游修订为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_jewelry_tools;
```

在目标数据库中安装扩展；如果上游提供了最小示例，请运行该示例，并在集成到应用 SQL 前验证安装版本和返回值。

### 重要对象

- `calculate_gemstone_carat_to_gr(carat double precision)` 是扩展函数，返回 `double`。
- `calculate_gemstone_gr_to_carat(carat double precision)` 是扩展函数，返回 `double`。
- `calculate_metal_weight_gr(metal TEXT, purity INTEGER, volume_mm3 double precision)` 是扩展函数，返回 `double`。
- `public.test_add_one(i integer)` 是扩展函数，返回 `integer`。
- `gemstone` 是扩展定义的类型。
- `jw_item` 是扩展定义的类型。
- `precious_metal` 是扩展定义的类型。

### 要求与注意事项

- 审阅的控制文件声明默认版本为 `0.0.1`。
- 控制文件将该扩展标记为可重定位。
- 生产使用前，请根据固定版本源码确认权限、支持的 PostgreSQL 版本、升级行为和失败情形。
