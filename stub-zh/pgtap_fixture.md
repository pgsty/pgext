## 用法

来源：

- [官方上游 README](https://github.com/tyrusj/pgtap_fixture/blob/5696f7dabdcce205efb78879df9b8de7123e0071/README.md)
- [官方扩展控制文件 (pgtap_fixture.control)](https://github.com/tyrusj/pgtap_fixture/blob/5696f7dabdcce205efb78879df9b8de7123e0071/install/pgtap_fixture.control)

`pgtap_fixture` — 一个 PostgreSQL 扩展，允许为 pgTAP 测试创建更复杂的测试数据集。当应用程序需要此特定数据库功能时，请使用它。在安装此扩展之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION pgtap_fixture;
```

在目标数据库中安装该扩展，在可用时运行上游最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 该扩展记录了其版本信息 `0.0.1`。
- 请首先安装已确认的扩展依赖项：`plpgsql`, `pgtap`。
- 控制文件将该扩展标记为可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源保持一致。
