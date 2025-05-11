import type { FC } from "react";

export const Spinner: FC = () => {
  return (
    <div className="h-4 w-4 animate-spin rounded-full border-t-2 border-b-2 border-amber-200/50" />
  );
};
