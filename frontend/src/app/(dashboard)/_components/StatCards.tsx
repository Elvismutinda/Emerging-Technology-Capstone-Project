"use client";

import SkeletonWrapper from "@/components/SkeletonWrapper";
import { Card } from "@/components/ui/card";
import useCurrentUser from "@/hooks/use-current-user";
import { useQuery } from "@tanstack/react-query";
import axios from "axios";
import { TrendingDown, TrendingUp, Wallet } from "lucide-react";
import { ReactNode } from "react";
import CountUp from "react-countup";

interface Props {
  from: Date;
  to: Date;
}

function StatCards({ from, to }: Props) {
  const { token, user } = useCurrentUser();
  const userId = user?.id;

  const statsQuery = useQuery({
    queryKey: ["overview", "stats", from, to],
    queryFn: async () => {
      const response = await axios.get(
        "http://localhost:8000/transaction/get-overview",
        {
          headers: {
            Authorization: `Bearer ${token}`,
            userId: userId,
          }
          // params: {
          //   startDate: from.toISOString(),
          //   endDate: to.toISOString(),
          // }
        }
      );
      return response.data.data;
    },
  });

  const income = statsQuery.data?.income || 0;
  const expense = statsQuery.data?.expense || 0;
  const balance = income - expense;

  return (
    <div className="relative flex w-full flex-wrap gap-2 md:flex-nowrap">
      <SkeletonWrapper isLoading={false}>
        <StatCard
          currency="Ksh"
          value={income}
          title="Income"
          icon={
            <TrendingUp className="h-12 w-12 items-center rounded-lg p-2 text-emerald-500 bg-emerald-400/10" />
          }
        />
      </SkeletonWrapper>

      <SkeletonWrapper isLoading={false}>
        <StatCard
          currency="Ksh"
          value={expense}
          title="Expense"
          icon={
            <TrendingDown className="h-12 w-12 items-center rounded-lg p-2 text-rose-500 bg-rose-400/10" />
          }
        />
      </SkeletonWrapper>

      <SkeletonWrapper isLoading={false}>
        <StatCard
          currency="Ksh"
          value={balance}
          title="Balance"
          icon={
            <Wallet className="h-12 w-12 items-center rounded-lg p-2 text-violet-500 bg-violet-400/10" />
          }
        />
      </SkeletonWrapper>
    </div>
  );
}

export default StatCards;

function StatCard({
  currency,
  value,
  title,
  icon,
}: {
  currency: string;
  value: number;
  title: string;
  icon: ReactNode;
}) {
  return (
    <Card className="flex h-24 w-full items-center gap-2 p-4">
      {icon}
      <div className="flex flex-col items-start gap-0">
        <p className="text-muted-foreground">{title}</p>
        <CountUp
          preserveValue
          redraw={false}
          end={value}
          decimals={2}
          formattingFn={(value) => `${currency} ${value}`}
          className="text-2xl"
        />
      </div>
    </Card>
  );
}
