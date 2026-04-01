import { NextResponse } from 'next/server'
import type { NextRequest } from 'next/server'

export function middleware(request: NextRequest) {
  const token = request.cookies.get('jwt')?.value;
  const { pathname } = request.nextUrl;

  // 1. ถ้า Login แล้ว (มี Token) แต่พยายามจะไปหน้า /login หรือ /register
  if (token && (pathname === '/login' || pathname === '/register' || pathname === '/')) {
    // ส่งกลับไปหน้า Dashboard (หรือหน้าแรกที่เหมาะสม)
    return NextResponse.redirect(new URL('/patient/dashboard', request.url));
  }

  // 2. ถ้ายังไม่ได้ Login (ไม่มี Token) แต่พยายามจะเข้าหน้าที่มีความสำคัญ (Protected Routes)
  const protectedPaths = ['/patient', '/dentist']; // ระบุ Path ที่ต้อง Login ก่อน
  const isProtected = protectedPaths.some(path => pathname.startsWith(path));

  if (!token && isProtected) {
    return NextResponse.redirect(new URL('/login', request.url));
  }

  return NextResponse.next();
}

// กำหนดว่าให้ Middleware ทำงานที่ Path ไหนบ้าง
export const config = {
  matcher: ['/','/login', '/register', '/patient/:path*', '/dentist/:path*'],
};