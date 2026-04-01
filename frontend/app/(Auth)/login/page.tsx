import Link from "next/link";
import LoginForm from "./loginfrom";

export const metadata = {
  title: "Sign In | Dental Management",
};

export default function LoginPage() {
  return (
    <div className="min-h-screen bg-[#eeeeee] flex items-center justify-center p-4">
      <div className="bg-white w-full max-w-md rounded-[2rem] shadow-xl p-10">
        <div className="text-center mb-8">
          <h1 className="text-3xl font-black text-[#8f1eae] mb-2">SIGN IN</h1>
          <p className="text-gray-500 italic">ยินดีต้อนรับกลับมาอีกครั้ง</p>
        </div>

        <LoginForm />

        {/* <div className="mt-8 text-center">
          <p className="text-sm text-gray-600">
            new user?{" "}
            <Link
              href="/register"
              className="text-[#8f1eae] font-bold underline"
            >
              Sign up
            </Link>
          </p>
        </div> */}
      </div>
    </div>
  );
}
