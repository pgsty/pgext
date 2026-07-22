## 用法

来源：

- [官方 README](https://github.com/okbob/pltoolbox/blob/6d9efc22d1071a071d54d6db9bd22d1672bf28de/README.md)
- [官方控制文件](https://github.com/okbob/pltoolbox/blob/6d9efc22d1071a071d54d6db9bd22d1672bf28de/pltoolbox.control)
- [官方安装 SQL](https://github.com/okbob/pltoolbox/blob/6d9efc22d1071a071d54d6db9bd22d1672bf28de/pltoolbox--1.0.3.sql)
- [官方回归测试](https://github.com/okbob/pltoolbox/blob/6d9efc22d1071a071d54d6db9bd22d1672bf28de/sql/pltoolbox.sql)

`pltoolbox` 是一组供 PostgreSQL 存储过程使用的 C 工具函数。`1.0.3` 版本创建不可迁移的 `pst` 模式，把格式化、字符串、日期、记录、差异和位图集合辅助函数组织在其中。

### 核心流程

安装扩展后，应使用模式限定的函数名，避免与 PostgreSQL 内置函数产生歧义。

```sql
CREATE EXTENSION pltoolbox;

SELECT pst.format('Hello %s! Now it is %s', 'world', now());
SELECT pst.concat_ws('-', 'alpha', 42, NULL, 'omega');
SELECT pst.record_get_field(ROW(10, 'ten'), 'f2');

WITH s AS (
  SELECT pst.intarray_to_bms(ARRAY[1, 3, 5]) AS bits
)
SELECT pst.bitmapset_num_members(bits), pst.bms_to_intarray(bits)
FROM s;
```

`pst.format` 与 `pst.sprintf` 实现百分号风格的格式化。`pst.concat`、`pst.concat_ws`、`pst.concat_js` 和 `pst.concat_sql` 使用不同的引用规则拼接值。

### 对象索引

- 字符串辅助函数：`pst.reverse`、`pst.left`、`pst.right` 与 `pst.chars_to_array`。
- 日期辅助函数：`pst.next_day` 与 `pst.last_day`。
- 差异辅助函数：`pst.diff_string` 与 `pst.lc_substring`。
- 记录辅助函数：`pst.record_expand`、`pst.record_get_field`、`pst.record_get_fields` 与 `pst.record_set_fields`。
- 位图辅助对象：`pst.bitmapset` 类型、成员与集合操作函数、整数数组转换、`pst.bitmapset_collect` 及并集运算符 `||`。
- 格式化运算符：`%%` 同时安装在 `public` 与 `pst` 中。

### 兼容性与维护

上游 README 表明文档仍不完整，并指出部分函数后来进入 PostgreSQL 核心，但行为并不完全相同。特别是项目的 `left` 与 `right` 对负长度的语义曾经发生变化。当核心中存在同名函数时应始终使用模式限定名，并测试应用实际依赖的边界情况。

尽管名称如此，该扩展并不是通用过程语言。目录将其标记为已废弃，而且其 C 实现依赖位图集合和元组布局等 PostgreSQL 内部结构。应固定经过测试的服务器构建，不能假设它能跨 PostgreSQL 大版本保持二进制或行为兼容。
