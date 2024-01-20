export const dynamic = 'force-dynamic' // defaults to auto
import { cookies } from 'next/headers'
 


  export async function POST(request:Request) {
    //const cookieStore = cookies()
    const formData = await request.json();
    const formDataJson = JSON.stringify(formData)

    const res = await fetch("http://localhost:8080/workout/addworkoutlog", {
      method: "POST",
      body: formDataJson,
    }).then((res) => {
        if (!res.ok) {
          return res.text().then((text) => {
            JSON.stringify({flashMessage: new Error(text)})
          });
        }
        
        return res.json();
      })
      .then((message) => {

        //const sessionId = cookieStore.get('test-session')?.value
        return JSON.stringify({flashMessage: message})
      });

    return Response.json( res )
  }