

## 用法

> [cube: 多维立方体数据类型](https://www.postgresql.org/docs/current/cube.html)

`cube` 扩展提供了一种数据类型，用于表示多维立方体（最多支持 100 维）。

```sql
CREATE EXTENSION cube;
```

### 输入格式

```sql
SELECT '(1,2),(3,4)'::cube;         -- 二维立方体
SELECT '(1,2,3)'::cube;              -- 三维点（零体积）
SELECT cube(1.0, 2.0);               -- 从 1 到 2 的一维立方体
SELECT cube(ARRAY[1,2], ARRAY[3,4]); -- 从数组创建的二维立方体
```

### 运算符

| 运算符 | 描述 |
|--------|------|
| `&&` | 重叠 |
| `@>` | 包含 |
| `<@` | 被包含 |
| `->` | 提取第 n 个坐标 |
| `<->` | 欧氏距离 |
| `<#>` | 曼哈顿距离（L-1） |
| `<=>` | 切比雪夫距离（L-inf） |

### 函数

```sql
SELECT cube_dim('(1,2),(3,4)'::cube);                  -- 2
SELECT cube_ll_coord('(1,2),(3,4)'::cube, 1);          -- 1
SELECT cube_ur_coord('(1,2),(3,4)'::cube, 1);          -- 3
SELECT cube_is_point(cube(1,1));                        -- true
SELECT cube_distance('(1,2)'::cube, '(3,4)'::cube);    -- 2.828...
SELECT cube_union('(1,2)'::cube, '(3,4)'::cube);       -- (1,2),(3,4)
SELECT cube_inter('(0,0),(2,2)'::cube, '(1,1),(3,3)'::cube);
SELECT cube_subset(cube('(1,3,5),(6,7,8)'), ARRAY[2]); -- (3),(7)
SELECT cube_enlarge('(1,2),(3,4)'::cube, 0.5, 2);      -- 按半径扩展
```

### GiST 索引与 KNN 搜索

```sql
CREATE INDEX idx ON points USING gist (location);

-- 查找最近的点（距 (0.5, 0.5, 0.5) 最近）
SELECT * FROM points
ORDER BY location <-> cube(ARRAY[0.5, 0.5, 0.5])
LIMIT 1;
```

### 维度处理

当对不同维度的立方体进行运算时，低维立方体的缺失坐标将被视为零。
