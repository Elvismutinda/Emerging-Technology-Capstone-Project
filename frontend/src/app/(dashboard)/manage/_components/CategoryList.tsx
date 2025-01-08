"use client";

import SkeletonWrapper from "@/components/SkeletonWrapper";
import { Card, CardHeader, CardTitle } from "@/components/ui/card";
import { TransactionType } from "@/types";
import { useQuery } from "@tanstack/react-query";
import { PlusSquare, TrendingDown, TrendingUp } from "lucide-react";
import React from "react";
import CreateCategoryDialog from "../../_components/CreateCategoryDialog";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { cn } from "@/lib/utils";
import axios from "axios";
import useCurrentUser from "@/hooks/use-current-user";

function CategoryList({ type }: { type: TransactionType }) {
  const { token } = useCurrentUser();

  const categoriesQuery = useQuery({
    queryKey: ["categories", type],
    queryFn: async () => {
      const response = await axios.get(
        `http://localhost:8000/category/${type}`,
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );
      return response.data;
    },
  });

  const categories = categoriesQuery.data || [];

  return (
    <SkeletonWrapper isLoading={categoriesQuery.isLoading}>
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center justify-between gap-2">
            <div className="flex items-center gap-2">
              {type === "expense" ? (
                <TrendingDown className="h-12 w-12 items-center rounded-lg bg-rose-400/10 text-rose-500 p-2" />
              ) : (
                <TrendingUp className="h-12 w-12 items-center rounded-lg bg-emerald-400/10 text-emerald-500 p-2" />
              )}
              <div>
                {type === "income" ? "Incomes" : "Expenses"} categories
                <div className="text-sm text-muted-foreground">
                  Sorted by name
                </div>
              </div>
            </div>

            <CreateCategoryDialog
              type={type}
              trigger={
                <Button className="gap-2 text-sm">
                  <PlusSquare className="w-4 h-4" />
                  Create category
                </Button>
              }
            />
          </CardTitle>
        </CardHeader>

        <Separator />

        <div>
          {categories.length > 0 ? (
            <ul className="p-4 space-y-2">
              {categories.map((category: any) => (
                <li
                  key={category.id}
                  className="flex justify-between items-center border rounded-md px-4 py-2 bg-muted/10"
                >
                  <span className="font-medium">{category.name}</span>
                  <span
                    className={cn(
                      "text-sm",
                      type === "income" ? "text-emerald-500" : "text-rose-500"
                    )}
                  >
                    {category.type}
                  </span>
                </li>
              ))}
            </ul>
          ) : (
            <div className="flex h-40 w-full flex-col items-center justify-center">
              <p>
                No
                <span
                  className={cn(
                    "m-1",
                    type === "income" ? "text-emerald-500" : "text-rose-500"
                  )}
                >
                  {type}
                </span>
                categories yet
              </p>

              <p className="text-sm text-muted-foreground">
                Create one to get started
              </p>
            </div>
          )}
        </div>
      </Card>
    </SkeletonWrapper>
  );
}

export default CategoryList;
