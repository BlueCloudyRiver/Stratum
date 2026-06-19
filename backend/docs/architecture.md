# Stratum

Stratum is a business collaboration platform that helps startups and small businesses manage projects, teams, clients, and workflows from a single intuitive workspace.

## 1.0 
#### Expected to release by October 2026
Version 1.0 will deliver a modern, user-friendly platform designed to help businesses manage and streamline their operations. The platform will enable teams to collaborate effectively by managing employees, projects, deadlines, and basic client management from a unified dashboard.

This version focuses on reliability, performance, and intuitive user experience, the platform aims to integrate seamlessly with the tools and technologies businesses already use, helping organizations improve productivity, coordination, and growth.

## Technologies used

### Backend:
- Golang
- Chi
- PostgreSQL
- JWT
- Bcrypt
- Pgx
- Docker
- Redis

### Frontend:
- TypeScript
- HTML
- TailwindCSS

## Core Entities

### User
Represents an account on the platform.

### Workspace
Represents a company or team.

### Workspace Member
Links users to workspaces.

### Project
Represents a business initiative.

### Task
Represents work within a project.

### Client
Represents a customer of a workspace.

## Relationships

User
 ↓
WorkspaceMember
 ↓
Workspace
 ├── Projects
 │     └── Tasks
 │
 ├── Tasks
 │
 └── Clients


 ## User Flow

1. User registers.
2. User creates a workspace.
3. User invites team members.
4. User creates a project.
5. User creates tasks.
6. Team members complete tasks.
7. Workspace owner tracks progress.


## Design Decisions

- Tasks may belong directly to a workspace.
- Tasks may optionally belong to a project.
- Clients belong to workspaces.
- Users may belong to multiple workspaces.