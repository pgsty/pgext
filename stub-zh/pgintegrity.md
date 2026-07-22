## 用法

来源：

- [0.0.1 版扩展 SQL](https://github.com/liang7878/pgintegrity/blob/7f812066f20c5d3052708612ee823bdd1e2ec663/pgintegrity--0.0.1.sql)
- [触发器实现](https://github.com/liang7878/pgintegrity/blob/7f812066f20c5d3052708612ee823bdd1e2ec663/pgintegrity.c)
- [上游实现笔记](https://github.com/liang7878/pgintegrity/blob/7f812066f20c5d3052708612ee823bdd1e2ec663/doc/README.md)

`pgintegrity` 是一个未完成的 2018 年触发器原型，试图通过 `dblink` 在单独的本地数据库中维护逐行“水印”。其发布 SQL 只定义触发器函数 `pg_integrity`，不会创建所需远端数据库、表、角色、授权、配置或触发器。它不是可部署的完整性或审计方案。

### 安全检查

若在一次性旧环境中构建了历史二进制，初始工作只应限于目录检查：

```sql
CREATE EXTENSION pgintegrity;

SELECT p.oid::regprocedure, p.prorettype::regtype, l.lanname
FROM pg_proc AS p
JOIN pg_language AS l ON l.oid = p.prolang
WHERE p.proname = 'pg_integrity';
```

不要把触发器挂到应用表。C 函数假定自己由 `AFTER` 行级触发器调用，却没有先验证触发器调用上下文；外部配套设置也没有安装，更缺少足以安全复现的文档。

### 源码确认的隐含假设

实现要求扩展契约之外的所有下列条件：

- `dblink` 函数已经可用，但控制文件没有声明该依赖；
- 名为 `pgintegrity.password` 的自定义设置保存明文密码；
- 已存在名为 `pgintegrity` 的本地数据库、硬编码角色 `umsuser` 以及 `t_privilege`、`t_watermark` 表；
- 应用表暴露行 OID，而现代 PostgreSQL 已移除该能力；
- 每次触发变更时服务器可以打开额外数据库连接。

连接串和 SQL 通过字符串拼接构造，内容包括角色、数据库、表、密码与行数据。这会带来正确性、秘密处理和注入风险。代码还对没有明确字段边界的文本表示拼接使用非密码学哈希，因此“水印”不是可信的防篡改原语。

### 兼容性与替代方案

0.0.1 版可重定位且不要求预加载。仓库包含已编译对象/共享库产物，却没有回归测试配置、设置 DDL、升级路径、兼容矩阵或运维恢复流程。它依赖行 OID 与旧服务器结构，无法原样用于当前 PostgreSQL。

真实完整性需求首先应定义威胁模型。按需要使用受支持的审计日志、受限追加审计表、带规范序列化的密码学哈希或签名、受保护密钥管理以及外部可验证存储。绝不能把数据库密码放进用户可设置的 GUC，也不要用未引用的标识符和数据拼接高权限 `dblink` SQL。
