## 用法

来源：

- [官方文档](https://github.com/lgrosz/pg_climb/blob/8dad0d7239e0cb6761a729f5be45e6698ecae733/README.md)
- [扩展控制文件](https://github.com/lgrosz/pg_climb/blob/8dad0d7239e0cb6761a729f5be45e6698ecae733/pg_climb.control)
- [官方仓库](https://github.com/lgrosz/pg_climb)

`pg_climb` 提供攀岩难度等级转换与实用函数。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `pg_climb`：

```sql
CREATE EXTENSION pg_climb;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
CREATE TABLE routes (grade grade);
INSERT INTO routes VALUES ('V5'::grade), ('F7A+'::grade);

SELECT grade, grade_type(grade)
FROM routes
ORDER BY grade;
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `btree_grade_ops` | OPERATOR CLASS | 供索引使用的操作符类。 |
| `grade` | TYPE | 扩展创建的用户数据类型。 |
| `grade_cmp` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `grade_eq` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `grade_ge` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `grade_gt` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `grade_in` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `grade_le` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 已验证的 PostgreSQL 主版本证据覆盖 17；不要推断未列出的主版本。
