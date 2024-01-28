"use server";
import { permanentRedirect } from "next/navigation";

export async function POST(request: Request) {
  process.env.NODE_TLS_REJECT_UNAUTHORIZED = "0";

  const loginDetails = await request.json();

  const res = await fetch("https://localhost:8080/user/login", {
    method: "POST",
    body: JSON.stringify(loginDetails),
    agent: new (require("https").Agent)({ rejectUnauthorized: false }), // Ignore certificate verification
  });
  process.env.NODE_TLS_REJECT_UNAUTHORIZED = "1";

  if (!res.ok) {
    const responseBody = await res.json();
    return new Response(JSON.stringify(responseBody), {
      status: res.status,
    });
  }

  return permanentRedirect(`/dashboard`);
}
