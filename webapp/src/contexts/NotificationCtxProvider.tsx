import { useState, type FC, type PropsWithChildren } from "react";
import { NotificationCtx } from "./notificationCtx";

export const NotificationProvider: FC<PropsWithChildren> = ({ children }) => {
  const [message, setMessage] = useState("");

  return (
    <NotificationCtx.Provider
      value={{
        message,
        notify: (msg) => setMessage(msg),
        clear: () => setMessage(""),
      }}
    >
      {children}
    </NotificationCtx.Provider>
  );
};
