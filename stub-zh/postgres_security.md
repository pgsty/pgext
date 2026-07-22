## 用法

来源：

- [扩展 control 文件](https://github.com/MasaoFujii/postgres_security/blob/8a5d11461b5381b4cd3a7cae70d1d97bb73f2711/postgres_security.control)
- [扩展 SQL](https://github.com/MasaoFujii/postgres_security/blob/8a5d11461b5381b4cd3a7cae70d1d97bb73f2711/postgres_security--1.0.0.sql)
- [C 实现](https://github.com/MasaoFujii/postgres_security/blob/8a5d11461b5381b4cd3a7cae70d1d97bb73f2711/postgres_security.c)

`postgres_security` 1.0.0 是最小化的类型开发示例，并非安全功能。它创建模式内基础类型 `postgres_security.text`，存储与输入输出例程模仿 PostgreSQL 内置文本类型。它不提供加密、脱敏、验证、访问控制、认证或审计。

### 检查流程

只能在隔离开发数据库安装，并始终用模式限定自定义类型：

```sql
CREATE EXTENSION postgres_security;

CREATE TABLE security_type_demo (
  value postgres_security.text
);

INSERT INTO security_type_demo VALUES ('example');
SELECT pg_typeof(value), value FROM security_type_demo;
```

不要在 `search_path` 中把 `postgres_security` 模式放到 `pg_catalog` 前面，否则未限定的 `text` 可能意外解析为自定义类型。

### 对象范围

安装脚本创建 `postgres_security` 模式、自定义 `text` 基础类型，以及四个 C 输入输出函数：`textin`、`textout`、`textrecv` 和 `textsend`。它不创建比较操作符、转换、操作符类、约束或安全策略。自定义类型拥有不同的类型 OID，不会自动继承内置文本函数或操作符行为。

### 运维说明

control 文件称扩展可重定位，但 SQL 硬编码创建并使用 `postgres_security` 模式，因此不能依赖重定位。安装还可能与同名现有模式冲突。不要在应用模式、ORM 或安全设计中使用该扩展。若现有数据库中已经出现它，应盘点依赖列，并通过受控转换或数据导出导入迁移到 `pg_catalog.text`，不能假定二进制或操作符兼容。
