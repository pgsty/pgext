## 用法

来源：

- [v2.0.21 README](https://github.com/FractalSQLabs/fractalsql-postgresql/blob/v2.0.21/README.md)
- [扩展控制文件](https://github.com/FractalSQLabs/fractalsql-postgresql/blob/v2.0.21/fractalsql.control)
- [SQL 定义](https://github.com/FractalSQLabs/fractalsql-postgresql/blob/v2.0.21/sql/fractalsql--1.0.sql)
- [v2.0.21 发行版](https://github.com/FractalSQLabs/fractalsql-postgresql/releases/tag/v2.0.21)

`fractalsql` 是上游 v2.0.21 发行版中的基础 PostgreSQL 扩展，SQL 版本为 1.0，面向 PostgreSQL 14–18，通过 C 函数提供向量搜索、推理、嵌入和代理工作流基础能力。

### 安装与检查

```sql
CREATE EXTENSION fractalsql;
SELECT fractal_search(ARRAY[0.1, 0.2, 0.3]::float8[]);
```

具体函数面由版本化 SQL 文件定义。可选的 `fractalsql_agents` 扩展依赖此基础扩展，并安装带参数的 PL/pgSQL 代理引擎。

### 运行边界

推理和嵌入函数可能调用配置的外部模型服务。应将服务凭据、模型输入、生成的 SQL 和模型输出视为不可信应用数据，严格授予执行权限并审查网络出口。应记录生成的计划和决策，并测试服务失败、超时、重试及事务回滚。上游发行版提供 PostgreSQL 14–18 的平台包，但这些上游工件不等于本地打包基线已收录。

### 2.0.21 版本变化

特征存储记录写入者来源信息；覆盖其他角色的向量时，要求调用者属于该角色或为超级用户。`fractal_store_morphology` 新增可选的来源标签。安装 SQL 撤销了 `fractal_audit_log` 的公共执行权限，已有数据库需要单独复核授权。`fractal_lock_destructive()` 在当前连接中不可逆地禁用账本修改和审计写入。当调用者有删除队列行的权限时，已完成的向量化任务在一小时后清理，失败任务在一天后清理。

安装定义发生了变化，但扩展 SQL 版本仍是 1.0。仅替换库文件不会更新已有 SQL 对象或权限；采用上游迁移说明前，应检查并备份已安装的定义。
