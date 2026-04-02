🦷 Dental Appointment Queue Management System
ระบบจัดการคิวและนัดหมายคลินิกทันตกรรม พัฒนาด้วย Go (Fiber) และ Next.js (App Router)

🚀 Tech Stack
Backend
Language: Go 1.2x

Framework: Fiber v2

ORM: GORM

Database: PostgreSQL

Authentication: JWT (JSON Web Token) with HttpOnly Cookies

Frontend
Framework: Next.js 14+ (App Router)

Styling: Tailwind CSS

HTTP Client: Axios / Fetch API

State Management: React Hooks (useState, useEffect)

📂 Project Structure
Plaintext
.
├── backend/
│   ├── internal/
│   │   ├── controller/   # Handling HTTP Requests
│   │   ├── middleware/   # JWT & CORS Middlewares
│   │   ├── models/       # GORM Entities (User, Patient, Role)
│   │   ├── repository/   # Database Operations (Transactions)
│   │   └── service/      # Business Logic (Hashing, JWT Signing)
│   ├── dto/              # Data Transfer Objects
│   ├── .env              # Environment Variables (JWT_SECRET, DB_URL)
│   └── main.go           # Entry Point
│
└── frontend/
    ├── app/
    │   ├── (Auth)/       # Login & Register Pages
    │   ├── dentist/      # Dentist Dashboard (Protected)
    │   └── patient/      # Patient Dashboard (Protected)
    ├── components/       # Shared UI Components (Forms, Navbar)
    └── middleware.ts     # Next.js Middleware for Auth Guard
🛠️ Installation & Setup
1. Backend Setup
เข้าไปที่โฟลเดอร์ backend:

Bash
cd backend
สร้างไฟล์ .env และกำหนดค่า:

ข้อมูลโค้ด
DB_URL=host=localhost user=postgres password=yourpassword dbname=dental_db port=5432 sslmode=disable
JWT_SECRET=your_super_secret_key
รันโปรเจกต์:

Bash
go run main.go
2. Frontend Setup
เข้าไปที่โฟลเดอร์ frontend:

Bash
cd frontend
ติดตั้ง Dependencies:

Bash
npm install
สร้างไฟล์ .env.local:

ข้อมูลโค้ด
NEXT_PUBLIC_API_URL=http://localhost:8080/api
รันโปรเจกต์:

Bash
npm run dev
🔐 Features Implemented
User Authentication: * ระบบสมัครสมาชิก (Register) พร้อมระบบ Transaction ป้องกันข้อมูลซ้ำ (Unique Constraint Handling)

ระบบเข้าสู่ระบบ (Login) ด้วย Username/Email และจัดเก็บ JWT ใน HttpOnly Cookie

Role-based Access Control:

แบ่งระดับผู้ใช้งานเป็น Dentist (Admin) และ Patient

ระบบ Middleware ฝั่ง Backend สำหรับตรวจสอบ Token

ระบบ Middleware ฝั่ง Frontend สำหรับป้องกันการเข้าหน้า Dashboard โดยไม่ได้รับอนุญาต

CORS & Security:

ตั้งค่า CORS ให้รองรับการรับ-ส่ง Credentials (Cookies) ระหว่าง Next.js และ Fiber

📝 Database Schema (Brief)
Users: เก็บข้อมูลพื้นฐาน (username, email, password_hash, phone, role_id)

Patients: เก็บข้อมูลเพิ่มเติมสำหรับคนไข้ (user_id, medical_condition, allergic_medication)

Roles: นิยามสิทธิ์ (1: Dentist, 2: Patient)