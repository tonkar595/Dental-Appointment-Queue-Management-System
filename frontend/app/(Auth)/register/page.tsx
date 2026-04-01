import Link from "next/link";
import RegisterForm from "./registerform";

export const metadata = {
  title: "Sign Up | Dental Management",
};

export default function RegisterPage() {
  return (
    <div className="min-h-screen bg-[#eeeeee] flex items-center justify-center p-4">
      <div className="bg-white w-full max-w-md rounded-[2rem] shadow-xl p-10 my-8">
        <div className="text-center mb-8">
          <h1 className="text-3xl font-black text-[#8f1eae] mb-2">SIGN UP</h1>
          <p className="text-gray-500 italic">สร้างบัญชีผู้ใช้ใหม่</p>
        </div>

        <RegisterForm/>

        <div className="mt-8 text-center">
          <p className="text-sm text-gray-600">
            have a account?{" "}
            <Link href="/login" className="text-[#8f1eae] font-bold underline">
              Sign in
            </Link>
          </p>
        </div>
      </div>
    </div>
  );
}
