import { Button } from "@/components/ui/button";
import React from "react";
import CreateTransactionDialogue from "./_components/CreateTransactionDialogue";
import Overview from "./_components/Overview";

async function page() {
  // const user = await currentUser();
  // if (!user) {
  //   redirect('/login');
  // }

  return (
    <div className="h-full bg-background">
      <div className="border-b bg-card">
        <div className="container flex flex-wrap items-center justify-between gap-6 py-8">
          <p className="text-3xl font-bold">Hello, Elvis</p>

          <div className="flex items-center gap-3">
            <CreateTransactionDialogue
              type="income"
              trigger={
                <Button
                  variant="outline"
                  className="border-emerald-500 bg-emerald-950 text-white hover:bg-emerald-700 hover:text-white"
                >
                  New Income
                </Button>
              }
            />
            <CreateTransactionDialogue
              type="expense"
              trigger={
                <Button
                  variant="outline"
                  className="border-rose-500 bg-rose-950 text-white hover:bg-rose-700 hover:text-white"
                >
                  New Expense
                </Button>
              }
            />
          </div>
        </div>
      </div>
      <Overview />
    </div>
  );
}

export default page;
