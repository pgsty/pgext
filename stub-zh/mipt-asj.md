## 用法

来源：

- [已复核 commit 的 mipt-asj README](https://github.com/leskin-in/mipt-asj/blob/ad19f53317f1ab189e6615d5a5c5a5df77609c37/README.md)
- [已复核 commit 的 mipt-asj.control](https://github.com/leskin-in/mipt-asj/blob/ad19f53317f1ab189e6615d5a5c5a5df77609c37/mipt-asj.control)
- [版本 0.1 的 SQL 接口](https://github.com/leskin-in/mipt-asj/blob/ad19f53317f1ab189e6615d5a5c5a5df77609c37/mipt-asj--0.1.sql)

`mipt-asj` 基于 Tao-Deng-Stonebraker 算法实现支持缩写规则的近似字符串匹配。它把函数安装到 `mipt_asj` 模式中：`calc_dict` 推导全称与缩写规则，`calc_pairs` 生成候选字符串对，`cmp` 执行最终的度量比较。

### 基本比较

```sql
CREATE EXTENSION "mipt-asj";

CREATE TABLE abbreviation_rules (
  full_form varchar,
  abbreviation varchar
);

INSERT INTO abbreviation_rules VALUES ('Massachusetts', 'MA');

SELECT mipt_asj.cmp(
  'Massachusetts',
  'MA',
  'abbreviation_rules'::regclass,
  'full_form',
  'abbreviation',
  0.7
);
```

完整连接流程应先生成或整理规则表，调用 `calc_pairs` 取得候选项，再用 `cmp` 去除误报。两个阶段应使用相同的 `0.7` 一类精确度参数。

### 注意事项

- 仓库已归档，上游明确把版本 `0.1` 标为 API 与内部实现都可能变化的 alpha 版本。
- `calc_pairs` 可能返回误报；它是候选生成器，不是最终等值谓词。
- API 接收关系 OID 和列名，并面向较旧的 PostgreSQL 扩展接口。在现代服务器上使用或接收不可信标识符前，应复核并测试源码。
