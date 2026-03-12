


## Usage

> [omni_txn: Transaction management](https://docs.omnigres.org/omni_txn/linearize/)

The `omni_txn` extension provides transaction linearization for stronger consistency guarantees beyond standard PostgreSQL isolation levels.

### Enable Linearization

```sql
BEGIN TRANSACTION ISOLATION LEVEL SERIALIZABLE;
SELECT omni_txn.linearize();
-- Perform operations; conflicts raise serialization errors
COMMIT;
```

### Check Linearization Status

```sql
SELECT omni_txn.linearized();  -- returns boolean
```

### Behavior

Linearizability ensures every operation appears to take place atomically in an order consistent with real-time ordering. The extension intercepts mutating statements (INSERT, UPDATE, DELETE, MERGE) and enforces:

- Write-after-read conflict detection across serializable linearized transactions
- Commit-after-read conflict detection
- Snapshot conflict detection

Raises serialization error exceptions when conflicts are detected. False positives are possible; use with retry mechanisms (compatible with `omni_txn.retry`).
