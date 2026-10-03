## 用法

来源：

- [packs/companion/extension/README.md](https://github.com/cpuchip/pg-ai-stewards/blob/c78212cf4dbca4a711cbc6416e06dad25d576b1d/packs/companion/extension/README.md)
- [packs/companion/extension/stewards_companion.control](https://github.com/cpuchip/pg-ai-stewards/blob/c78212cf4dbca4a711cbc6416e06dad25d576b1d/packs/companion/extension/stewards_companion.control)
- [packs/companion/extension/stewards_companion--0.3.0.sql](https://github.com/cpuchip/pg-ai-stewards/blob/c78212cf4dbca4a711cbc6416e06dad25d576b1d/packs/companion/extension/stewards_companion--0.3.0.sql)

`stewards_companion` 0.3.0 是 `pg_ai_stewards` 的 SQL 伴随扩展，在 `companion` 与 `forge` 模式中增加提醒、审批辅助功能及运行时工具注册。其运行范围沿用核心的 PostgreSQL 18 部署环境。

### 基本用法

```sql
CREATE EXTENSION pg_ai_stewards CASCADE;
CREATE EXTENSION stewards_companion;
SELECT * FROM companion.reminders LIMIT 5;
SELECT * FROM forge.forged_tools LIMIT 5;
```

须单独安装该包的 control 与版本化 SQL 文件；普通核心镜像构建不会自动包含它们。创建需要高权限的扩展安装角色。该包自身没有原生库或预加载项。

### 工具与卸载

提醒辅助函数负责创建、列出、取消及认领提醒，由独立客户端负责投递。Forge 注册流程在外层核心工作流中应用获批 SQL 工具。应限制这些权限，并在审批前审核完整函数定义。0.3.0 版本在解除作业卡住状态时重置循环计数。

卸载前先调用 `companion.companion_uninstall()`，停用该包的工具定义并收缩写入白名单，再执行 `DROP EXTENSION stewards_companion`。存在运行时创建的工具函数时，普通删除会拒绝；`CASCADE` 还会删除这些函数。流水线与历史记录仍保留，但扩展拥有的提醒表会被删除。其配置备份登记用于在备份／恢复时携带提醒内容，并不能使内容在删除扩展后继续存在。
