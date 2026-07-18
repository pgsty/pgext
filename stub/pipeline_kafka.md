## Usage

Sources:

- [Official extension control file](https://github.com/pipelinedb/pipeline_kafka/blob/2c2e6e9d71463427dca641b0657e12e35b255734/pipeline_kafka.control)
- [Official upstream documentation](https://github.com/pipelinedb/pipeline_kafka/blob/2c2e6e9d71463427dca641b0657e12e35b255734/README.md)

`pipeline_kafka` — Integrate PipelineDB with Kafka for background-worker consumption, production, offsets, and lag monitoring

The reviewed catalog snapshot records version `1.0.0`, kind `preload`, and implementation language `C`.
Install and validate the declared extension dependencies first: `pipelinedb`.

```sql
CREATE EXTENSION "pipeline_kafka";
SELECT extversion
FROM pg_extension
WHERE extname = 'pipeline_kafka';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
