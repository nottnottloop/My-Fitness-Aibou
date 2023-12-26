"use client";
import { PhotoIcon, UserCircleIcon } from "@heroicons/react/24/solid";
import { useState } from "react";

export default function Page() {
  const [bmrResult, setBmrResult] = useState();

  function getBmr(formData: FormData) {
    const rawFormData = {
      age: formData.get("age"),
      weight: formData.get("weight"),
      height: formData.get("height"),
      gender: formData.get("gender"),
    };
    //console.log(rawFormData);

    const res = fetch(
      "http://localhost:8080/kitchen/bmr?age=" +
        rawFormData.age +
        "&weight=" +
        rawFormData.weight +
        "&height=" +
        rawFormData.height +
        "&gender=" +
        rawFormData.gender,
      { method: "POST" }
    )
      .then((res) => {
        if (!res.ok) {
          return res.text().then((text) => {
            throw new Error(text);
          });
        }
        return res.json();
      })
      .then((data) => {
        setBmrResult(data);
      });
  }

  return (
    <div className="mx-auto max-w-7xl px-2 py-14 sm:px-6 lg:px-8">
      <form action={getBmr}>
        <div className="space-y-12">
          <div className="border-b border-gray-900/10 pb-12">
            <h1 className="text-base font-semibold leading-7 text-gray-900">
              BMR Calculator
            </h1>
            <p className="mt-1 text-sm leading-6 text-gray-600">
              The Basal Metabolic Rate (BMR) Calculator estimates your basal
              metabolic rate—the amount of energy expended while at rest in a
              neutrally temperate environment, and in a post-absorptive state
              (meaning that the digestive system is inactive, which requires
              about 12 hours of fasting).
            </p>
          </div>

          <div className="border-b border-gray-900/10 pb-12">
            <h2 className="text-base font-semibold leading-7 text-gray-900">
              Personal Information
            </h2>
            <p className="mt-1 text-sm leading-6 text-gray-600">Metric Units</p>

            <div className="mt-10 grid grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-6">
              <div className="sm:col-span-3">
                <label
                  htmlFor="weight"
                  className="block text-sm font-medium leading-6 text-gray-900"
                >
                  Weight (in kg)
                </label>
                <div className="mt-2">
                  <input
                    type="number"
                    name="weight"
                    id="weight"
                    autoComplete="weight"
                    className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                  />
                </div>
              </div>

              <div className="sm:col-span-3">
                <label
                  htmlFor="height"
                  className="block text-sm font-medium leading-6 text-gray-900"
                >
                  Height (in cm)
                </label>
                <div className="mt-2">
                  <input
                    type="number"
                    name="height"
                    id="height"
                    className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                  />
                </div>
              </div>

              <div className="sm:col-span-4">
                <label
                  htmlFor="age"
                  className="block text-sm font-medium leading-6 text-gray-900"
                >
                  Age
                </label>
                <div className="mt-2">
                  <input
                    id="age"
                    name="age"
                    type="number"
                    autoComplete="age"
                    className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                  />
                </div>
              </div>

              <div className="sm:col-span-3">
                <label
                  htmlFor="gender"
                  className="block text-sm font-medium leading-6 text-gray-900"
                >
                  Gender
                </label>
                <div className="mt-2">
                  <select
                    id="gender"
                    name="gender"
                    className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:max-w-xs sm:text-sm sm:leading-6"
                  >
                    <option>Male</option>
                    <option>Female</option>
                  </select>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div className="mt-6 flex items-center justify-end gap-x-6">
          <button
            type="button"
            className="text-sm font-semibold leading-6 text-gray-900"
          >
            Clear
          </button>
          <button
            type="submit"
            className="rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
          >
            Calculate
          </button>
        </div>
      </form>
      <div>
        <h1 className="text-base font-semibold leading-7 text-gray-900">
          {bmrResult ?? bmrResult}
        </h1>
      </div>
    </div>
  );
}
