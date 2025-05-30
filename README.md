# GoldenOwl-Test
## 🌍 Deployment Links

| Service               | URL                                                              |
|------------------------|------------------------------------------------------------------|
| Backend                | [Backend Host](https://go-backend-rc59.onrender.com) |
| Backend - Search       |  [Tra cứu điểm SBD cụ thể](https://go-backend-rc59.onrender.com/student-scores/?sbd=01000005) |
| Backend - Top10 GroupA |  [Trả về Top 10 điểm cao nhất khối A](https://go-backend-rc59.onrender.com/student-scores/top-10-groupA)         |
| Backend - Report       |  [Thống kê số điểm của từng môn theo level](https://go-backend-rc59.onrender.com/student-scores/report)         |
| Frontend               | [Frontend Host](https://golden-owl-gscore.netlify.app) |


## Tech Stack
- Backend: Golang
- Frontend: React.js (Vite/CRA)
- Database: PostgreSQL
- 
---

### Run Locally
#### Golang
```bash
# 1. Move to backend directory
cd backend

# 2. Download dependencies
go mod tidy

# 3. Run the server
go run main.go
```

### Reactjs
```bash
# 1. Move to frontend directory
cd frontend

# 2. Install dependencies
npm install

# 3. Start the development server
npm run dev
```

