# 🦷 Dental Appointment Queue Management System
ระบบจัดการคิวและนัดหมายคลินิกทันตกรรม พัฒนาด้วยเทคโนโลยี Full-stack สมัยใหม่

---

## 🚀 Tech Stack

### Backend
* **Language:** Go 1.2x
* **Framework:** [Fiber v2](https://gofiber.io/)
* **ORM:** [GORM](https://gorm.io/)
* **Database:** PostgreSQL
* **Authentication:** JWT (JSON Web Token) with **HttpOnly Cookies**

### Frontend
* **Framework:** [Next.js 14+ (App Router)](https://nextjs.org/)
* **Styling:** Tailwind CSS
* **HTTP Client:** Axios / Fetch API
* **State Management:** React Hooks (useState, useEffect)

---

## 📂 Project Structure

```plaintext
.
├── backend/
│   ├── internal/
│   │   ├── controller/   # Handling HTTP Requests
│   │   ├── middleware/   # JWT & CORS Middlewares
│   │   ├── models/       # GORM Entities (User, Patient, Role)
│   │   ├── repository/   # Database Operations (Transactions)
│   │   └── service/      # Business Logic (Hashing, JWT Signing)
│   ├── dto/              # Data Transfer Objects
│   ├── .env              # Environment Variables
│   └── main.go           # Entry Point
│
└── frontend/
    ├── app/
    │   ├── (Auth)/       # Login & Register Pages
    │   ├── dentist/      # Dentist Dashboard (Protected)
    │   └── patient/      # Patient Dashboard (Protected)
    ├── components/       # Shared UI Components
    └── middleware.ts     # Next.js Middleware for Auth Guard

## 🛠️ Installation & Setup
###1. Backend Setup
1. เข้าไปที่โฟลเดอร์ `backend`:

```
cd backend

2. สร้างไฟล์ .env และกำหนดค่า:

```
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
User Authentication
ระบบสมัครสมาชิก (Register): มาพร้อมระบบ Transaction ป้องกันข้อมูลผิดพลาด และดักจับ Unique Constraint Handling (Username, Email, Phone)

ระบบเข้าสู่ระบบ (Login): รองรับทั้ง Username/Email และจัดเก็บ JWT ในรูปแบบ HttpOnly Cookie เพื่อความปลอดภัยสูงสุด

Role-based Access Control
User Roles: แบ่งระดับผู้ใช้งานชัดเจนเป็น Dentist (Admin) และ Patient

Backend Security: ระบบ Middleware สำหรับตรวจสอบและยืนยันความถูกต้องของ Token ก่อนเข้าถึง API

Frontend Guard: ใช้ Next.js Middleware ป้องกันการเข้าถึงหน้า Dashboard โดยไม่ได้รับอนุญาต (Unauthorized Access)

CORS & Security
Cross-Origin Resource Sharing: ตั้งค่า CORS ให้รองรับการรับ-ส่ง Credentials (Cookies) ระหว่าง Next.js และ Fiber อย่างปลอดภัย

📝 Database Schema (Brief)
Users: เก็บข้อมูลพื้นฐาน (username, email, password_hash, phone, role_id)

Patients: เก็บข้อมูลเพิ่มเติมเฉพาะคนไข้ (user_id, medical_condition, allergic_medication)

Roles: นิยามสิทธิ์การใช้งาน (1: Dentist, 2: Patient)