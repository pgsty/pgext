## 用法

来源：

- [1.0 版本 SQL 对象](https://github.com/zandmitrij/pgchy_similarity/blob/76e4b9ba0530ba16a577141b0375268261626a37/sql/pgchy_similarity--1.0.sql)
- [Tanimoto 实现](https://github.com/zandmitrij/pgchy_similarity/blob/76e4b9ba0530ba16a577141b0375268261626a37/src/pgchy_similarity.c)
- [基准说明](https://github.com/zandmitrij/pgchy_similarity/blob/76e4b9ba0530ba16a577141b0375268261626a37/benchmark/BENCHMARK.md)

`pgchy_similarity` 计算分子指纹的阈值 Tanimoto 相似度。SQL 辅助函数 `smiles_to_fingerprint(text)` 通过不可信 PL/Python、Chython 与 NumPy 生成 2048 位 Morgan 指纹；`is_tanimoto_similar(bytea, bytea, float4)` 比较两个打包指纹。

```sql
CREATE EXTENSION plpython3u;
CREATE EXTENSION pgchy_similarity;
WITH q AS (SELECT smiles_to_fingerprint('CCO') AS fp)
SELECT molecule_id
FROM molecule, q
WHERE is_tanimoto_similar(q.fp, molecule.fingerprint, 0.7);
```

指纹语义取决于精确 Chython/NumPy 版本、半径、长度、位序与打包布局。应固定 Python 环境，并验证已知分子对；任何变化都要求重算所有已存指纹。

不可信 PL/Python 由超级用户控制，可访问服务器操作系统。C 函数会直接解释打包 `bytea` 布局，因此进入原生代码前必须拒绝畸形或长度错误的输入。Tanimoto 阈值只是筛选方法，不代表化学同一性或安全验证。应基准误报/漏报及内存/CPU，并把生成与比较函数只授予受审查角色。
