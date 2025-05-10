import { type FC, type PropsWithChildren } from "react";
import { Header } from "../components/Header";
import { Footer } from "../components/Footer";

export const Layout: FC<PropsWithChildren> = ({ children }) => (
  <div className="flex h-screen flex-col text-center">
    <Header />
    <main className="grow">
      <div className="wrapper debug h-full py-6">{children}</div>
    </main>
    <Footer />
  </div>
);
