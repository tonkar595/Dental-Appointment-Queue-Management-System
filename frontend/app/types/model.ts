export interface Appointment {
  id: string;
  patientName: string;
  gender: "Male" | "Female" | "Other";
  reason: string;
  status: "pending" | "confirmed" | "cancelled" | "no-show";
}

export interface StatCard {
  title: string;
  value: number;
  change: number;
  subStats: { label: string; value: number }[];
  icon: string;
}

export interface NavItem {
  label: string;
  icon: string;
  href: string;
  active?: boolean;
}

export interface UpcomingEvent {
  id: string;
  title: string;
  date: string;
  time: string;
}
