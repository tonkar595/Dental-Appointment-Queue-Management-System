import { Appointment, StatCard, UpcomingEvent } from "@/app/types/model";

export const stats: StatCard[] = [
  {
    title: "นัดหมายวันนี้",
    value: 24,
    change: 12,
    icon: "calendar",
    subStats: [
      { label: "เสร็จแล้ว", value: 14 },
      { label: "รอดำเนินการ", value: 7 },
      { label: "ยกเลิก", value: 3 },
    ],
  },
  {
    title: "ผู้ป่วยใหม่วันนี้",
    value: 8,
    change: 5,
    icon: "users",
    subStats: [
      { label: "ผู้ป่วยเก่า", value: 16 },
      { label: "Walk-in", value: 3 },
    ],
  },
  {
    title: "ยกเลิก / No-show วันนี้",
    value: 3,
    change: 8,
    icon: "x-circle",
    subStats: [
      { label: "ยกเลิก", value: 2 },
      { label: "No-show", value: 1 },
    ],
  },
];

export const appointments: Appointment[] = [
  {
    id: "1safsfaa",
    patientName: "tonkar apiwat",
    gender: "Male",
    reason: "meet dentist",
    status: "pending",
  },
  {
    id: "2brtghyx",
    patientName: "Somchai Jaidee",
    gender: "Male",
    reason: "teeth cleaning",
    status: "confirmed",
  },
  {
    id: "3cvuiopq",
    patientName: "Napa Wongsa",
    gender: "Female",
    reason: "tooth extraction",
    status: "pending",
  },
];

export const upcomingEvents: UpcomingEvent[] = [
  {
    id: "1",
    title: "Staff Training",
    date: "2024-12-20",
    time: "09:00",
  },
  {
    id: "2",
    title: "Equipment Maintenance",
    date: "2024-12-22",
    time: "14:00",
  },
];
