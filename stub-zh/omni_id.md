


## 用法

> [omni_id: 标识类型](https://docs.omnigres.org/omni_id/identity_type/)

`omni_id` 扩展引入自定义标识类型，防止不同 ID 类型之间的比较错误，在查询时即可捕获此类缺陷。

### 创建标识类型

```sql
CREATE EXTENSION omni_id;
SELECT identity_type('user_id');
```

这将创建一个由 `bigint`（默认）支持的 `user_id` 类型，自动创建序列和构造函数。

### 参数

| 参数 | 默认值 | 描述 |
|:-----|:-------|:-----|
| `type` | `bigint` | 基础类型（`int`、`smallint`、`uuid`） |
| `sequence` | `<type>_seq` | 序列标识符 |
| `create_sequence` | `true` | 自动创建序列 |
| `increment` | `1` | 序列步长 |
| `cache` | -- | 预分配序列号数量 |
| `cycle` | -- | 达到限制时启用循环 |
| `constructor` | -- | 构造函数名称 |
| `operator_schema` | `public` | 运算符模式 |

### 自动生成的函数

- `user_id(1)` -- 构造函数
- `user_id_nextval()` -- 获取下一个值
- `user_id_currval()` -- 获取当前值
- `user_id_setval(user_id, bool)` -- 设置值
