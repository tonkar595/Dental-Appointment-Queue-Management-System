// import Image from "next/image";

// export default function Home() {
//   return (
//     <div >
//      home page
//     </div>
//   );
// }

import React from 'react';
import { Calendar, User, Clock, ChevronRight, LogIn } from 'lucide-react';
import ServiceCard from './components/Ui/servicecard';

export default function DentalLandingPage() {
  return (
    <div className="min-h-screen bg-[#eeeeee] text-slate-800 font-sans">
      {/* Navigation */}
      <nav className="bg-white border-b border-gray-200">
        <div className="max-w-7xl mx-auto px-6 py-4 flex justify-between items-center">
          <div className="flex items-center gap-2">
            <div className="w-10 h-10 bg-[#8f1eae] rounded-lg flex items-center justify-center">
              <span className="text-white font-bold text-xl">D</span>
            </div>
            <span className="text-xl font-bold tracking-tight text-[#8f1eae]">DentalPlus</span>
          </div>
          
          <div className="flex items-center gap-4">
            <button className="text-sm font-medium hover:text-[#8f1eae] transition">สำหรับเจ้าหน้าที่</button>
            <button className="bg-[#8f1eae] text-white px-6 py-2 rounded-lg text-sm font-medium hover:opacity-90 transition flex items-center gap-2">
              <LogIn size={22} /> เข้าสู่ระบบ
            </button>
          </div>
        </div>
      </nav>

      {/* Hero Section */}
      <header className="max-w-7xl mx-auto px-6 py-16 md:py-24 grid md:grid-cols-2 gap-12 items-center">
        <div>
          <h1 className="text-4xl md:text-6xl font-black leading-tight mb-6">
            จองคิวทำฟัน <br />
            <span className="text-[#8f1eae]">ง่ายเพียงปลายนิ้ว</span>
          </h1>
          <p className="text-lg text-gray-600 mb-8 leading-relaxed">
            ระบบจัดการคิวทันตกรรมอัจฉริยะ ตรวจสอบสถานะคิวแบบ Real-time 
            และจัดการประวัติการรักษาของคุณได้ในที่เดียว
          </p>
          <div className="flex flex-wrap gap-4">
            <button className="bg-[#8f1eae] text-white px-8 py-4 rounded-2xl font-bold text-lg shadow-lg shadow-purple-200 hover:-translate-y-1 transition-all flex items-center gap-2">
              เริ่มจองการรักษาใหม่ <ChevronRight />
            </button>
            <button className="bg-white border-2 border-gray-200 px-8 py-4 rounded-2xl font-bold text-lg hover:bg-gray-50 transition-all">
              ดูประวัติการจอง
            </button>
          </div>
        </div>
        
        {/* Quick Status Card - อ้างอิงจากสถานะคิวในหน้า Dashboard [cite: 58, 64] */}
        {/* <div className="bg-white p-8 rounded-[2rem] shadow-xl border border-gray-100">
          <div className="flex justify-between items-center mb-6">
            <h3 className="font-bold text-xl">สถานะคิวของคุณ</h3>
            <span className="bg-purple-100 text-[#8f1eae] text-xs font-bold px-3 py-1 rounded-full">LIVE</span>
          </div>
          <div className="space-y-4">
            <div className="flex items-center gap-4 p-4 bg-[#eeeeee] rounded-xl font-medium">
              <div className="w-12 h-12 bg-white rounded-lg flex items-center justify-center text-[#8f1eae]">
                <Clock />
              </div>
              <div>
                <p className="text-sm text-gray-500">รหัสคิว [cite: 63]</p>
                <p className="text-lg font-bold">DT-001</p>
              </div>
            </div>
            <div className="flex items-center gap-4 p-4 bg-[#eeeeee] rounded-xl font-medium">
              <div className="w-12 h-12 bg-white rounded-lg flex items-center justify-center text-[#8f1eae]">
                <User />
              </div>
              <div>
                <p className="text-sm text-gray-500">สถานะปัจจุบัน [cite: 64]</p>
                <p className="text-lg font-bold text-[#8f1eae]">รอเรียกคิว</p>
              </div>
            </div>
          </div>
        </div> */}
      </header>

      {/* Services Section - อ้างอิงจากหน้าบริการ [cite: 13, 203] */}
      <section className="bg-white py-20">
        <div className="max-w-7xl mx-auto px-6">
          <div className="text-center mb-16">
            <h2 className="text-3xl font-bold mb-4">บริการของเรา</h2>
            <div className="w-20 h-1 bg-[#8f1eae] mx-auto"></div>
          </div>
          
          <div className="grid md:grid-cols-3 gap-8">
            <ServiceCard title="จัดฟัน" icon={<Calendar />} />
            <ServiceCard title="อุดฟัน" icon={<User />} />
            <ServiceCard title="ตรวจสุขภาพฟัน" icon={<Clock />} />
          </div>
        </div>
      </section>

      {/* Footer */}
      {/* <footer className="py-12 bg-[#eeeeee] border-t border-gray-200 text-center">
        <p className="text-gray-500 text-sm">© 2026 Dental Management System. Built with Next.js & Tailwind CSS</p>
      </footer> */}
    </div>
  );
}

