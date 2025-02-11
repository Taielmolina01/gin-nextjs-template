"use client";
import Link from 'next/link'


export default function Home() {
  return (
    <main className="flex w-full flex-col px-2">
      <h1 className="my-5 mx-2">
        Welcome to my Gin + NextJS template. To run all the application you should use this manual.
      </h1>
      <section className="flex flex-col flex-start justify-center items-center">
        <div className="flex flex-col w-[50%] border-[1px] border-gray-500 p-2 text-gray-300 rounded-md">
          <ul>
            <li>
              Lorem ipsum dolor sit amet consectetur adipisicing elit. Quasi eius accusamus atque ipsum enim odit sequi voluptates magni velit incidunt error sint, impedit iure delectus architecto voluptatum dolor. Magni, vitae! 
            </li>
          </ul>
        </div>
      </section>
    </main>
  );
}
