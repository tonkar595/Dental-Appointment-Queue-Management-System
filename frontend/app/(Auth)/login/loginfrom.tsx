// 'use client';

// import React, { useState } from 'react';
// import Link from 'next/link';
// import { useRouter } from 'next/router';

// export default function LoginForm() {
//     const [identity, setIdentity] = useState('');
//   const [password, setPassword] = useState('');
//   const [error, setError] = useState('');
//   const router = useRouter();

//   const handleSubmit = async (e: React.FormEvent) => {
//     e.preventDefault();
//     setError('');

//     try {
//       const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/auth/login`, {
//         method: 'POST',
//         headers: {
//           'Content-Type': 'application/json',
//         },
//         body: JSON.stringify({ identity, password }),
//         // สำคัญมาก: ต้องมีบรรทัดนี้เพื่อให้ Browser เก็บ Cookie
//         credentials: 'include', 
//       });

//       const data = await response.json();

//       if (!response.ok) {
//         throw new Error(data.error || 'เข้าสู่ระบบไม่สำเร็จ');
//       }

//       // ถ้าสำเร็จ (ได้ Status 200)
//       console.log('Login Success:', data);
      
//       // ตรวจสอบ Role เพื่อแยกหน้า Dashboard (ส่งมาจาก DTO ที่เราแก้ไว้)
//       if (data.role_name === 'Dentist') {
//         router.push('/dentist/dashboard');
//       } else {
//         router.push('/patient/dashboard');
//       }
      
//     } catch (err: any) {
//       setError(err.message);
//     }
//   };
//   return (
//     <form className="space-y-6">
//           <div>
//             <label className="block text-sm font-bold text-gray-700 mb-2">Username or Email</label>
//             <input 
//               type="text" 
//               className="w-full px-5 py-3 rounded-xl border border-gray-200 focus:border-[#8f1eae] focus:ring-1 focus:ring-[#8f1eae] outline-none transition"
//               placeholder="กรอกชื่อผู้ใช้งาน หรือ อีเมล"
//             />
//           </div>

//           <div>
//             <label className="block text-sm font-bold text-gray-700 mb-2">Password</label>
//             <input 
//               type="password" 
//               className="w-full px-5 py-3 rounded-xl border border-gray-200 focus:border-[#8f1eae] focus:ring-1 focus:ring-[#8f1eae] outline-none transition"
//               placeholder="กรอกรหัสผ่าน"
//             />
//           </div>

//           <button 
//             type="submit"
//             className="w-full bg-[#8f1eae] text-white font-bold py-4 rounded-xl hover:opacity-90 transition-all shadow-lg shadow-purple-100"
//           >
//             Submit
//           </button>
//         </form>
//   );
// }

"use client";

import axios from "axios";
import Link from "next/link";
import { useRouter } from "next/navigation"; // แก้จาก next/router เป็น next/navigation
import { useState } from "react";

// ตั้งค่า Axios ให้ส่ง Cookie ไปด้วยเสมอ (Credentials)
const api = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080/api",
  withCredentials: true, 
});

type LoginFormProps = { redirectTo?: string };

export default function LoginForm({ redirectTo }: LoginFormProps) {
  const [showPassword, setShowPassword] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [isSubmitting, setIsSubmitting] = useState(false);
  const router = useRouter();

  async function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    setError(null);

    const formData = new FormData(e.currentTarget);
    // ใช้ identity แทน email เพื่อให้ตรงกับ Backend ที่เรารองรับทั้ง Username/Email
    const identity = (formData.get("identity") as string)?.trim() ?? "";
    const password = (formData.get("password") as string) ?? "";

    if (!identity || !password) {
      setError("Please fill in username/email and password.");
      return;
    }

    setIsSubmitting(true);
    try {      
      // ยิงไปที่ Endpoint /auth/login ของ Go Fiber
      const response = await api.post("/login", { identity, password });
      const data = response.data;

      console.log("Login Success:", data);

      // จัดการการเปลี่ยนหน้าตาม Role ที่ได้จาก DTO
      if (redirectTo) {
        router.push(redirectTo);
      } else {
        if (data.role_name === "Dentist") {
          router.push("/dentist/dashboard");
        } else {
          router.push("/patient/dashboard");
        }
      }
      
      router.refresh();
    } catch (err: any) {
      if (axios.isAxiosError(err) && err.response?.data?.error) {
        setError(err.response.data.error);
      } else {
        setError("Something went wrong. Please try again.");
      }
    } finally {
      setIsSubmitting(false);
    }
  }

  return (
    <form onSubmit={handleSubmit} className="flex flex-col gap-6">
      {error && (
        <p
          className="text-sm text-red-600 bg-red-50 px-3 py-2 rounded border border-red-100"
          role="alert"
        >
          {error}
        </p>
      )}

      <div className="flex flex-col gap-1">
        <label htmlFor="identity" className="text-[13px] font-bold text-gray-700">
          Username or Email
        </label>
        <input
          id="identity"
          type="text"
          name="identity"
          placeholder="Enter username or email"
          className="h-12 w-full px-4 border border-gray-200 rounded-xl text-[15px] focus:outline-none focus:border-[#8f1eae] focus:ring-1 focus:ring-[#8f1eae] transition"
          autoComplete="username"
          required
        />
      </div>

      <div className="flex flex-col gap-1">
        <label htmlFor="password" className="text-[13px] font-bold text-gray-700">
          Password
        </label>
        <div className="flex h-12 w-full items-center justify-between gap-2 rounded-xl border border-gray-200 bg-white px-4 focus-within:border-[#8f1eae] focus-within:ring-1 focus-within:ring-[#8f1eae]">
          <input
            id="password"
            type={showPassword ? "text" : "password"}
            name="password"
            placeholder="••••••••"
            className="flex-1 min-w-0 bg-transparent text-[15px] focus:outline-none"
            autoComplete="current-password"
            required
          />
          <button
            type="button"
            onClick={() => setShowPassword((p) => !p)}
            className="text-[13px] text-gray-500 hover:text-[#8f1eae] font-medium shrink-0"
          >
            {showPassword ? "Hide" : "Show"}
          </button>
        </div>
      </div>

      <button
        type="submit"
        disabled={isSubmitting}
        className="h-12 w-full rounded-xl bg-[#8f1eae] text-white font-bold text-base hover:opacity-90 transition-all shadow-lg shadow-purple-100 disabled:opacity-70 disabled:cursor-not-allowed"
      >
        {isSubmitting ? "Signing in…" : "Submit"}
      </button>

      <p className="text-center text-sm text-gray-600">
        No account?{" "}
        <Link
          className="text-[#8f1eae] font-bold hover:underline"
          href={
            !redirectTo
              ? "/register"
              : `/register?redirect=${encodeURIComponent(redirectTo)}`
          }
        >
          Create one
        </Link>
      </p>
    </form>
  );
}