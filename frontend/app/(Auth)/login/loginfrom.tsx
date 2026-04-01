'use client';

import React from 'react';
import Link from 'next/link';

export default function LoginForm() {
  return (
    <form className="space-y-6">
          <div>
            <label className="block text-sm font-bold text-gray-700 mb-2">Username or Email</label>
            <input 
              type="text" 
              className="w-full px-5 py-3 rounded-xl border border-gray-200 focus:border-[#8f1eae] focus:ring-1 focus:ring-[#8f1eae] outline-none transition"
              placeholder="กรอกชื่อผู้ใช้งาน หรือ อีเมล"
            />
          </div>

          <div>
            <label className="block text-sm font-bold text-gray-700 mb-2">Password</label>
            <input 
              type="password" 
              className="w-full px-5 py-3 rounded-xl border border-gray-200 focus:border-[#8f1eae] focus:ring-1 focus:ring-[#8f1eae] outline-none transition"
              placeholder="กรอกรหัสผ่าน"
            />
          </div>

          <button 
            type="submit"
            className="w-full bg-[#8f1eae] text-white font-bold py-4 rounded-xl hover:opacity-90 transition-all shadow-lg shadow-purple-100"
          >
            Submit
          </button>
        </form>
  );
}