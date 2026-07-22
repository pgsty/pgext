## 用法

来源：

- [官方上游仓库的归档快照](https://github.pkg.st/dmtkfs/pg-tamper-log)
- [最后已知的官方仓库地址](https://github.com/dmtkfs/pg-tamper-log)

`pg_tamperlog` 是一个教学用途的纯 SQL/PL/pgSQL 追加式审计事件扩展。每个 `audit_log` 行保存 JSON 事件、时间戳、前一个哈希和当前 SHA-256 哈希；`BEFORE INSERT` 触发器延伸哈希链，`tamper_log_verify` 则报告完整性失败。

```sql
CREATE EXTENSION pgcrypto;
CREATE EXTENSION pg_tamperlog;

INSERT INTO audit_log (event)
VALUES ('{"user":"alice","action":"login"}');

SELECT *
FROM tamper_log_verify
WHERE integrity_check IS NOT NULL;
```

1.1 版本可选用 `pg_tamperlog_rust` 加速哈希。验证会扫描整张表，因此调度时要明确预算 I/O 与延迟。只有在存在可信检查点或预期链头副本时，哈希链才能让编辑或删除变得可检测；数据库超级用户仍能修改表、函数、触发器和验证代码。

上游明确把项目定位为教学用途。应隔离日志写入者与管理员，撤销更新/删除和 DDL 权限，串行化并发追加，把签名检查点和备份导出到数据库主机之外，并测试回滚、恢复、截断、重排、故障切换以及无密钥哈希的威胁假设。

官方 GitHub 仓库目前返回 HTTP 404；上述来源是该仓库的归档渲染。不要安装未经验证的二进制或 fork。生产使用前必须恢复内容寻址的源码树，并审计准确 SQL。
