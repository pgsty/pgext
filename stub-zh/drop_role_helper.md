## 用法

来源：

- [官方 README](https://github.com/cybertec-postgresql/drop_role_helper/blob/master/README.md)
- [扩展控制文件](https://github.com/cybertec-postgresql/drop_role_helper/blob/master/drop_role_helper.control)
- [1.0 版扩展 SQL](https://github.com/cybertec-postgresql/drop_role_helper/blob/master/drop_role_helper--1.0.sql)
- [PostgreSQL 角色删除文档](https://www.postgresql.org/docs/current/role-removal.html)

drop_role_helper 提供一个返回集合的函数，用于生成删除 PostgreSQL 角色之前所需的撤权语句。它不会执行这些语句，也不会转移所有权、删除角色拥有的对象、访问其他数据库或删除角色本身。

### 核心流程

在目标角色存在依赖的每个数据库中执行所有权处理与生成的撤权语句。在 psql 中，下面的元命令会把每一行返回值作为 SQL 执行。

```sql
CREATE EXTENSION drop_role_helper SCHEMA public;

REASSIGN OWNED BY dummy TO otheruser;
DROP OWNED BY dummy;

SELECT * FROM public.drop_role_helper('dummy') \gexec
```

在集群的每个数据库中重复这一流程后，再删除集群级角色：

```sql
DROP ROLE dummy;
```

如果客户端没有 psql 的执行功能，应先检查生成的文本：

```sql
SELECT * FROM public.drop_role_helper('dummy');
```

### 生成的语句

该函数通过 PostgreSQL 的角色 OID 类型接收角色名，并按受影响对象组返回语句。1.0 版覆盖关系与序列、列、函数与过程、模式、大对象、类型、语言、数据库、表空间、外部服务器和外部数据包装器上的显式权限，也会为匹配的默认权限生成语句。

### 权限与注意事项

调用生成器本身不需要特殊权限，但执行其输出必须有权撤销所有受影响的授权，通常需要超级用户权限或对象所有权。对于默认权限，执行者还必须是其默认权限将被修改的角色的成员。

所有权处理仍是独立步骤。重新分配所有权并不覆盖所有对象类别，因此官方流程仍包含清理角色拥有对象的步骤。角色依赖属于数据库本地状态，而角色本身属于集群级状态；在一个数据库中成功处理，并不能证明该角色已可从整个集群删除。
