# Nats Terminal User Interface

- No Ai
- Use bubble tea
- Interactive CLI

## Commands

- Publish
- Subscribe
- Request
- Reply

## UI Flow

```mermaid
flowchart TD
    A[Initial screen] -- Enter nats url --> B[Show commands]
    B -- Enter publish --> C[Show `subject` and `message`]
    B -- Enter subscribe --> D[Show `subject`]
    B -- Enter request --> E[Show `subject` and `message`]
    B -- Enter reply --> F[Show `subject` and `message`]
```

## Tasks

- ✅ Setup initial terminal UI
- ✅ Get nats url
- ▶️ Show commands
- ⬜ Publish
- ⬜ Subscribe
- ⬜ Request
- ⬜ Reply
