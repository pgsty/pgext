## 用法

来源：

- [上游 README](https://github.com/alexandersamoylov/pg_audit_tools/blob/25ec4c7f216980825451a2352005567d03e9bbf4/README.md)
- [0.1 版控制文件](https://github.com/alexandersamoylov/pg_audit_tools/blob/25ec4c7f216980825451a2352005567d03e9bbf4/pg_audit_tools.control)
- [0.1 版安装 SQL](https://github.com/alexandersamoylov/pg_audit_tools/blob/25ec4c7f216980825451a2352005567d03e9bbf4/sql/pg_audit_tools--0.1.sql)
- [MIT 许可证](https://github.com/alexandersamoylov/pg_audit_tools/blob/25ec4c7f216980825451a2352005567d03e9bbf4/LICENSE)

pg_audit_tools 0.1 是一个 2017 年纯 SQL 原型，意图为已有表添加审计列、历史表、序列与行触发器。已发布辅助函数既不够安全，也不够完整，不能原样运行。

### 限制安装

如果在可丢弃数据库中检查，应立即撤销默认公共执行权限：

```sql
BEGIN;
CREATE EXTENSION pg_audit_tools;
REVOKE EXECUTE ON FUNCTION
  audit_tools.table_history_create(character varying, character varying)
  FROM PUBLIC;
REVOKE EXECUTE ON FUNCTION audit_tools.table_history_tf() FROM PUBLIC;
COMMIT;
```

调用 table_history_create 前必须修复并审查这些函数。

### 注意事项

- 两个函数都是 SECURITY DEFINER，默认向 public 开放，而且没有固定搜索路径。若由超级用户安装，会形成严重的权限提升与对象解析风险。
- table_history_create 将 aud_user 定义为 name 类型，却使用 now() 时间戳作为默认值。在正常 PostgreSQL 上该表达式类型不兼容，导致辅助函数无法完成。
- 辅助函数会修改源表并创建多个名称可预测的对象，却没有幂等回滚路径。部分失败可能遗留列或对象。
- 只有 UPDATE 与 DELETE 操作的旧版本会写入历史。已有行与 INSERT 事件不会复制到历史表。
- 这是可修改的应用历史，不是防篡改审计。表所有者与特权角色可以更改或删除它，行触发器也可被禁用或绕过。
- 上游只有一行 README，没有测试、升级脚本、并发指南、保留策略或当前兼容声明。
