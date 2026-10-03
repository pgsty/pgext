## 用法

来源：

- [官方 README.md](https://github.com/jgsn13/piipg/blob/661ae53fa2a110cc535cf9b33a15d9f8ce28f026/README.md)
- [官方 pg_pii_audit.control](https://github.com/jgsn13/piipg/blob/661ae53fa2a110cc535cf9b33a15d9f8ce28f026/pg_pii_audit.control)
- [官方 pg_pii_audit--1.0.sql](https://github.com/jgsn13/piipg/blob/661ae53fa2a110cc535cf9b33a15d9f8ce28f026/sql/pg_pii_audit--1.0.sql)
- [官方 pg_pii_audit.c](https://github.com/jgsn13/piipg/blob/661ae53fa2a110cc535cf9b33a15d9f8ce28f026/src/pg_pii_audit.c)
- [官方 Dockerfile](https://github.com/jgsn13/piipg/blob/661ae53fa2a110cc535cf9b33a15d9f8ce28f026/Dockerfile)

`pg_pii_audit` 1.0 是面向 PostgreSQL 16 的教学原型，用于采样 INSERT 执行计划中形似电子邮件地址的字面量，并记录计数，不能完整识别数据库中的个人信息。

### 基本流程

以超级用户安装扩展，再在需要观察的每个会话中加载钩子：

```sql
CREATE EXTENSION pg_pii_audit;
LOAD 'pg_pii_audit';

CREATE TABLE pii_demo (email text);
INSERT INTO pii_demo VALUES ('alice@example.org');

SELECT table_name, column_name, total_samples,
       matched_samples, column_probability
FROM pii_audit.identified_table_columns;
```

control 文件将模式固定为 `pii_audit`。SQL 脚本创建 `identified_table_columns`，但不会调用 C 函数来加载执行器钩子，因此必须显式执行 `LOAD`。

### 结果解读与限制

`total_samples` 统计检查过的计划常量，`matched_samples` 统计正则匹配次数，`column_probability` 为二者的比值。`table_probability` 使用相同的逐列比值，并非独立的整表估计值。

钩子只检查其遍历到的 text、varchar 和 char 字面量表达式，不能假定它覆盖多行 VALUES、参数、INSERT SELECT、既有数据或非 INSERT 操作。审计写入使用调用者权限；缺少审计表或序列权限可能导致原写入失败。

仅使用可丢弃的测试数据。源码直接将表名和列名拼入 SQL，未作转义；递归保护标记也未在所有异常路径上复位。此修订未提供配置 GUC、其他 PostgreSQL 版本的兼容性保证或明确许可。
