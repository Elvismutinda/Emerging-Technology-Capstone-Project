"use client";

import Link from "next/link";
import React, { useState } from "react";
import { Input } from "./ui/input";
import { Button } from "./ui/button";
import { Loader2 } from "lucide-react";
import Logo from "./Logo";

const LoginForm = () => {
  const [isLoading, setIsLoading] = useState(false);

  async function handleLogin(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    setIsLoading(true);
  }
  return (
    <form
      onSubmit={handleLogin}
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
          Enter your email and password to log in to your account.
        </p>
      </div>
      <Input
        aria-label="email"
        name="email"
        type="email"
        placeholder="Email"
        defaultValue=""
      />

      <Input
        aria-label="password"
        name="password"
        type="password"
        placeholder="Password"
        defaultValue=""
      />

      {/* <p className="text-sm underline text-[#32c06b] cursor-pointer">
        Forgot password?
      </p> */}

      <div>
        <Button type="submit" className="w-full bg-[#61bd73] text-[#130c49]">
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
  );
};

export default LoginForm;
