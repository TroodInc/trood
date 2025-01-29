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
