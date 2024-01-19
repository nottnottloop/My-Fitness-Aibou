import SideNav from "@/app/(post_login)/dashboard/components/sidenav";

export default function Page() {
  return (
    <div className="flex h-screen flex-col md:flex-row md:overflow-hidden">
      <div className="w-full flex-none md:w-64">
        <SideNav />
      </div>
      <div className="flex-grow p-6 md:overflow-y-auto md:p-12">
        <p className="text-gray-100">Dashboard Page</p>
      </div>
    </div>
  );
}
