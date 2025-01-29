## Capabilities of Custodian

### Broad data search and selection capabilities
- RQL queries (filters, sorting, pagination).
- Built-in `only/exclude` parameters for selective field output.
- Depth management for nested relationships.

These mechanisms allow efficient collection of complex data in a "tree" structure in one request, saving effort on writing manual joins.

### Relationship management (including onDelete and generic relationships)
- Multiple deletion modes (cascade, restrict, setNull).
- Support for "generic" references, where one record can reference multiple potential object types.

This provides great flexibility in modeling: you can describe polymorphic relationships between objects without additional "manual" code.

### Calling external services on data changes (actions)
- Ability to "subscribe" to create/update/delete events and trigger arbitrary HTTP endpoints.
- Inclusion of required fields (`includeValues`) when sending, including support for generic logic (various types of related objects).

This is convenient when you need to automatically integrate with other systems (notify CRM, trigger calculations, etc.) without writing your own "middleware."

### Fine support for bulk operations
- POST, PATCH, DELETE multiple records in one request.
- Full transactionality (either all records are changed, or a rollback occurs).

For large volumes of data, this reduces network overhead and simplifies code.

### Optimistic locking (CAS)
- Prevents conflicts when editing the same record in parallel.
- If CAS is not needed, it can be disabled.

### Convenient migrations
- Schema management through a full-fledged migration mechanism (history, branching, rollbacks).
- Generation and application of migrations without manual SQL writing.
- Ability to combine migrations with metadata (when you need to quickly change the structure on the fly).
