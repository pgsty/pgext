## 用法

来源：

- [README.md](https://github.com/Khraben/AutoAudit-Extension/blob/0ceda962dcc56f2c35fcd563a5814c2e9bfc6f13/README.md)
- [autoaudit--1.0.sql](https://github.com/Khraben/AutoAudit-Extension/blob/0ceda962dcc56f2c35fcd563a5814c2e9bfc6f13/autoaudit--1.0.sql)
- [autoaudit.control](https://github.com/Khraben/AutoAudit-Extension/blob/0ceda962dcc56f2c35fcd563a5814c2e9bfc6f13/autoaudit.control)

`autoaudit` 1.0 是课程项目中的 SQL 扩展，将行的插入、更新和删除记录到审计表。安装会为现有用户表注册触发器，并为之后创建的表安装事件触发器，因此启用前应审查其全库影响。

### 核心用法

```sql
CREATE EXTENSION autoaudit;
SELECT operationtype, tablename, username, databefore, dataafter
FROM autoaudit.audit_log ORDER BY id DESC LIMIT 20;
```

### 运行边界

`autoaudit.audit_log` 保存操作、表名、时间、用户、客户端地址及变更前后的 JSONB 值。`autoaudit.audit_function` 是行触发器，`autoaudit.ddl_handler` 在建表后创建触发器。审计与业务变更处于同一事务，业务回滚也会回滚对应审计行。

创建事件触发器要求超级用户。函数以调用者权限执行，安装脚本撤销了模式的公共访问权；其他应用角色需要明确的审计模式及审计表权限，否则写入可能失败。应先在隔离数据库中检查标识符处理、已有触发器名称和表所有者。

SQL 不使用预加载或共享库，控制文件中的库名未被使用。仓库未声明许可证或 PostgreSQL 主版本矩阵。这是教学实现，并非经过独立验证、防篡改的合规审计系统。
