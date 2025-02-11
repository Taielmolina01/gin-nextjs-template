import FormInput from "../components/formInput";
import Link from 'next/link'

export default function Register() {

    return (
        <main className="flex h-full w-full font-sans">
            <section className="flex flex-col justify-start w-1/2 min-w-[50vw] bg-gradient-to-b from-purple-900 to-black p-10">
                <div className="mx-auto my-[70px] w-[70%] justify-center">
                    <h1 className="text-2xl my-[10px] text-start font-bold">
                        Create your account
                    </h1>
                    <p>
                        Sign up now to access all features of our platform.  
                        <br/>
                        Fill out the form and start enjoying a secure and personalized experience.
                    </p>
                </div>
            </section>

            <section className="flex w-1/2 min-w-[50vw] justify-center bg-white text-black p-10">
                <div className="my-[80px] w-[50%]">
                    <h1 className="text-xl font-semibold">
                        Sign up to the Gin + NextJS template
                    </h1>
                    <form className="flex flex-col py-3">
                        <FormInput label="First name" type="text" fontColor="text-gray-800"/>
                        <FormInput label="Last name" type="text" fontColor="text-gray-800"/>
                        <FormInput label="Email" type="email" fontColor="text-gray-800"/>
                        <FormInput label="Password" type="password" help="Password must have at least 8 characters" fontColor="text-gray-800" />
                        <button className="bg-black text-white py-2 my-2 rounded-md" type="submit">
                            Continue >
                        </button>
                    </form>
                    <span>
                        Already have an account?  
                        <Link href="/login" className="ml-1 text-purple-900 underline underline-offset-2">
                            Login
                        </Link>
                    </span>
                </div>
            </section>
        </main>
    );
}