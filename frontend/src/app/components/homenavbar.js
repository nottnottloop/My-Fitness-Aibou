'use client'

import { Bars3Icon } from '@heroicons/react/24/outline'


export function HomeNavBar({ props }) {

  return (
    <nav
    className="flex items-center justify-between p-6 lg:px-8"
    aria-label="Global"
  >
    <div className="flex lg:flex-1">
      <a href="/" className="-m-1.5 p-1.5">
        <span className="sr-only">Your Company</span>
        <img
          className="h-8 w-auto"
          src="https://tailwindui.com/img/logos/mark.svg?color=amber&shade=300"
          alt=""
        />
      </a>
    </div>
    <div className="flex lg:hidden">
      <button
        type="button"
        className="-m-2.5 inline-flex items-center justify-center rounded-md p-2.5 text-gray-100"
        onClick={() => setMobileMenuOpen(true)}
      >
        <span className="sr-only">Open main menu</span>
        <Bars3Icon className="h-6 w-6" aria-hidden="true" />
      </button>
    </div>
    <div className="hidden lg:flex lg:gap-x-12">
      {props.map((item) => (
        <a
          key={item.name}
          href={item.href}
          className="text-sm font-semibold leading-6 text-gray-100"
        >
          {item.name}
        </a>
      ))}
    </div>
    <div className="hidden lg:flex lg:flex-1 lg:justify-end">
      <a
        href="/user/login"
        className="text-sm font-semibold leading-6 text-amber-300"
      >
        Log in <span aria-hidden="true">&rarr;</span>
      </a>
    </div>
  </nav>
  )
}