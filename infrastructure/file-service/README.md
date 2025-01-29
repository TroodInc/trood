## Capabilities of File Service

### Flexible integration with cloud storage
- Support for S3 (Amazon, MinIO), DigitalOcean Spaces, etc.
- Local disk (in development environments or if there is no cloud).
- Developers don’t need to set up integration themselves — File Service already knows how to store and serve files via URL.

### Document generation from templates
- Django-template (or similar) logic inside, support for PDF/PNG.
- Automatic data substitution from JSON.
- Can cover many use cases for generating acts, invoices, contracts "on the fly."

### Soft deletion mechanism
- When a file is deleted, a `deleted=true` flag is set.
- There is a periodic/manual mechanism for final cleanup (`cleanup`).

Useful for "safe" deletion and the possibility of "restoration" if needed.