## 用法

来源：

- [官方上游 README](https://github.com/evancarroll/pg-dump-fcinfo/blob/3defa8fe6769e15fdca9185bbfda3eb707a9962d/README.md)
- [官方扩展控制文件 (dump_fcinfo.control)](https://github.com/evancarroll/pg-dump-fcinfo/blob/3defa8fe6769e15fdca9185bbfda3eb707a9962d/dump_fcinfo.control)
- [官方扩展 SQL (dump_fcinfo--0.0.1.sql)](https://github.com/evancarroll/pg-dump-fcinfo/blob/3defa8fe6769e15fdca9185bbfda3eb707a9962d/dump_fcinfo--0.0.1.sql)

`dump_fcinfo` — 你是否曾经好奇过 fcinfo 的 FunctionCallInfoBaseData 在 PostgreSQL 中存储了 C 函数的上下文？使用它来对应 SQL 或数据库实用工具的工作流。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION dump_fcinfo;

SELECT dump_fcinfo();
SELECT * FROM dump_fcinfo();
SELECT * FROM dump_fcinfo() LIMIT 5;
SELECT dump_fcinfo() FROM ( VALUES (1) ) AS t(x);
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `dump_fcinfo()` 是一个扩展函数并返回 `TABLE`。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.0.1`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
