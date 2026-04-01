// 'use client';

// import React from 'react';
// import Link from 'next/link';

// export default function RegisterForm() {
//   return (
//         <form className="space-y-5">
//           <div>
//             <label className="block text-sm font-bold text-gray-700 mb-1">Username</label>
//             <input type="text" className="w-full px-5 py-3 rounded-xl border border-gray-200 focus:border-[#8f1eae] outline-none" placeholder="Username" />
//           </div>

//           <div>
//             <label className="block text-sm font-bold text-gray-700 mb-1">Email</label>
//             <input type="email" className="w-full px-5 py-3 rounded-xl border border-gray-200 focus:border-[#8f1eae] outline-none" placeholder="Email" />
//           </div>

//           <div>
//             <label className="block text-sm font-bold text-gray-700 mb-1">Password</label>
//             <input type="password" className="w-full px-5 py-3 rounded-xl border border-gray-200 focus:border-[#8f1eae] outline-none" placeholder="Password" />
//           </div>

//           <div>
//             <label className="block text-sm font-bold text-gray-700 mb-1">Phone</label>
//             <input type="text" className="w-full px-5 py-3 rounded-xl border border-gray-200 focus:border-[#8f1eae] outline-none" placeholder="Phone number" />
//           </div>

//           <button 
//             type="submit"
//             className="w-full bg-[#8f1eae] text-white font-bold py-4 rounded-xl hover:opacity-90 transition-all shadow-lg mt-4"
//           >
//             Submit
//           </button>
//         </form>

//   );
// }

"use client";

import axios from "axios";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { useState } from "react";

// สร้าง Axios Instance (เหมือนกับ Login)
const api = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080/api",
});

export default function RegisterForm() {
  const [error, setError] = useState<string | null>(null);
  const [isSubmitting, setIsSubmitting] = useState(false);
  const router = useRouter();

  async function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    setError(null);

    const formData = new FormData(e.currentTarget);
    const user_name = (formData.get("username") as string)?.trim();
    const email = (formData.get("email") as string)?.trim();
    const password = (formData.get("password") as string);
    const phone = (formData.get("phone") as string)?.trim();

    // Validation เบื้องต้นฝั่ง Client
    if (!user_name || !email || !password || !phone) {
      setError("Please fill in all fields.");
      return;
    }

    setIsSubmitting(true);
    try {
      // ยิงไปที่ /auth/register ของ Go Fiber
      await api.post("/register", {
        user_name,
        email,
        password,
        phone,
      });

      // ถ้าสมัครสำเร็จ ให้เด้งไปหน้า Login
      alert("Registration successful! Please login.");
      router.push("/login");
      
    } catch (err: any) {
      // ดักจับ Error Message ที่เราเขียนไว้ใน Go (เช่น "email already exists")
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
    <form onSubmit={handleSubmit} className="space-y-5">
      {/* ส่วนแสดง Error */}
      {error && (
        <div className="bg-red-50 text-red-500 p-3 rounded-xl border border-red-100 text-sm font-medium">
          {error}
        </div>
      )}

      <div>
        <label className="block text-sm font-bold text-gray-700 mb-1">Username</label>
        <input 
          name="username"
          type="text" 
          className="w-full px-5 py-3 rounded-xl border border-gray-200 focus:border-[#8f1eae] focus:ring-1 focus:ring-[#8f1eae] outline-none transition" 
          placeholder="Username" 
          required
        />
      </div>

      <div>
        <label className="block text-sm font-bold text-gray-700 mb-1">Email</label>
        <input 
          name="email"
          type="email" 
          className="w-full px-5 py-3 rounded-xl border border-gray-200 focus:border-[#8f1eae] focus:ring-1 focus:ring-[#8f1eae] outline-none transition" 
          placeholder="Email" 
          required
        />
      </div>

      <div>
        <label className="block text-sm font-bold text-gray-700 mb-1">Password</label>
        <input 
          name="password"
          type="password" 
          className="w-full px-5 py-3 rounded-xl border border-gray-200 focus:border-[#8f1eae] focus:ring-1 focus:ring-[#8f1eae] outline-none transition" 
          placeholder="Password" 
          required
        />
      </div>

      <div>
        <label className="block text-sm font-bold text-gray-700 mb-1">Phone</label>
        <input 
          name="phone"
          type="text" 
          className="w-full px-5 py-3 rounded-xl border border-gray-200 focus:border-[#8f1eae] focus:ring-1 focus:ring-[#8f1eae] outline-none transition" 
          placeholder="Phone number" 
          required
        />
      </div>

      <button 
        type="submit"
        disabled={isSubmitting}
        className="w-full bg-[#8f1eae] text-white font-bold py-4 rounded-xl hover:opacity-90 transition-all shadow-lg mt-4 disabled:opacity-70 disabled:cursor-not-allowed"
      >
        {isSubmitting ? "Creating account..." : "Submit"}
      </button>

      <p className="text-center text-sm text-gray-600 mt-4">
        Already have an account?{" "}
        <Link href="/login" className="text-[#8f1eae] font-bold hover:underline">
          Login here
        </Link>
      </p>
    </form>
  );
}