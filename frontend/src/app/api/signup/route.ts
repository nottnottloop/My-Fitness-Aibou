"use server";

export async function POST(request: Request) {
  process.env.NODE_TLS_REJECT_UNAUTHORIZED = "0"; // Set to "0" to disable certificate verification. set back to "1" when in prod

  const formData = await request.json();

  const res = await fetch("https://localhost:8080/user/signup", {
    method: "POST",
    body: JSON.stringify(formData),
    agent: new (require("https").Agent)({ rejectUnauthorized: false }), // Ignore certificate verification
  })
    .then((res) => {
      if (!res.ok) {
        return res.json().then((json) => {
          throw json;
        });
      }
      return res.json();
    })
    .catch((error) => {
      return error; // Catch errors thrown by the server
    });

  process.env.NODE_TLS_REJECT_UNAUTHORIZED = "1";

  return Response.json(res);
}
