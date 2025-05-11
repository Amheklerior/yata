import type { FC } from "react";

export const Header: FC = () => (
  <header>
    <div className="wrapper pt-12 pb-8 md:pt-18 md:pb-12" data-width="wide">
      <h1 className="text-6xl text-amber-100 md:text-7xl">YATA</h1>
      <h2 className="text-xl text-amber-100 md:text-3xl">
        Yet Another Todo App
      </h2>
    </div>
  </header>
);
