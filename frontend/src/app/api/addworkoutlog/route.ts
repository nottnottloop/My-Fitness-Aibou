export const dynamic = 'force-dynamic' // defaults to auto
import { cookies } from 'next/headers'
 


  export async function POST(request:Request) {
    const cookieStore = cookies()
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
        console.log(res)

      
        return res.json();
      })
      .then((data) => {

        const flash = cookieStore.get('flash')?.value

        if (flash){
        return JSON.stringify({flashMessage: flash})
        }
        else{
        return JSON.stringify({flashMessage: "no flash message but successful post. Workoutlog ID: "+data})

        }
      });

    return Response.json( res )
  }