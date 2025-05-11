import type { FC } from "react";

export const Header: FC = () => (
  <header>
    <div className="wrapper md:pt-18 pb-8 pt-12 md:pb-12" data-width="wide">
      <h1 className="text-primary-100 text-6xl md:text-7xl">YATA</h1>
      <h2 className="text-primary-100 text-xl md:text-3xl">
        Yet Another Todo App
      </h2>
    </div>
  </header>
);
