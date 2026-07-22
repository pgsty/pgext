## 用法

来源：

- [扩展 SQL](https://github.com/nate1001/chess_index/blob/4d6d473ea9c8238f06d8759f65862a382621f33e/sql/chess_index.sql)
- [C 实现](https://github.com/nate1001/chess_index/blob/4d6d473ea9c8238f06d8759f65862a382621f33e/src/chess_index.c)
- [回归测试示例](https://github.com/nate1001/chess_index/blob/4d6d473ea9c8238f06d8759f65862a382621f33e/test/sql/setup.sql)
- [构建文件](https://github.com/nate1001/chess_index/blob/4d6d473ea9c8238f06d8759f65862a382621f33e/Makefile)
- [扩展控制文件](https://github.com/nate1001/chess_index/blob/4d6d473ea9c8238f06d8759f65862a382621f33e/chess_index.control)

`chess_index` 为国际象棋局面、棋子与格子定义紧凑的 PostgreSQL 类型，以及用于建立索引的比较与哈希/B-tree 运算符类。该项目仍是未完成的预发布实现；只能用于可丢弃数据上的研究，不要用于依赖正确性的棋局存储或索引。

### 核心流程

`board` 输入是一种缩短为四字段的 FEN：棋子布局、行棋方、易位权和吃过路兵目标；其存储表示不包含半回合与完整回合计数。

```sql
CREATE EXTENSION chess_index;

CREATE TABLE positions (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    position board NOT NULL
);

INSERT INTO positions (position)
VALUES ('rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq -');

SELECT position, pcount(position), side(position), pretty(position)
FROM positions;
```

SQL 为 `board`、`square` 和 `pindex` 声明了哈希及 B-tree 运算符类，因此语法上可以创建索引：

```sql
CREATE INDEX positions_board_hash ON positions USING hash (position);
```

在修复并验证下面的实现缺陷之前，不要依赖这种索引。

### 重要对象

- `board`：紧凑局面值，提供 `pcount(board)`、`side(board)` 和 `pieces(board, side)` 访问函数。
- `side`、`piece`、`pindex`、`square` 和 `piecesquare`：紧凑的棋局组件类型。
- `cfile`、`rank` 和 `diagonal`：从格子派生的坐标类型。
- `pretty(board, boolean, boolean)`：把局面格式化为文本或 Unicode 棋子。
- `invert(board)`：在文本表示中交换棋子大小写与行棋方。
- 哈希与 B-tree 运算符/运算符类为部分类型提供相等和排序语义。

### 正确性与打包注意事项

在复核的提交中，Makefile 指向 `sql/chess_index--0.0.1.sql`，但仓库只有 `sql/chess_index.sql`；因此未经补丁的源码安装可能在能够执行 `CREATE EXTENSION` 前就失败。

更关键的是，C 比较器在比较值时使用了指针大小，而不是已存储 `board` 结构的大小。相等与排序可能忽略局面的相关字节，而哈希采用另一种表示，可能导致错误比较和不安全的索引语义。代码还在固定 100 字节数组中保留原始 FEN，边界检查却允许达到数组名义上限的输入。请把解析和索引都视为不可信的实验代码，保留备份，并在使用前以对抗性 FEN 输入验证任何本地修补版本。
