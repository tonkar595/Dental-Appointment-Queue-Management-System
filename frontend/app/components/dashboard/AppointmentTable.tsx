"use client";

import {
  useReactTable,
  getCoreRowModel,
  getSortedRowModel,
  getFilteredRowModel,
  flexRender,
  createColumnHelper,
  SortingState,
  getPaginationRowModel,
} from "@tanstack/react-table";
import { useState } from "react";
import { Plus, ChevronDown, ChevronUp, ChevronsUpDown } from "lucide-react";
import { Appointment } from "../../types/model";

const statusConfig: Record<string, { bg: string; text: string }> = {
  pending: { bg: "#8F1EAE", text: "#ffffff" },
  confirmed: { bg: "#059669", text: "#ffffff" },
  cancelled: { bg: "#ef4444", text: "#ffffff" },
  "no-show": { bg: "#6b7280", text: "#ffffff" },
};

const columnHelper = createColumnHelper<Appointment>();

interface AppointmentTableProps {
  appointments: Appointment[];
}

export default function AppointmentTable({
  appointments,
}: AppointmentTableProps) {
  const [sorting, setSorting] = useState<SortingState>([]);
  const [globalFilter, setGlobalFilter] = useState("");

  const columns = [
    columnHelper.accessor("id", {
      header: "No.",
      cell: (info) => (
        <span className="text-xs text-gray-400 font-mono">
          {info.getValue()}
        </span>
      ),
    }),
    columnHelper.accessor("patientName", {
      header: "Pation name",
      cell: (info) => (
        <span className="text-xs font-semibold text-gray-700">
          {info.getValue()}
        </span>
      ),
    }),
    columnHelper.accessor("gender", {
      header: "Gender",
      cell: (info) => (
        <span className="text-xs text-gray-600">{info.getValue()}</span>
      ),
    }),
    columnHelper.accessor("reason", {
      header: "Reason",
      cell: (info) => (
        <span className="text-xs text-gray-600">{info.getValue()}</span>
      ),
    }),
    columnHelper.accessor("status", {
      header: "Action",
      cell: (info) => {
        const cfg = statusConfig[info.getValue()] ?? statusConfig.pending;
        return (
          <span
            className="text-[11px] font-bold px-4 py-1.5 rounded-lg inline-block"
            style={{ backgroundColor: cfg.bg, color: cfg.text }}
          >
            {info.getValue()}
          </span>
        );
      },
    }),
  ];

  const table = useReactTable({
    data: appointments,
    columns,
    state: { sorting, globalFilter },
    getPaginationRowModel: getPaginationRowModel(),
    initialState: {
      pagination: { pageSize: 5, pageIndex: 0 },
    },
    onSortingChange: setSorting,
    onGlobalFilterChange: setGlobalFilter,
    getCoreRowModel: getCoreRowModel(),
    getSortedRowModel: getSortedRowModel(),
    getFilteredRowModel: getFilteredRowModel(),
  });

  return (
    <div className="bg-white rounded-2xl border border-[#EEEEEE] p-6 h-full flex flex-col">
      {/* Header */}
      <div className="flex items-center justify-between mb-5">
        <h2 className="text-base font-bold" style={{ color: "#8F1EAE" }}>
          appointment list
        </h2>
        <div className="flex items-center gap-3">
          <button className="text-xs text-gray-400 hover:text-purple-600 font-semibold transition-colors">
            view all
          </button>
          <button
            className="flex items-center gap-1.5 text-white text-xs font-bold px-4 py-2 rounded-lg transition-all hover:opacity-90 active:scale-95"
            style={{ backgroundColor: "#8F1EAE" }}
          >
            <Plus size={12} />
            APPOINTMENT
          </button>
        </div>
      </div>

      {/* Table */}
      <div className="flex-1 overflow-auto">
        <table className="w-full">
          <thead>
            {table.getHeaderGroups().map((headerGroup) => (
              <tr key={headerGroup.id} className="border-b border-[#EEEEEE]">
                {headerGroup.headers.map((header) => (
                  <th
                    key={header.id}
                    className="text-left py-3 px-3 text-[11px] font-bold text-gray-500 tracking-wide cursor-pointer select-none whitespace-nowrap"
                    onClick={header.column.getToggleSortingHandler()}
                  >
                    <span className="flex items-center gap-1">
                      {flexRender(
                        header.column.columnDef.header,
                        header.getContext(),
                      )}
                      {header.column.getCanSort() && (
                        <span className="text-gray-300">
                          {header.column.getIsSorted() === "asc" ? (
                            <ChevronUp size={12} style={{ color: "#8F1EAE" }} />
                          ) : header.column.getIsSorted() === "desc" ? (
                            <ChevronDown
                              size={12}
                              style={{ color: "#8F1EAE" }}
                            />
                          ) : (
                            <ChevronsUpDown size={12} />
                          )}
                        </span>
                      )}
                    </span>
                  </th>
                ))}
              </tr>
            ))}
          </thead>
          <tbody>
            {table.getRowModel().rows.map((row) => (
              <tr
                key={row.id}
                className="border-b border-[#EEEEEE] hover:bg-[#faf5ff] transition-colors"
              >
                {row.getVisibleCells().map((cell) => (
                  <td key={cell.id} className="py-3.5 px-3">
                    {flexRender(cell.column.columnDef.cell, cell.getContext())}
                  </td>
                ))}
              </tr>
            ))}
            {table.getRowModel().rows.length === 0 && (
              <tr>
                <td
                  colSpan={5}
                  className="py-12 text-center text-sm text-gray-300"
                >
                  No appointments found
                </td>
              </tr>
            )}
          </tbody>
        </table>
      </div>
    </div>
  );
}
