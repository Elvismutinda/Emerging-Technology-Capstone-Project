import { z } from "zod";

export const CreateCategorySchema = z.object({
  name: z.string().min(3).max(20),
  type: z.enum(["INCOME", "EXPENSE"]),
});

export type CreateCategoryRequest = z.infer<typeof CreateCategorySchema>;
