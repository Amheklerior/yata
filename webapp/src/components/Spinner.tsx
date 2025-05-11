import type { FC } from "react";

export const Spinner: FC = () => {
  return (
    <div className="relative inset-0 z-10 h-4 w-4 animate-spin rounded-full border-t-2 border-b-2 border-amber-200/60" />
  );
};
