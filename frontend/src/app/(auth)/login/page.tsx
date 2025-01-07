import { Metadata } from "next";

import LoginForm from "@/components/LoginForm";

export const metadata: Metadata = {
  title: "Login",
  description: "Login to your account",
}

const LoginPage = () => {
  return (
    <div className="flex h-screen w-screen flex-col items-center justify-center">
      <LoginForm />
    </div>
  );
};

export default LoginPage;
