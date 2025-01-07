"use client";

import Link from "next/link";
import React, { useState } from "react";
import { Input } from "./ui/input";
import { Button } from "./ui/button";
import { Loader2 } from "lucide-react";
import Logo from "./Logo";

const RegisterForm = () => {
  const [isLoading, setIsLoading] = useState(false);

  async function handleSubmit(e) {
    e.preventDefault();
    setIsLoading(true);
  }
  return (
    <form
      onSubmit={handleSubmit}
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
      <Input
        aria-label="username"
        name="username"
        type="text"
        placeholder="Username"
        defaultValue=""
        required
      />

      <Input
        aria-label="email"
        name="email"
        type="email"
        placeholder="Email"
        defaultValue=""
        required
      />

      <Input
        aria-label="password"
        name="password"
        type="password"
        placeholder="Password"
        defaultValue=""
        required
      />

      <div>
        <Button type="submit" className="w-full bg-[#61bd73] text-[#130c49]">
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
  );
};

export default RegisterForm;
