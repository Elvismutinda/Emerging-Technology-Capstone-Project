import { Metadata } from "next";

import RegisterForm from "@/components/RegisterForm";

export const metadata: Metadata = {
  title: "Sign up",
  description: "Create an account",
};

const SignupPage = () => {
  return (
    <div className="flex h-screen w-screen flex-col items-center justify-center">
      <RegisterForm />
    </div>
  );
};

export default SignupPage;
