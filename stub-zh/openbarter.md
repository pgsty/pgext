## 用法

来源：

- [官方扩展控制文件](https://api.pgxn.org/src/openbarter/openbarter-0.8.2/openbarter.control)
- [官方项目页面](https://pgxn.org/dist/openbarter/)

`openbarter` — 多边协议引擎。当应用程序需要此特定数据库功能时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION openbarter;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `’0.8.0’`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况。
