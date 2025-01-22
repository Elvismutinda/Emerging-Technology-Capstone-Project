"use client";

import { DateRangePicker } from "@/components/ui/date-range-picker";
import { MAX_DATE_RANGE_DAYS } from "@/config/constants";
import { differenceInDays, startOfMonth } from "date-fns";
import React, { useState } from "react";
import { toast } from "sonner";
import TransactionTable from "./_components/TransactionTable";

function TransactionsPage() {
  return (
    <>
      <div className="border-b bg-card">
        <div className="container flex items-center py-8">
          <div>
            <p className="text-3xl font-bold">Transactions history</p>
          </div>
        </div>
      </div>
      <div className="container">
        <TransactionTable />
      </div>
    </>
  );
}

export default TransactionsPage;
