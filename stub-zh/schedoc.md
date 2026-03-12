

## 用法

> [schedoc: 通过 PostgreSQL 对象的 COMMENT 实现模式文档化](https://github.com/ZeroGachis/pg_schedoc)

从 PostgreSQL 对象的 `COMMENT` 构建自动文档。需要 `ddl_historization` 扩展。

### 安装

```sql
CREATE EXTENSION schedoc CASCADE;
SELECT schedoc_start();
```

### 添加列文档

以 JSON 格式在列上设置注释，包含预定义字段：

```sql
COMMENT ON COLUMN my_table.id IS '{"status": "private"}';
COMMENT ON COLUMN my_table.email IS '{"status": "public"}';
COMMENT ON COLUMN my_table.name IS '{"status": "internal"}';
```

### 查询文档

查询已解析的列注释：

```sql
SELECT * FROM schedoc_column_comments;
```

结果：

```
 databasename | tablename | columnname | status
--------------+-----------+------------+---------
 mydb         | my_table  | id         | private
 mydb         | my_table  | email      | public
 mydb         | my_table  | name       | internal
```

### 使用场景

将列元数据与其他系统（如 Django `db_comment`、DBT docs）进行交叉引用，在开发人员和数据分析师之间定义数据契约。
