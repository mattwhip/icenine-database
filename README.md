# icenine-database
Common database access code for IceNine services

This is a common services library for the IceNine project, a scalable cloud-based multiplayer server. It contains ORM code using Buffalo's Pop library for all database access in the IceNine system. This separation is made to allow multiple entities in the IceNine system to access the same objects (e.g. Admin and any other service).

Concepts/technologies used:
- Golang with Buffalo's Pop library for Object Relational Mapping
