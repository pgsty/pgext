
## 用法

> 语法：
>
> ```bash
> regresql init postgres://localhost/mydb
> regresql add src/sql/
> regresql update
> regresql test
> ```
>
> 来源：[README](https://github.com/boringsql/regresql)，[产品页](https://boringsql.com/products/regresql/)

`RegreSQL` 在上游被定位为一个与语言无关的 PostgreSQL SQL 回归测试工具，而不是数据库内的 `CREATE EXTENSION` 式模块。它会发现 `.sql` 文件、在 PostgreSQL 上执行这些文件、保存期望输出，并跟踪查询计划变化。

## 快速上手

README 给出的基本流程是：

```bash
regresql init postgres://localhost/mydb
regresql discover
regresql add src/sql/
regresql update
regresql test
```

这会初始化测试套件、发现查询文件、生成计划定义、捕获期望输出，并执行回归检查。

## 跟踪内容

上游文档重点关注：

- 期望的查询输出快照
- `EXPLAIN` 计划基线
- 顺序扫描告警
- 与迁移相关的查询回归
- 面向 CI 的输出格式，例如 `junit`、`json`、`pgtap` 和 `github-actions`

## 查询文件与计划

RegreSQL 使用普通 SQL 文件，并支持通过 `-- name:` 注释在单个文件中编排多个查询：

```sql
-- name: get-user-by-id
SELECT * FROM users WHERE id = :id;
```

计划文件用于提供测试参数：

```yaml
"1":
  id: 42
"2":
  id: 100
```

## 快照与迁移

该工具可以构建和恢复数据库快照，并比较迁移前后的查询行为：

```bash
regresql snapshot build
regresql snapshot restore
regresql migrate --script db/migrations/001_add_column.sql
```

## 安装

README 说明可通过 Homebrew 或 Go 安装：

```bash
brew tap boringsql/boringsql
brew install regresql
```

或：

```bash
go install github.com/boringsql/regresql@latest
```

快照相关命令需要 `pg_dump`、`pg_restore` 和 `psql` 等 PostgreSQL 客户端工具。
