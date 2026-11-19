## 用法

来源：

- [官方来源](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/contrib/kingres_ledger/README.md)
- [官方来源](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/contrib/kingres_ledger/kingres_ledger.control)
- [官方来源](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/contrib/kingres_ledger/kingres_ledger--0.3.sql)
- [官方来源](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/kingres/README.md)

`kingres_ledger` 0.3 为追加式或可更新账本表记录哈希链与 Merkle 块。只有与事先保存在数据库外部的可信摘要比较，才能检测篡改；内部一致但被整体重写的链不能证明完整性。

### 启用账本

所审查的 KingresDB 源码支持 PostgreSQL 18 与 19 分支。预加载后重启，扩展所有者必须一直保留超级用户身份。仅创建 SQL 对象而不预加载，无法启用账本。以下示例会永久启用追加模式。

```ini
shared_preload_libraries = 'kingres_ledger'
```

```sql
CREATE EXTENSION kingres_ledger;
CREATE TABLE public.events (id bigint, payload text);
SELECT kingres.ledger_enable('public.events'::regclass);
INSERT INTO public.events VALUES (1, 'record');
SELECT kingres.ledger_generate_digest();
SELECT * FROM kingres.ledger_verify();
```

### 验证与权限

`kingres.ledger_enable(relation, 'updatable')` 会改为为更新和删除的行版本创建历史。启用操作不可逆，之后不能切换模式。受保护表无法通过普通管理操作删除或截断。管理 API 需要显式执行授权，调用者须拥有目标表。历史访问与基础表的 RLS、列权限是不同的权限边界。

`kingres.ledger_inclusion_proof` 与 `kingres.ledger_verify_proof` 验证行的包含关系。应将 `kingres.ledger_generate_digest()` 的结果导出到独立可信存储，数据库内记录的副本不是信任锚。新增 `kingres.ledger_head`、`kingres.ledger_blocks_entries` 与 `kingres.ledger_block_preimage` 暴露发布输入；完整外部签名与锚定属于其他组件。

### 维护边界

提交的行进入块且摘要离开数据库之前，仍存在未受保护的时间窗口。特权目录或文件修改、禁用保护机制等行为不在防护边界内。从 0.2 升级时部署匹配库、重启并执行 `ALTER EXTENSION kingres_ledger UPDATE TO '0.3'`；提供的迁移保留已有块与证明。
