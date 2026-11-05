## 用法

来源：

- [上游来源 1](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_graphql/alohadb_graphql--1.0.sql)
- [上游来源 2](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_graphql/alohadb_graphql.control)

`alohadb_graphql` 的 SQL 版本 1.0 是 AlohaDB v20.0.0 内置的 GraphQL 辅助扩展。

### 核心流程

`graphql_schema` 从 public 模式中的表生成 SDL，`graphql_schema_json` 返回 JSONB 表示。`graphql_execute(query, variables, operation_name)` 解析请求，将其转换为 SQL，并以 JSONB 数据封装返回。

```sql
CREATE EXTENSION alohadb_graphql;
SELECT graphql_schema();
SELECT graphql_schema_json();
```

### 运行边界

应使用匹配的 AlohaDB 运行环境，内置控制文件不能证明其兼容标准 PostgreSQL。安装需要特权角色。这些函数不是 HTTP 身份认证网关。允许不受信任请求前，应核对暴露的表、调用者权限及解析器行为，并测试支持的查询与变更语义，不能假定完全符合 GraphQL 规范。
