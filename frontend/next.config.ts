import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  /* config options here */
  reactCompiler: true,

  experimental: {
    typedRoutes: true,
  },
  
  // ป้องกันปัญหา Runtime Error เวลาลืมใส่ Environment Variables (แนะนำใช้ร่วมกับ t3-env)
  // แต่เบื้องต้นใส่ไว้เพื่อให้ Next.js จัดการเรื่อง Output ให้ดีขึ้น
  output: 'standalone',
};

export default nextConfig;
