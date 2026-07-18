## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/softvisio/postgresql-softvisio-admin/blob/67e923d3761b0ede5adb16a3ec9cf47779d8099e/README.md)
- [1.2.9 版本 SQL 对象](https://github.com/softvisio/postgresql-softvisio-admin/blob/67e923d3761b0ede5adb16a3ec9cf47779d8099e/softvisio_admin--1.2.9.sql)
- [扩展 control 文件](https://github.com/softvisio/postgresql-softvisio-admin/blob/67e923d3761b0ede5adb16a3ec9cf47779d8099e/softvisio_admin.control)

`softvisio_admin` 是一个纯 SQL 管理扩展，用于创建应用数据库和角色、报告整个实例中过期扩展，以及跨数据库更新已安装扩展。

```sql
CREATE EXTENSION softvisio_admin CASCADE;
CALL create_database('appdb', 'C.UTF-8');
SELECT * FROM outdated_extensions();
```

`create_database` 会在需要时创建带随机密码的角色，再创建数据库并授权。`update_extensions()` 会执行实例范围的扩展更新，并非只读检查。

这些操作需要广泛管理权限，并可能影响所有数据库。应安装在专用管理数据库，从 `PUBLIC` 撤销执行权，只授予受控运维角色，并通过认可的秘密流程接收生成凭据。应审计 SQL 中的动态标识符与远程数据库连接。调用 `update_extensions()` 前必须备份并测试所有扩展升级路径；扩展升级可运行任意迁移 SQL、获取锁、修改对象，或在实例内部分失败。不要向应用角色开放这些例程。
