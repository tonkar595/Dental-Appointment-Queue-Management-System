'use client';

import React from 'react';
import Link from 'next/link';

export default function RegisterForm() {
  return (
        <form className="space-y-5">
          <div>
            <label className="block text-sm font-bold text-gray-700 mb-1">Username</label>
            <input type="text" className="w-full px-5 py-3 rounded-xl border border-gray-200 focus:border-[#8f1eae] outline-none" placeholder="Username" />
          </div>

          <div>
            <label className="block text-sm font-bold text-gray-700 mb-1">Email</label>
            <input type="email" className="w-full px-5 py-3 rounded-xl border border-gray-200 focus:border-[#8f1eae] outline-none" placeholder="Email" />
          </div>

          <div>
            <label className="block text-sm font-bold text-gray-700 mb-1">Password</label>
            <input type="password" className="w-full px-5 py-3 rounded-xl border border-gray-200 focus:border-[#8f1eae] outline-none" placeholder="Password" />
          </div>

          <div>
            <label className="block text-sm font-bold text-gray-700 mb-1">Phone</label>
            <input type="text" className="w-full px-5 py-3 rounded-xl border border-gray-200 focus:border-[#8f1eae] outline-none" placeholder="Phone number" />
          </div>

          <button 
            type="submit"
            className="w-full bg-[#8f1eae] text-white font-bold py-4 rounded-xl hover:opacity-90 transition-all shadow-lg mt-4"
          >
            Submit
          </button>
        </form>

  );
}