"use client";
import React from "react";
import { Plus, Check } from "lucide-react";
import Link from "next/link";

const services = ["บริการการจัดฟัน", "บริการการอุดฟัน", "ตรวจสุขภาพฟัน"]; // อ้างอิงจากบริการที่มี [cite: 535, 536]

export default function SelectService({
  selected,
  onSelect,
}: {
  selected: string;
  onSelect: (s: string) => void;
}) {
  return (
    <div className="flex-1 p-8">
      <h2 className="text-2xl font-bold mb-6">บริการที่เลือก</h2> 
      <div className="space-y-4">
        {services.map((service) => (
          <div
            key={service}
            onClick={() => onSelect(service)}
            className={`flex justify-between items-center p-6 rounded-2xl cursor-pointer transition-all border-2 ${
              selected === service
                ? "border-[#8f1eae] bg-white"
                : "border-transparent bg-white"
            }`}
          >
            <span className="text-lg font-medium">{service}</span> 
            <div
              className={`w-8 h-8 rounded-full flex items-center justify-center ${
                selected === service
                  ? "bg-[#8f1eae] text-white"
                  : "bg-gray-200 text-gray-400"
              }`}
            >
              {selected === service ? <Check size={20} /> : <Plus size={20} />}{" "}
              
            </div>
          </div>
        ))}
      </div>
      <Link href={"/patient/dashboard"}>
        <button className="mt-8 text-gray-500 font-medium hover:underline">
          ยังไม่ต้องการจองตอนนี้
        </button>
      </Link>
    </div>
  );
}
