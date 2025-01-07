"use client";

import Link from "next/link";
import React, { useState } from "react";
import { Input } from "./ui/input";
import { Button } from "./ui/button";
import { Loader2 } from "lucide-react";
import Logo from "./Logo";
import { Controller, useForm } from "react-hook-form";
import { registerUserForm, registerUserSchema } from "@/lib/validations/auth";
import { zodResolver } from "@hookform/resolvers/zod";
import axios from "axios";
import { toast } from "sonner";
import { useRouter } from "next/navigation";

const RegisterForm = () => {
  const {
    control,
    handleSubmit,
    formState: { errors },
  } = useForm<registerUserForm>({
    resolver: zodResolver(registerUserSchema),
  });

  const [isLoading, setIsLoading] = useState<boolean>(false);
  const router = useRouter();

  async function onSubmit(data: registerUserForm) {
    setIsLoading(true);

    console.log(data);

    try {
      const response = await axios.post(
        "http://localhost:8000/user/signup",
        data,
        {
          headers: {
            "Content-Type": "application/json",
          },
        }
      );

      console.log("User registered successfully: ", response.data);
      toast.success("User registered successfully");

      router.push("/login");
    } catch (error: any) {
      console.log(
        "Error registering user: ",
        error.response?.data || error.message
      );
      toast.error(error.response?.data?.message || "Something went wrong!");
    } finally {
      setIsLoading(false);
    }
  }
  return (
    <div>
      <form
        onSubmit={handleSubmit(onSubmit)}
        className="mx-auto flex w-full flex-col justify-center space-y-6 max-w-[360px] min-w-[320px]"
      >
        <div className="flex flex-col">
          <header className="block space-y-3">
            <Logo />
            <h1 className="text-2xl font-bold text-white tracking-tight">
              Create an account
            </h1>
          </header>
          <p className="pt-1 text-base text-slate-400">
            Create an account to start tracking your espenses.
          </p>
        </div>

        <Controller
          name="username"
          control={control}
          defaultValue=""
          rules={{ required: true }}
          render={({ field }) => (
            <div>
              <Input
                aria-label="username"
                autoComplete="off"
                type="text"
                {...field}
                placeholder="Username"
                className="bg-white text-black"
              />
              {errors.username && (
                <p className="px-1 text-md text-red-600">
                  {errors.username.message}
                </p>
              )}
            </div>
          )}
        />

        <Controller
          name="email"
          control={control}
          defaultValue=""
          rules={{ required: true }}
          render={({ field }) => (
            <div>
              <Input
                aria-label="email"
                autoComplete="off"
                type="email"
                {...field}
                placeholder="Email"
                className="bg-white text-black"
              />
              {errors.email && (
                <p className="px-1 text-md text-red-600">
                  {errors.email.message}
                </p>
              )}
            </div>
          )}
        />

        <Controller
          name="password"
          control={control}
          defaultValue=""
          rules={{ required: true }}
          render={({ field }) => (
            <div>
              <Input
                autoComplete="off"
                type="password"
                {...field}
                placeholder="Password"
                className="bg-white text-black"
              />
              {errors.password && (
                <p className="px-1 text-md text-red-600">
                  {errors.password.message}
                </p>
              )}
            </div>
          )}
        />

        <div>
          <Button
            type="submit"
            className="w-full bg-[#61bd73] text-[#130c49]"
            disabled={isLoading}
          >
            Sign Up
            {isLoading && <Loader2 className="ml-2 h-4 w-4 animate-spin" />}
          </Button>
        </div>
        <p className="text-sm text-slate-400">
          Already have an account?{" "}
          <Link href="/login" className="text-[#32c06b] underline">
            Login
          </Link>
        </p>
      </form>
    </div>
  );
};

export default RegisterForm;
