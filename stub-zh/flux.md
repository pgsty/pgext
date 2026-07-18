## 用法

来源：

- [扩展 control 文件](https://github.com/benchub/flux/blob/836e16f6631907c1b176586e960c460e1d225140/flux.control)
- [使用指南](https://github.com/benchub/flux/blob/836e16f6631907c1b176586e960c460e1d225140/docs/howto.md)
- [已记录的限制](https://github.com/benchub/flux/blob/836e16f6631907c1b176586e960c460e1d225140/docs/limitations.md)
- [SQL 实现](https://github.com/benchub/flux/blob/836e16f6631907c1b176586e960c460e1d225140/flux--0.3.sql)

`flux` 是一个纯 SQL、基于触发器的行历史扩展。它为每个已启用的表记录 `INSERT`、`UPDATE` 和 `DELETE` 的反向变更，从而可以重建某个时间点的行，或返回完整行历史。它要求 `hstore`、主键以及固定的 `_flux` schema。

### 启用并查询历史

```sql
CREATE EXTENSION hstore;
CREATE EXTENSION flux;

CREATE TABLE public.accounts (
  id bigint PRIMARY KEY,
  status text NOT NULL
);

SELECT _flux.enable_change_logging(
  'public', 'accounts', '_accounts_history'
);

INSERT INTO public.accounts VALUES (1, 'new');
UPDATE public.accounts SET status = 'active' WHERE id = 1;

SELECT *
FROM _flux.get_row_history('public', 'accounts', '"id"=>"1"'::hstore);
```

`_flux.enable_change_logging()` 还支持 `include` 或 `exclude` 列清单以及可选的保留周期。`_flux.get_row_from_history()` 按表当前的复合类型重建单行。`_flux.disable_change_logging()` 会移除触发器，而 `_flux.cleanup()` 随后删除已停用的日志表并执行保留策略。

Flux 保存的是反向差异，而不是初始快照。因此，`TRUNCATE` 会丢失重建当前存活行所需的最新状态。它拒绝修改主键值；重命名被跟踪表会破坏元数据；修改主键或列类型需要谨慎规划停用、清理和重新启用流程；被删除的列也不会出现在强类型历史重建结果中。停用操作需要表级排他锁，而清理会破坏性删除选中的历史表。版本 `0.3` 较老且没有当前兼容矩阵，投入运维前应验证锁行为、存储增长、schema 迁移和恢复效果。
