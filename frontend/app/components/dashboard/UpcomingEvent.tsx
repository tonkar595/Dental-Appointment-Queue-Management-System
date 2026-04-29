import { UpcomingEvent } from "../../types/model";
import { CalendarDays } from "lucide-react";

interface UpcomingEventsProps {
  events: UpcomingEvent[];
}

export default function UpcomingEvents({ events }: UpcomingEventsProps) {
  return (
    <div className="bg-white rounded-2xl border border-[#EEEEEE] p-6 flex flex-col h-full">
      <h2 className="text-base font-bold text-gray-800 mb-5">Upcoming event</h2>

      <div className="flex-1 space-y-3">
        {events.map((event) => (
          <div
            key={event.id}
            className="flex items-center gap-3 p-3 rounded-xl"
            style={{ backgroundColor: "#faf5ff" }}
          >
            <div
              className="w-9 h-9 rounded-lg flex items-center justify-center flex-shrink-0"
              style={{ backgroundColor: "#f3e8ff" }}
            >
              <CalendarDays size={16} style={{ color: "#8F1EAE" }} />
            </div>
            <div>
              <p className="text-xs font-semibold text-gray-700">{event.title}</p>
              <p className="text-[10px] text-gray-400 font-medium mt-0.5">
                {new Date(event.date).toLocaleDateString("th-TH", {
                  day: "numeric",
                  month: "short",
                })}{" "}
                • {event.time}
              </p>
            </div>
          </div>
        ))}

        {events.length === 0 && (
          <div className="flex-1 flex items-center justify-center text-gray-300 text-sm py-10">
            No upcoming events
          </div>
        )}
      </div>

      <div className="mt-4 pt-3 border-t border-[#EEEEEE] flex justify-end">
        <button className="text-xs text-gray-400 hover:text-purple-600 font-semibold transition-colors">
          view all
        </button>
      </div>
    </div>
  );
}
