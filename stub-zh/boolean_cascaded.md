## 用法

来源：

- [上游 README](https://github.com/ozum/boolean_cascaded/blob/88502bb17dd17b0bb590eb29c328343010d64a7a/README.md)
- [2.0 版控制文件](https://github.com/ozum/boolean_cascaded/blob/88502bb17dd17b0bb590eb29c328343010d64a7a/boolean_cascaded.control)
- [2.0 版安装 SQL](https://github.com/ozum/boolean_cascaded/blob/88502bb17dd17b0bb590eb29c328343010d64a7a/sql/boolean_cascaded--2.0.sql)
- [转换测试](https://github.com/ozum/boolean_cascaded/blob/88502bb17dd17b0bb590eb29c328343010d64a7a/test/001-conversions.sql)

boolean_cascaded 2.0 是一个纯 SQL 的整数后端类型及一组触发器函数，用于在父子表之间传播禁用状态。2.0 版将本地状态与继承的 false 计数编码到一个整数中；README 部分内容所述的旧复合表示已与安装 SQL 不符。

### 检查数值语义

```sql
CREATE EXTENSION boolean_cascaded WITH SCHEMA public;
SELECT 0::public.boolean_cascaded::boolean AS active,
       1::public.boolean_cascaded::boolean AS inactive;
```

零转换为 true，普通非零值转换为 false。布尔输入会先变成负数哨兵，再由所提供的 BEFORE 触发器在保留继承状态的同时进行规范化。

### 注意事项

- 即使请求其他扩展模式，该类型、若干操作符与触发器函数仍会创建在 public 中。安装前应审查全局名称冲突与所有权。
- 不要直接更新编码整数。负数哨兵与十进制位计数都是实现细节，绕过完整触发器集合可能使父子状态不一致。
- 级联是同步的行触发器工作。大扇出树会放大写入、锁、死锁与复制量；循环或跨模式设计尤其需要谨慎。
- 触发器参数动态指定表和列，但关系名称没有模式限定。调用者搜索路径可能选中非预期的同名关系。
- 非默认状态列的分支使用字面文本 p_status_col 构造 JSON 键，而不是使用变量值，因此所宣称的自定义列名可能不会得到更新。
- 项目最后修改于 2016 年，没有当前 PostgreSQL 兼容矩阵。使用前必须测试插入、父级迁移、并发更新、删除、循环、转储恢复与触发器顺序。
