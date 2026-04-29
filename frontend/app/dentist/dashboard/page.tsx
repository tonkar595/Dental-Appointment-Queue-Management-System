"use client"


import { ChevronDown } from "lucide-react";
import StatCardComponent from "../../components/dashboard/StatCard";
import AppointmentTable from "../../components/dashboard/AppointmentTable";
import UpcomingEvents from "../../components/dashboard/UpcomingEvent";
import { stats, appointments, upcomingEvents } from "../../data/mockData";
export default function DentistDashboard(){
    return (
    <main className="ml-[240px] pt-[64px] min-h-screen" style={{ backgroundColor: "#EEEEEE" }}>
      <div className="p-7">
        {/* TODAY label */}
        <div className="flex items-center gap-1.5 mb-5">
          <span className="text-xs font-bold text-gray-600 tracking-widest">TODAY</span>
          <ChevronDown size={14} className="text-gray-400" />
        </div>

        {/* Stats cards */}
        <div className="bg-white rounded-2xl p-5 border border-[#EEEEEE] mb-5">
          <div className="flex gap-4">
            {stats.map((stat, i) => (
              <StatCardComponent key={i} stat={stat} />
            ))}
          </div>
        </div>

        {/* Bottom row */}
        <div className="flex gap-5" style={{ minHeight: "380px" }}>
          {/* Appointment table */}
          <div className="flex-1 min-w-0">
            <AppointmentTable appointments={appointments} />
          </div>

          {/* Upcoming events */}
          <div className="w-[260px] flex-shrink-0">
            <UpcomingEvents events={upcomingEvents} />
          </div>
        </div>
      </div>
    </main>
  );
}