## Capabilities of Auth Service (ABAC)

### Attribute-Based Access Control
- You can check access not only by roles but also by user attributes, object attributes, and environment (sbj, obj, ctx).
- Complex conditions with AND/OR, operators (in, not, lt/gt), etc.

For designing "granular" security: "The user sees objects where `object.owner == user.id` or `object.isPublic == true`."

### 2-Factor Authentication
- Ready-made logic for SMS/phone verification (can be extended).
- Management of factors and temporary tokens.

If additional security measures are needed, everything is built-in and available out of the box.

### Field masking mechanism
- You can hide certain fields in the response (e.g., "salary" for regular users).
- Convenient when the same entity is used by different user categories.
