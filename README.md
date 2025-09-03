# TelegramClerk 📝

Personal Telegram bot to organize everyday life: notes (journal), tasks with deadlines and reminders, simple schedule, shopping and checklist lists. Planned: lightweight AI text & voice analysis (parsing / structuring only, no content generation).

## Core Features ✨
- Notes / journal
- Tasks & reminders
- Schedule (basic time slots / events)
- Lists (shopping, checklists)
- Text & voice message analysis (parse lists and commands)

## Tech Stack 🛠️
- Backend: Go 
- Integrations: Python 
- Database: PostgreSQL 
- Cache / fast access: Redis
- Async messaging: RabbitMQ
- Containerization: Docker → Kubernetes

## Repository Structure 📁
```
services/        # Microservices (Go + Python)
  api/
  telegram/
  ai/
  tasks/
  schedule/
  lists/
shared/          # Shared proto / configs
database/        # DB migrations
deployments/     # docker-compose & k8s manifests
scripts/         # Utility scripts
```

## Status 🚧
Early development. Defining base architecture and data schema.

## License 📄
See LICENSE

## Author 👤
shishmarevv
