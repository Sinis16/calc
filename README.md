# Full-Stack Calculator

React + TypeScript frontend (Vite) consuming a Go REST microservice.
Supports add, subtract, multiply, divide, power, square root and percentage.

## Setup

Requirements: **Go 1.27+**, **Node.js 20+**, and optionally **Docker**.

```bash
git clone https://github.com/Sinis16/calc.git
cd calc/frontend && npm install
```

## Running

**Docker** (both services at `http://localhost:3000`):

```bash
docker compose up --build
```

**Backend** (`http://localhost:8080`):

```bash
cd backend
go run ./cmd/server
```

**Frontend** (`http://localhost:5173`, proxies `/api` to the backend):

```bash
cd frontend
npm run dev
```

## Tests

```bash
cd backend && go test ./... -cover        # calculator: 100%, api: 100%
cd frontend && npm run coverage           # 100% lines, 96% branches
```

## API

`POST /api/v1/calculate` with `{ "operation": string, "a": number, "b": number }`
(`b` is omitted for `sqrt`). Returns `200 { "result": number }` or `400 { "error": string }`.

```bash
curl -X POST http://localhost:8080/api/v1/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation":"add","a":2,"b":3}'
# {"result":5}

curl -X POST http://localhost:8080/api/v1/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation":"divide","a":10,"b":0}'
# {"error":"division by zero"}

curl -i http://localhost:8080/health
```

## Design Decisions

- **Layered backend:** `calculator` holds pure math with no HTTP, `api` handles JSON and validation,
  and `cmd/server` only starts the server. Each layer is tested independently.
- **Frontend split by responsibility:** `api.ts` is the only module that calls `fetch`; validation
  and formatting are pure functions, so the UI is tested by mocking the API.
- **Single endpoint** with an `operation` field; adding an operation is one map entry.
- **`*float64` operands** distinguish a missing value from an explicit `0`.
- **Overflow and NaN are returned as errors**, since they can't be encoded in JSON.
- **Go standard library only**; Go 1.22+ routing supports HTTP methods natively.
- **Validation in both layers:** the frontend gives instant feedback, the backend is the source of truth.
- **Proxy instead of CORS:** Vite (dev) and nginx (Docker) serve everything from one origin.
- **Responsive UI:** a horizontal `A · operation · B` layout on desktop that stacks vertically on mobile.
- **Multi-stage Docker builds** keep runtime images small (static Go binary on distroless, nginx for the frontend).

## Assumptions

- Percentage means "A percent of B": `percentage(20, 50) = 10`.
- Results are displayed with up to 12 significant digits.

## AI Usage

Claude was used for guidance, code review and debugging. See [PROMPTS.md](PROMPTS.md).
