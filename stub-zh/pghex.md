## 用法

来源：

- [官方 README](https://github.com/k2bd/pghex/blob/0afc072f18799e36071ed6a14e44478c84db7a63/README.md)
- [SQL API 实现](https://github.com/k2bd/pghex/blob/0afc072f18799e36071ed6a14e44478c84db7a63/src/lib.rs)
- [六边形网格算法](https://github.com/k2bd/pghex/blob/0afc072f18799e36071ed6a14e44478c84db7a63/src/hex_alg.rs)

`pghex` 增加 `hex` 坐标类型与常用六边形网格算法。它适用于棋盘游戏、地图单元格、视野计算，以及其他需要在 SQL 中完成邻居、距离、直线、环、范围或螺旋操作的应用。

### 核心流程

坐标写作二轴值 `[q, r]`；实现会按 `s = -q - r` 推导第三个立方体坐标。

```sql
CREATE EXTENSION pghex;

CREATE TABLE units (
    name text PRIMARY KEY,
    position hex NOT NULL,
    vision_range integer NOT NULL
);

INSERT INTO units VALUES
    ('Hero', '[0, 1]', 3),
    ('Goblin', '[5, 1]', 2);

SELECT name, tile
FROM units
CROSS JOIN LATERAL hexes_in_range(position, vision_range) AS tile;

SELECT hex_distance('[0, 1]'::hex, '[5, 1]'::hex);
SELECT * FROM linedraw('[0, 1]'::hex, '[5, 1]'::hex);
```

`+`、`-` 与 `=` 操作符分别对 `hex` 值执行加法、减法与比较。

### 主要函数

- `neighbors(hex)` 与 `diagonals(hex)` 返回六个相邻或对角单元格。
- `hex_distance(hex, hex)` 返回立方体网格距离。
- `linedraw(hex, hex)` 返回直线经过的单元格，并包含两个端点。
- `hexes_in_range(hex, integer)` 返回指定距离内的所有单元格。
- `ring_path(hex, integer)` 返回单个环；`spiral_path(hex, integer)` 返回中心与逐层环。

路径与邻域函数都是集合返回函数；按表中每行调用时，应把它们放在 `FROM` 中或配合 `LATERAL` 使用。

### 注意事项

已复核软件包报告版本 `0.0.0`；上游只说明开发环境安装，并未给出稳定兼容保证。其 Cargo 特性覆盖 PostgreSQL `12` 到 `17`。扩展不可重定位，control 文件要求超级用户安装。应用应校验半径与距离参数；大范围的返回行数按平方增长。整个应用必须保持统一的坐标约定；不需要预加载或重启。
