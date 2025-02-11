import FormInput from "../components/formInput";
import Link from "next/link";

export default function Login() {
    return (
        <main className="flex flex-col min-h-screen min-w-screen justify-items-center items-start">
            <div className="flex mx-auto flex-col justify-items-center my-[80px] w-[20%]">
                    <h1 className="text-xl font-semibold">
                        Login to the Gin + NextJS template
                    </h1>
                    <form className="flex flex-col py-3">
                        <FormInput label="Email" type="email" fontColor="text-gray-200"/>
                        <FormInput label="Password" type="password" fontColor="text-gray-200"/>
                        <button className="bg-black border-[1px] border-gray-500 text-white py-2 my-2 rounded-md" type="submit">
                            Continue >
                        </button>
                    </form>
                    <span>
                        Don't have an account? 
                        <Link href="/register" className="ml-1 text-purple-900 underline underline-offset-2">
                            Register
                        </Link>
                    </span>
            </div>
        </main>
    );
}