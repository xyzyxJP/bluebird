import { Toaster } from "@/components/ui/toaster";
import WithApollo from "@/components/with-apollo";
import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Bluebird App",
  description: "",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ja">
      <WithApollo>
        <body className={inter.className}>
          <main className="flex min-h-screen flex-col items-center justify-between p-24">
            {children}
          </main>
          <Toaster />
        </body>
      </WithApollo>
    </html>
  );
}
