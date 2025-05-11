import type { FC } from "react";

export const Spinner: FC = () => {
  return (
    <div className="border-primary-200/60 relative inset-0 z-10 h-4 w-4 animate-spin rounded-full border-b-2 border-t-2" />
  );
};
