---
title: "扩展列表"
breadcrumbs: false
excludeSearch: true
comments: false
weight: 900
---

共有 461 个可用的 PostgreSQL 扩展：

| 扩展 | PG 版本 | 属性 | 分类 | 描述 |
|:----------|:------------|:---------:|:---------:|:--------------|
| {{< ext "timescaledb" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "TIME" >}} | 时序数据库扩展插件 |
| {{< ext "timescaledb_toolkit" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "TIME" >}} | 超表分析查询，时间序列流式处理，以及其他SQL工具 |
| {{< ext "timeseries" "pg_timeseries" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "TIME" >}} | 时序数据API封装 |
| {{< ext "periods" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "TIME" >}} | 为 PERIODs 和 SYSTEM VERSIONING 提供标准 SQL 功能 |
| {{< ext "temporal_tables" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TIME" >}} | 时态表功能支持 |
| {{< ext "emaj" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "TIME" >}} | 让数据库的子集具有细粒度日志和时间旅行功能 |
| {{< ext "table_version" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "TIME" >}} | PostgreSQL 版本控制表扩展 |
| {{< ext "pg_cron" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "TIME" >}} | 定时任务调度器 |
| {{< ext "pg_task" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | {{< category "TIME" >}} | 在特定时间点在后台执行SQL命令 |
| {{< ext "pg_later" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "TIME" >}} | 执行查询，并在稍后异步获取查询结果 |
| {{< ext "pg_background" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TIME" >}} | 在后台运行 SQL 查询 |
| {{< ext "postgis" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "GIS" >}} | PostGIS 几何和地理空间扩展 |
| {{< ext "postgis_topology" "postgis" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "GIS" >}} | PostGIS 拓扑空间类型和函数 |
| {{< ext "postgis_raster" "postgis" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "GIS" >}} | PostGIS 光栅类型和函数 |
| {{< ext "postgis_sfcgal" "postgis" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "GIS" >}} | PostGIS SFCGAL 函数 |
| {{< ext "postgis_tiger_geocoder" "postgis" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "GIS" >}} | PostGIS tiger 地理编码器和反向地理编码器 |
| {{< ext "address_standardizer" "postgis" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "GIS" >}} | 地址标准化函数。 |
| {{< ext "address_standardizer_data_us" "postgis" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "GIS" >}} | 地址标准化函数：美国数据集示例 |
| {{< ext "pgrouting" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "GIS" >}} | 提供寻路能力 |
| {{< ext "pointcloud" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "GIS" >}} | 提供激光雷达点云数据类型支持 |
| {{< ext "pointcloud_postgis" "pointcloud" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "GIS" >}} | 将激光雷达点云与PostGIS几何类型相集成 |
| {{< ext "h3" "pg_h3" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "GIS" >}} | H3六边形层级索引支持 |
| {{< ext "h3_postgis" "pg_h3" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "GIS" >}} | H3与PostGIS集成的扩展插件 |
| {{< ext "q3c" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "GIS" >}} | Q3C天空索引插件 |
| {{< ext "ogr_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "GIS" >}} | GIS 数据外部数据源包装器 |
| {{< ext "geoip" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "GIS" >}} | IP 地理位置扩展（围绕 MaxMind GeoLite 数据集的包装器） |
| {{< ext "pg_polyline" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "GIS" >}} | Google快速Polyline编码解码扩展 |
| {{< ext "pg_geohash" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "GIS" >}} | 使用GeoHash处理空间坐标的函数包 |
| {{< ext "mobilitydb" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "GIS" >}} | MobilityDB地理空间投影数据管理分析平台 |
| {{< ext "mobilitydb_datagen" "mobilitydb" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | {{< category "GIS" >}} | MobilityDB随机数据生成函数 |
| {{< ext "tzf" "pg_tzf" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "GIS" >}} | 快速根据GPS经纬度坐标查找时区 |
| {{< ext "earthdistance" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "GIS" >}} | 计算地球表面上的大圆距离 |
| {{< ext "vector" "pgvector" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "RAG" >}} | 向量数据类型和 ivfflat / hnsw 访问方法 |
| {{< ext "vchord" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--sLd-r" color="blue" >}} | {{< category "RAG" >}} | 使用Rust重写的高性能向量扩展 |
| {{< ext "vectorscale" "pgvectorscale" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "RAG" >}} | 使用DiskANN算法对向量进行高效索引 |
| {{< ext "vectorize" "pg_vectorize" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "RAG" >}} | 在PostgreSQL中封装RAG向量检索服务 |
| {{< ext "pg_similarity" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "RAG" >}} | 提供17种距离度量函数 |
| {{< ext "smlar" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "RAG" >}} | 高效的相似度搜索函数 |
| {{< ext "pg_summarize" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "RAG" >}} | 使用LLM对文本字段进行总结 |
| {{< ext "pg_tiktoken" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "RAG" >}} | 在PostgreSQL中计算OpenAI使用的Token数 |
| {{< ext "pg4ml" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dtr" color="blue" >}} | {{< category "RAG" >}} | PG4ML是一个机器学习框架 |
| {{< ext "pgml" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,r" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "RAG" >}} | PostgresML：用SQL运行机器学习算法并训练模型 |
| {{< ext "pg_search" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "FTS" >}} | ParadeDB BM25算法全文检索插件，ES全文检索 |
| {{< ext "pgroonga" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dtr" color="blue" >}} | {{< category "FTS" >}} | 使用Groonga，面向所有语言的高速全文检索平台 |
| {{< ext "pgroonga_database" "pgroonga" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dtr" color="blue" >}} | {{< category "FTS" >}} | PGGroonga 数据库管理模块 |
| {{< ext "pg_bigm" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FTS" >}} | 基于二字组的多语言全文检索扩展 |
| {{< ext "zhparser" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FTS" >}} | 中文分词，全文搜索解析器 |
| {{< ext "pg_bestmatch" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "FTS" >}} | 在数据库内生成BM25稀疏向量 |
| {{< ext "vchord_bm25" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "FTS" >}} | BM25排序算法 |
| {{< ext "pg_tokenizer" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "FTS" >}} | 用于全文检索的分词器 |
| {{< ext "biscuit" "pg_biscuit" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,r,r,r" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FTS" >}} | 使用IAM的高性能文本模式匹配 |
| {{< ext "pg_textsearch" >}} | {{< pgver "18,17,16,15,14,13" "g,g,r,r,r,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FTS" >}} | 带有BM25排序的全文搜索扩展 |
| {{< ext "hunspell_cs_cz" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | {{< category "FTS" >}} | Hunspell捷克语全文检索词典 |
| {{< ext "hunspell_de_de" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | {{< category "FTS" >}} | Hunspell德语全文检索词典 |
| {{< ext "hunspell_en_us" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | {{< category "FTS" >}} | Hunspell英语全文检索词典 |
| {{< ext "hunspell_fr" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | {{< category "FTS" >}} | Hunspell法语全文检索词典 |
| {{< ext "hunspell_ne_np" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | {{< category "FTS" >}} | Hunspell尼泊尔语全文检索词典 |
| {{< ext "hunspell_nl_nl" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | {{< category "FTS" >}} | Hunspell荷兰语全文检索词典 |
| {{< ext "hunspell_nn_no" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | {{< category "FTS" >}} | Hunspell挪威语全文检索词典 |
| {{< ext "hunspell_pt_pt" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | {{< category "FTS" >}} | Hunspell葡萄牙语全文检索词典 |
| {{< ext "hunspell_ru_ru" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | {{< category "FTS" >}} | Hunspell俄语全文检索词典 |
| {{< ext "hunspell_ru_ru_aot" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | {{< category "FTS" >}} | Hunspell俄语全文检索词典（来自AOT.ru小组） |
| {{< ext "fuzzystrmatch" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | {{< category "FTS" >}} | 确定字符串之间的相似性和距离 |
| {{< ext "pg_trgm" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "FTS" >}} | 文本相似度测量函数与模糊检索 |
| {{< ext "citus" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,r,r,r" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "OLAP" >}} | Citus 分布式数据库 |
| {{< ext "citus_columnar" "citus" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,r,r,r" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "OLAP" >}} | Citus 列式存储引擎 |
| {{< ext "columnar" "hydra" >}} | {{< pgver "18,17,16,15,14,13" "r,r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "OLAP" >}} | 开源列式存储扩展 |
| {{< ext "pg_analytics" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "OLAP" >}} | 由 DuckDB 驱动的数据分析引擎 |
| {{< ext "pg_duckdb" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "OLAP" >}} | 在PostgreSQL中的嵌入式DuckDB扩展 |
| {{< ext "pg_mooncake" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="---Ld--" color="blue" >}} | {{< category "OLAP" >}} | PostgreSQL列式存储表 |
| {{< ext "pg_clickhouse" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "OLAP" >}} | 从PostgreSQL中查询ClickHouse的接口 |
| {{< ext "duckdb_fdw" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "OLAP" >}} | DuckDB 外部数据源包装器 |
| {{< ext "pg_parquet" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--sLdt-" color="blue" >}} | {{< category "OLAP" >}} | 在PostgreSQL与本地/S3中的Parquet文件复制数据 |
| {{< ext "pg_fkpart" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "OLAP" >}} | 按外键实用程序进行表分区的扩展 |
| {{< ext "pg_partman" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "OLAP" >}} | 用于按时间或 ID 管理分区表的扩展 |
| {{< ext "plproxy" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "OLAP" >}} | 作为过程语言实现的数据库分区 |
| {{< ext "pg_strom" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "OLAP" >}} | 使用GPU与NVMe加速大数据处理 |
| {{< ext "tablefunc" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | {{< category "OLAP" >}} | 交叉表函数 |
| {{< ext "age" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Apache AGE，图数据库扩展 （Deb可用） |
| {{< ext "hll" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | hyperloglog 数据类型 |
| {{< ext "rum" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FEAT" >}} | RUM 索引访问方法 |
| {{< ext "pg_ai_query" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | AI驱动的 Postgres SQL 查询生成 |
| {{< ext "pg_ttl_index" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="--sLd-r" color="blue" >}} | {{< category "FEAT" >}} | 基于TTL索引的自动数据过期清理 |
| {{< ext "pg_graphql" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | PG内的GraphQL支持 |
| {{< ext "pg_jsonschema" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "FEAT" >}} | 提供JSON Schema校验能力 |
| {{< ext "jsquery" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FEAT" >}} | 用于内省 JSONB 数据类型的查询类型 |
| {{< ext "pg_hint_plan" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | 添加强制指定执行计划的能力 |
| {{< ext "hypopg" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FEAT" >}} | 假设索引，用于创建一个虚拟索引检验执行计划 |
| {{< ext "index_advisor" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | {{< category "FEAT" >}} | 查询索引建议器 |
| {{< ext "plan_filter" "pg_plan_filter" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | {{< category "FEAT" >}} | 使用执行计划代价过滤阻止特定查询语句 |
| {{< ext "imgsmlr" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FEAT" >}} | 使用Haar小波分析计算图片相似度 |
| {{< ext "pg_ivm" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | 增量维护的物化视图 |
| {{< ext "pg_incremental" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,r,r,r" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | 增量处理流式事件 |
| {{< ext "pgmb" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "FEAT" >}} | 一个简单的PostgreSQL消息代理系统 |
| {{< ext "pgmq" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "FEAT" >}} | 基于Postgres实现类似AWS SQS/RSMQ的消息队列 |
| {{< ext "pgq" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | 通用队列的PG实现 |
| {{< ext "orioledb" >}} | {{< pgver "18,17,16,15,14,13" "r,g,r,r,r,r" >}} | {{< badge content="--sLd-r" color="blue" >}} | {{< category "FEAT" >}} | OrioleDB，下一代事务处理引擎 |
| {{< ext "pg_cardano" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Cardano相关工具包：加密函数，地址编解码，区块链处理 |
| {{< ext "rdkit" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FEAT" >}} | 在PostgreSQL化学领域数据管理功能 |
| {{< ext "omni" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "FEAT" >}} | PostgreSQL即平台，Omnigres主扩展与加载器 |
| {{< ext "omni_auth" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "FEAT" >}} | Omnigres 基础会话认证管理模块 |
| {{< ext "omni_aws" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dt-" color="blue" >}} | {{< category "FEAT" >}} | Omnigres AWS S3 API封装 |
| {{< ext "omni_cloudevents" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dt-" color="blue" >}} | {{< category "FEAT" >}} | Omnigres CloudEvents 支持 |
| {{< ext "omni_containers" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Omnigres Docker容器管理模块 |
| {{< ext "omni_credentials" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "FEAT" >}} | Omnigres 应用密钥管理模块 |
| {{< ext "omni_csv" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Omnigres CSV 工具箱 |
| {{< ext "omni_datasets" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "FEAT" >}} | Omnigres 数据库置备工具 |
| {{< ext "omni_email" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "FEAT" >}} | Omnigres Email 框架 |
| {{< ext "omni_http" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "FEAT" >}} | Omnigres 基本HTTP类型 |
| {{< ext "omni_httpc" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Omnigres HTTP客户端 |
| {{< ext "omni_httpd" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Omnigres HTTP服务器 |
| {{< ext "omni_id" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FEAT" >}} | Omnigres ID身份数据类型 |
| {{< ext "omni_json" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dt-" color="blue" >}} | {{< category "FEAT" >}} | Omnigres JSON工具箱 |
| {{< ext "omni_kube" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Omnigres Kubernetes集成模块 |
| {{< ext "omni_ledger" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Omnigres 金融账本模块 |
| {{< ext "omni_manifest" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "FEAT" >}} | Omnigres 包管理清单模块 |
| {{< ext "omni_mimetypes" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "FEAT" >}} | Omnigres MIME数据类型 |
| {{< ext "omni_os" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Omnigres 操作系统集成模块 |
| {{< ext "omni_polyfill" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Omnigres Postgres多态API |
| {{< ext "omni_python" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Omnigres 第一类Python支持模块 |
| {{< ext "omni_regex" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FEAT" >}} | Omnigres PCRE兼容正则表达式模块 |
| {{< ext "omni_rest" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "FEAT" >}} | Omnigres REST API 工具包 |
| {{< ext "omni_schema" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "FEAT" >}} | Omnigres 高级模式管理组件 |
| {{< ext "omni_seq" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Omnigres 分布式整型序列号 |
| {{< ext "omni_service" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "FEAT" >}} | Omnigres 服务管理器 |
| {{< ext "omni_session" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "FEAT" >}} | Omnigres 会话管理器 |
| {{< ext "omni_shmem" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Omnigres 共享内存管理 |
| {{< ext "omni_sql" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Omnigres SQL编程组件 |
| {{< ext "omni_sqlite" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Omnigres 嵌入的SQLite支持 |
| {{< ext "omni_test" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "FEAT" >}} | Omnigres 测试框架 |
| {{< ext "omni_txn" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Omnigres 事务管理器模块 |
| {{< ext "omni_types" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Omnigres 高级数据类型模块 |
| {{< ext "omni_var" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Omnigres 局部变量模块 |
| {{< ext "omni_vfs" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Omnigres 虚拟文件系统 |
| {{< ext "omni_vfs_types_v1" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "FEAT" >}} | Omnigres 虚拟文件系统（v1） |
| {{< ext "omni_web" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Omnigres Web工具箱 |
| {{< ext "omni_worker" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Omnigres 通用Worker池 |
| {{< ext "omni_xml" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Omnigres XML工具包 |
| {{< ext "omni_yaml" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Omnigres YAML工具包 |
| {{< ext "bloom" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "FEAT" >}} | bloom 索引-基于指纹的索引 |
| {{< ext "pg_tle" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "LANG" >}} | AWS 可信语言扩展 |
| {{< ext "plv8" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "LANG" >}} | PL/JavaScript (v8) 可信过程程序语言 |
| {{< ext "pljs" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "LANG" >}} | PL/JS 可信过程程序语言 |
| {{< ext "pllua" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "LANG" >}} | Lua 程序语言 |
| {{< ext "hstore_pllua" "pllua" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "LANG" >}} | Lua 程序语言的Hstore适配扩展 |
| {{< ext "plluau" "pllua" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "LANG" >}} | Lua 程序语言（不受信任的） |
| {{< ext "hstore_plluau" "pllua" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "LANG" >}} | Lua 程序语言的Hstore适配扩展（不受信任的） |
| {{< ext "plprql" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "LANG" >}} | 在PostgreSQL使用PRQL——管线式关系查询语言 |
| {{< ext "pldbgapi" "pldebugger" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "LANG" >}} | 用于调试 PL/pgSQL 函数的服务器端支持 |
| {{< ext "plpgsql_check" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "LANG" >}} | 对 plpgsql 函数进行扩展检查 |
| {{< ext "plprofiler" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "LANG" >}} | 剖析 PL/pgSQL 函数 |
| {{< ext "plsh" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "LANG" >}} | PL/sh 程序语言 |
| {{< ext "pljava" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "LANG" >}} | Java 程序语言 |
| {{< ext "plr" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "LANG" >}} | 从数据库中加载R语言解释器并执行R脚本 |
| {{< ext "plxslt" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "LANG" >}} | XSLT 存储过程语言 |
| {{< ext "pgtap" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "LANG" >}} | PostgreSQL单元测试框架 |
| {{< ext "faker" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "LANG" >}} | 插入生成的测试伪造数据，Python库的包装 |
| {{< ext "dbt2" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "LANG" >}} | OSDL-DBT-2 测试组件 |
| {{< ext "pltcl" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "LANG" >}} | PL/TCL 存储过程语言 |
| {{< ext "pltclu" "pltcl" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---d--" color="blue" >}} | {{< category "LANG" >}} | PL/TCL 存储过程语言（未受信/高权限） |
| {{< ext "plperl" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "LANG" >}} | PL/Perl 存储过程语言 |
| {{< ext "bool_plperl" "plperl" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "LANG" >}} | 在 bool 和 plperl 之间转换 |
| {{< ext "hstore_plperl" "plperl" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "LANG" >}} | 在 hstore 和 plperl 之间转换适配类型 |
| {{< ext "jsonb_plperl" "plperl" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---d--" color="blue" >}} | {{< category "LANG" >}} | 在 jsonb 和 plperl 之间转换 |
| {{< ext "plperlu" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "LANG" >}} | PL/PerlU 存储过程语言（未受信/高权限） |
| {{< ext "bool_plperlu" "plperlu" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---d--" color="blue" >}} | {{< category "LANG" >}} | 在 bool 和 plperlu 之间转换 |
| {{< ext "jsonb_plperlu" "plperlu" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---d--" color="blue" >}} | {{< category "LANG" >}} | 在 jsonb 和 plperlu 之间转换 |
| {{< ext "hstore_plperlu" "plperlu" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---d--" color="blue" >}} | {{< category "LANG" >}} | 在 hstore 和 plperlu 之间转换适配类型 |
| {{< ext "plpgsql" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "LANG" >}} | PL/pgSQL 程序设计语言 |
| {{< ext "plpython3u" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "LANG" >}} | PL/Python3 存储过程语言（未受信/高权限） |
| {{< ext "jsonb_plpython3u" "plpython3u" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---d-r" color="blue" >}} | {{< category "LANG" >}} | 在 jsonb 和 plpython3u 之间转换 |
| {{< ext "ltree_plpython3u" "plpython3u" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d-r" color="blue" >}} | {{< category "LANG" >}} | 在 ltree 和 plpython3u 之间转换 |
| {{< ext "hstore_plpython3u" "plpython3u" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---d-r" color="blue" >}} | {{< category "LANG" >}} | 在 hstore 和 plpython3u 之间转换 |
| {{< ext "prefix" "pg_prefix" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | 前缀树数据类型 |
| {{< ext "semver" "pg_semver" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | 语义版本号数据类型 |
| {{< ext "unit" "pgunit" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "TYPE" >}} | SI 国标单位扩展 |
| {{< ext "pgpdf" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLdtr" color="blue" >}} | {{< category "TYPE" >}} | PDF数据类型，管理函数与全文检索 |
| {{< ext "pglite_fusion" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "TYPE" >}} | 在PG表中嵌入SQLite数据库作为数据类型 |
| {{< ext "md5hash" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | 提供128位MD5的原生数据类型 |
| {{< ext "asn1oid" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | ASN1OID数据类型支持 |
| {{< ext "roaringbitmap" "pg_roaringbitmap" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | 支持RoaringBitmap数据类型 |
| {{< ext "pgfaceting" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dt-" color="blue" >}} | {{< category "TYPE" >}} | 使用倒排索引的高速切面查询 |
| {{< ext "pg_sphere" "pgsphere" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | 球面对象函数、运算符与索引支持 |
| {{< ext "country" "pg_country" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | 国家代码数据类型，遵循ISO 3166-1标准 |
| {{< ext "pg_xenophile" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "TYPE" >}} | PostgreSQL i8n与l10n工具包 |
| {{< ext "l10n_table_dependent_extension" "pg_xenophile" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dtr" color="blue" >}} | {{< category "TYPE" >}} | PostgreSQL l10n 工具包 |
| {{< ext "currency" "pg_currency" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | 使用1字节表示的货币数据类型 |
| {{< ext "collection" "pgcollection" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | 在PlPGSQL中使用的内存优化高性能集合数据结构 |
| {{< ext "pgmp" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | 多精度算术扩展 |
| {{< ext "numeral" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | 数值类型扩展 |
| {{< ext "pg_rational" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "TYPE" >}} | 使用BIGINT表示的有理数数据类型 |
| {{< ext "uint" "pguint" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | 无符号整型数据类型 |
| {{< ext "uint128" "pg_uint128" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | 原生128位无符号整型数据类型 |
| {{< ext "hashtypes" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "TYPE" >}} | 包括SHA1，MD5在内的多种哈希数据类型 |
| {{< ext "ip4r" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | PostgreSQL 的 IPv4/v6 和 IPv4/v6 范围索引类型 |
| {{< ext "pg_duration" >}} | {{< pgver "18,17,16,15,14,13" "g,g,r,r,r,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | 用于表示时间段的强化数据类型 |
| {{< ext "uri" "pg_uri" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | URI数据类型 |
| {{< ext "emailaddr" "pg_emailaddr" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "TYPE" >}} | Email地址数据类型 |
| {{< ext "acl" "pg_acl" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | ACL数据类型 |
| {{< ext "debversion" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | Debian版本号数据类型 |
| {{< ext "pg_rrule" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | 日历重复规则RRULE数据类型 |
| {{< ext "timestamp9" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "TYPE" >}} | 纳秒分辨率时间戳 |
| {{< ext "chkpass" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "TYPE" >}} | 数据类型：自动加密的密码 |
| {{< ext "isn" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | {{< category "TYPE" >}} | 用于国际产品编号标准的数据类型 |
| {{< ext "seg" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | {{< category "TYPE" >}} | 表示线段或浮点间隔的数据类型 |
| {{< ext "cube" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "TYPE" >}} | 用于存储多维立方体的数据类型 |
| {{< ext "ltree" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | {{< category "TYPE" >}} | 用于表示分层树状结构的数据类型 |
| {{< ext "hstore" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | {{< category "TYPE" >}} | 用于存储（键，值）对集合的数据类型 |
| {{< ext "citext" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | {{< category "TYPE" >}} | 提供大小写不敏感的字符串类型 |
| {{< ext "xml2" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---d--" color="blue" >}} | {{< category "TYPE" >}} | XPath 查询和 XSLT |
| {{< ext "gzip" "pg_gzip" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "UTIL" >}} | 使用SQL执行Gzip压缩与解压缩 |
| {{< ext "bzip" "pg_bzip" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "UTIL" >}} | BZIP压缩解压缩函数包 |
| {{< ext "zstd" "pg_zstd" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "UTIL" >}} | ZSTD压缩解压缩函数包 |
| {{< ext "http" "pg_http" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "UTIL" >}} | HTTP客户端，允许在数据库内收发HTTP请求 (supabase) |
| {{< ext "pg_net" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "UTIL" >}} | 用 SQL 进行异步非阻塞HTTP/HTTPS 请求的扩展 (supabase) |
| {{< ext "pg_curl" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "UTIL" >}} | 封装CURL，执行各种用URL传输数据的操作 |
| {{< ext "pg_retry" >}} | {{< pgver "18,17,16,15,14,13" "g,g,r,r,r,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "UTIL" >}} | 在临时错误中使用指数退避重试语句 |
| {{< ext "pgjq" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-dtr" color="blue" >}} | {{< category "UTIL" >}} | 在Postgres中使用jq查询JSON |
| {{< ext "pgjwt" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "UTIL" >}} | JSON Web Token API 的PG实现 (supabase) |
| {{< ext "pg_smtp_client" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "UTIL" >}} | 使用SMTP从PostgreSQL内发送邮件的客户端扩展 |
| {{< ext "pg_html5_email_address" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | {{< category "UTIL" >}} | 验证Email是否符合HTML5规范的扩展 |
| {{< ext "url_encode" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "UTIL" >}} | 提供URL编码解码函数 |
| {{< ext "pgsql_tweaks" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "UTIL" >}} | 一些日常会用到的便利函数与视图 |
| {{< ext "pg_extra_time" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dtr" color="blue" >}} | {{< category "UTIL" >}} | 一些关于日期与时间的扩展函数 |
| {{< ext "pgpcre" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "UTIL" >}} | PCRE/Perl风格的正则表达式支持 |
| {{< ext "icu_ext" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "UTIL" >}} | 访问ICU库提供的函数 |
| {{< ext "pgqr" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "UTIL" >}} | 从数据库中直接生成QR二维码 |
| {{< ext "pg_protobuf" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "UTIL" >}} | 提供Protobuf函数支持 |
| {{< ext "envvar" "pg_envvar" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "UTIL" >}} | 获取环境变量的函数 |
| {{< ext "floatfile" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "UTIL" >}} | 将浮点数组存储到文件中而不是堆表中 |
| {{< ext "pg_render" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "UTIL" >}} | 使用SQL渲染HTML页面 |
| {{< ext "pg_readme" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dtr" color="blue" >}} | {{< category "UTIL" >}} | 为模式与扩展生成Markdown文档 |
| {{< ext "pg_readme_test_extension" "pg_readme" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dtr" color="blue" >}} | {{< category "UTIL" >}} | 为模式与扩展生成Markdown文档 |
| {{< ext "ddl_historization" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "UTIL" >}} | 用SQL将所有DDL变更写入到数据库表中 |
| {{< ext "data_historization" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dt-" color="blue" >}} | {{< category "UTIL" >}} | 用SQL将数据变更历史保存到分区表中 |
| {{< ext "schedoc" "pg_schedoc" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "UTIL" >}} | 在Django与DBT之间通过注释文档交换元数据 |
| {{< ext "hashlib" "pg_hashlib" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dtr" color="blue" >}} | {{< category "UTIL" >}} | 稳定哈希函数包 |
| {{< ext "xxhash" "pg_xxhash" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dtr" color="blue" >}} | {{< category "UTIL" >}} | xxhash哈希函数包 |
| {{< ext "shacrypt" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "UTIL" >}} | 实现SHA256-CRYPT与SHA512-CRYPT密钥加密算法 |
| {{< ext "cryptint" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "UTIL" >}} | 加密INT与BIGINT类型 |
| {{< ext "pguecc" "pg_ecdsa" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "UTIL" >}} | PostgreSQL的uECC绑定，椭圆曲线加解密函数包 |
| {{< ext "sparql" "pgsparql" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "UTIL" >}} | 使用SQL查询SPARQL数据源 |
| {{< ext "pg_idkit" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FUNC" >}} | 生成各式各样的唯一标识符：UUIDv6, ULID, KSUID |
| {{< ext "pgx_ulid" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "FUNC" >}} | ULID数据类型与函数 |
| {{< ext "pg_uuidv7" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FUNC" >}} | UUIDv7 支持 |
| {{< ext "permuteseq" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FUNC" >}} | 伪随机数ID置换生成器 |
| {{< ext "pg_hashids" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FUNC" >}} | 加盐将整型ID转为短字符串ID |
| {{< ext "sequential_uuids" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FUNC" >}} | 生成连续生成的UUID |
| {{< ext "typeid" "pg_typeid" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FUNC" >}} | PG原生TypeID类型与函数 |
| {{< ext "snowflake" >}} | {{< pgver "18,17,16,15,14,13" "r,g,r,r,r,r" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FUNC" >}} | Snowflake 风格 64 位 ID 生成与序列工具 |
| {{< ext "topn" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FUNC" >}} | top-n JSONB 的类型 |
| {{< ext "quantile" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FUNC" >}} | Quantile聚合函数 |
| {{< ext "lower_quantile" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FUNC" >}} | Lower Quantile 聚合函数 |
| {{< ext "count_distinct" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FUNC" >}} | COUNT(DISTINCT …) 聚合的替代方案 |
| {{< ext "omnisketch" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FUNC" >}} | 实现OmniSketch数据结构，实现近似摘要聚合 |
| {{< ext "ddsketch" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FUNC" >}} | 实现DDSketch数据结构，实现在线的Quantile聚合 |
| {{< ext "vasco" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FUNC" >}} | 使用MIC发现数据中隐含的关联 |
| {{< ext "xicor" "pgxicor" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dtr" color="blue" >}} | {{< category "FUNC" >}} | 在PG中计算XI相关系数 |
| {{< ext "weighted_statistics" "pg_weighted_statistics" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dtr" color="blue" >}} | {{< category "FUNC" >}} | 针对稀疏数据的高性能加权统计量计算 |
| {{< ext "tdigest" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FUNC" >}} | tdigest 聚合函数 |
| {{< ext "first_last_agg" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FUNC" >}} | first() 与 last() 聚合函数 |
| {{< ext "extra_window_functions" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FUNC" >}} | 额外的窗口函数 |
| {{< ext "floatvec" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FUNC" >}} | 数组类型数学运算扩展 |
| {{< ext "aggs_for_vecs" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FUNC" >}} | 针对数组类型的聚合函数集合扩展 |
| {{< ext "aggs_for_arrays" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FUNC" >}} | 计算数组聚合统计值的函数包 |
| {{< ext "pg_csv" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-dtr" color="blue" >}} | {{< category "FUNC" >}} | 灵活的CSV聚合处理函数 |
| {{< ext "arraymath" "pg_arraymath" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FUNC" >}} | 数组逐元素数学运算符包 |
| {{< ext "pg_math" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FUNC" >}} | 使用GSL库的数学统计函数 |
| {{< ext "random" "pg_random" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FUNC" >}} | 随机数生成器 |
| {{< ext "base36" "pg_base36" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FUNC" >}} | Base36编码解码扩展 |
| {{< ext "base62" "pg_base62" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FUNC" >}} | Base62编码解码扩展 |
| {{< ext "pg_base58" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FUNC" >}} | Base58 编码/解码函数 |
| {{< ext "financial" "pg_financial" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FUNC" >}} | 金融领域聚合函数 |
| {{< ext "convert" "pg_convert" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "FUNC" >}} | 用于空间里程等的公英制转换函数 |
| {{< ext "refint" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "FUNC" >}} | 实现引用完整性的函数 |
| {{< ext "autoinc" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "FUNC" >}} | 用于自动递增字段的函数 |
| {{< ext "insert_username" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "FUNC" >}} | 用于跟踪谁更改了表的函数 |
| {{< ext "moddatetime" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "FUNC" >}} | 跟踪最后修改时间 |
| {{< ext "tsm_system_time" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | {{< category "FUNC" >}} | 接受毫秒数限制的 TABLESAMPLE 方法 |
| {{< ext "dict_xsyn" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "FUNC" >}} | 用于扩展同义词处理的文本搜索字典模板 |
| {{< ext "tsm_system_rows" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | {{< category "FUNC" >}} | 接受行数限制的 TABLESAMPLE 方法 |
| {{< ext "tcn" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | {{< category "FUNC" >}} | 用触发器通知变更 |
| {{< ext "uuid-ossp" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | {{< category "FUNC" >}} | 生成通用唯一标识符（UUIDs） |
| {{< ext "btree_gist" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | {{< category "FUNC" >}} | 用GiST索引常见数据类型 |
| {{< ext "btree_gin" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | {{< category "FUNC" >}} | 用GIN索引常见数据类型 |
| {{< ext "intarray" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---dt-" color="blue" >}} | {{< category "FUNC" >}} | 1维整数数组的额外函数、运算符和索引支持 |
| {{< ext "intagg" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---d--" color="blue" >}} | {{< category "FUNC" >}} | 整数聚合器和枚举器（过时） |
| {{< ext "dict_int" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | {{< category "FUNC" >}} | 用于整数的文本搜索字典模板 |
| {{< ext "unaccent" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | {{< category "FUNC" >}} | 删除重音的文本搜索字典 |
| {{< ext "pg_repack" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="-bs-d--" color="blue" >}} | {{< category "ADMIN" >}} | 在线垃圾清理与表膨胀治理 |
| {{< ext "pg_rewrite" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | {{< category "ADMIN" >}} | 在线重写整表，不阻塞读写 |
| {{< ext "pg_squeeze" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "ADMIN" >}} | 从关系中删除未使用空间 |
| {{< ext "pg_dirtyread" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "ADMIN" >}} | 从表中读取尚未垃圾回收的行 |
| {{< ext "pgfincore" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "ADMIN" >}} | 检查和管理操作系统缓冲区缓存 |
| {{< ext "pg_cooldown" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "ADMIN" >}} | 从缓冲区中移除特定关系的页面 |
| {{< ext "ddlx" "pg_ddlx" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "ADMIN" >}} | 提取数据库对象的DDL |
| {{< ext "pglinter" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "ADMIN" >}} | PG数据库规则检查插件 |
| {{< ext "prioritize" "pg_prioritize" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "ADMIN" >}} | 获取和设置 PostgreSQL 后端的优先级 |
| {{< ext "pg_checksums" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s---r" color="blue" >}} | {{< category "ADMIN" >}} | 在离线模式下激活/启用/禁用数据库集群的校验和功能 |
| {{< ext "pg_readonly" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "ADMIN" >}} | 将集群设置为只读 |
| {{< ext "pgdd" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "ADMIN" >}} | 提供通过标准SQL查询数据库目录集簇的能力 |
| {{< ext "pg_permissions" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "ADMIN" >}} | 查看对象权限并将其与期望状态进行比较 |
| {{< ext "pgautofailover" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "ADMIN" >}} | PG 自动故障迁移 |
| {{< ext "pg_catcheck" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "ADMIN" >}} | 用于诊断系统目录是否损坏的工具 |
| {{< ext "pre_prepare" "preprepare" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "ADMIN" >}} | 在服务端预先准备好PreparedStatement备用 |
| {{< ext "pg_upless" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dt-" color="blue" >}} | {{< category "ADMIN" >}} | 检测表上的无用UPDATE |
| {{< ext "pgcozy" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "ADMIN" >}} | 根据先前的pg_buffercache快照预热内存缓冲区 |
| {{< ext "pg_orphaned" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "ADMIN" >}} | 处理孤儿文件的扩展插件 |
| {{< ext "pg_crash" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | {{< category "ADMIN" >}} | 向数据库进程随机发送信号模拟故障 |
| {{< ext "pg_cheat_funcs" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "ADMIN" >}} | 一些超级实用的作弊函数 |
| {{< ext "fio" "pg_fio" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "ADMIN" >}} | PostgreSQL文件IO函数包 |
| {{< ext "pg_savior" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "ADMIN" >}} | 阻止不带条件的全表更新以避免意外事故 |
| {{< ext "safeupdate" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--sL---" color="blue" >}} | {{< category "ADMIN" >}} | 强制在 UPDATE 和 DELETE 时提供 Where 条件 |
| {{< ext "pg_strict" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "ADMIN" >}} | 防止不带WHERE条件的危险UPDATE和DELETE操作 |
| {{< ext "pg_drop_events" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "ADMIN" >}} | 记录删表删列删视图的事务号，辅助PITR确定时间点 |
| {{< ext "table_log" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "ADMIN" >}} | 记录某张表的修改日志并做表/行级时间点恢复 |
| {{< ext "pgagent" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "ADMIN" >}} | PostgreSQL任务调度工具，与PGADMIN配合使用 |
| {{< ext "pg_prewarm" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-sLd--" color="blue" >}} | {{< category "ADMIN" >}} | 预热关系数据 |
| {{< ext "pgpool_adm" "pgpool" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "ADMIN" >}} | PGPool 管理函数 |
| {{< ext "pgpool_recovery" "pgpool" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "ADMIN" >}} | PGPool辅助扩展，从v4.3提供的恢复函数 |
| {{< ext "pgpool_regclass" "pgpool" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "ADMIN" >}} | PGPool辅助扩展，RegClass替代 |
| {{< ext "lo" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | {{< category "ADMIN" >}} | 大对象维护 |
| {{< ext "basic_archive" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="c-s----" color="blue" >}} | {{< category "ADMIN" >}} | 归档模块样例 |
| {{< ext "basebackup_to_shell" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="c-s----" color="blue" >}} | {{< category "ADMIN" >}} | 添加一种备份到Shell终端到基础备份方式 |
| {{< ext "old_snapshot" >}} | {{< pgver "18,17,16,15,14,13" "r,r,g,g,g,r" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "ADMIN" >}} | 支持 old_snapshot_threshold 的实用程序 |
| {{< ext "adminpack" >}} | {{< pgver "18,17,16,15,14,13" "r,r,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "ADMIN" >}} | PostgreSQL 管理函数集合 |
| {{< ext "amcheck" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "ADMIN" >}} | 校验关系完整性 |
| {{< ext "pg_surgery" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "ADMIN" >}} | 对损坏的关系进行手术 |
| {{< ext "pg_profile" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "STAT" >}} | PostgreSQL 数据库负载记录与AWR报表工具 |
| {{< ext "pg_tracing" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--sLd-r" color="blue" >}} | {{< category "STAT" >}} | PostgreSQL分布式Tracing |
| {{< ext "pg_show_plans" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "STAT" >}} | 打印所有当前正在运行查询的执行计划 |
| {{< ext "pg_stat_kcache" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | {{< category "STAT" >}} | 内核统计信息收集 |
| {{< ext "pg_stat_monitor" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | {{< category "STAT" >}} | 提供查询聚合统计、客户端信息、执行计划详细信息和直方图 |
| {{< ext "pg_qualstats" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "STAT" >}} | 收集有关 quals 的统计信息的扩展 |
| {{< ext "pg_store_plans" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | {{< category "STAT" >}} | 跟踪所有执行的 SQL 语句的计划统计信息 |
| {{< ext "pg_track_settings" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "STAT" >}} | 跟踪设置更改 |
| {{< ext "pg_track_optimizer" >}} | {{< pgver "18,17,16,15,14,13" "g,g,r,r,r,r" >}} | {{< badge content="--sLd-r" color="blue" >}} | {{< category "STAT" >}} | 跟踪规划器决策与实际执行的差距 |
| {{< ext "pg_wait_sampling" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | {{< category "STAT" >}} | 基于采样的等待事件统计 |
| {{< ext "pgsentinel" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | {{< category "STAT" >}} | 活跃会话历史 |
| {{< ext "system_stats" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "STAT" >}} | PostgreSQL 的系统统计函数 |
| {{< ext "meta" "pg_meta" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | {{< category "STAT" >}} | 标准化，更友好的PostgreSQL系统目录视图 |
| {{< ext "pgnodemx" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "STAT" >}} | 使用SQL查询获取操作系统指标 |
| {{< ext "pg_proctab" "pgnodemx" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "STAT" >}} | 通过SQL接口访问操作系统进程表 |
| {{< ext "pg_sqlog" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "STAT" >}} | 提供访问PostgreSQL日志的SQL接口 |
| {{< ext "bgw_replstatus" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | {{< category "STAT" >}} | 用于汇报本机主从状态的后台工作进程 |
| {{< ext "pgmeminfo" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "STAT" >}} | 显示内存使用情况 |
| {{< ext "toastinfo" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "STAT" >}} | 显示TOAST字段的详细信息 |
| {{< ext "explain_ui" "pg_explain_ui" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "STAT" >}} | 快速跳转至PEV查阅可视化执行计划 |
| {{< ext "pg_relusage" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | {{< category "STAT" >}} | 打印查询引用的表与列 |
| {{< ext "pagevis" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | {{< category "STAT" >}} | 使用ASCII字符可视化数据库物理页面布局 |
| {{< ext "powa" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "STAT" >}} | PostgreSQL 工作负载分析器-核心 |
| {{< ext "pg_overexplain" >}} | {{< pgver "18,17,16,15,14,13" "g,r,r,r,r,r" >}} | {{< badge content="c-sL---" color="blue" >}} | {{< category "STAT" >}} | 允许 EXPLAIN 转储更多详细 |
| {{< ext "pg_logicalinspect" >}} | {{< pgver "18,17,16,15,14,13" "g,r,r,r,r,r" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "STAT" >}} | 检视逻辑解码组件详情 |
| {{< ext "pageinspect" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "STAT" >}} | 检查数据库页面二进制内容 |
| {{< ext "pgrowlocks" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "STAT" >}} | 显示行级锁信息 |
| {{< ext "sslinfo" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "STAT" >}} | 关于 SSL 证书的信息 |
| {{< ext "pg_buffercache" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "STAT" >}} | 检查共享缓冲区缓存 |
| {{< ext "pg_walinspect" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "STAT" >}} | 用于检查 PostgreSQL WAL 日志内容的函数 |
| {{< ext "pg_freespacemap" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "STAT" >}} | 检查自由空间映射的内容（FSM） |
| {{< ext "pg_visibility" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "STAT" >}} | 检查可见性图（VM）和页面级可见性信息 |
| {{< ext "pgstattuple" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "STAT" >}} | 显示元组级统计信息 |
| {{< ext "auto_explain" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-sL---" color="blue" >}} | {{< category "STAT" >}} | 提供一种自动记录执行计划的手段 |
| {{< ext "pg_stat_statements" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-sLd--" color="blue" >}} | {{< category "STAT" >}} | 跟踪所有执行的 SQL 语句的计划和执行统计信息 |
| {{< ext "passwordcheck_cracklib" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | {{< category "SEC" >}} | 使用cracklib加固PG用户密码 |
| {{< ext "supautils" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | {{< category "SEC" >}} | 用于在云环境中确保数据库集群的安全 |
| {{< ext "pgsodium" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "SEC" >}} | 表数据加密存储 TDE |
| {{< ext "supabase_vault" "pg_vault" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "SEC" >}} | 在 Vault 中存储加密凭证的扩展 (supabase) |
| {{< ext "pg_session_jwt" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "SEC" >}} | 使用JWT进行会话认证 |
| {{< ext "anon" "pg_anon" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "SEC" >}} | 数据匿名化处理工具 |
| {{< ext "pgsmcrypto" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "SEC" >}} | 为PostgreSQL提供商密算法支持：SM2,SM3,SM4 |
| {{< ext "pg_enigma" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "SEC" >}} | PostgreSQL 加密数据类型 |
| {{< ext "pgaudit" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | {{< category "SEC" >}} | 提供审计功能 |
| {{< ext "pgauditlogtofile" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "SEC" >}} | pgAudit 子扩展，将审计日志写入单独的文件中 |
| {{< ext "pg_auditor" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | {{< category "SEC" >}} | 审计数据变更并提供闪回能力 |
| {{< ext "logerrors" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "SEC" >}} | 用于收集日志文件中消息统计信息的函数 |
| {{< ext "pg_auth_mon" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "SEC" >}} | 监控每个用户的连接尝试 |
| {{< ext "pg_jobmon" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "SEC" >}} | 记录和监控函数 |
| {{< ext "credcheck" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "SEC" >}} | 明文凭证检查器 |
| {{< ext "pgcryptokey" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "SEC" >}} | PG密钥管理 |
| {{< ext "pg_pwhash" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "SEC" >}} | PostgreSQL 高级密码哈希扩展（Argon2/scrypt/yescrypt） |
| {{< ext "login_hook" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "SEC" >}} | 在用户登陆时执行login_hook.login()函数 |
| {{< ext "set_user" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "SEC" >}} | 增加了日志记录的 SET ROLE |
| {{< ext "pg_snakeoil" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL--r" color="blue" >}} | {{< category "SEC" >}} | PostgreSQL动态链接库反病毒功能 |
| {{< ext "pgextwlist" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | {{< category "SEC" >}} | PostgreSQL扩展白名单功能 |
| {{< ext "sslutils" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "SEC" >}} | 使用SQL管理SSL证书 |
| {{< ext "noset" "pg_noset" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | {{< category "SEC" >}} | 阻止非超级用户使用SET/RESET设置变量 |
| {{< ext "pg_tde" >}} | {{< pgver "18,17,16,15,14,13" "g,g,r,r,r,r" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "SEC" >}} | Percona加密存储引擎 |
| {{< ext "sepgsql" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-sL---" color="blue" >}} | {{< category "SEC" >}} | 基于SELinux标签的强制访问控制 |
| {{< ext "auth_delay" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-sL---" color="blue" >}} | {{< category "SEC" >}} | 在返回认证失败前暂停一会，避免爆破 |
| {{< ext "pgcrypto" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | {{< category "SEC" >}} | 实用加解密函数 |
| {{< ext "passwordcheck" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-sL---" color="blue" >}} | {{< category "SEC" >}} | 用于强制拒绝修改弱密码的扩展 |
| {{< ext "wrappers" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "FDW" >}} | Supabase提供的外部数据源包装器捆绑包 |
| {{< ext "multicorn" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FDW" >}} | 用Python编写自定义的外部数据源包装器 |
| {{< ext "odbc_fdw" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FDW" >}} | 访问ODBC可访问的任何外部数据源 |
| {{< ext "jdbc_fdw" >}} | {{< pgver "18,17,16,15,14,13" "r,r,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FDW" >}} | 访问JDBC可访问的任何外部数据源 |
| {{< ext "pgspider_ext" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,r,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FDW" >}} | 使用多种FDW访问远程数据库服务器 |
| {{< ext "mysql_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FDW" >}} | MySQL外部数据包装器 |
| {{< ext "oracle_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FDW" >}} | 提供对Oracle的外部数据源包装器 |
| {{< ext "tds_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FDW" >}} | TDS 数据库（Sybase/SQL Server）外部数据包装器 |
| {{< ext "db2_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FDW" >}} | 提供对DB2的外部数据源包装器 |
| {{< ext "sqlite_fdw" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FDW" >}} | SQLite 外部数据包装器 |
| {{< ext "pgbouncer_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FDW" >}} | 用SQL查询pgbouncer统计信息，并执行pgbouncer命令 |
| {{< ext "etcd_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FDW" >}} | etcd分布式键值存储外部数据包装器 |
| {{< ext "informix_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FDW" >}} | Informix 外部数据包装器 |
| {{< ext "nominatim_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FDW" >}} | Nominatim 地理编码接口的 FDW 扩展 |
| {{< ext "mongo_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FDW" >}} | MongoDB 外部数据包装器 |
| {{< ext "redis_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FDW" >}} | 查询外部Redis数据源 |
| {{< ext "redis" "pg_redis_pubsub" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FDW" >}} | 从PG向Redis发送Pub/Sub消息 |
| {{< ext "kafka_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FDW" >}} | Kafka外部数据源包装器 |
| {{< ext "hdfs_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FDW" >}} | hdfs 外部数据包装器 |
| {{< ext "firebird_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FDW" >}} | Firebird外部数据源包装器 |
| {{< ext "aws_s3" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | {{< category "FDW" >}} | 从S3导入导出数据的外部数据源包装器 |
| {{< ext "log_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FDW" >}} | 访问PostgreSQL日志文件的FDW |
| {{< ext "dblink" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "FDW" >}} | 从数据库内连接到其他 PostgreSQL 数据库 |
| {{< ext "file_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "FDW" >}} | 访问外部文件的外部数据包装器 |
| {{< ext "postgres_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "FDW" >}} | 用于远程 PostgreSQL 服务器的外部数据包装器 |
| {{< ext "documentdb" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "SIM" >}} | 微软DocumentDB的API层 |
| {{< ext "documentdb_core" "documentdb" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "SIM" >}} | 微软DocumentDB的核心API层实现 |
| {{< ext "documentdb_distributed" "documentdb" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "SIM" >}} | DocumentDB多节点模式的API层 |
| {{< ext "documentdb_extended_rum" "documentdb" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="--sLd-r" color="blue" >}} | {{< category "SIM" >}} | DocumentDB扩展RUM索引访问方法 |
| {{< ext "orafce" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "SIM" >}} | 模拟 Oracle RDBMS 的一部分函数和包的函数和运算符 |
| {{< ext "pgtt" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "SIM" >}} | 类似Oracle的全局临时表功能 |
| {{< ext "session_variable" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "SIM" >}} | Oracle兼容的会话变量/常量操作函数 |
| {{< ext "pg_statement_rollback" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | {{< category "SIM" >}} | 在服务端提供类似Oracle/DB2的语句级回滚能力 |
| {{< ext "ivorysql_ora" "ivorysql" >}} | {{< pgver "18,17,16,15,14,13" "g,r,r,r,r,r" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "SIM" >}} | Oracle 兼容扩展 |
| {{< ext "ora_btree_gin" "ivorysql" >}} | {{< pgver "18,17,16,15,14,13" "g,r,r,r,r,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "SIM" >}} | Oracle 数据类型 GIN 索引支持 |
| {{< ext "ora_btree_gist" "ivorysql" >}} | {{< pgver "18,17,16,15,14,13" "g,r,r,r,r,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "SIM" >}} | Oracle 数据类型 GiST 索引支持 |
| {{< ext "pg_get_functiondef" "ivorysql" >}} | {{< pgver "18,17,16,15,14,13" "g,r,r,r,r,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "SIM" >}} | 获取函数定义 |
| {{< ext "plisql" "ivorysql" >}} | {{< pgver "18,17,16,15,14,13" "g,r,r,r,r,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "SIM" >}} | PL/iSQL 过程语言 |
| {{< ext "gb18030_2022" "ivorysql" >}} | {{< pgver "18,17,16,15,14,13" "g,r,r,r,r,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "SIM" >}} | 支持 GB18030-2022 与 UTF-8 编码转换 |
| {{< ext "pg_dbms_metadata" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "SIM" >}} | 添加 Oracle DBMS_METADATA 兼容性支持的扩展 |
| {{< ext "pg_dbms_lock" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "SIM" >}} | 为PG添加对 Oracle DBMS_LOCK 的完整兼容性支持 |
| {{< ext "pg_dbms_job" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "SIM" >}} | 添加 Oracle DBMS_JOB 兼容性支持的扩展 |
| {{< ext "pg_dbms_errlog" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "SIM" >}} | 模仿 Oracle DBMS_ERRLOG 模块来记录特定表的DML错误 |
| {{< ext "pg_utl_smtp" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "SIM" >}} | Oracle UTL_SMTP 兼容扩展（基于 plperlu） |
| {{< ext "babelfishpg_common" "babelfish" >}} | {{< pgver "18,17,16,15,14,13" "r,g,r,r,r,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "SIM" >}} | SQL Server 数据类型兼容扩展 |
| {{< ext "babelfishpg_tsql" "babelfish" >}} | {{< pgver "18,17,16,15,14,13" "r,g,r,r,r,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "SIM" >}} | SQL Server SQL语法兼容性扩展 |
| {{< ext "babelfishpg_tds" "babelfish" >}} | {{< pgver "18,17,16,15,14,13" "r,g,r,r,r,r" >}} | {{< badge content="--sLd-r" color="blue" >}} | {{< category "SIM" >}} | SQL Server TDS线缆协议兼容扩展 |
| {{< ext "babelfishpg_money" "babelfish" >}} | {{< pgver "18,17,16,15,14,13" "r,g,r,r,r,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "SIM" >}} | SQL Server 货币数据类型兼容扩展 |
| {{< ext "spat" >}} | {{< pgver "18,17,16,15,14,13" "r,g,r,r,r,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "SIM" >}} | 在PG中嵌入Redis风格的内存数据库 |
| {{< ext "pgmemcache" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "SIM" >}} | 为PG提供memcached兼容接口 |
| {{< ext "aux_mysql" "openhalo" >}} | {{< pgver "18,17,16,15,14,13" "r,r,r,r,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "SIM" >}} | MySQL兼容辅助扩展模块 |
| {{< ext "pglogical" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "ETL" >}} | PostgreSQL逻辑复制：三方扩展实现 |
| {{< ext "pglogical_origin" "pglogical" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "ETL" >}} | 用于从 Postgres 9.4 升级时的兼容性虚拟扩展 |
| {{< ext "pglogical_ticker" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "ETL" >}} | pglogical复制延迟以秒计的精确视图 |
| {{< ext "pgl_ddl_deploy" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "ETL" >}} | 使用 pglogical 执行自动 DDL 部署 |
| {{< ext "pg_failover_slots" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL--r" color="blue" >}} | {{< category "ETL" >}} | 在Failover过程中保留复制槽 |
| {{< ext "db_migrator" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dt-" color="blue" >}} | {{< category "ETL" >}} | 使用FDW从其他DBMS迁移到PostgreSQL |
| {{< ext "pgactive" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="-bsLd--" color="blue" >}} | {{< category "ETL" >}} | PostgreSQL多主逻辑复制 |
| {{< ext "spock" >}} | {{< pgver "18,17,16,15,14,13" "r,g,r,r,r,r" >}} | {{< badge content="-bsLd--" color="blue" >}} | {{< category "ETL" >}} | PostgreSQL 多主逻辑复制扩展 |
| {{< ext "lolor" >}} | {{< pgver "18,17,16,15,14,13" "r,g,r,r,r,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "ETL" >}} | 让 PostgreSQL 大对象兼容逻辑复制的扩展 |
| {{< ext "wal2json" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s----" color="blue" >}} | {{< category "ETL" >}} | 用逻辑解码捕获 JSON 格式的 CDC 变更 |
| {{< ext "wal2mongo" >}} | {{< pgver "18,17,16,15,14,13" "r,r,g,g,g,g" >}} | {{< badge content="--s----" color="blue" >}} | {{< category "ETL" >}} | 使用逻辑解码捕获MongoDB JSON格式的CDC变更 |
| {{< ext "decoderbufs" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s----" color="blue" >}} | {{< category "ETL" >}} | 将WAL逻辑解码为ProtocolBuffer协议的消息 |
| {{< ext "decoder_raw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s----" color="blue" >}} | {{< category "ETL" >}} | 逻辑复制解码输出插件：RAW SQL格式 |
| {{< ext "mimeo" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "ETL" >}} | 在PostgreSQL实例间进行表级复制 |
| {{< ext "repmgr" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "ETL" >}} | PostgreSQL复制管理组件 |
| {{< ext "pg_fact_loader" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "ETL" >}} | 在 Postgres 中构建事实表 |
| {{< ext "pg_bulkload" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="-bs-d--" color="blue" >}} | {{< category "ETL" >}} | 向 PostgreSQL 中高速加载数据 |
| {{< ext "test_decoding" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s----" color="blue" >}} | {{< category "ETL" >}} | 基于SQL的WAL逻辑解码样例 |
| {{< ext "pgoutput" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s----" color="blue" >}} | {{< category "ETL" >}} | PG内置的逻辑解码输出插件 |

**属性说明**: `c`:contrib `b`:二进制 `s`:动态库 `l`:需加载 `d`:需DDL `t`:无需特权 `r`:可重定位