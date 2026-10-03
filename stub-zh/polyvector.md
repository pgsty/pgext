## 用法

来源：

- [polyvector.control](https://github.com/mynelis/polyvector/blob/1393fc58eee67894d2bd5850eef2dd12717422da/polyvector.control)
- [README.md](https://github.com/mynelis/polyvector/blob/1393fc58eee67894d2bd5850eef2dd12717422da/README.md)
- [Makefile](https://github.com/mynelis/polyvector/blob/1393fc58eee67894d2bd5850eef2dd12717422da/Makefile)
- [sql/polyvector--1.0.sql](https://github.com/mynelis/polyvector/blob/1393fc58eee67894d2bd5850eef2dd12717422da/sql/polyvector--1.0.sql)

`polyvector` 是基于 pgvector 的纯 SQL 封装，在复合值中保存一个有效向量槽。实际安装的 1.0 脚本使用 3、4、5、6 维，槽名称中较大的数字并不代表其真实维度。

### 核心用法

```sql
CREATE EXTENSION vector;
CREATE EXTENSION polyvector;
SELECT polyvec_get_embedding(polyvec_convert('[1,2,3]'::vector));
CREATE TABLE polyvector_sample (
  id bigint PRIMARY KEY,
  embedding polyvec CHECK (polyvec_validate(embedding))
);
```

### 运行边界

先启用 `vector`，再由超级用户安装；内部使用 PL/pgSQL，无需预加载。`polyvec_validate` 检查单槽约束，转换、读取和设置函数按维度选择槽。扩展还提供 L2、余弦及内积检索辅助函数。应用表应添加校验约束。独立的生产示例不是 Makefile 选中的安装 SQL，不能据此声称本版本支持 384/768/1024/1536 维。
