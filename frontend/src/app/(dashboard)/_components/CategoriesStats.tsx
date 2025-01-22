"use client";

import SkeletonWrapper from "@/components/SkeletonWrapper";
import { Card, CardHeader, CardTitle } from "@/components/ui/card";
import { Progress } from "@/components/ui/progress";
import { ScrollArea } from "@/components/ui/scroll-area";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import useCurrentUser from "@/hooks/use-current-user";
import { TransactionType } from "@/types";
import { useQuery } from "@tanstack/react-query";
import axios from "axios";
import React, { useState } from "react";

interface Props {
  from: Date;
  to: Date;
}

function formattedDate(date: Date): string {
  return date.toISOString().split("T")[0];
}

function CategoriesStats({ from, to }: Props) {
  const { token, user } = useCurrentUser();
  const userId = user?.id;
  const [incomeCategory, setIncomeCategory] = useState<string>("");
  const [expenseCategory, setExpenseCategory] = useState<string>("");

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

  const incomeStatsQuery = useQuery({
    queryKey: ["incomeStats", from, to, incomeCategory],
    queryFn: async () => {
      const response = await axios.get(
        `http://localhost:8000/transaction/get-category-stats/${incomeCategory}`,
        {
          headers: {
            Authorization: `Bearer ${token}`,
            userId: userId,
          },
          params: {
            start_date: formattedDate(from),
            end_date: formattedDate(to),
          },
        }
      );
      return response.data.data;
    },
    enabled: !!incomeCategory,
  });

  const expenseStatsQuery = useQuery({
    queryKey: ["expenseStats", from, to, expenseCategory],
    queryFn: async () => {
      const response = await axios.get(
        `http://localhost:8000/transaction/get-category-stats/${expenseCategory}`,
        {
          headers: {
            Authorization: `Bearer ${token}`,
            userId: userId,
          },
          params: {
            start_date: formattedDate(from),
            end_date: formattedDate(to),
          },
        }
      );
      return response.data.data;
    },
    enabled: !!expenseCategory,
  });

  return (
    <div className="flex w-full flex-wrap gap-2 md:flex-nowrap">
      <SkeletonWrapper isLoading={incomeStatsQuery.isLoading}>
        <CategoriesCard
          currency="Ksh"
          type="INCOME"
          data={incomeStatsQuery.data || { daily_incomes: [] }}
          categories={categories.filter(
            (cat: { type: string }) => cat.type === "INCOME"
          )}
          selectedCategory={incomeCategory}
          setSelectedCategory={setIncomeCategory}
        />
      </SkeletonWrapper>

      <SkeletonWrapper isLoading={expenseStatsQuery.isLoading}>
        <CategoriesCard
          currency="Ksh"
          type="EXPENSE"
          data={expenseStatsQuery.data || { daily_expenses: [] }}
          categories={categories.filter(
            (cat: { type: string }) => cat.type === "EXPENSE"
          )}
          selectedCategory={expenseCategory}
          setSelectedCategory={setExpenseCategory}
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
  categories,
  selectedCategory,
  setSelectedCategory,
}: {
  currency: string;
  type: TransactionType;
  data: any;
  categories: any[];
  selectedCategory: string;
  setSelectedCategory: React.Dispatch<React.SetStateAction<string>>;
}) {
  const dailyData =
    type === "INCOME" ? data?.daily_incomes || [] : data?.daily_expenses || [];

  const total = dailyData.reduce(
    (acc: number, amount: number) => acc + amount,
    0
  );

  return (
    <Card className="h-80 w-full col-span-6">
      <CardHeader>
        <CardTitle className="grid grid-flow-row justify-between gap-2 text-muted-foreground md:grid-flow-col">
          {type === "INCOME" ? "Incomes" : "Expenses"} by category
        </CardTitle>
        <Select value={selectedCategory} onValueChange={setSelectedCategory}>
          <SelectTrigger>
            <SelectValue
              placeholder={`Select an ${type.toLowerCase()} category`}
            />
          </SelectTrigger>
          <SelectContent>
            {categories.map((category) => (
              <SelectItem key={category.id} value={category.name}>
                {category.name}
              </SelectItem>
            ))}
          </SelectContent>
        </Select>
      </CardHeader>

      <div className="flex items-center justify-between gap-2">
        {dailyData.length === 0 && (
          <div className="flex h-60 w-full flex-col items-center justify-center">
            No data for the selected period
            <p className="text-sm text-muted-foreground">
              Try selecting a different period or try adding new{" "}
              {type === "INCOME" ? "income" : "expense"}
            </p>
          </div>
        )}

        {dailyData.length > 0 && (
          <ScrollArea className="h-60 w-full px-4">
            <div className="flex w-full flex-col gap-4 p-4">
              {dailyData.map((amount: number, index: number) => {
                const percentage = (amount * 100) / (total || amount);

                return (
                  <div key={index} className="flex flex-col gap-2">
                    <div className="flex items-center justify-between">
                      <span className="flex items-center text-gray-400">
                        {selectedCategory} {index + 1}
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
