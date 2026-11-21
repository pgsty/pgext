## 用法

来源：

- [pg_gaps.control](https://github.com/gburd/pg_gaps/blob/d35680cc3303e9791465256bfa9f5d8feb5615c2/pg_gaps.control)
- [pg_gaps--0.1.0.sql](https://github.com/gburd/pg_gaps/blob/d35680cc3303e9791465256bfa9f5d8feb5615c2/pg_gaps--0.1.0.sql)
- [ROADMAP.md](https://github.com/gburd/pg_gaps/blob/d35680cc3303e9791465256bfa9f5d8feb5615c2/ROADMAP.md)
- [doc/DESIGN.md](https://github.com/gburd/pg_gaps/blob/d35680cc3303e9791465256bfa9f5d8feb5615c2/doc/DESIGN.md)
- [LICENSE](https://github.com/gburd/pg_gaps/blob/d35680cc3303e9791465256bfa9f5d8feb5615c2/LICENSE)

`pg_gaps` 是面向关系数据缺失值的实验性贝叶斯模型扩展。所审阅的 0.1.0 源码包含模型目录、精确推断、可选采样和条件概率学习。路线图记录已实现阶段，其状态优先于 README 中“尚无功能”的过时描述。它面向 PostgreSQL 17–19，安装在固定的 `gaps` 模式中。

### 定义简单模型

以超级用户创建扩展。模型函数使用调用者的表权限，目录表按角色实施行级安全控制。

```sql
CREATE EXTENSION pg_gaps;
CREATE TABLE observations (id integer PRIMARY KEY, state text);
INSERT INTO observations VALUES (1, 'yes'), (2, NULL);
SELECT gaps.create_model('example');
SELECT gaps.declare_attribute(
  'example', 'observations', 'state', ARRAY['yes', 'no']);
SELECT gaps.add_stripe('example', 'state', 'observations', 'state');
SELECT gaps.add_factor(
  'example', 'prior', 'state', ARRAY[]::text[], ARRAY[0.6,0.4]::float8[]);
SELECT * FROM gaps.validate('example');
SELECT * FROM gaps.marginals('example');
SELECT * FROM gaps.impute('example');
```

声明的属性有有限取值域，实体键默认为表主键。条带表示一组随机变量，因子关联条件概率表，父映射连接相关实体。`gaps.set_cpt()` 替换因子概率，`gaps.learn_cpts()` 根据完整观测估计概率，并可使用加性平滑。

### 理解结果

`gaps.marginals()` 返回每个缺失值的分布。同一连通分量内的变量可能相关，因此不能通过相乘边缘概率得到联合分布。对于精确计算的分量，`gaps.impute()` 返回联合概率最大的赋值；使用采样时，它返回逐变量的边缘最优选择，不保证联合最优，联合概率字段为 NULL。

推断默认执行精确计算，超过资源边界时报告错误。将 `gaps.samples` 设置为正数，可对超出边界的分量启用采样；采样链之间一致并不能证明结果正确。`gaps.max_factor_cells` 和 `gaps.max_total_cells` 限制推断内存分配。

### 边界

目录表及其序列已登记到备份恢复机制中。读取数据遵循调用者权限；学得的概率对获准读取模型的角色可见，因此应将模型视为可能泄露训练观测信息的数据。

模式变更可能破坏模型假设，受会话设置影响的常量输入仍是待解决事项。验证采取保守策略，结构证明失败不一定意味着实际展开后的模型无效。该早期源码版本尚无稳定升级保证；应保留模型备份，并在隔离数据库中测试变更。
