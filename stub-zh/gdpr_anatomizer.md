## 用法

来源：

- [Official README](https://github.com/pavilkau/postgresql_anatomy/blob/431772c17f1a7931414d8066b744150d41b00770/README.md)
- [Extension control file](https://github.com/pavilkau/postgresql_anatomy/blob/431772c17f1a7931414d8066b744150d41b00770/gdpr_anatomizer.control)
- [Extension SQL implementation](https://github.com/pavilkau/postgresql_anatomy/blob/431772c17f1a7931414d8066b744150d41b00770/gdpr_anatomizer--0.0.1.sql)

gdpr_anatomizer 是一个使用 PL/Python 实现 Anatomy 隐私发布算法的研究原型。它把准标识符和敏感属性拆分到通过关系连接的表中。上游明确说明它处理大数据时效率低，目标是学习和调试；扩展本身既不能证明符合 GDPR，也不构成隐私保证。

### 核心流程

扩展依赖不受信任语言 `plpython3u`。使用前必须检查安装 SQL：安装过程会调用初始化函数，删除并重建未限定模式名的示例表和相似度表。只能在一次性数据库中运行，然后分析数据集副本并生成 Anatomy 表：

```sql
CREATE EXTENSION gdpr_anatomizer CASCADE;

SELECT analyze_dataset('public.input_copy', 'sensitive_value');

SELECT anatomy(
  'public.input_copy',
  'sensitive_value',
  ARRAY['age_band', 'region'],
  3
);
```

默认输出名为 `qi_table` 和 `sa_table`。公开结果前要检查生成的等价类，并确认抑制效果满足应用的威胁模型。

### 重要对象

- `analyze_dataset` 报告敏感值分布及候选多样性等级。
- `anatomy` 构造准标识符表和敏感属性表；其参数还会控制模式、输出名称和引用列。
- `apply_sa_tag` 和 `remove_sa_tags` 管理敏感属性所用标签。
- `set_del_eq_class_trigger` 安装等价类删除处理。
- `init_csv_dataset` 从服务端文件导入，并重建演示表。

### 安全与数据保护说明

PL/Python 是不受信任语言，其代码拥有数据库服务端进程的能力。该原型会把大部分数据集载入内存、使用随机分组，并用字符串格式化拼接动态 SQL，因此不可信标识符或值可导致 SQL 注入。若干辅助函数会删除或创建固定且未限定模式名的表，`init_csv_dataset` 还会读取服务端本地路径。只能使用受限的一次性数据库，检查每个生成表，绝不能把输出当作经过独立验证的匿名化结果。
