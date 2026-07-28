## 用法

来源：

- [Official database.dev 包页面](https://database.dev/dev/tienlen_validation)

`dev@tienlen_validation` — 提供天伦游戏验证函数的扩展。当 SQL 需要这些特殊函数或聚合时使用它。使用上游提供的固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION "dev@tienlen_validation";
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `can_beat(picked_cards TEXT[], played_cards TEXT[], ignore_rules BOOLEAN DEFAULT FALSE)` 是一个扩展函数，返回 `TABLE`。
- `contains_2(cards TEXT[])` 是一个扩展函数，返回 `BOOLEAN`。
- `get_card_face_rank(card TEXT)` 是一个扩展函数，返回 `INTEGER`。
- `get_card_suit_rank(card TEXT)` 是一个扩展函数，返回 `INTEGER`。
- `get_face_and_suit(card TEXT)` 是一个扩展函数，返回 `TABLE`。
- `get_hand_info(cards TEXT[])` 是一个扩展函数，返回 `TABLE`。
- `get_rank(card TEXT)` 是一个扩展函数，返回 `INTEGER`。
- `get_total_ranks(cards TEXT[])` 是一个扩展函数，返回 `INTEGER`。
- `is_double_sequence(cards TEXT[])` 是一个扩展函数，返回 `BOOLEAN`。
- `is_four_of_a_kind(cards TEXT[])` 是一个扩展函数，返回 `BOOLEAN`。
- `is_pair(cards TEXT[])` 是一个扩展函数，返回 `BOOLEAN`。
- `is_sequence(cards TEXT[])` 是一个扩展函数，返回 `BOOLEAN`。
- `is_three_of_a_kind(cards TEXT[])` 是一个扩展函数，返回 `BOOLEAN`。

### 要求与注意事项

- 该目录记录版本 `0.1.0`。
- 这是一个 database.dev/pg_tle 包；在发出引用的 `CREATE EXTENSION` 身份之前，请先注册或生成其包迁移。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行验证。
