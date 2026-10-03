## 用法

来源：

- [pgquant.control](https://github.com/pranshu05/pgQuant/blob/73850d036f98757c65b3715c6c1f39524734c187/pgquant.control)
- [README.md](https://github.com/pranshu05/pgQuant/blob/73850d036f98757c65b3715c6c1f39524734c187/README.md)
- [Cargo.toml](https://github.com/pranshu05/pgQuant/blob/73850d036f98757c65b3715c6c1f39524734c187/Cargo.toml)
- [src/lib.rs](https://github.com/pranshu05/pgQuant/blob/73850d036f98757c65b3715c6c1f39524734c187/src/lib.rs)
- [docs/returns.md](https://github.com/pranshu05/pgQuant/blob/73850d036f98757c65b3715c6c1f39524734c187/docs/returns.md)
- [docs/risk.md](https://github.com/pranshu05/pgQuant/blob/73850d036f98757c65b3715c6c1f39524734c187/docs/risk.md)

`pgquant` 可计算简单／对数收益率、累计与年化收益率，以及历史、正态和 Student-t 分布下的 VaR/ES。0.1.0 属于早期版本，波动率模型与投资组合优化仍在计划中。

### 核心用法

```sql
CREATE EXTENSION pgquant;
SELECT pgquant_cumulative_simple_return(ARRAY[0.02,-0.01,0.03]::float8[]);
SELECT pgquant_var_historical(ARRAY[-0.03,0.01,-0.02,0.02]::float8[], 0.95);
SELECT pgquant_es_t(0.001, 0.015, 5.0, 0.95);
```

### 运行边界

上游文档列出 PostgreSQL 15–17。控制文件允许非超级用户安装，但未将扩展标记为可信；没有预加载要求。基于查询的函数要求按顺序返回标的、日期和双精度价格三列，每个标的的首条记录不产生收益率。查询文本应来自可信来源。置信水平、分布假设和输入校验会影响结果；Student-t ES 要求自由度大于一。
