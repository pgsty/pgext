## 用法

来源：

- [v2.0.20 README](https://github.com/FractalSQLabs/fractalsql-postgresql/blob/v2.0.20/README.md)
- [扩展控制文件](https://github.com/FractalSQLabs/fractalsql-postgresql/blob/v2.0.20/fractalsql.control)
- [SQL 定义](https://github.com/FractalSQLabs/fractalsql-postgresql/blob/v2.0.20/sql/fractalsql--1.0.sql)
- [v2.0.20 发行版](https://github.com/FractalSQLabs/fractalsql-postgresql/releases/tag/v2.0.20)

`fractalsql` 是上游 v2.0.20 发行版中的基础 PostgreSQL 扩展，SQL 版本为 1.0，面向 PostgreSQL 14–18，通过 C 函数提供向量搜索、推理、嵌入和代理工作流基础能力。

### 安装与检查

```sql
CREATE EXTENSION fractalsql;
SELECT fractal_search(ARRAY[0.1, 0.2, 0.3]::float8[]);
```

具体函数面由版本化 SQL 文件定义。可选的 `fractalsql_agents` 扩展依赖此基础扩展，并安装带参数的 PL/pgSQL 代理引擎。

### 运行边界

推理和嵌入函数可能调用配置的外部模型服务。应将服务凭据、模型输入、生成的 SQL 和模型输出视为不可信应用数据，严格授予执行权限并审查网络出口。应记录生成的计划和决策，并测试服务失败、超时、重试及事务回滚。上游发行版提供 PostgreSQL 14–18 的平台包，但这些上游工件不等于本地打包基线已收录。
