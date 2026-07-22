## 用法

来源：

- [官方 README](https://github.com/matlads/mapcode-postgres/blob/c74956e43492bcf9a4761af641753ede07ce122b/README.md)
- [SQL 函数定义](https://github.com/matlads/mapcode-postgres/blob/c74956e43492bcf9a4761af641753ede07ce122b/sql/mapcode.sql)
- [函数实现](https://github.com/matlads/mapcode-postgres/blob/c74956e43492bcf9a4761af641753ede07ce122b/src/funcs.c)
- [官方 PGXN 发行页面](https://pgxn.org/dist/mapcode/)

`mapcode` 封装 Mapcode C++ 库，可将经纬度编码成相对于领地的短 mapcode，并将 mapcode 解码回坐标。软件包名是 `mapcode-postgres`，但 PostgreSQL 扩展名是 `mapcode`。

### 编码与解码

```sql
CREATE EXTENSION mapcode;

SELECT mapcode_version();
SELECT mapcode_encode(0.341637, 32.593781, 'UGA');

SELECT mapcode_decode('N7.FR', 'UGA');
```

`mapcode_encode(float8, float8, cstring)` 接受纬度、经度与领地代码，返回包含一个或多个逗号分隔 mapcode 的文本；没有结果时返回 `NULL`。`mapcode_decode(cstring, cstring)` 返回 `latitude,longitude` 文本，而不是几何类型或结构化 SQL 类型。`mapcode_version()` 报告捆绑的 mapcode 库版本。

### 数据处理

调用前应校验坐标：实现会规范化越界值，而不是拒绝它们。不要假设第一个编码结果就是唯一有效代码；应显式解析逗号分隔的结果与解码坐标。每个 mapcode 都应连同领地一起存储，因为解码依赖领地。

这个旧封装捆绑了特定的 Mapcode 库快照，且不与 PostGIS 集成。应固定并记录 `mapcode_version()`，使用代表性领地与边界坐标对消费代码的应用做一致性测试，并定义 mapcode 库数据更新时的迁移方式。
