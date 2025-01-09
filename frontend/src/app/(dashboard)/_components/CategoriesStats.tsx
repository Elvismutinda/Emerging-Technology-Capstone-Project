"use client";

import SkeletonWrapper from "@/components/SkeletonWrapper";
import { Card, CardHeader, CardTitle } from "@/components/ui/card";
import { Progress } from "@/components/ui/progress";
import { ScrollArea } from "@/components/ui/scroll-area";
import { TransactionType } from "@/types";
import { useQuery } from "@tanstack/react-query";
import axios from "axios";
import React from "react";

interface Props {
  from: Date;
  to: Date;
}

function CategoriesStats({ from, to }: Props) {
  const statsQuery = useQuery({
    queryKey: ["overview", "categories", from, to],
    queryFn: async () => {
      const response = await axios.get(
        `http://localhost:8000/transaction/get-category-stats`,
        {
          params: {
            from: from.toUTCString(),
            to: to.toUTCString(),
          },
        }
      );
      return response.data;
    },
  });

  return (
    <div className="flex w-full flex-wrap gap-2 md:flex-nowrap">
      <SkeletonWrapper isLoading={false}>
        <CategoriesCard
          currency="Ksh"
          type="INCOME"
          data={statsQuery.data || []}
        />
      </SkeletonWrapper>

      <SkeletonWrapper isLoading={false}>
        <CategoriesCard
          currency="Ksh"
          type="EXPENSE"
          data={statsQuery.data || []}
        />
      </SkeletonWrapper>
    </div>
  );
}

export default CategoriesStats;

function CategoriesCard({
  currency,
  type,
  data,
}: {
  currency: string;
  type: TransactionType;
  data: any[];
}) {
  const filteredData = data.filter((el) => el.type === type);
  const total = filteredData.reduce(
    (acc, el) => acc + (el._sum?.amount || 0),
    0
  );
  return (
    <Card className="h-80 w-full col-span-6">
      <CardHeader>
        <CardTitle className="grid grid-flow-row justify-between gap-2 text-muted-foreground md:grid-flow-col">
          {type === "INCOME" ? "Incomes" : "Expenses"} by category
        </CardTitle>
      </CardHeader>

      <div className="flex items-center justify-between gap-2">
        {filteredData.length === 0 && (
          <div className="flex h-60 w-full flex-col items-center justify-center">
            No data for the selected period
            <p className="text-sm text-muted-foreground">
              Try selecting a different period or try adding new{" "}
              {type === "INCOME" ? "income" : "expense"}
            </p>
          </div>
        )}

        {filteredData.length > 0 && (
          <ScrollArea className="h-60 w-full px-4">
            <div className="flex w-full flex-col gap-4 p-4">
              {filteredData.map((item) => {
                const amount = item._sum.amount || 0;
                const percentage = (amount * 100) / (total || amount);

                return (
                  <div key={item.category} className="flex flex-col gap-2">
                    <div className="flex items-center justify-between">
                      <span className="flex items-center text-gray-400">
                        {item.category}
                        <span className="ml-2 text-xs text-muted-foreground">
                          ({percentage.toFixed(0)}%)
                        </span>
                      </span>

                      <span className="text-sm text-gray-400">
                        {currency} {amount.toFixed(2)}
                      </span>
                    </div>

                    <Progress
                      value={percentage}
                      indicator={
                        type === "INCOME" ? "bg-emerald-500" : "bg-rose-500"
                      }
                    />
                  </div>
                );
              })}
            </div>
          </ScrollArea>
        )}
      </div>
    </Card>
  );
}
