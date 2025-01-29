## What is it and why is it needed?

Trood Core is a set of microservices that provide various functional blocks for typical business applications. The project is focused on a quick start when creating corporate solutions: CRM systems, internal portals, document management systems, etc. Instead of writing typical blocks from scratch, you get ready-made services:

### Custodian (Data structure management and interaction with the database):
- Generation of REST endpoints and database tables (PostgreSQL) from JSON descriptions.
- Migration mechanism (with history and branching).
- Bulk operations (mass addition, updating, deletion).
- Integration with external systems via Webhooks.

### File Service (File management and processing):
- Support for metadata, extensions, types, and templates.
- Document templates and on-the-fly file generation.
- Integration with ABAC for file access management.

### Auth Service (Authorization and access management):
- High flexibility in access management thanks to ABAC.
- Easy integration with other systems via REST API.
- Detailed control over users and roles.
- Built-in security mechanisms (multi-factor authentication, tokens).

### Mail Service (Email handling):
- Mailbox management (SMTP/IMAP).
- Sending and receiving emails, attachments.
- Working with folders, email threads, contacts.

Together, these services cover a wide range of application needs: data storage, file handling, email management, and access rights configuration — all without reinventing the wheel or manually writing repetitive code.

## How can this help developers?

### Quick start and reduced routine
- **CRUD out of the box**: Describe objects in JSON — Custodian creates tables and REST endpoints.
- **Bulk operations and search (RQL)**: No need to manually write SQL queries or tweak filter logic.
- **Minimal code duplication**: Quickly define fields and relationships through metadata. Custodian automatically synchronizes the database schema.

### Ready-made mail and file modules
- **Mail Service** simplifies sending/receiving emails, including attachments, folders, and threads. It speeds up typical use cases (mailings, notifications, capturing incoming emails).
- **File Service** solves file storage and document generation from templates. Developers don’t need to write logic for uploading/storing files, etc.

### Flexible authorization (ABAC)
- Fine-tuned access control without manual "if user in role" checks: Auth Service allows setting rules based on user attributes (subject), object attributes (object), and context (context).
- Example: "A manager can only see their orders, while an administrator can see all."

### Scalable microservice architecture
- Each service (Custodian, File, Mail, Auth) has its own Docker container/repository.
- You can scale narrowly, for example, if you need to handle high email load without affecting other blocks.

### Convenient integrations
- REST API in all services — easily integrates with any Frontend framework (React, Vue, Angular) or mobile applications.
- External calls on create/update/delete events (hooks, actions).

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

## Capabilities of Mail Service

### Connecting external mail servers
- SMTP/IMAP can be configured for any mail provider (Gmail, corporate Exchange, etc.).
- Developers don’t need to solve the problem of "how to communicate with the mail server," everything is implemented "out of the box."

### Grouping emails into threads
- Analysis of `reply-to` headers → merging emails into discussion threads.

This simplifies communication logic and searching for the entire history on a topic.

### Working with folders and contacts
- Creating folders, moving emails.
- Bulk operations, linking emails and contacts.

Mechanics very similar to a "standard email client," but available via API.

### Attachments
- Works in conjunction with File Service (an uploaded file can be immediately attached to an email).
- Management of large files/documents without manual workarounds.

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

## Additional capabilities

### Metadata cache management
- Custodian has a caching mechanism, which avoids recreating object structures on every request.
- The cache is only reset when the schema changes (when you change metadata).

This provides a balance between flexibility (dynamic schemas) and performance.

### Event model and notifications
- When performing CRUD operations, Custodian can call "actions," Mail Service can process new email triggers, and File Service can update readiness status.
- You can set up chains like "after creating a record → send a notification to an external service" without writing manual "observer" code.

## Example usage scenario

### Describing a model in Custodian:
```json
{
  "name": "order",
  "key": "id",
  "fields": [
    {
      "name": "id",
      "type": "number",
      "optional": true,
      "default": { "func": "nextval" }
    },
    { "name": "clientName", "type": "string", "optional": false },
    { "name": "status", "type": "string" }
  ]
}
```
- Sent this to `/custodian/meta/order` → the table and endpoints `/custodian/data/order` are created automatically.

### Files:
- Uploaded an invoice template to File Service → `/api/v1.0/templates/`.
- When generating an invoice: `/api/v1.0/files/from_template` → the response contains a ready-made PDF with order data.

### Mail:
- Created a Mailbox `/mailboxes` (SMTP), sent an email to the client along with the PDF from File Service.
- Received incoming mail via `/mailboxes/<id>/fetch`, processed it in an external service.

### Rights via Auth Service (ABAC):
- Set up a rule: "A user with the manager role can only retrieve and update their own orders."
- When accessing `/custodian/data/order`, filtering/checking is automatically applied.

**Result**: You don’t manually write CRUD, SQL, file uploads, SMTP scripts, etc. — everything is ready to use and integrate.

## Comparison with other options

### Classic frameworks (Django, Laravel, Rails)
- **Pros of frameworks**: Large ecosystems, many plugins, extensive communities.
- **Cons**: You need to manually (or via generators) describe models, migrations, routes. Although there is CRUD generation, there is no single "out of the box" mail module with threads or a full-fledged file generator (you need to add third-party plugins and deal with their support).
- ABAC will have to be tweaked through additional libraries (most often the standard is RBAC).
- Trood Core is more "ready-made" for complex "corporate" services: mail, files, metadata, migrations, ABAC — all under one roof.

### Custom microservice architecture + external services
- **Pros**: Complete freedom — choose any database, authorization, file handling library, mail service.
- **Cons**: A lot of integration code is needed (Kafka, HTTP service, IAM like Keycloak/Auth0, S3/MinIO, Postgres + Liquibase, etc.).
- Support and updates are distributed across many solutions → potential version conflicts, complex DevOps setup.
- Trood Core — fewer integrations: one ecosystem, one approach to metadata, migrations, endpoints.

### Low-code/No-code platforms
- **Pros**: You can click around, quickly create prototypes.
- **Cons**: For "serious" corporate requirements (complex processes, strict security requirements, extensibility), you often hit limitations.
- Trood Core allows writing full-fledged code while eliminating routine at the level of CRUD, files, mail, and authorization. Much more flexible for customization than No-code.

### Using AI tools (e.g., Cursor, ChatGPT/Copilot)
- **What do AI tools offer?**: Boilerplate code generation, SQL hints, etc.
- **Comparison with Trood Core**:
  - AI tools help write code faster, but you are still responsible for integration, structure, and testing.
  - Trood Core doesn’t just generate code. It already contains ready-made services and API endpoints that don’t require constant generation and manual maintenance.
  - Instead of "generating CRUD on Node.js + Express + Sequelize," you describe fields in metadata, and Custodian provides everything, while AI can assist in integration scenarios (logic, business rules) but doesn’t replace the platform’s model level.
- **Win**: Trood Core + AI tools can be combined for even faster results (AI writes part of the business logic, but you don’t need to write basic CRUD/files/mail code).
- If you only use AI, you still have to design the architecture manually, decide how to store files, where to get mail modules, etc. Trood already solves these issues.

## Performance and scaling considerations
- Custodian can handle load using PostgreSQL as a reliable backend.
- File Service and Mail Service can be scaled independently (vertically or horizontally).
- The database (PostgreSQL) can be scaled either to a more powerful machine or sharded (this is already a detail of corporate infrastructure).
- In most cases, the bottleneck is a properly configured Postgres cluster/caching, not Custodian or File Service.

## Brief summary
- Trood Core simplifies the developer’s life: provides a ready-made foundation (data, files, mail, authorization).
- Less boilerplate — you can focus on project specifics (business logic, UI, integrations).
- Microservice model: easy to add your own services/modules.
- ABAC, 2FA, REST "out of the box": convenient for corporate systems.
- **Main benefit**: Rapid development and prototyping (reduced timelines).
- Stable architecture: a unified ecosystem for CRUD, mail, files, and authorization.
- Flexibility for future changes (migrations, dynamic model, access rights configuration).

Comparing Trood Core with classic frameworks, custom microservices, or AI code generators, the main takeaway is: Trood services already solve standard tasks at an advanced level, rather than just generating templates. This reduces development and maintenance costs without limiting you in further customizations.

## What’s next?
If you want to test Trood Core in practice:
1. Launch the services (Docker or locally), connect Postgres.
2. Describe one or two entities in Custodian (via JSON metadata) and try adding a few records.
3. Set up File Service, upload a file, and insert its link into a Custodian object.
4. Connect Auth Service (2FA, ABAC), test access restrictions.
5. Mail Service: set up a mailbox, send yourself an email, process incoming mail.

You’ll see that many things that usually require several libraries and significant effort "just work" almost "out of the box."

If you need to convince clients that this is faster and cheaper than custom development or limited no-code solutions, just demonstrate:
- **Speed**: How quickly you can set up basic CRUD and file/mail logic.
- **Flexibility**: If needed, you can write additional code, create services around it, without being limited by a "black box."

Use Trood Core as a foundation, and AI tools (Cursor, ChatGPT, etc.) — to speed up writing auxiliary logic, tests, frontend, etc. This combination gives developers the maximum advantage in time and quality.

## Enterprise Manager (EM) Admin Panel
Enterprise Manager is a built-in admin panel that allows managing key objects and settings of Trood Core "out of the box." Instead of developing your own interface for configuration and CRUD operations, you get a ready-made SPA application.

### Main features
- **Business object management**: Visualization of structure and relationships, migration history.
- **Access policy definition**: Interface for setting access policies at the level of subject objects, fields, and conditions.
- **File list access**.

### Why is this useful?
- **Unified "control center"**: You don’t write your own admin panel for operations with business objects, files, or rights; everything is available through EM.
- **Accelerated configuration**: You can experiment with a new Custodian schema (metadata, fields) and immediately check how the data looks in tables.
- **Easy debugging and visualization**: The relationship diagram helps quickly understand how objects interact without diving deep into SQL or API calls.

## Enterprise Manager 2 (EM 2)
EM 2 is an updated version of the Trood admin panel, providing a centralized interface for managing all key aspects of the project’s infrastructure.

EM 2 was created as a SAAS solution for teams that want to deploy their projects based on Trood Core.

In addition to managing Trood Core services (as in EM 1), the new admin panel allows configuring the project’s network infrastructure:
- Connecting domains.
- Connecting custom services based on Docker images.
- Storing backups.
- Linking to different types of applications (Web, mobile).
- Creating front-end using Front Composer.

As one of the options for implementing business logic, EM 2 assumes integration with the n8n service.

### Why is EM 2 useful for developers?
- **Centralized management**: Everything in one place — services, applications, domains, access, and settings.
- **Quick deployment**: Templates, automation, and flexible configuration tools.
- **Flexibility and scalability**: Ability to connect custom services, work with various applications, and backups.
- **Easy integration**: REST API for working with external systems and flexible route and authorization configuration.

## Front Composer
Front Composer is a visual tool that extends Trood Core’s capabilities, allowing you to build user interfaces (forms, pages, lists, buttons) literally "from blocks" without manual coding. It is built into EM 2 and tightly integrates with Custodian, File Service, Mail Service, and Auth Service.

### Main idea
Instead of writing React components and linking them to the API, a developer (or advanced user) can drag and drop ready-made blocks onto a "canvas" area and configure them (text, styles, data sources) in a convenient interface. The resulting structure is stored in JSON and can be loaded/updated dynamically without recompiling the code.

### What does this offer?
- **Drag-n-Drop editor**: Select the necessary components (blocks, buttons, fields, containers, routes) from the panel and place them on the page with a single mouse movement.
- **Declarative configuration**: The entire interface (components, their nesting, data bindings) is described in JSON, which Composer can read and edit. This speeds up changes: edits become "configuration" rather than "code."
- **Easy data connection**:
  - **Custodian**: Display lists of records, forms for editing business objects.
  - **File Service**: Quickly embed a file list or upload form.
  - **Mail Service**: Configure email, mailbox, and template output.
  - **Auth Service**: Configure access to components, show or hide elements depending on roles.
- **Logic management**:
  - **Expressions**: Through a small "expression" syntax, Composer allows conditional displays, field calculations, etc.
  - **Actions**: Buttons/links can call REST methods, open modal windows, switch pages.
- **Flexible architecture**: The entire configuration can be split into templates/fragments, loaded lazily (lazy load). Composer integrates with versioning and drafts in Enterprise Manager (stores templates as objects in Custodian).

### How does this help developers?
- **Reduces code**: Many typical UI tasks (forms, lists, navigation) are solved "out of the box" — just visually arrange and configure components.
- **Quick prototypes**: You can experiment with screen layout, test ideas — all without writing React/JSX manually.
- **Convenient for support**: Managers, business analysts, or other team members can make UI edits (texts, layout, even logic) without diving into the codebase.
- **Security and rights**: Composer considers ABAC, so access to editing or viewing individual components is configured through Auth Service.

### Example scenario
- **Creating a new page**: The user opens Front Composer, sees an empty "canvas."
- Drags a container, then several fields (Input, Button) for a form, adds a "Table" to display a list of orders (data — from Custodian).
- Configures properties: API path, expressions for display/filtering, click handler (calls a `$action` like "create a new record").
- Saves to Custodian: the configuration goes into a template object.
- Front Core automatically loads the JSON and builds the UI when this page is launched.

### When is it especially relevant?
- **Quick UI changes**: Need to quickly add fields or buttons without touching the main code.
- **Multi-form applications**: Many forms, tables, different relationships; Composer simplifies maintaining order and visual style.
- **Task delegation**: Part of the team can work with Composer while developers focus on complex business logic.
