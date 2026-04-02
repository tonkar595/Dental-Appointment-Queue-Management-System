'use client';
import React from 'react';

export default function ServiceSidebar({ service, date, time, onNext }: any) {
  return (
    <div className="w-full md:w-[400px] bg-white border-l border-gray-200 p-8 flex flex-col justify-between">
      <div className="space-y-8">
        <div>
          <h3 className="text-lg font-bold text-gray-400 mb-4 uppercase tracking-wider">รายละเอียดบริการ</h3> 
          <div className="p-4 bg-[#eeeeee] rounded-xl">
            <p className="font-bold text-[#8f1eae]">{service || 'ยังไม่ได้เลือกบริการ'}</p> 
          </div>
        </div>

        {date && time && (
          <div>
            <h3 className="text-lg font-bold text-gray-400 mb-4 uppercase tracking-wider">รายละเอียดเวลา</h3> 
            <div className="p-4 bg-[#eeeeee] rounded-xl space-y-1">
              <p className="font-bold">{date} ธันวาคม 2568</p> 
              <p className="text-[#8f1eae] font-bold">เวลา {time} น.</p> 
            </div>
          </div>
        )}
      </div>

      <button 
        onClick={onNext}
        disabled={!service}
        className="w-full bg-[#8f1eae] text-white py-4 rounded-2xl font-bold text-lg shadow-lg hover:opacity-90 disabled:bg-gray-300 disabled:shadow-none transition-all mt-8"
      >
        ถัดไป 
      </button>
    </div>
  );
}