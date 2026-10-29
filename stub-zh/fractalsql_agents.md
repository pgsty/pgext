## 用法

来源：

- [v2.0.20 README](https://github.com/FractalSQLabs/fractalsql-postgresql/blob/v2.0.20/README.md)
- [扩展控制文件](https://github.com/FractalSQLabs/fractalsql-postgresql/blob/v2.0.20/fractalsql_agents/fractalsql_agents.control)
- [SQL 定义](https://github.com/FractalSQLabs/fractalsql-postgresql/blob/v2.0.20/fractalsql_agents/sql/fractalsql_agents--1.0.sql)
- [v2.0.20 发行版](https://github.com/FractalSQLabs/fractalsql-postgresql/releases/tag/v2.0.20)

`fractalsql_agents` 是 FractalSQL v2.0.20 随附的可选依赖扩展，基于 `fractalsql` 的搜索、推理和分析原语，为 PostgreSQL 14–18 安装十六个带参数的 PL/pgSQL 代理引擎。

### 安装

```sql
CREATE EXTENSION fractalsql;
CREATE EXTENSION fractalsql_agents;
```

控制文件要求先安装 `fractalsql`；应先完成基础扩展的安装和升级。SQL 文件定义各代理引擎函数及其参数契约，向应用角色授予执行权限前应先阅读这些定义。

### 安全与服务边界

部分引擎调用推理或嵌入函数，其他引擎只组合检索或分析能力。服务凭据和网络出口属于部署配置。生成的 SQL 和模型输出均不可信；应使用最小权限、明确的数据访问授权、语句和服务超时及审计日志。启用会写入数据的引擎前，应测试部分推理失败和事务回滚。
