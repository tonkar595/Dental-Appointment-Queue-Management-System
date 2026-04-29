"use client"

import { Bell, Search } from "lucide-react";

export default function Header(){
    return (
    <header className="fixed top-0 left-[240px] right-0 h-[64px] bg-white border-b border-[#EEEEEE] flex items-center px-8 gap-4 z-20">
      <div className="flex-1 max-w-[420px]">
        <div className="flex items-center gap-2 bg-white border border-[#EEEEEE] rounded-lg px-4 py-2.5 shadow-sm">
          <Search size={14} style={{ color: "#8F1EAE" }} />
          <input
            type="text"
            placeholder="SEARCH..."
            className="flex-1 text-xs font-semibold text-gray-400 placeholder-gray-400 bg-transparent outline-none tracking-widest"
          />
        </div>
      </div>

      <div className="flex-1" />

      <button className="w-9 h-9 bg-white border border-[#EEEEEE] rounded-lg flex items-center justify-center hover:border-purple-300 transition-colors relative shadow-sm">
        <Bell size={16} style={{ color: "#8F1EAE" }} />
      </button>
    </header>
  );
}