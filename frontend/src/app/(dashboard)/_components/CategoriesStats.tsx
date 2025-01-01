"use client";

import SkeletonWrapper from "@/components/SkeletonWrapper";
import { Card, CardHeader, CardTitle } from "@/components/ui/card";
import { TransactionType } from "@/types";
import React from "react";

interface Props {
  from: Date;
  to: Date;
}

function CategoriesStats({ from, to }: Props) {
  return (
    <div className="flex w-full flex-wrap gap-2 md:flex-nowrap">
      <SkeletonWrapper isLoading={false}>
        <CategoriesCard type="income" data={1} />
      </SkeletonWrapper>

      <SkeletonWrapper isLoading={false}>
        <CategoriesCard type="expense" data={1} />
      </SkeletonWrapper>
    </div>
  );
}

export default CategoriesStats;

function CategoriesCard({
  type,
  data,
}: {
  type: TransactionType;
  data: number;
}) {
  return (
    <Card className="h-80 w-full col-span-6">
      <CardHeader>
        <CardTitle className="grid grid-flow-row justify-between gap-2 text-muted-foreground md:grid-flow-col">
          {type === "income" ? "Incomes" : "Expenses"} by category
        </CardTitle>
      </CardHeader>

      <div className="flex items-center justify-between gap-2">
        <div className="flex h-60 w-full flex-col items-center justify-center">
          No data for the selected period
          <p className="text-sm text-muted-foreground">
            Try selecting a different period or try adding new{" "}
            {type === "income" ? "transactions" : "expenses"}
          </p>
        </div>
      </div>
    </Card>
  );
}
