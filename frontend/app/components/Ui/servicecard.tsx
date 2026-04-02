"use client";

export default function ServiceCard({ title, icon }: { title: string, icon: React.ReactNode }) {
  return (
    <div className="p-8 bg-[#eeeeee] rounded-2xl hover:bg-[#8f1eae] hover:text-white transition-all cursor-pointer group">
      <div className="w-14 h-14 bg-white rounded-xl flex items-center justify-center mb-6 text-[#8f1eae] group-hover:bg-purple-400 group-hover:text-white transition-colors">
        {icon}
      </div>
      <h3 className="text-xl font-bold mb-2">{title}</h3>
      <p className="opacity-70">บริการของทันตะกรรมที่มีคุณภาพและดูแลโดยผู้เชี่ยวชาญ [cite: 13]</p>
    </div>
  );
}

