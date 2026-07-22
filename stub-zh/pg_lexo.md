## 用法

来源：

- [官方 README](https://github.com/blad3mak3r/pg_lexo/blob/521605b23f5c5cfa80f7a743c536b8040bb81e5a/README.md)
- [0.6 版本 SQL 说明](https://github.com/blad3mak3r/pg_lexo/blob/521605b23f5c5cfa80f7a743c536b8040bb81e5a/sql/pg_lexo--0.6.0.sql)
- [公共函数实现](https://github.com/blad3mak3r/pg_lexo/blob/521605b23f5c5cfa80f7a743c536b8040bb81e5a/src/schema.rs)

`pg_lexo` 0.6.0 提供 Base62 `lexo` 类型，用于稳定的用户可控排序。它可以在已有位置之前、之后或之间生成短值，使项目移动时不必重新编号全部行。该类型提供比较操作符，以及默认 B-tree 和哈希操作符类。

### 核心流程

```sql
CREATE EXTENSION pg_lexo;

CREATE TABLE items (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title text NOT NULL,
    position lexo NOT NULL UNIQUE
);

INSERT INTO items (title, position)
VALUES ('first', lexo_first());

INSERT INTO items (title, position)
SELECT 'last', lexo_after(max(position)) FROM items;

INSERT INTO items (title, position)
VALUES ('middle', lexo_between('H'::lexo, 'I'::lexo));

SELECT * FROM items ORDER BY position;
```

`lexo_first()` 从 `H` 开始。`lexo_after(lexo)`、`lexo_before(lexo)` 和 `lexo_between(lexo, lexo)` 生成相邻位置；`lexo_between` 的任一参数可以为 NULL，表示开放的一端。输入只能包含 `0-9`、`A-Z` 和 `a-z`，其 ASCII 顺序与该类型的排序一致。

### 表辅助函数

`lexo_next(table, position_column, key_column, key_value)` 查找当前最大值并返回下一个位置，可选限定到某个分组。`lexo_add_column(table, column)` 添加 `lexo` 列。`lexo_rebalance(table, position_column, key_column, key_value)` 用均匀分布的位置重写匹配行，并返回更新行数。

这些辅助函数执行动态 SQL，调用者必须拥有相应表权限。重平衡会更新每条选中行，应与应用写入协调执行。位置生成本身不会预留值：并发会话可能得到相同位置，因此应建立合适的唯一约束，并在冲突时重试或串行化移动。当相邻位置之间（或最小值之前）没有可用 Base62 值时，应先重平衡。

### 0.6 版本边界

0.6 版移除了专用 `lexo` 模式。原来的 `lexo.lexorank` 类型改名为 `lexo`，`lexo.first()` 等模式限定函数改为 `lexo_first()` 等名称。旧示例必须迁移后才能使用。此版本的 Cargo 功能通过 pgrx 0.16.1 面向 PostgreSQL 16–18。
