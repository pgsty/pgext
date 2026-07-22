## 用法

来源：

- [官方 README](https://github.com/hjort/pgafis/blob/98096d694a33414ace34ca51ff194890717ef8da/README.md)
- [1.2 版本扩展 SQL](https://github.com/hjort/pgafis/blob/98096d694a33414ace34ca51ff194890717ef8da/src/pgafis--1.2.sql)
- [官方 PGXS Makefile](https://github.com/hjort/pgafis/blob/98096d694a33414ace34ca51ff194890717ef8da/src/Makefile)

`pgafis` 为 PostgreSQL 增加自动指纹识别系统所需的基础能力。它可以压缩指纹图像、提取与检查细节点模板、评估图像质量，并对一对一或一对多模板匹配进行评分。该扩展封装了旧版 NIST Biometric Image Software (NBIS) 库；它既不提供索引搜索方法，也不会替应用决定身份判定阈值。

### 核心流程

安装扩展后，分别保存原始图像和派生数据。上游示例使用原始 PGM 图像、WSQ 压缩图像、二进制 MDT 模板和文本形式的 XYT 细节点：

```sql
CREATE EXTENSION pgafis;

CREATE TABLE fingerprints (
    id  text PRIMARY KEY,
    pgm bytea,
    wsq bytea,
    mdt bytea,
    xyt text
);

UPDATE fingerprints
SET wsq = cwsq(pgm, 2.25, 300, 300, 8, NULL)
WHERE wsq IS NULL;

UPDATE fingerprints
SET mdt = mindt(wsq, true)
WHERE mdt IS NULL;
```

使用 `bz_match` 比较两个模板。其整数结果是相似度分数；应用必须针对自己的数据选择并验证合适的阈值：

```sql
SELECT bz_match(probe.mdt, candidate.mdt) AS score
FROM fingerprints AS probe
CROSS JOIN fingerprints AS candidate
WHERE probe.id = '101_1'
  AND candidate.id = '101_6';
```

执行一对多识别时，需要显式地为候选行评分并排序或过滤。上游示例执行顺序扫描，因此在将该操作放到请求路径之前，应先针对实际候选集规模做基准测试。

### 函数索引

- `cwsq(bytea, real, int, int, int, int) -> bytea` 将原始指纹图像编码为 WSQ。
- `nfiq(bytea) -> int` 计算 NFIQ 图像质量分数。
- `mindt(bytea)` 与 `mindt(bytea, boolean) -> bytea` 提取二进制细节点模板。
- `mdt2text(bytea) -> text` 将模板渲染为文本，`mdt_mins(bytea) -> int` 返回其中的细节点数量。
- `bz_match(bytea, bytea) -> int` 与 `bz_match(text, text) -> int` 分别计算二进制或文本模板的 Bozorth3 匹配分数。

### 运维说明

版本 `1.2` 可重定位，且未声明预加载要求。但其固定版本的构建文件将 NBIS `5.0.0` 硬编码到 `/opt/nbis-5.0.0`，静态链接 NBIS 归档，传入 `-m32`，并保留 PostgreSQL `9.3` 的源码树内构建默认值。这些都是显著的可移植性约束：应针对实际使用的 PostgreSQL、编译器、架构、NBIS 与图像库组合重新构建并测试。

指纹模板与阈值属于生物识别安全决策，而不是通用常量。应使用有代表性的数据验证误接受与误拒绝行为，限制对原始图像和模板的访问，并避免把匹配分数表述为身份保证。
