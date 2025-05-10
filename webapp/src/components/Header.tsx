import type { FC } from "react";

export const Header: FC = () => (
  <header>
    <div className="wrapper debug py-12" data-width="wide">
      <h1 className="text-6xl">YATA</h1>
      <h2 className="text-4xl">Yet Another Todo App</h2>
    </div>
  </header>
);
