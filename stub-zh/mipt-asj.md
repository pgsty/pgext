## 用法

来源：

- [已复核 commit 的 mipt-asj README](https://github.com/leskin-in/mipt-asj/blob/ad19f53317f1ab189e6615d5a5c5a5df77609c37/README.md)
- [版本 0.1 的 SQL 接口](https://github.com/leskin-in/mipt-asj/blob/ad19f53317f1ab189e6615d5a5c5a5df77609c37/mipt-asj--0.1.sql)
- [近似字符串连接论文](http://www.vldb.org/pvldb/vol11/p53-tao.pdf)

`mipt-asj` 基于 Tao-Deng-Stonebraker 方法实现支持缩写规则的近似字符串连接。它安装在固定的 `mipt_asj` schema 中，并提供三阶段工作流：`calc_dict` 推导缩写规则，`calc_pairs` 生成候选字符串对，`cmp` 执行最终相似度判断。它只适合允许近似实体匹配的场景，不能提供精确的关系等值语义。

### 核心工作流

先构建或整理双列规则表，再从两个源关系生成候选项，最后使用同一精确度阈值过滤每一对候选项。

```sql
CREATE EXTENSION "mipt-asj";

CREATE TABLE rules(f varchar, a varchar);
INSERT INTO rules
SELECT * FROM mipt_asj.calc_dict(
  'public.full_forms'::regclass::oid, 'name',
  'public.abbreviations'::regclass::oid, 'name'
);

SELECT p.s1, p.s2
FROM mipt_asj.calc_pairs(
  'public.left_names'::regclass::oid, 'name',
  'public.right_names'::regclass::oid, 'name',
  'public.rules'::regclass::oid, 'f', 'a', 0.7
) AS p
WHERE mipt_asj.cmp(
  p.s1, p.s2,
  'public.rules'::regclass::oid, 'f', 'a', 0.7
);
```

`exactness` 参数范围从零到一；README 以 `0.7` 为例。阈值越低，进入候选集的字符串对越多，通常误报也越多。生产环境选定阈值前，应使用带标签数据测量召回率和精确率。

### 函数索引

- `calc_dict` 按关系 OID 和列名读取全称关系与缩写关系，并返回规则对 `(f, a)`。
- `calc_pairs` 读取两个源关系和规则关系，返回候选对 `(s1, s2)`。候选项可能是误报。
- `cmp` 使用相同的规则关系与阈值比较一个字符串对，并返回最终布尔判断。

### 运维注意事项

- 仓库已经归档，上游把版本 `0.1` 标为 alpha。其 API、内部实现和现代 PostgreSQL 兼容性均不再维护。
- 接口接收关系 OID 和列名。应使用带 schema 限定的 `regclass` 表达式，只接受可信标识符；引用表被删除并以新 OID 重建后，需要重新生成输入。
- 候选生成不是最终谓词。必须始终应用 `cmp`，并确保两个阶段使用一致的规则表与 `exactness` 值。
- 相似度质量取决于缩写语料和阈值。自动执行合并或更新前，应使用领域专用的已标注字符串对进行基准测试，并复核错误匹配。
