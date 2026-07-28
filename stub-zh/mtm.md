## 用法

来源：

- [官方上游 README](https://github.com/skilyazhnev/mtm/blob/470bc005b180469c1246e0ead8cfeed703a8a6e3/README.md)
- [官方扩展控制文件 (mtm.control)](https://github.com/skilyazhnev/mtm/blob/470bc005b180469c1246e0ead8cfeed703a8a6e3/mtm/mtm.control)
- [官方扩展 SQL (mtm--0.1.sql)](https://github.com/skilyazhnev/mtm/blob/470bc005b180469c1246e0ead8cfeed703a8a6e3/mtm/mtm--0.1.sql)

`mtm` — 聚合以查找最小值和最大值。当 SQL 需要这些特殊函数或聚合时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION mtm;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `final_mtm(state state_mtm)` 是一个扩展函数，返回 `text`。
- `final_mtm(state state_mtm_dp)` 是一个扩展函数，返回 `text`。
- `transition_mtm(state state_mtm, val numeric)` 是一个扩展函数，返回 `state_mtm`。
- `transition_mtm(state state_mtm_dp, val double precision)` 是一个扩展函数，返回 `state_mtm_dp`。
- `max_to_min` 是由扩展公开的聚合。
- `state_mtm` 是一个扩展定义的类型。
- `state_mtm_dp` 是一个扩展定义的类型。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.1`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况，与固定源进行比对。
