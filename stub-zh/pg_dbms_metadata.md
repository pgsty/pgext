## 用法

来源：

- [1.1.0 README](https://github.com/HexaCluster/pg_dbms_metadata/blob/v1.1.0/README.md)
- [SQL 定义](https://github.com/HexaCluster/pg_dbms_metadata/blob/v1.1.0/sql/pg_dbms_metadata--1.1.0.sql)
- [升级脚本](https://github.com/HexaCluster/pg_dbms_metadata/blob/v1.1.0/updates/pg_dbms_metadata--1.0.0--1.1.0.sql)
- [发布说明](https://github.com/HexaCluster/pg_dbms_metadata/releases/tag/v1.1.0)

`pg_dbms_metadata` 1.1.0 通过 Oracle 风格的 API 提取 PostgreSQL 对象 DDL。它由 SQL/PL/pgSQL 实现，无需预加载；只实现 GET 系列函数，不提供 Oracle FETCH 或 XML 接口。

### 启用

```sql
CREATE EXTENSION pg_dbms_metadata;
```

### GET_DDL

提取数据库对象的 DDL。支持的类型：TABLE、VIEW、SEQUENCE、PROCEDURE、FUNCTION、TRIGGER、INDEX、CONSTRAINT、CHECK_CONSTRAINT、REF_CONSTRAINT、TYPE、ENUM。

```sql
SELECT dbms_metadata.get_ddl('TABLE', 'employees', 'public');
SELECT dbms_metadata.get_ddl('VIEW', 'active_users', 'myschema');
SELECT dbms_metadata.get_ddl('FUNCTION', 'calculate_tax', 'public');
SELECT dbms_metadata.get_ddl('INDEX', 'idx_name', 'public');

-- Schema is optional; uses search_path when omitted
SELECT dbms_metadata.get_ddl('SEQUENCE', 'my_seq');
```

### GET_DEPENDENT_DDL

提取基础对象的所有指定类型依赖对象的 DDL。支持：SEQUENCE、TRIGGER、CONSTRAINT、REF_CONSTRAINT、INDEX、ENUM。

```sql
SELECT dbms_metadata.get_dependent_ddl('CONSTRAINT', 'employees', 'public');
SELECT dbms_metadata.get_dependent_ddl('INDEX', 'orders', 'sales');
SELECT dbms_metadata.get_dependent_ddl('TRIGGER', 'accounts');
```

### GET_GRANTED_DDL

提取用于重建已授予权限和角色的 SQL 语句。支持：ROLE_GRANT。

```sql
SELECT dbms_metadata.get_granted_ddl('ROLE_GRANT', 'user_test');
```

### SET_TRANSFORM_PARAM

通过会话级转换参数自定义 DDL 输出：

```sql
-- Append SQL terminator (;) to each DDL statement
CALL dbms_metadata.set_transform_param('SQLTERMINATOR', true);

-- Exclude constraints from TABLE DDL
CALL dbms_metadata.set_transform_param('CONSTRAINTS', false);

-- Exclude foreign keys from TABLE DDL
CALL dbms_metadata.set_transform_param('REF_CONSTRAINTS', false);

-- Exclude partitioning clauses
CALL dbms_metadata.set_transform_param('PARTITIONING', false);

-- Exclude storage parameters
CALL dbms_metadata.set_transform_param('STORAGE', false);

-- Reset all params to defaults
CALL dbms_metadata.set_transform_param('DEFAULT', true);
```

### 1.1.0 行为与升级

该版本新增 ENUM 提取；显式指定模式时，枚举名称会带模式限定。函数或过程存在重载时，不带参数列表的名称会返回全部匹配重载；在名称后附上参数类型可以选择其中一个：

```sql
SELECT dbms_metadata.get_ddl('FUNCTION', 'my_function(char, char)', 'my_schema');
ALTER EXTENSION pg_dbms_metadata UPDATE TO '1.1.0';
```

1.1.0 还修复了含引号的表注释和布尔转换参数默认值。执行 SQL 更新前须先安装新扩展文件。上游文档列出 PostgreSQL 9.1–18，并对本版本在 12–19 上的升级路径进行测试；软件包支持范围另行记录。

DDL 可见性不局限于调用者拥有的对象。省略模式时按 `search_path` 查找，并要求具有模式 USAGE 权限；TRIGGER 和 REF_CONSTRAINT 调用必须显式指定模式。对象定义敏感时应限制 API 访问。会话转换设置会影响后续提取调用，可以用 DEFAULT 转换重置；该 API 不提供完整数据库备份。
