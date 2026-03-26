// app/api/login/route.ts
const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';
import { NextResponse } from "next/server";
import { cookies } from 'next/headers';
import { apiFetch } from "@/lib/api";

export async function POST(req: Request) {
  const endpoint = "/user/login"
  const url_api = `${API_URL}${endpoint}`
  const body = await req.json();


  try {
    const data = await apiFetch<{ AccesToken: string; RefreshToken: string }>('/user/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(body),
    });

    var check = (data.AccesToken && data.RefreshToken) ? true : false;
    console.log("Login check: ", check)
    if (check) {
      const response = NextResponse.json({ success: check, message: "" });
      response.cookies.set("accessToken", data.AccesToken, {
        httpOnly: true,
        secure: process.env.NODE_ENV === "production",
        path: "/",
      });
      response.cookies.set("refreshToken", data.RefreshToken, {
        httpOnly: true,
        secure: process.env.NODE_ENV === "production",
        path: "/",
      });
      return response;
    } else {
      return NextResponse.json({ success: false, message: "Invalid credentials" }, { status: 400 });
    }
  } catch (error: any) {
    return NextResponse.json({ success: false, message: error.message || 'Login failed. Please try again.' }, { status: 500 });
  }


}