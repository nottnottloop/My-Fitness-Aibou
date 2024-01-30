"use server";

import { BmrFormData } from "@/app/lib/types";

export async function POST(request: Request) {
  process.env.NODE_TLS_REJECT_UNAUTHORIZED = "0"; // Set to "0" to disable certificate verification. set back to "1" when in prod
  const go_server_url = await process.env.GO_SERVER_URL;

  const userID = await request.json();

  const res = await fetch(go_server_url + "/user/details", {
    method: "POST",
    body: JSON.stringify(userID),
    agent: new (require("https").Agent)({ rejectUnauthorized: false }), // Ignore certificate verification
  }).then((res) => {
    if (!res.ok) {
      return res.text().then((text) => {
        throw new Error(text);
      });
    }
    return res.json();
  });

  process.env.NODE_TLS_REJECT_UNAUTHORIZED = "1";

  return Response.json(res);
}
