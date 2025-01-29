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
