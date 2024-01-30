"use client";
import SideNav from "@/app/(post_login)/dashboard/components/sidenav";
import { useSearchParams } from "next/navigation";
import { useEffect, useState } from "react";

export default function Page() {
  const searchParams = useSearchParams();

  const userid = searchParams.get("userid");

  const [user, setUser] = useState({
    userName: "",
    about: "",
    firstName: "",
    lastName: "",
    email: "",
    country: "",
    password: "",
    dob: "", // You might want to initialize it based on your date format
    gender: "",
  });

  useEffect(() => {
    const getUniqueUser = async () => {
      try {
        const response = await fetch(`/api/get_user_details`, {
          method: "POST",
          body: userid,
        });
        const userDetails: UserDetails = await response.json();
        setUser(userDetails);
      } catch (error) {
        console.error("Error fetching user:", error);
      }
    };

    if (userid) {
      getUniqueUser();
    }
  }, [userid]);

  return (
    <div className="flex h-screen flex-col md:flex-row md:overflow-hidden">
      <div className="w-full flex-none md:w-64">
        <SideNav />
      </div>
      <div className="flex-grow p-6 md:overflow-y-auto md:p-12">
        <p className="text-gray-100">{user.firstName}'s Dashboard Page</p>
        <br></br>

        <p className="text-gray-100">User details:</p>
        <p className="text-gray-100">First name: {user.firstName}</p>
        <p className="text-gray-100">Last name: {user.lastName}</p>
        <p className="text-gray-100">Gender: {user.gender}</p>
        <p className="text-gray-100">Username: {user.userName}</p>
        <p className="text-gray-100">Date of Birth: {user.dob}</p>
        <p className="text-gray-100">Country: {user.country}</p>
      </div>
    </div>
  );
}
