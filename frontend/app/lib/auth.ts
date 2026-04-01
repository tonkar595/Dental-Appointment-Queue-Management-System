import { cookies } from "next/headers";

export async function getAuthSession() {
  const cookieStore = await cookies();
  const token = cookieStore.get("jwt")?.value;

  if (!token) return null;

  try {
    return { token }; 
  } catch (err) {
    return null;
  }
}