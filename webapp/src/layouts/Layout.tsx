import { type FC, type PropsWithChildren } from "react";
import { Header } from "../components/Header";
import { Footer } from "../components/Footer";

export const Layout: FC<PropsWithChildren> = ({ children }) => (
  <div className="flex flex-col text-center h-screen">
    <Header />
    <main className="grow">
      <div className="wrapper debug py-6 h-full">{children}</div>
    </main>
    <Footer />
  </div>
);
