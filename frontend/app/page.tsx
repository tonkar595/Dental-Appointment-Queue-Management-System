// import Image from "next/image";

// export default function Home() {
//   return (
//     <div >
//      home page
//     </div>
//   );
// }

import React from 'react';
import { ArrowRight, Code2, Layout, Zap } from 'lucide-react';

export default function HomePage() {
  return (
    <div className="min-h-screen bg-white text-slate-900">
      {/* Navigation */}
      <nav className="flex items-center justify-between px-8 py-6 max-w-7xl mx-auto">
        <div className="text-2xl font-bold tracking-tight text-blue-600">DevStudio</div>
        <div className="hidden md:flex space-x-8 font-medium">
          <a href="#" className="hover:text-blue-600 transition">Solutions</a>
          <a href="#" className="hover:text-blue-600 transition">Showcase</a>
          <a href="#" className="hover:text-blue-600 transition">About us</a>
        </div>
        <button className="bg-slate-900 text-white px-5 py-2 rounded-full font-medium hover:bg-slate-800 transition">
          Contact Us
        </button>
      </nav>

      {/* Hero Section */}
      <header className="px-8 py-20 max-w-7xl mx-auto text-center md:py-32">
        <h1 className="text-5xl md:text-7xl font-extrabold tracking-tight mb-6">
          Build faster with <span className="text-blue-600">Modern Stack</span>
        </h1>
        <p className="text-lg md:text-xl text-slate-600 mb-10 max-w-2xl mx-auto">
          สัมผัสประสบการณ์การพัฒนาเว็บไซต์ที่รวดเร็วและมีประสิทธิภาพสูงสุด 
          ด้วยโครงสร้างที่ออกแบบมาเพื่อรองรับการขยายตัวในอนาคต
        </p>
        <div className="flex flex-col md:flex-row justify-center gap-4">
          <button className="bg-blue-600 text-white px-8 py-4 rounded-xl font-semibold flex items-center justify-center gap-2 hover:bg-blue-700 transition shadow-lg shadow-blue-200">
            Get Started <ArrowRight size={20} />
          </button>
          <button className="border border-slate-200 px-8 py-4 rounded-xl font-semibold hover:bg-slate-50 transition">
            View Documentation
          </button>
        </div>
      </header>

      {/* Features Section */}
      <section className="bg-slate-50 py-24 px-8">
        <div className="max-w-7xl mx-auto">
          <div className="grid md:grid-cols-3 gap-12">
            <FeatureCard 
              icon={<Zap className="text-blue-600" />}
              title="High Performance"
              description="รีโหลดหน้าเว็บได้อย่างรวดเร็วด้วยการทำ Server-side Rendering และ Optimization ที่ยอดเยี่ยม"
            />
            <FeatureCard 
              icon={<Layout className="text-blue-600" />}
              title="Responsive Design"
              description="แสดงผลได้อย่างสวยงามในทุกหน้าจอ ไม่ว่าจะเป็นมือถือ แท็บเล็ต หรือเดสก์ท็อป"
            />
            <FeatureCard 
              icon={<Code2 className="text-blue-600" />}
              title="Clean Code"
              description="เขียนโค้ดได้ง่ายและเป็นระเบียบด้วย TypeScript พร้อมการจัดการสไตล์ด้วย Tailwind CSS"
            />
          </div>
        </div>
      </section>

      {/* Footer */}
      <footer className="py-12 text-center border-t border-slate-100 text-slate-500 text-sm">
        <p>© 2026 DevStudio Inc. All rights reserved.</p>
      </footer>
    </div>
  );
}

// Sub-component สำหรับการแสดง Feature
function FeatureCard({ icon, title, description }: { icon: React.ReactNode, title: string, description: string }) {
  return (
    <div className="bg-white p-8 rounded-2xl border border-slate-100 hover:border-blue-200 transition-all hover:shadow-xl group">
      <div className="w-12 h-12 bg-blue-50 rounded-lg flex items-center justify-center mb-6 group-hover:scale-110 transition-transform">
        {icon}
      </div>
      <h3 className="text-xl font-bold mb-3">{title}</h3>
      <p className="text-slate-600 leading-relaxed">{description}</p>
    </div>
  );
}