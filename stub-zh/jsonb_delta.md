## 用法

来源：

- [官方 README](https://github.com/evoludigit/jsonb_delta/blob/e3ed58ca75b22cdef187674ff0447ab47d77091b/README.md)
- [控制文件](https://github.com/evoludigit/jsonb_delta/blob/e3ed58ca75b22cdef187674ff0447ab47d77091b/jsonb_delta.control)
- [版本化 SQL](https://github.com/evoludigit/jsonb_delta/blob/e3ed58ca75b22cdef187674ff0447ab47d77091b/sql/jsonb_delta--0.3.1.sql)
- [变更日志](https://github.com/evoludigit/jsonb_delta/blob/e3ed58ca75b22cdef187674ff0447ab47d77091b/CHANGELOG.md)

`jsonb_delta` 0.3.1 提供 JSONB 变更集、按键编辑数组及合并函数。函数返回新的 JSONB 值；保存结果仍然使用普通 PostgreSQL 行更新与 MVCC。

### 核心流程

启用已安装的扩展，然后按键更新数组元素：

```sql
CREATE EXTENSION jsonb_delta;

SELECT jsonb_array_update_where(
  '{"orders":[{"id":1,"status":"new"}]}'::jsonb,
  'orders', 'id', '1'::jsonb,
  '{"status":"shipped"}'::jsonb
);
```

返回结果包含更新后的订单。要保存变更，应在 UPDATE 表达式中使用返回值；单独执行 SELECT 不会修改已存储的文档。

### 主要函数

- `jsonb_apply_changeset(doc, ops)` 按顺序应用 JSONB 编辑数组，支持设置、删除、合并、增量和数组编辑；路径可使用字符串或路径段数组。
- `jsonb_array_update_where` 与 `jsonb_array_delete_where` 根据键及 JSONB 值匹配元素。`jsonb_array_contains_id` 检查元素是否存在；`jsonb_array_insert_where` 支持排序插入。
- `jsonb_array_update_where_batch` 合并单个文档的多次编辑；`jsonb_array_update_multi_row` 接收文档数组并返回结果集。
- `jsonb_smart_patch_scalar`、`jsonb_smart_patch_nested` 与 `jsonb_smart_patch_array` 在不同范围内合并。`jsonb_merge_shallow` 与 `jsonb_merge_at_path` 提供显式合并函数。
- `jsonb_extract_id` 在提取标识符时保留数值精度。

### 兼容性与边界

0.3.x 支持 PostgreSQL 14–18：变更日志明确移除了 PostgreSQL 13，README 徽章与 Cargo 功能标记尚未同步。普通使用不需要共享预加载或重启。控制文件将 superuser 设为 false，但未标记为可信扩展；安装 SQL 会创建 C 语言函数，因此仍需相应权限。

扩展不可重定位。大多数变换函数具有不可变、严格及并行安全属性；具体签名应以版本化 SQL 为准，尤其注意可选参数。每个变更集最多包含 10,000 次操作，路径最多包含 1,000 个段；整数增量溢出会报错。

0.3.0 修复了按数值匹配、文本/UUID 批量键及标识符精度问题。0.3.1 保持相同的编译行为，SQL 升级只推进版本标签。更早的安装可经过 0.2.0 和 0.3.0 升级。执行 SQL 升级前应安装匹配的二进制，并替换已加载旧库的会话。
