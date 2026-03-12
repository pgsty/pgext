

## 用法

> [pg_dbms_metadata: 为 PostgreSQL 添加 Oracle DBMS_METADATA 兼容性的扩展](https://github.com/HexaCluster/pg_dbms_metadata)

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

-- 模式可选；省略时使用 search_path
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
-- 在每条 DDL 语句后追加 SQL 终结符（;）
CALL dbms_metadata.set_transform_param('SQLTERMINATOR', true);

-- 从 TABLE DDL 中排除约束
CALL dbms_metadata.set_transform_param('CONSTRAINTS', false);

-- 从 TABLE DDL 中排除外键
CALL dbms_metadata.set_transform_param('REF_CONSTRAINTS', false);

-- 排除分区子句
CALL dbms_metadata.set_transform_param('PARTITIONING', false);

-- 排除存储参数
CALL dbms_metadata.set_transform_param('STORAGE', false);

-- 将所有参数重置为默认值
CALL dbms_metadata.set_transform_param('DEFAULT', true);
```
