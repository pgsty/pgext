## 用法

来源：

- [官方扩展控制文件](https://github.com/credativ/plsh-mtb/blob/ec5de1202ba185acb96561d799797811c909a389/plsh_mtb.control)
- [官方上游文档](https://github.com/credativ/plsh-mtb/blob/ec5de1202ba185acb96561d799797811c909a389/README.md)

`plsh_mtb` — plsh_mtb 是一个 PostgreSQL 多租户备份控制器，使用 PL/SH 调用随扩展安装的 Korn Shell 命令。

已复核目录快照记录的版本为 `1.0`、类型为 `puresql`、实现语言为 `Shell`。
应先安装并验证声明的扩展依赖：`plsh`。

```sql
CREATE EXTENSION "plsh_mtb";
SELECT extversion
FROM pg_extension
WHERE extname = 'plsh_mtb';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
