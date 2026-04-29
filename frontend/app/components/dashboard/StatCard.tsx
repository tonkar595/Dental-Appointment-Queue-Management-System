import { CalendarDays, Users, XCircle, TrendingUp } from "lucide-react";
import { StatCard } from "../../types/model";

const iconMap: Record<string, React.ReactNode> = {
  calendar: <CalendarDays size={20} style={{ color: "#8F1EAE" }} />,
  users: <Users size={20} style={{ color: "#8F1EAE" }} />,
  "x-circle": <XCircle size={20} style={{ color: "#8F1EAE" }} />,
};

interface StatCardProps {
  stat: StatCard;
}

export default function StatCardComponent({ stat }: StatCardProps) {
  return (
    <div className="bg-white rounded-2xl p-5 flex-1 min-w-0 border border-[#EEEEEE]">
      {/* Header row */}
      <div className="flex items-start justify-between mb-4">
        <div className="w-10 h-10 rounded-xl flex items-center justify-center" style={{ backgroundColor: "#f3e8ff" }}>
          {iconMap[stat.icon]}
        </div>
        <span className="flex items-center gap-1 text-[11px] font-bold text-emerald-600 bg-emerald-50 px-2.5 py-1 rounded-full">
          <TrendingUp size={10} />
          +{stat.change}%
        </span>
      </div>

      {/* Number */}
      <div className="text-4xl font-black mb-1" style={{ color: "#1a0a2e" }}>
        {stat.value}
      </div>
      <p className="text-sm font-semibold mb-5" style={{ color: "#8F1EAE" }}>
        {stat.title}
      </p>

      {/* Sub stats */}
      <div className="flex gap-6 pt-3 border-t border-[#EEEEEE]">
        {stat.subStats.map((sub) => (
          <div key={sub.label}>
            <p className="text-[10px] text-gray-400 font-medium mb-0.5">{sub.label}</p>
            <p className="text-base font-black" style={{ color: "#1a0a2e" }}>{sub.value}</p>
          </div>
        ))}
      </div>
    </div>
  );
}
