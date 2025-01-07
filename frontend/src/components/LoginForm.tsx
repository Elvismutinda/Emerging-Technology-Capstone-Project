"use client";

import Link from "next/link";
import React, { useState } from "react";
import { Input } from "./ui/input";
import { Button } from "./ui/button";
import { Loader2 } from "lucide-react";
import Logo from "./Logo";
import { Controller, useForm } from "react-hook-form";
import { loginUserForm, loginUserSchema } from "@/lib/validations/auth";
import { zodResolver } from "@hookform/resolvers/zod";
import axios from "axios";
import { toast } from "sonner";
import { useRouter } from "next/navigation";

const LoginForm = () => {
  const {
    control,
    handleSubmit,
    formState: { errors },
  } = useForm<loginUserForm>({
    resolver: zodResolver(loginUserSchema),
  });

  const [isLoading, setIsLoading] = useState<boolean>(false);
  const router = useRouter();

  async function onSubmit(data: loginUserForm) {
    setIsLoading(true);

    try {
      const response = await axios.post(
        "http://localhost:8000/auth/login",
        data,
        {
          headers: {
            "Content-Type": "application/json",
          },
        }
      );

      // Save token and user data in local storage
      const { token, user_data } = response.data.data;
      localStorage.setItem("authToken", token);
      localStorage.setItem("userData", JSON.stringify(user_data));

      toast.success(response.data.message);

      router.push("/");
    } catch (error: any) {
      console.log("Error logging in: ", error.response?.data || error.message);
      toast.error(
        error.response?.data?.message || "Login failed. Please try again!"
      );
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
              Login
            </h1>
          </header>
          <p className="pt-1 text-base text-slate-400">
            Enter your email/username and password to log in to your account.
          </p>
        </div>
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

        {/* <p className="text-sm underline text-[#32c06b] cursor-pointer">
        Forgot password?
      </p> */}

        <div>
          <Button
            type="submit"
            className="w-full bg-[#61bd73] text-[#130c49]"
            disabled={isLoading}
          >
            Login
            {isLoading && <Loader2 className="ml-2 h-4 w-4 animate-spin" />}
          </Button>
        </div>
        <p className="text-sm text-slate-400">
          Don&apos;t have an account?{" "}
          <Link href="/signup" className="text-[#32c06b] underline">
            Sign Up
          </Link>
        </p>
      </form>
    </div>
  );
};

export default LoginForm;
