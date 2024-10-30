"use client";

import SkeletonWrapper from "@/components/SkeletonWrapper";
import { Card } from "@/components/ui/card";
import { TrendingDown, TrendingUp, Wallet } from "lucide-react";
import { ReactNode } from "react";
import CountUp from "react-countup";

interface Props {
  from: Date;
  to: Date;
}

function StatCards({ from, to }: Props) {
  return (
    <div className="relative flex w-full flex-wrap gap-2 md:flex-nowrap">
      <SkeletonWrapper isLoading={false}>
        <StatCard
          value={1000}
          title="Income"
          icon={
            <TrendingUp className="h-12 w-12 items-center rounded-lg p-2 text-emerald-500 bg-emerald-400/10" />
          }
        />
      </SkeletonWrapper>

      <SkeletonWrapper isLoading={false}>
        <StatCard
          value={1000}
          title="Expense"
          icon={
            <TrendingDown className="h-12 w-12 items-center rounded-lg p-2 text-rose-500 bg-rose-400/10" />
          }
        />
      </SkeletonWrapper>

      <SkeletonWrapper isLoading={false}>
        <StatCard
          value={1000}
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
  value,
  title,
  icon,
}: {
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
          className="text-2xl"
        />
      </div>
    </Card>
  );
}
