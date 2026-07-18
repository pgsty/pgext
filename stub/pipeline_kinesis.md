## Usage

Sources:

- [Upstream README](https://github.com/pipelinedb/pipeline_kinesis/blob/87b8ab05ca848fe68ed29701e1d35d9e1a441b7c/README.md)
- [Extension control file](https://github.com/pipelinedb/pipeline_kinesis/blob/87b8ab05ca848fe68ed29701e1d35d9e1a441b7c/pipeline_kinesis.control)
- [SQL installation script](https://github.com/pipelinedb/pipeline_kinesis/blob/87b8ab05ca848fe68ed29701e1d35d9e1a441b7c/pipeline_kinesis--0.9.0.sql)
- [Extension implementation](https://github.com/pipelinedb/pipeline_kinesis/blob/87b8ab05ca848fe68ed29701e1d35d9e1a441b7c/pipeline_kinesis.c)

`pipeline_kinesis` version `0.9.0` is a historical PipelineDB extension that starts dynamic background workers to consume AWS Kinesis shards and copy records into a PipelineDB stream. It persists endpoint, consumer, and shard-sequence metadata in the `pipeline_kinesis` schema.

### Historical workflow

The following assumes old PipelineDB semantics such as `CREATE STREAM`; it is not a current PostgreSQL example:

```sql
CREATE EXTENSION pipeline_kinesis;
SELECT pipeline_kinesis.add_endpoint('endpoint', 'us-west-2');
SELECT pipeline_kinesis.consume_begin('endpoint', 'test', 'foo_stream');
SELECT pipeline_kinesis.consume_end('endpoint', 'test', 'foo_stream');
```

The 2016 code targets PipelineDB and an old AWS C++ SDK, has no current PostgreSQL compatibility path, and starts networked background workers inside the server. Credential-file paths, endpoint metadata, worker limits, replay position, and grants all require explicit control. The implementation also contains unsafe argument-count handling in consumer lookup. Keep it out of modern or production clusters; use a maintained external Kinesis consumer instead.
