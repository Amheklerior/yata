import { createContext } from "react";

/* 
  A globally accessible context to read and fire notifications
*/

export const NotificationCtx = createContext<{
  message: string;
  notify: (message: string) => void;
  clear: () => void;
}>({
  message: "",
  notify: () => {},
  clear: () => {},
});
