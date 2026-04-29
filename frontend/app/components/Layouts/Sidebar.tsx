"use client";

import Link from "next/link";
import {
  LayoutDashboard,
  Stethoscope,
  CalendarDays,
  CalendarCheck2,
  Users,
  BarChart2,
  Settings,
  HelpCircle,
  LogOut,
  Box,
} from "lucide-react";

interface NavItem {
  label: string;
  icon: React.ReactNode;
  href: string;
}

const navItems: NavItem[] = [
  { label: "OVERVIEW", icon: <LayoutDashboard size={16} />, href: "/dentist/dashboard" },
  { label: "SERVICE", icon: <Stethoscope size={16} />, href: "/dentist/service" },
  { label: "APPOINTMENT", icon: <CalendarDays size={16} />, href: "/dentist/appointment" },
  { label: "EVENT", icon: <CalendarCheck2 size={16} />, href: "/dentist/event" },
  { label: "PATIENTS", icon: <Users size={16} />, href: "/dentist/patients" },
  { label: "ANALYTICS", icon: <BarChart2 size={16} />, href: "/dentist/analytics" },
  { label: "SETTINGS", icon: <Settings size={16} />, href: "/dentist/settings" },
];

const bottomItems: NavItem[] = [
  { label: "HELP CENTER", icon: <HelpCircle size={16} />, href: "/dentist/help" },
  { label: "LOG OUT", icon: <LogOut size={16} />, href: "/dentist/logout" },
];

interface SidebarProps {
  activeItem?: string;
}

export default function Sidebar({ activeItem = "OVERVIEW" }: SidebarProps) {
  return (
    <aside className="w-[240px] min-h-screen flex flex-col fixed left-0 top-0 z-30" style={{ backgroundColor: "#1a0a2e" }}>
      <div className="flex items-center gap-3 px-6 py-5">
        <div className="w-9 h-9 rounded-lg flex items-center justify-center border border-purple-400/40" style={{ backgroundColor: "#8F1EAE" }}>
          <Box size={18} className="text-white" />
        </div>
        <span className="text-sm font-bold text-white tracking-widest">UNTITLE CLINIC</span>
      </div>

      <nav className="flex-1 px-3 py-2">
        {navItems.map((item) => {
          const isActive = activeItem === item.label;
          return (
            <Link
              key={item.label}
              href={item.href}
              className={`flex items-center gap-3 px-4 py-3 rounded-xl mb-0.5 text-[11px] font-bold tracking-widest transition-all ${
                isActive ? "text-white" : "text-gray-400 hover:text-white hover:bg-white/5"
              }`}
              style={isActive ? { backgroundColor: "#8F1EAE" } : {}}
            >
              <span
                className="flex-shrink-0 p-1.5 rounded-md"
                style={{ backgroundColor: isActive ? "rgba(255,255,255,0.15)" : "rgba(239,68,68,0.12)" }}
              >
                <span className={isActive ? "text-white" : "text-red-400"}>{item.icon}</span>
              </span>
              {item.label}
            </Link>
          );
        })}
      </nav>

      <div className="px-3 pb-3">
        {bottomItems.map((item) => (
          <Link
            key={item.label}
            href={item.href}
            className="flex items-center gap-3 px-4 py-3 rounded-xl mb-0.5 text-[11px] font-bold tracking-widest text-gray-400 hover:text-white hover:bg-white/5 transition-all"
          >
            <span className="flex-shrink-0 p-1.5 rounded-md" style={{ backgroundColor: "rgba(239,68,68,0.12)" }}>
              <span className="text-red-400">{item.icon}</span>
            </span>
            {item.label}
          </Link>
        ))}
      </div>

      <div className="border-t border-white/10 px-5 py-4 flex items-center gap-3">
        <div className="w-10 h-10 rounded-full overflow-hidden flex-shrink-0 ring-2 ring-yellow-400/40">
          <div className="w-full h-full bg-gradient-to-br from-amber-400 to-orange-500 flex items-center justify-center text-white text-xs font-bold">
            JD
          </div>
        </div>
        <div>
          <p className="text-[11px] font-bold text-white tracking-wide">JOHN DOE</p>
          <p className="text-[10px] text-gray-400">D. IN MEDICINE</p>
        </div>
        <div className="ml-auto w-2 h-2 rounded-full bg-emerald-400 shadow-lg shadow-emerald-400/50" />
      </div>
    </aside>
  );
}
