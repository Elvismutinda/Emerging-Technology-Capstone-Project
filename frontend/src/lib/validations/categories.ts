import { z } from "zod";

export const CreateCategorySchema = z.object({
  name: z.string().min(3).max(20),
  icon: z.string().max(20),
  type: z.enum(["transaction", "expense"]),
});

export type CreateCategoryRequest = z.infer<typeof CreateCategorySchema>;
