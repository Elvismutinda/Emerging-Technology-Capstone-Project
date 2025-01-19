import { z } from "zod";

export const CreateTransactionSchema = z.object({
  amount: z.coerce.number().positive().multipleOf(0.01),
  description: z.string().optional(),
  transaction_date: z.coerce.date(),
  category: z.string(),
  type: z.union([z.literal("INCOME"), z.literal("EXPENSE")]),
});

export type CreateTransactionRequest = z.infer<typeof CreateTransactionSchema>;
