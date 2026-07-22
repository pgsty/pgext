## 用法

来源：

- [官方 README](https://github.com/segasai/pg_coordconvert/blob/8acfae816cf373204918c8231eabae26c575dbc3/README)
- [1.0.0 版本扩展 SQL](https://github.com/segasai/pg_coordconvert/blob/8acfae816cf373204918c8231eabae26c575dbc3/scripts/coordconv--1.0.0.sql)
- [PostgreSQL 包装实现](https://github.com/segasai/pg_coordconvert/blob/8acfae816cf373204918c8231eabae26c575dbc3/pgcoord.c)

`coordconv` 在 IAU 1958 银道坐标系与 J2000 FK5 赤道坐标系之间转换坐标。所有输入和输出均为十进制度数；每个输出分量由单独的函数返回。

### 核心流程

```sql
CREATE EXTENSION coordconv;

-- FK5 right ascension/declination to Galactic longitude/latitude.
SELECT fk52gall(266.4051, -28.936175),
       fk52galb(266.4051, -28.936175);

-- Galactic longitude/latitude back to FK5 right ascension/declination.
SELECT gal2fk5ra(0.0, 0.0),
       gal2fk5dec(0.0, 0.0);
```

将同一对坐标传给某一转换方向的两个函数，再组合两个标量结果作为输出位置。

### 对象索引

- `fk52gall(ra, dec)` 返回银经。
- `fk52galb(ra, dec)` 返回银纬。
- `gal2fk5ra(l, b)` 返回 J2000 FK5 赤经。
- `gal2fk5dec(l, b)` 返回 J2000 FK5 赤纬。

四个函数的参数和返回值均为 `double precision`，并且都是 `IMMUTABLE`、`STRICT` 和 `PARALLEL SAFE`。

### 运维说明

版本 `1.0.0` 未声明扩展依赖或预加载要求。输入角度单位为度，不是弧度、时分秒字符串或 PostGIS 几何。应随存储值保留坐标系与历元信息：这些函数专门实现 IAU 1958 银道坐标与 J2000 FK5 的转换，不提供任意参考系变换或自行修正。
