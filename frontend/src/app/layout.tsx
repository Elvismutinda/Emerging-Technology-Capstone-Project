import type { Metadata } from "next";
import { Inter } from "next/font/google"
import "./globals.css";
import { siteConfig } from "@/config/site";
import Providers from "@/components/Providers";

const inter = Inter({ subsets: ["latin"]})

export const metadata: Metadata = {
  title: {
    default: siteConfig.name,
    template: `%s | ${siteConfig.name}`,
  },
  description: siteConfig.description,
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" className="dark" style={{
      colorScheme: "dark"
    }}>
      <body
        className={inter.className}
      >
        <Providers>{children}</Providers>
      </body>
    </html>
  );
}
