import * as z from "zod";

export const registerUserSchema = z.object({
  username: z
    .string({
      required_error: "Username is required",
    })
    .min(1, "Username is required"),
  email: z
    .string({ required_error: "Email is required" })
    .min(1, "Email is required")
    .email("Invalid email address"),
  password: z
    .string({ required_error: "Password is required" })
    .min(1, "Password is required")
    .min(8, "Password must be at least 8 characters"),
});

export const loginUserSchema = z.object({
  email: z
    .string({ required_error: "Email is required" })
    .min(1, "Email is required")
    .email("Invalid email address"),
  password: z
    .string({ required_error: "Password is required" })
    .min(1, "Password is required"),
});

export type registerUserForm = z.infer<typeof registerUserSchema>;
export type loginUserForm = z.infer<typeof loginUserSchema>;
