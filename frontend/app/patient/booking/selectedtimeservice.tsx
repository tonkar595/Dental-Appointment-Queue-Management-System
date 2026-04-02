// 'use client';
// import React from 'react';

// const timeSlots = ["09.00", "09.15", "09.30", "09.45", "10.00", "10.15"]; // อ้างอิงจากตารางเวลา [cite: 380, 556, 657]

// export default function SelectedTimeService({ date, onDateChange, time, onTimeChange }: any) {
//   return (
//     <div className="flex-1 p-8">
//       <h2 className="text-2xl font-bold mb-2">เลือกวันที่และเวลา</h2> 
//       <p className="text-sm text-gray-500 mb-8 font-medium">เวลาแสดงใน Bangkok (GMT+07:00)</p> 
//       <div className="mb-10">
//         <h3 className="font-bold mb-4">วันพฤหัสบดีที่ 4 ธันวาคม</h3> 
//         <div className="flex gap-3 overflow-x-auto pb-2">
//           {[4, 5, 6, 7, 8, 9, 10, 11].map((d) => (
//             <button 
//               key={d}
//               onClick={() => onDateChange(d)}
//               className={`min-w-[70px] py-4 rounded-xl flex flex-col items-center transition-all ${
//                 date === d ? 'bg-[#8f1eae] text-white shadow-lg' : 'bg-white'
//               }`}
//             >
//               <span className="text-xs opacity-70">ธ.ค.</span> 
//               <span className="text-xl font-bold">{d < 10 ? `0${d}` : d}</span> 
//             </button>
//           ))}
//         </div>
//       </div>

//       <div>
//         <h3 className="font-bold mb-4">ช่วงเวลา</h3>
//         <div className="grid grid-cols-3 gap-3">
//           {timeSlots.map((t) => (
//             <button 
//               key={t}
//               onClick={() => onTimeChange(t)}
//               className={`py-3 rounded-xl font-bold border-2 transition-all ${
//                 time === t ? 'border-[#8f1eae] text-[#8f1eae] bg-purple-50' : 'border-transparent bg-white'
//               }`}
//             >
//               {t} 
//             </button>
//           ))}
//         </div>
//       </div>
//     </div>
//   );
// }

'use client';
import React, { useState, useEffect, useRef } from 'react';
import { ChevronLeft, ChevronRight } from 'lucide-react';

export default function SelectedTimeService({ date, onDateChange, time, onTimeChange }: any) {
  const scrollRef = useRef<HTMLDivElement>(null);
  
  // ตั้งค่าวันที่ปัจจุบัน
  const now = new Date();
  const currentMonth = now.toLocaleString('th-TH', { month: 'long' });
  const currentYear = now.getFullYear() + 543; // แปลงเป็น พ.ศ.

  // ฟังก์ชันสร้างรายการวันที่ทั้งหมดในเดือนปัจจุบัน
  const [daysInMonth, setDaysInMonth] = useState<{ d: number, label: string }[]>([]);

  useEffect(() => {
    const year = now.getFullYear();
    const month = now.getMonth();
    const dateCount = new Date(year, month + 1, 0).getDate();
    const daysArray = [];

    for (let i = 1; i <= dateCount; i++) {
      const dayName = new Date(year, month, i).toLocaleDateString('th-TH', { weekday: 'short' });
      daysArray.push({ d: i, label: dayName });
    }
    setDaysInMonth(daysArray);
  }, []);

  // ฟังก์ชันสำหรับการเลื่อน (Scroll)
  const scroll = (direction: 'left' | 'right') => {
    if (scrollRef.current) {
      const { scrollLeft, clientWidth } = scrollRef.current;
      const scrollTo = direction === 'left' ? scrollLeft - 200 : scrollLeft + 200;
      scrollRef.current.scrollTo({ left: scrollTo, behavior: 'smooth' });
    }
  };

  // สร้าง Time Slots 09.00 - 17.00 
  const generateTimeSlots = () => {
    const slots = [];
    for (let h = 9; h <= 17; h++) {
      for (let m = 0; m < 60; m += 15) {
        if (h === 17 && m > 0) break;
        slots.push(`${h.toString().padStart(2, '0')}.${m.toString().padStart(2, '0')}`);
      }
    }
    return slots;
  };

  return (
    <div className="flex-1 p-8 overflow-y-auto max-h-screen">
      <h2 className="text-2xl font-bold mb-2">เลือกวันที่และเวลา</h2>
      <p className="text-sm text-gray-500 mb-8 font-medium">
        เวลาแสดงใน Bangkok (GMT+07:00) [cite: 354, 631]
      </p>

      {/* Header ปฏิทินพร้อมปุ่มเลื่อน */}
      <div className="mb-6">
        <div className="flex justify-between items-center mb-4">
          <h3 className="font-bold text-gray-700 text-lg">เดือน{currentMonth} {currentYear}</h3>
          <div className="flex gap-2">
            <button onClick={() => scroll('left')} className="p-2 rounded-full bg-white shadow-sm hover:bg-gray-50 border border-gray-100">
              <ChevronLeft size={20} />
            </button>
            <button onClick={() => scroll('right')} className="p-2 rounded-full bg-white shadow-sm hover:bg-gray-50 border border-gray-100">
              <ChevronRight size={20} />
            </button>
          </div>
        </div>

        {/* Calendar Strip พร้อม ScrollRef */}
        <div 
          ref={scrollRef}
          className="flex gap-3 overflow-x-auto pb-4 scrollbar-hide no-scrollbar"
          style={{ scrollSnapType: 'x mandatory' }}
        >
          {daysInMonth.map((item) => (
            <button
              key={item.d}
              onClick={() => onDateChange(item.d)}
              className={`min-w-[80px] py-5 rounded-[1.5rem] flex flex-col items-center transition-all border ${
                date === item.d 
                  ? 'bg-[#8f1eae] text-white border-[#8f1eae] shadow-lg scale-105' 
                  : 'bg-white border-gray-100 text-gray-600 hover:border-purple-200'
              }`}
            >
              <span className={`text-xs mb-2 ${date === item.d ? 'text-purple-100' : 'text-gray-400'}`}>
                {item.label}
              </span>
              <span className="text-2xl font-black">{item.d < 10 ? `0${item.d}` : item.d}</span>
            </button>
          ))}
        </div>
      </div>

      {/* Time Grid  */}
      <div className="mb-10">
        <h3 className="font-bold mb-4 text-gray-700">ช่วงเวลาที่ว่าง</h3>
        <div className="grid grid-cols-3 sm:grid-cols-4 md:grid-cols-5 gap-3">
          {generateTimeSlots().map((t) => (
            <button
              key={t}
              onClick={() => onTimeChange(t)}
              className={`py-3 rounded-xl font-bold border-2 transition-all ${
                time === t 
                  ? 'border-[#8f1eae] text-[#8f1eae] bg-purple-50' 
                  : 'border-transparent bg-white text-gray-400 hover:bg-gray-50'
              }`}
            >
              {t}
            </button>
          ))}
        </div>
      </div>
    </div>
  );
}