"use server";
import { permanentRedirect } from "next/navigation";

export async function POST(request: Request) {
  process.env.NODE_TLS_REJECT_UNAUTHORIZED = "0";
  const go_server_url = await process.env.GO_SERVER_URL;

  const loginDetails = await request.json();

  const res = await fetch(go_server_url + "/user/login", {
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

  // Extracting cookies from the response headers
  const setCookieHeader = res.headers.get("Set-Cookie");
  const parsedCookies = parseSetCookieHeader(setCookieHeader);

  // Get the "userID" cookie
  const useridCookie = parsedCookies.find(
    (cookie: { name: string }) => cookie.name === "userID"
  );

  if (useridCookie) {
    return permanentRedirect(
      `/dashboard/${useridCookie?.value}?userid=${useridCookie?.value}`
    );
  } else {
    return permanentRedirect(`/dashboard/0`);
  }
}

// Helper function to parse Set-Cookie header
function parseSetCookieHeader(setCookieHeader: string | null) {
  if (!setCookieHeader) {
    return [];
  }

  return setCookieHeader.split(";").map((cookie) => {
    const [name, ...options] = cookie.trim().split("=");
    const [value] = options.join("=").split(";");
    return { name, value };
  });
}
