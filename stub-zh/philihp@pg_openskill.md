## 用法

来源：

- [官方 README](https://github.com/philihp/pg_openskill/blob/v0.1.0/README.md)
- [版本化安装 SQL](https://github.com/philihp/pg_openskill/blob/v0.1.0/pg_openskill--0.1.0.sql)
- [database.dev 软件包](https://database.dev/philihp/pg_openskill)

`philihp@pg_openskill` 是 database.dev TLE 软件包，在 `openskill` schema 中以纯 SQL 与 PL/pgSQL 实现 OpenSkill Plackett–Luce 评分。

### 启用

虽然源码文件使用 `pg_openskill` 词干，正式安装身份是带命名空间的名称。通过 dbdev 安装，并创建带引号的扩展名：

```sql
SELECT dbdev.install('philihp@pg_openskill');
CREATE EXTENSION "philihp@pg_openskill" VERSION '0.1.0';
```

该软件包依赖 database.dev installer 与 `pg_tle`。其 CI 验证 PostgreSQL 16；官方来源没有声明其他大版本。

### 评分与排序

`openskill.rating` 返回默认 `(mu, sigma)` 对。`openskill.rate` 按完赛顺序更新各队，也可以用 rank 数组表示胜负与平局。`openskill.ordinal` 计算保守分数 `mu - 3 * sigma`。

```sql
SELECT openskill.rating();

SELECT openskill.rate(
  '[[{"mu":25,"sigma":8.333333333333334}],
    [{"mu":25,"sigma":8.333333333333334}]]'::jsonb
);

SELECT openskill.ordinal(openskill.rating());
```

每队一名选手时，`openskill.rate_1v1` 接收 `openskill.rating` 值数组并返回更新后的数组。

### 数值边界

实现保持 openskill.js 5.0.1 的运算顺序与指数函数行为，测试套件会比较浮点位模式。评分仍取决于输入顺序与历史结果；需要增量历史时，应确定性地应用比赛并保存返回状态。

