import Sidebar from "@/app/components/layouts/Sidebar";
import Header from "@/app/components/layouts/Header";

export default function DentistLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div style={{ backgroundColor: "#EEEEEE" }} className="min-h-screen">
      <Sidebar />
      <Header />
      {children}
    </div>
  );
}
