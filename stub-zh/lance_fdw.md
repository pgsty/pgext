## 用法

来源：

- [README.md](https://github.com/MisterRaindrop/cloudberry-lance/blob/0663d3062150896d9deb9073c4caecc5b1e911b7/README.md)
- [sql/lance_fdw--0.2.sql](https://github.com/MisterRaindrop/cloudberry-lance/blob/0663d3062150896d9deb9073c4caecc5b1e911b7/sql/lance_fdw--0.2.sql)
- [LICENSE](https://github.com/MisterRaindrop/cloudberry-lance/blob/0663d3062150896d9deb9073c4caecc5b1e911b7/LICENSE)
- [lance_fdw.control](https://github.com/MisterRaindrop/cloudberry-lance/blob/0663d3062150896d9deb9073c4caecc5b1e911b7/lance_fdw.control)

`lance_fdw` 0.2 在 Apache Cloudberry 的各个 segment 上并行读取 Lance 数据集。它针对 Cloudberry 实现，尚无普通 PostgreSQL 兼容性验证。扩展依赖 Lance C 绑定，并支持明确列出的 Arrow 类型子集。

### 核心用法

```sql
CREATE EXTENSION lance_fdw;
CREATE SERVER lance_local FOREIGN DATA WRAPPER lance_fdw
  OPTIONS (base_uri '/data/lance');
IMPORT FOREIGN SCHEMA lance LIMIT TO ("events.lance")
  FROM SERVER lance_local INTO public;
SELECT count(*) FROM public."events.lance";
```

### 运行边界

通过 `CREATE SERVER`、用户映射、外部表及 `IMPORT FOREIGN SCHEMA` 访问数据。扫描只读；语义一致时可下推支持的过滤条件和最近邻操作。不支持的类型会被拒绝，时间戳、排序规则和数值类型的下推有明确限制。参与扫描的 segment 必须能够访问数据集路径和相应凭据。安装匹配的原生文件后，使用 `ALTER EXTENSION lance_fdw UPDATE` 从 0.1 升级到 0.2；部分开发版 0.1 已包含向量距离运算符，0.2 对此记录新版本。原生安装需要超级用户。
