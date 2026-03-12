

## 用法

> [db_migrator: 将其他数据库迁移到 PostgreSQL 的工具](https://github.com/cybertec-postgresql/db_migrator)

使用外部数据包装器和特定源插件将数据库从其他数据源迁移到 PostgreSQL 的框架。

### 启用

```sql
CREATE EXTENSION db_migrator;
```

### 可用插件

- **ora_migrator** - Oracle 迁移
- **mysql_migrator** - MySQL/MariaDB 迁移
- **mssql_migrator** - Microsoft SQL Server 迁移

### 完整迁移示例（Oracle）

```sql
-- 设置（以超级用户身份）
CREATE EXTENSION oracle_fdw;
CREATE SERVER oracle FOREIGN DATA WRAPPER oracle_fdw
    OPTIONS (dbserver '//dbserver.mydomain.com/ORADB');
GRANT USAGE ON FOREIGN SERVER oracle TO migrator;
CREATE USER MAPPING FOR migrator SERVER oracle
    OPTIONS (user 'orauser', password 'orapwd');

-- 迁移（以 migrator 用户身份）
CREATE EXTENSION ora_migrator;

SELECT db_migrate(
    plugin => 'ora_migrator',
    server => 'oracle',
    only_schemas => '{TESTSCHEMA1,TESTSCHEMA2}'
);
```

### 分步迁移

需要更多控制时，按阶段执行迁移：

```sql
-- 1. 创建暂存模式并快照元数据
SELECT db_migrate_prepare(
    plugin => 'ora_migrator',
    server => 'oracle',
    only_schemas => '{SCHEMA1}'
);

-- 2. 审查和修改暂存数据
-- 编辑 pgsql_stage 表以自定义类型映射、重命名对象等。
UPDATE pgsql_stage.tables SET migrate = TRUE WHERE ...;

-- 3. 创建模式并迁移数据
SELECT db_migrate_mkforeign(plugin => 'ora_migrator', server => 'oracle');
SELECT db_migrate_tables(plugin => 'ora_migrator');

-- 4. 创建约束和索引
SELECT db_migrate_constraints(plugin => 'ora_migrator');
SELECT db_migrate_indexes(plugin => 'ora_migrator');

-- 5. 清理
SELECT db_migrate_finish();
```

### 关键特性

- 迁移表、序列、索引、约束、视图、函数
- 从源到 PostgreSQL 类型的数据类型映射（可自定义）
- 遇到错误时继续执行，报告哪些对象失败
- 使用 FDW 暂存模式在迁移前检查元数据
- 通过插件函数进行模式和对象名称转换
