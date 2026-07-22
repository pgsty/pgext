## Usage

Sources:

- [Version 1.0 SQL API](https://github.com/RekGRpth/pg_grpc/blob/adbdf3a00c68581b6691c48bb45e81f8295e442e/pg_grpc--1.0.sql)
- [gRPC C-core implementation](https://github.com/RekGRpth/pg_grpc/blob/adbdf3a00c68581b6691c48bb45e81f8295e442e/pg_grpc.c)
- [Extension control file](https://github.com/RekGRpth/pg_grpc/blob/adbdf3a00c68581b6691c48bb45e81f8295e442e/pg_grpc.control)

`pg_grpc` is an experimental, very low-level wrapper around a few gRPC C-core calls. It maintains one channel, call, and completion queue in backend-process globals and exposes only boolean step functions; it does not provide protobuf encoding, service stubs, response data, status details, deadlines, metadata, or a complete RPC abstraction.

### Exposed Sequence

The callable order is to create a channel, create one call for a method, submit a raw text payload, and advance the completion queue.

```sql
CREATE EXTENSION pg_grpc;

SELECT grpc_insecure_channel_create('127.0.0.1:50051');
SELECT grpc_channel_create_call('/example.Service/Method');
SELECT grpc_call_start_batch('raw request bytes');
```

The SQL API also declares `grpc_secure_channel_create(target)` and `grpc_completion_queue_next()`. Every function returns only `boolean`; there is no SQL function that returns the received message. Calling a step out of order raises an extension error.

### Prototype and Security Boundaries

The reviewed implementation passes a NULL credentials pointer to its secure-channel constructor, uses an infinite call deadline, stores mutable state per backend, and loops in `grpc_completion_queue_next()` until queue shutdown. These paths can fail, block a database session indefinitely, or leave state unsuitable for a subsequent call. The batch code does not expose protobuf serialization or the received buffer to SQL, so a successful `true` result is not proof of an application-level RPC response.

All functions can initiate outbound network activity from the PostgreSQL server process and are granted to `PUBLIC` unless administrators revoke them. Revoke execution, allow only trusted destinations, and do not use this prototype for production RPCs or untrusted input. Prefer a maintained external service or a fully specified extension that enforces TLS credentials, deadlines, payload limits, cancellation, and response/status handling.
