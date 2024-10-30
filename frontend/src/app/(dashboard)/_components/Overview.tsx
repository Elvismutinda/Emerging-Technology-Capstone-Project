"use client";

import { DateRangePicker } from "@/components/ui/date-range-picker";
import { MAX_DATE_RANGE_DAYS } from "@/config/constants";
import { toast } from "sonner";
import { differenceInDays, startOfMonth } from "date-fns";
import React, { useState } from "react";
import StatCards from "./StatCards";
import CategoriesStats from "./CategoriesStats";

function Overview() {
  const [dateRange, setDateRange] = useState<{ from: Date; to: Date }>({
    from: startOfMonth(new Date()),
    to: new Date(),
  });
  return (
    <>
      <div className="container flex flex-wrap items-center justify-between gap-2 py-6">
        <h2 className="text-3xl font-bold">Overview</h2>
        <div className="flex items-center gap-3">
          <DateRangePicker
            initialDateFrom={dateRange.from}
            initialDateTo={dateRange.to}
            showCompare={false}
            onUpdate={(values) => {
              const { from, to } = values.range;
              // update the date range only if both dates are set

              if (!from || !to) return;
              if (differenceInDays(to, from) > MAX_DATE_RANGE_DAYS) {
                toast.error(
                  `Date range cannot exceed ${MAX_DATE_RANGE_DAYS} days`
                );
                return;
              }

              setDateRange({ from, to });
            }}
          />
        </div>
      </div>
      <div className="container flex flex-col w-full gap-2">
        <StatCards from={dateRange.from} to={dateRange.to} />

        <CategoriesStats from={dateRange.from} to={dateRange.to} />
      </div>
    </>
  );
}

export default Overview;
