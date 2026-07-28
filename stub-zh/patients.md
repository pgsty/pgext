## 用法

来源：

- [官方上游 README](https://github.com/constructive-io/pgpm-demo/blob/7793d0dca37c96d92e0cd4c16e25c8140142039e/packages/patients/README.md)
- [官方扩展控制文件 (patients.control)](https://github.com/constructive-io/pgpm-demo/blob/7793d0dca37c96d92e0cd4c16e25c8140142039e/packages/patients/patients.control)

`patients` — **🛠 构建者团队打造 — Postgres 工具模块的创造者，致力于构建安全、可组合的后端。如您喜欢我们的工作，请在 GitHub 上贡献代码。**。当应用程序需要此特定数据库功能时使用它。在安装扩展之前，必须先安装并验证其依赖项。

### 核心工作流

```sql
CREATE EXTENSION patients;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 控制文件声明默认版本为 `0.1.0`。
- 请先安装确认的扩展依赖项：`plpgsql`。
- 控制文件标记该扩展为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源保持一致。
