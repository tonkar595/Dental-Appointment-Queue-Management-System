

type Appointment = {
  id: string;
  patientName: string;
  date: string;  
  time: string;  
  status: "BOOKED" | "CANCELLED" | "COMPLETED";
};

async function getAppointments(): Promise<Appointment[]> {
  const res = await fetch(`${process.env.API_BASE_URL}/appointments`, { cache: "no-store" });
  if (!res.ok) throw new Error("Failed to fetch appointments");


  return (await res.json()) as Appointment[];
}

export default async function AppointmentsPage() {
  const appointments = await getAppointments();

  return (
    <main>
      <h1>รายการนัดหมาย</h1>
      <ul>
        {appointments.map((a) => (
          <li key={a.id}>
            {a.patientName} — {a.date} {a.time} ({a.status})
          </li>
        ))}
      </ul>
    </main>
  );
}