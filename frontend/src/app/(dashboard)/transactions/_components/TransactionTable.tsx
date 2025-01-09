"use client";

import React, { useState } from "react";
import {
  ColumnDef,
  flexRender,
  getCoreRowModel,
  getPaginationRowModel,
  getSortedRowModel,
  SortingState,
  useReactTable,
} from "@tanstack/react-table";

import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { useQuery } from "@tanstack/react-query";
import SkeletonWrapper from "@/components/SkeletonWrapper";
import { DataTableColumnHeader } from "@/components/datatable/ColumnHeader";
import { cn } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import axios from "axios";
import useCurrentUser from "@/hooks/use-current-user";

interface Props {
  from: Date;
  to: Date;
}

const emptyData: any[] = [];

export type TransactionHistoryRow = {
  id: string;
  createdAt: Date;
  updatedAt: Date;
  amount: number;
  description: string;
  transaction_date: Date;
  userId: string;
  type: string;
  category_id: string;
};

function TransactionTable({ from, to }: Props) {
  const { token, user } = useCurrentUser();
  const userId = user?.id;

  const [sorting, setSorting] = useState<SortingState>([]);
  const history = useQuery({
    queryKey: ["transactions", "history", from, to],
    queryFn: async () => {
      const response = await axios.get(
        "http://localhost:8000/transaction/get-all",
        {
          params: {
            from: from.toUTCString(),
            to: to.toUTCString(),
          },
          headers: {
            Authorization: `Bearer ${token}`,
            userId: userId,
          },
        }
      );
      return response.data.data;
    },
  });

  const { data: categories = [] } = useQuery({
    queryKey: ["categories"],
    queryFn: async () => {
      const response = await axios.get("http://localhost:8000/category", {
        headers: {
          Authorization: `Bearer ${token}`,
          userId: userId,
        },
      });
      return response.data.data;
    },
  });

  const columns: ColumnDef<TransactionHistoryRow>[] = [
    {
      accessorKey: "category",
      header: ({ column }) => (
        <DataTableColumnHeader column={column} title="Category" />
      ),
      cell: ({ row }) => {
        const category = categories.find(
          (cat: { id: string; name: string }) => cat.id === row.original.category_id
        );
        const categoryName = category?.name || "Unknown";
        return <div className="capitalize">{categoryName}</div>;
      },
    },
    {
      accessorKey: "description",
      header: ({ column }) => (
        <DataTableColumnHeader column={column} title="Description" />
      ),
      cell: ({ row }) => (
        <div className="capitalize">{row.original.description}</div>
      ),
    },
    {
      accessorKey: "transaction_date",
      header: "Transaction Date",
      cell: ({ row }) => {
        const date = new Date(row.original.transaction_date);
        const formattedDate = date.toLocaleDateString("default", {
          timeZone: "UTC",
          year: "numeric",
          month: "2-digit",
          day: "2-digit",
        });
        return <div className="text-muted-foreground">{formattedDate}</div>;
      },
    },
    {
      accessorKey: "type",
      header: ({ column }) => (
        <DataTableColumnHeader column={column} title="Type" />
      ),
      cell: ({ row }) => (
        <div
          className={cn(
            "capitalize rounded-lg text-center p-2",
            row.original.type === "INCOME" &&
              "bg-emerald-400/10 text-emerald-500",
            row.original.type === "EXPENSE" && "bg-rose-400/10 text-rose-500"
          )}
        >
          {row.original.type}
        </div>
      ),
    },
    {
      accessorKey: "amount",
      header: ({ column }) => (
        <DataTableColumnHeader column={column} title="Amount" />
      ),
      cell: ({ row }) => (
        <p className="text-md rounded-lg bg-gray-400/5 p-2 text-center font-medium">
          {row.original.amount}
        </p>
      ),
    },
  ];

  const table = useReactTable({
    data: history.data || emptyData,
    columns,
    getCoreRowModel: getCoreRowModel(),
    state: {
      sorting,
    },
    onSortingChange: setSorting,
    getSortedRowModel: getSortedRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
  });

  return (
    <div className="w-full py-4">
      <SkeletonWrapper isLoading={history.isLoading}>
        <div className="rounded-md border">
          <Table>
            <TableHeader>
              {table.getHeaderGroups().map((headerGroup) => (
                <TableRow key={headerGroup.id}>
                  {headerGroup.headers.map((header) => {
                    return (
                      <TableHead key={header.id}>
                        {header.isPlaceholder
                          ? null
                          : flexRender(
                              header.column.columnDef.header,
                              header.getContext()
                            )}
                      </TableHead>
                    );
                  })}
                </TableRow>
              ))}
            </TableHeader>
            <TableBody>
              {table.getRowModel().rows?.length ? (
                table.getRowModel().rows.map((row) => (
                  <TableRow
                    key={row.id}
                    data-state={row.getIsSelected() && "selected"}
                  >
                    {row.getVisibleCells().map((cell) => (
                      <TableCell key={cell.id}>
                        {flexRender(
                          cell.column.columnDef.cell,
                          cell.getContext()
                        )}
                      </TableCell>
                    ))}
                  </TableRow>
                ))
              ) : (
                <TableRow>
                  <TableCell
                    colSpan={columns.length}
                    className="h-24 text-center"
                  >
                    No results.
                  </TableCell>
                </TableRow>
              )}
            </TableBody>
          </Table>
        </div>
        <div className="flex items-center justify-end space-x-2 py-4">
          <Button
            variant="outline"
            size="sm"
            onClick={() => table.previousPage()}
            disabled={!table.getCanPreviousPage()}
          >
            Previous
          </Button>
          <Button
            variant="outline"
            size="sm"
            onClick={() => table.nextPage()}
            disabled={!table.getCanNextPage()}
          >
            Next
          </Button>
        </div>
      </SkeletonWrapper>
    </div>
  );
}

export default TransactionTable;
