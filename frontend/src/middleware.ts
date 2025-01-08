import { NextResponse } from "next/server";
import { NextRequest } from "next/server";
import { authRoutes } from "../routes"; // Import your route exceptions

export function middleware(req: NextRequest) {
  const token = req.cookies.get("authToken");

  if (!token && !authRoutes.includes(req.nextUrl.pathname)) {
    return NextResponse.redirect(new URL("/login", req.url)); // Redirect to login
  }

  return NextResponse.next();
}

export const config = {
  matcher: ["/((?!.+\\.[\\w]+$|_next).*)", "/", "/(api|trpc)(.*)"],
};
