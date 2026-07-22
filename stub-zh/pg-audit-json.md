## 用法

来源：

- [官方 README](https://github.com/m-martinez/pg-audit-json/blob/68bbc4417c44dac145084ccb703a24e2a199977c/README.md)
- [1.0.2 版扩展 SQL](https://github.com/m-martinez/pg-audit-json/blob/68bbc4417c44dac145084ccb703a24e2a199977c/sql/pg-audit-json--1.0.2.sql)
- [官方审计测试](https://github.com/m-martinez/pg-audit-json/blob/68bbc4417c44dac145084ccb703a24e2a199977c/test/sql/audit_test.sql)

`pg-audit-json` 1.0.2 是基于触发器的表审计系统。它在 `audit.log` 中记录 INSERT、UPDATE、DELETE 和 TRUNCATE 事件，将旧行与递归 `jsonb` 差异连同会话、查询、事务、客户端和时间戳元数据一起保存。

### 启用审计

由 PostgreSQL 管理员安装，然后由 `audit.log` 所有者为每张目标表附加触发器：

```sql
CREATE EXTENSION "pg-audit-json";

CREATE TABLE app.customer (
  id bigint PRIMARY KEY,
  email text,
  profile jsonb,
  updated_at timestamptz
);

SELECT audit.audit_table(
  'app.customer'::regclass,
  true,
  true,
  ARRAY['updated_at']::text[]
);

SELECT action, row_data, changed_fields, session_user_name
FROM audit.log
WHERE relid = 'app.customer'::regclass
ORDER BY id DESC;
```

四参数 `audit.audit_table` 分别选择行级日志、查询文本日志和忽略列。单参数包装器会启用行值和查询文本，不忽略任何列。TRUNCATE 仍使用语句级事件；若关闭行审计，INSERT/UPDATE/DELETE 只记录事件而没有行值。

应用可以用 `SET LOCAL` 设置 `audit.application_name` 和 `audit.application_user_name`，附加应用身份。INSERT 把新行写入 `changed_fields`；UPDATE 把旧行写入 `row_data`，只把变更值写入 `changed_fields`；DELETE 保存旧行；只修改忽略列的 UPDATE 不会记录。

### 安全与覆盖边界

审计行与业务变更处于同一事务，因此回滚会同时删除二者。表所有者、超级用户或能够禁用/删除触发器的代码可以绕过审计。它不审计 SELECT、一般 DDL、角色变更、失败语句，也不审计未附加触发器的表；这些要求应使用服务器审计日志。

`client_query` 和完整行 JSON 可能包含凭据或个人数据。应撤销对 `audit.log` 的直接访问，定义保留与脱敏规则，并考虑存储、索引和写放大。扩展把 `audit.log` 配置为可随扩展转储的数据，因此备份也可能包含审计载荷。模式变化会改变行结构，删除并重建表会改变 `relid`；必要时应结合对象身份元数据和保留周期查询。
