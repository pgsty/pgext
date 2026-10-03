## 用法

来源：

- [README.md](https://github.com/sipesistemas/pgaudix/blob/b552a8afe614e403aec7e8784806d344782969e3/README.md)
- [pgaudix.control](https://github.com/sipesistemas/pgaudix/blob/b552a8afe614e403aec7e8784806d344782969e3/pgaudix.control)
- [pgaudix--1.0.0.sql](https://github.com/sipesistemas/pgaudix/blob/b552a8afe614e403aec7e8784806d344782969e3/pgaudix--1.0.0.sql)

`pgaudix` 1.0.0 将普通表的行镜像到审计表，并同步支持的 DDL 变更。此版本明确拒绝分区表、分区和使用继承的表。

### 核心工作流

```sql
CREATE EXTENSION pgaudix;
CREATE TABLE audit_example (id bigint PRIMARY KEY, amount numeric);
SELECT pgaudix.enable('audit_example'::regclass);
INSERT INTO audit_example VALUES (1, 10);
UPDATE audit_example SET amount = 20 WHERE id = 1;
SELECT audit_operation, id, amount FROM audit_example_audit ORDER BY audit_id;
SELECT * FROM pgaudix.status();
```

### 权限与身份

上游要求 PostgreSQL 16 及以上版本。安装需要超级用户，无需预加载。函数不向 `PUBLIC` 开放执行权限。对需要使用 `pgaudix.enable`、`pgaudix.disable` 与 `pgaudix.status` 的表所有者，应分别授予模式使用权和指定函数执行权；审计表读取权限另行授予。

`pgaudix.app_user` 与 `pgaudix.app_user_ip` 可在事务内记录应用提供的身份信息。这些值属于应用自行声明的信息，与已认证的数据库角色和客户端地址不同。

### 保留与升级

`pgaudix.disable(table)` 停止审计但保留历史；传入 `drop_data := true` 会删除审计表。删除源列或源表可能删除对应历史；如果需要保留，应先停用审计。更新操作记录新行，清空操作仅记录一次操作而不保留逐行删除内容。扩展不自动实施保留策略。

所核验发布仅包含 1.0.0 安装脚本，没有从 0.2.0 升级的路径。不能假定 `ALTER EXTENSION` 会无损迁移，应先保全已有审计数据并明确规划迁移。恢复扩展依赖及镜像表结构时，需要服务器已具备扩展共享库。未发现明确的上游许可证。
