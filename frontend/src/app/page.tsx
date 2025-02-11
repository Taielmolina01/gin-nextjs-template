"use client";
import Link from 'next/link'


export default function Home() {
  return (
    <main className="flex w-full flex-col items-center">
      <nav className="flex justify-between flex-row w-full items-center border-b-[0.5px] border-gray-900 px-4 py-2">
        Gin + NextJS Template
        <div className="flex w-[150px] justify-between">
          <button className="flex border-[0.5px] py-[5px] px-[10px] bg-white border-white text-black rounded-md">
            <Link href="/register">
              Register
            </Link>
          </button>
          <button className="flex border-[0.5px] py-[5px] px-[10px] border-gray-900 rounded-md">
            <Link href="login">
              Login
            </Link>
          </button>
        </div>
      </nav>
      <section className="flex min-h-screen flex-col items-center justify-center">
        Manual:
        <ul>
          <li>
            You must run the backend 
          </li>
        </ul>
      </section>
    </main>
  );
}
