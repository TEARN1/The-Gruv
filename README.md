# The Gruv — Frontend (Vite)

This folder contains the Vite-based frontend scaffold for The Gruv.

Quick start
1. Copy `.env.example` to `.env` and fill the Firebase + API values.
2. Install deps:
   - npm install
3. Run dev server:
   - npm run dev
4. Open http://localhost:3000

Notes
- Backend API: configured via VITE_API_BASE_URL
- WebSocket (dev): VITE_WS_URL
- Firebase config must be set in environment variables.

Next steps after this scaffold:
- Split remaining UI pieces into components & polish styles.
- Add WebSocket hook for real-time updates.
- Replace placeholder images and wire missing API endpoints.